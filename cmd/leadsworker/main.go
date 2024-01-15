package main

import (
	"context"
	"fmt"
	pkgLeadCottageRepository "gitlab.com/zharzhanov/mercury/internal/leadscottages/repository"
	pkgLeadCottageService "gitlab.com/zharzhanov/mercury/internal/leadscottages/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	pkgCronService "gitlab.com/zharzhanov/mercury/internal/cronlead/service"
	pkgHelper "gitlab.com/zharzhanov/mercury/internal/helpers"
	pkgLeadBuilderRepository "gitlab.com/zharzhanov/mercury/internal/leadsbuilders/repository"
	pkgLeadBuilderService "gitlab.com/zharzhanov/mercury/internal/leadsbuilders/service"
	pkgLeadResidenceRepository "gitlab.com/zharzhanov/mercury/internal/leadsresidences/repository"
	pkgLeadResidenceService "gitlab.com/zharzhanov/mercury/internal/leadsresidences/service"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
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

	// Repository layer.
	leadCottageRepository := pkgLeadCottageRepository.NewRepository(db)
	leadResidenceRepository := pkgLeadResidenceRepository.NewRepository(db)
	leadBuilderRepository := pkgLeadBuilderRepository.NewRepository(db)

	// Service layer.
	leadCottageService := pkgLeadCottageService.NewService(
		leadCottageRepository,
		logger,
	)
	leadResidenceService := pkgLeadResidenceService.NewService(
		leadResidenceRepository,
		logger,
	)
	leadBuilderService := pkgLeadBuilderService.NewService(
		leadBuilderRepository,
		logger,
	)
	cronService := pkgCronService.NewService(
		leadCottageService,
		leadResidenceService,
		leadBuilderService,
		logger,
	)

	// Set cron
	//
	cronStub := cron.New()

	// Schedule tasks
	cronStub.AddFunc("30 0 * * *", func() {
		cronService.RevokeExpiredLeadCottages(ctx, pkgHelper.CallerID(ctx))
	})

	cronStub.AddFunc("30 0 * * *", func() {
		cronService.RevokeExpiredLeadResidences(ctx, pkgHelper.CallerID(ctx))
	})

	// Schedule tasks
	cronStub.AddFunc("30 0 * * *", func() {
		cronService.RevokeExpiredLeadBuilders(ctx, pkgHelper.CallerID(ctx))
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
