package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	pkgAmoCrmService "gitlab.com/zharzhanov/mercury/internal/amocrm/service"
	pkgContactDetailsRepository "gitlab.com/zharzhanov/mercury/internal/contactdetails/repository"
	pkgContactDetailsService "gitlab.com/zharzhanov/mercury/internal/contactdetails/service"
	pkgCrmService "gitlab.com/zharzhanov/mercury/internal/crm/service"
	pkgCronService "gitlab.com/zharzhanov/mercury/internal/croncontactdetails/service"
	pkgHelper "gitlab.com/zharzhanov/mercury/internal/helpers"
	pkgHttpClientService "gitlab.com/zharzhanov/mercury/internal/httpclient/service"
	pkgMailService "gitlab.com/zharzhanov/mercury/internal/mail/service"
	pkgSenderService "gitlab.com/zharzhanov/mercury/internal/mailsender/service"
	pkgResidencesRepository "gitlab.com/zharzhanov/mercury/internal/residences/repository"
	pkgResidencesService "gitlab.com/zharzhanov/mercury/internal/residences/service"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
	pkgRedis "gitlab.com/zharzhanov/mercury/pkg/database/redis"
	pkgMail "gitlab.com/zharzhanov/mercury/pkg/mail"
	pkgTelemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	kitzapadapter "github.com/go-kit/kit/log/zap"
	"github.com/go-kit/log"
	"github.com/oklog/run"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	// Create a single logger, which we'll use and give to other components.
	//
	zapLogger, _ := zap.NewProduction()
	defer func() {
		_ = zapLogger.Sync()
	}()

	var logger log.Logger
	logger = kitzapadapter.NewZapSugarLogger(zapLogger, zapcore.InfoLevel)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	// Logging helper function
	logFatal := func(err error) {
		_ = logger.Log("err", err)
		os.Exit(1)
	}

	var (
		ctx = context.Background()
	)

	// Reads configuration from the environment variables.
	//
	cfg, err := pkgHelper.LoadConfig()
	if err != nil {
		logFatal(err)
	}

	if err := pkgHelper.ValidateConfig(cfg); err != nil {
		logFatal(err)
	}

	emailCfg, err := pkgHelper.LoadEmailConfig()
	if err != nil {
		logFatal(err)
	}
	emailTemplate, err := pkgHelper.LoadEmailTemplate()
	if err != nil {
		logFatal(err)
	}

	// Init telemetry provider.
	//
	shutdown, err := pkgTelemetry.InitTelemetryProvider(cfg.OpenTelemetryCollectorURL, cfg.Environment)
	if err != nil {
		logFatal(err)
	}

	// Cleanly shutdown and flush telemetry when the application exits.
	//
	defer func(ctx context.Context) {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := shutdown(ctx); err != nil {
			logFatal(err)
		}
	}(ctx)

	// Setup database connection
	//
	db, err := pkgPostgres.NewConnection(cfg.DSN)
	if err != nil {
		logFatal(err)
	}

	// Setup redis connection
	//
	redisClient, err := pkgRedis.NewRedisClient(cfg)
	if err != nil {
		logFatal(err)
	}

	// Setup new email authentication
	//
	smtpAuth := pkgMail.NewSMTPAuth(emailCfg)

	// Repository layer.
	contactDetailsRepository := pkgContactDetailsRepository.NewRepository(db)
	residencesRepository := pkgResidencesRepository.NewResidenceRepository(db)
	residenceRedisRepository := pkgResidencesRepository.NewRedisRepository(redisClient)

	// Service layer.
	httpClientService := pkgHttpClientService.NewService(logger)
	senderService := pkgSenderService.NewService(
		smtpAuth,
		emailCfg.Host+":"+emailCfg.Port,
		emailCfg.Email,
		logger,
	)
	mailService := pkgMailService.NewService(
		senderService,
		emailTemplate,
		logger,
	)
	crmService := pkgCrmService.NewService(
		cfg.CrmBaseURL,
		cfg.Environment,
		httpClientService,
		logger,
	)
	amoCrmService := pkgAmoCrmService.NewService(
		httpClientService, logger,
	)
	contactDetailsService := pkgContactDetailsService.NewService(
		mailService,
		contactDetailsRepository,
		logger,
	)
	residenceService := pkgResidencesService.NewResidenceService(
		residencesRepository,
		residenceRedisRepository,
		logger,
	)
	cronService := pkgCronService.NewService(
		contactDetailsService,
		crmService,
		amoCrmService,
		residenceService,
		logger,
	)

	// Set cron
	//
	cronStub := cron.New()

	// Schedule tasks
	cronStub.AddFunc("*/5 * * * *", func() {
		err := cronService.SendContactsToCRM(ctx, pkgHelper.CallerID(ctx))
		if err != nil {
			_ = logger.Log("err",
				fmt.Sprintf("cron returned error: %v", err),
			)
		}
	})

	// Start the scheduler
	cronStub.Start()

	var g run.Group
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	_ = logger.Log("exit", g.Run())
}
