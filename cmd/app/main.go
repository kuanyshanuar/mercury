package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	identityapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/identityserviceapi/v1"
	residencesapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	pkgBuildersRepository "gitlab.com/zharzhanov/mercury/internal/builder/repository"
	pkgBuildersService "gitlab.com/zharzhanov/mercury/internal/builder/service"
	pkgBuildersTransport "gitlab.com/zharzhanov/mercury/internal/builder/transport"
	pkgCodesRepository "gitlab.com/zharzhanov/mercury/internal/codeservice/repository"
	pkgCodesService "gitlab.com/zharzhanov/mercury/internal/codeservice/service"
	pkgCodesTransport "gitlab.com/zharzhanov/mercury/internal/codeservice/transport"
	pkgContactsRepository "gitlab.com/zharzhanov/mercury/internal/contactdetails/repository"
	pkgContactsService "gitlab.com/zharzhanov/mercury/internal/contactdetails/service"
	pkgContactsTransport "gitlab.com/zharzhanov/mercury/internal/contactdetails/transport"
	pkgCottageRepository "gitlab.com/zharzhanov/mercury/internal/cottages/repository"
	pkgCottageService "gitlab.com/zharzhanov/mercury/internal/cottages/service"
	pkgCottageTransport "gitlab.com/zharzhanov/mercury/internal/cottages/transport"
	pkgFileStorageRepository "gitlab.com/zharzhanov/mercury/internal/filestorage/repository"
	pkgFileStorageService "gitlab.com/zharzhanov/mercury/internal/filestorage/service"
	pkgFileStorageTransport "gitlab.com/zharzhanov/mercury/internal/filestorage/transport"
	pkgFiltersRepository "gitlab.com/zharzhanov/mercury/internal/filters/repository"
	pkgFiltersService "gitlab.com/zharzhanov/mercury/internal/filters/service"
	pkgFiltersTransport "gitlab.com/zharzhanov/mercury/internal/filters/transport"
	pkgIdentityRepository "gitlab.com/zharzhanov/mercury/internal/identitymanager/repository"
	pkgIdentityService "gitlab.com/zharzhanov/mercury/internal/identitymanager/service"
	pkgIdentityTransport "gitlab.com/zharzhanov/mercury/internal/identitymanager/transport"
	pkgLeadsBuildersRepository "gitlab.com/zharzhanov/mercury/internal/leadsbuilders/repository"
	pkgLeadsBuildersService "gitlab.com/zharzhanov/mercury/internal/leadsbuilders/service"
	pkgLeadsBuildersTransport "gitlab.com/zharzhanov/mercury/internal/leadsbuilders/transport"
	pkgLeadCottageRepository "gitlab.com/zharzhanov/mercury/internal/leadscottages/repository"
	pkgLeadCottageService "gitlab.com/zharzhanov/mercury/internal/leadscottages/service"
	pkgLeadCottageTransport "gitlab.com/zharzhanov/mercury/internal/leadscottages/transport"
	pkgLeadRepository "gitlab.com/zharzhanov/mercury/internal/leadsresidences/repository"
	pkgLeadService "gitlab.com/zharzhanov/mercury/internal/leadsresidences/service"
	pkgLeadTransport "gitlab.com/zharzhanov/mercury/internal/leadsresidences/transport"
	pkgMailService "gitlab.com/zharzhanov/mercury/internal/mail/service"
	pkgSenderService "gitlab.com/zharzhanov/mercury/internal/mailsender/service"
	pkgManagersRepository "gitlab.com/zharzhanov/mercury/internal/manager/repository"
	pkgManagersService "gitlab.com/zharzhanov/mercury/internal/manager/service"
	pkgManagersTransport "gitlab.com/zharzhanov/mercury/internal/manager/transport"
	pkgNewsRepository "gitlab.com/zharzhanov/mercury/internal/news/repository"
	pkgNewsService "gitlab.com/zharzhanov/mercury/internal/news/service"
	pkgNewsTransport "gitlab.com/zharzhanov/mercury/internal/news/transport"
	pkgPermissionsRepository "gitlab.com/zharzhanov/mercury/internal/permissions/repository"
	pkgPermissionsService "gitlab.com/zharzhanov/mercury/internal/permissions/service"
	pkgPermissionsTransport "gitlab.com/zharzhanov/mercury/internal/permissions/transport"
	pkgProfileRepository "gitlab.com/zharzhanov/mercury/internal/profile/repository"
	pkgProfileService "gitlab.com/zharzhanov/mercury/internal/profile/service"
	pkgProfileTransport "gitlab.com/zharzhanov/mercury/internal/profile/transport"
	pkgResidencesRepository "gitlab.com/zharzhanov/mercury/internal/residences/repository"
	pkgResidencesService "gitlab.com/zharzhanov/mercury/internal/residences/service"
	pkgResidencesTransport "gitlab.com/zharzhanov/mercury/internal/residences/transport"
	pkgSmsRepository "gitlab.com/zharzhanov/mercury/internal/sms/repository"
	pkgSmsService "gitlab.com/zharzhanov/mercury/internal/sms/service"
	pkgUserBuilderRepository "gitlab.com/zharzhanov/mercury/internal/userbuilders/repository"
	pkgUserBuilderService "gitlab.com/zharzhanov/mercury/internal/userbuilders/service"
	pkgUserBuilderTransport "gitlab.com/zharzhanov/mercury/internal/userbuilders/transport"
	pkgUserCottagesRepository "gitlab.com/zharzhanov/mercury/internal/usercottages/repository"
	pkgUserCottagesService "gitlab.com/zharzhanov/mercury/internal/usercottages/service"
	pkgUserCottagesTransport "gitlab.com/zharzhanov/mercury/internal/usercottages/transport"
	pkgUserResidenceRepository "gitlab.com/zharzhanov/mercury/internal/userresidences/repository"
	pkgUserResidenceService "gitlab.com/zharzhanov/mercury/internal/userresidences/service"
	pkgUserResidencesTransport "gitlab.com/zharzhanov/mercury/internal/userresidences/transport"
	pkgUsersRepository "gitlab.com/zharzhanov/mercury/internal/users/repository"
	pkgUserService "gitlab.com/zharzhanov/mercury/internal/users/service"
	pkgUserTransport "gitlab.com/zharzhanov/mercury/internal/users/transport"
	pkgMail "gitlab.com/zharzhanov/mercury/pkg/mail"

	pkgDomain "gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
	"gitlab.com/zharzhanov/mercury/pkg/database/postgres"
	"gitlab.com/zharzhanov/mercury/pkg/database/redis"
	"gitlab.com/zharzhanov/mercury/pkg/minio"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	kitzapadapter "github.com/go-kit/kit/log/zap"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/oklog/run"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

// AppVersion - application version.
var AppVersion = "unversioned"

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
	logger = log.With(logger, "ver", AppVersion)
	// Logging helper function
	logFatal := func(err error) {
		_ = logger.Log("err", err)
		os.Exit(1)
	}

	// Reads configuration from the environment variables.
	//
	cfg, err := helpers.LoadConfig()
	if err != nil {
		logFatal(err)
	}

	if err := helpers.ValidateConfig(cfg); err != nil {
		logFatal(err)
	}

	emailCfg, err := helpers.LoadEmailConfig()
	if err != nil {
		logFatal(err)
	}
	emailTemplate, err := helpers.LoadEmailTemplate()
	if err != nil {
		logFatal(err)
	}

	// Define our flags.
	//
	fs := flag.NewFlagSet("", flag.ExitOnError)
	grpcAddr := fs.String("grpc-addr", ":"+cfg.Port, "gRPC listen address")
	err = fs.Parse(os.Args[1:])
	if err != nil {
		logFatal(err)
	}

	// Define service secret key.
	//
	var (
		serviceSecretKey = pkgDomain.ServiceSecretKey(cfg.GRPCDefaultBearerToken)
	)

	var (
		ctx = context.Background()
	)

	// Init telemetry provider.
	//
	shutdown, err := telemetry.InitTelemetryProvider(cfg.OpenTelemetryCollectorURL, cfg.Environment)
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
	db, err := postgres.NewConnection(cfg.DSN)
	if err != nil {
		logFatal(err)
	}

	// Setup redis connection
	//
	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		logFatal(err)
	}

	// Setup storage client
	//
	minioClient, err := minio.NewStorageClient(cfg)
	if err != nil {
		logFatal(err)
	}

	// Setup new email authentication
	//
	smtpAuth := pkgMail.NewSMTPAuth(emailCfg)

	// Repository layer.
	//
	leadCottageRepository := pkgLeadCottageRepository.NewRepository(db)
	cottageRepository := pkgCottageRepository.NewCottageRepository(db)
	newsRepository := pkgNewsRepository.NewRepository(db)
	residenceRepository := pkgResidencesRepository.NewResidenceRepository(db)
	residenceRedisRepository := pkgResidencesRepository.NewRedisRepository(redisClient)
	usersRepository := pkgUsersRepository.NewRepository(db)
	codesRedisRepository := pkgCodesRepository.NewRepository(redisClient)
	userResidenceRepository := pkgUserResidenceRepository.NewRepository(db)
	smsRepository := pkgSmsRepository.NewRepository(db)
	permissionsRepository := pkgPermissionsRepository.NewRepository(db)
	profileRepository := pkgProfileRepository.NewRepository(db)
	leadResidenceRepository := pkgLeadRepository.NewRepository(db)
	filtersRepository := pkgFiltersRepository.NewRepository(db)
	contactsRepository := pkgContactsRepository.NewRepository(db)
	userBuildersRepository := pkgUserBuilderRepository.NewRepository(db)
	builderRepository := pkgBuildersRepository.NewRepository(db)
	fileStorageRepository := pkgFileStorageRepository.NewFileStorageRepository(minioClient)
	leadsBuildersRepository := pkgLeadsBuildersRepository.NewRepository(db)
	identityRepository := pkgIdentityRepository.NewIdentityManagerRedisRepository(redisClient)
	managersRepository := pkgManagersRepository.NewRepository(db)
	profileRedisRepository := pkgProfileRepository.NewProfileRedisRepository(redisClient)
	userCottagesRepository := pkgUserCottagesRepository.NewRepository(db)

	// Service layer.
	//
	leadsCottageService := pkgLeadCottageService.NewService(
		leadCottageRepository,
		logger,
	)

	cottageService := pkgCottageService.NewService(
		cottageRepository,
		logger,
	)
	newsService := pkgNewsService.NewService(
		newsRepository,
		logger,
	)
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
	residenceService := pkgResidencesService.NewResidenceService(
		residenceRepository,
		residenceRedisRepository,
		logger,
	)
	smsService := pkgSmsService.NewService(
		smsRepository,
		cfg.SmsServiceURL,
		cfg.SmsAPIKey,
		cfg.SmsAlphaName,
		logger,
	)
	codeService := pkgCodesService.NewService(
		smsService,
		codesRedisRepository,
		usersRepository,
		logger,
	)
	identityService := pkgIdentityService.NewService(
		codeService,
		mailService,
		usersRepository,
		identityRepository,
		logger,
	)
	userResidenceService := pkgUserResidenceService.NewService(
		residenceService,
		userResidenceRepository,
		logger,
	)
	permissionsService := pkgPermissionsService.NewService(
		permissionsRepository,
		logger,
	)
	profileService := pkgProfileService.NewService(
		codeService,
		profileRepository,
		profileRedisRepository,
		logger,
	)
	leadResidenceService := pkgLeadService.NewService(
		leadResidenceRepository,
		logger,
	)
	filtersService := pkgFiltersService.NewService(
		filtersRepository,
		logger,
	)
	contactsService := pkgContactsService.NewService(
		mailService,
		contactsRepository,
		logger,
	)
	builderService := pkgBuildersService.NewService(
		builderRepository,
		logger,
	)
	userBuilderService := pkgUserBuilderService.NewService(
		userBuildersRepository,
		logger,
	)
	fileStorageService := pkgFileStorageService.NewService(
		fileStorageRepository,
		cfg.BucketName,
		cfg.MinioBaseURL,
		logger,
	)
	leadBuilderService := pkgLeadsBuildersService.NewService(
		leadsBuildersRepository,
		logger,
	)
	managersService := pkgManagersService.NewService(
		managersRepository,
		logger,
	)
	usersService := pkgUserService.NewService(
		usersRepository,
		logger,
	)
	userCottagesService := pkgUserCottagesService.NewService(
		userCottagesRepository,
		cottageService,
		logger,
	)

	// Endpoints layer.
	//
	leadsCottageEndpoint := pkgLeadCottageTransport.NewEndpoint(leadsCottageService, serviceSecretKey, logger)
	cottageEndpoint := pkgCottageTransport.NewEndpoint(cottageService, serviceSecretKey, logger)
	newsEndpoint := pkgNewsTransport.NewEndpoint(newsService, serviceSecretKey, logger)
	residenceEndpoint := pkgResidencesTransport.NewEndpoint(residenceService, serviceSecretKey, logger)
	identityEndpoint := pkgIdentityTransport.NewEndpoint(identityService, serviceSecretKey, logger)
	userResidenceEndpoint := pkgUserResidencesTransport.NewEndpoint(userResidenceService, serviceSecretKey, logger)
	permissionEndpoint := pkgPermissionsTransport.NewEndpoint(permissionsService, serviceSecretKey, logger)
	profileEndpoint := pkgProfileTransport.NewEndpoint(profileService, serviceSecretKey, logger)
	leadResidenceEndpoint := pkgLeadTransport.NewEndpoint(leadResidenceService, serviceSecretKey, logger)
	filtersEndpoint := pkgFiltersTransport.NewEndpoint(filtersService, serviceSecretKey, logger)
	contactsEndpoint := pkgContactsTransport.NewEndpoint(contactsService, serviceSecretKey, logger)
	buildersEndpoint := pkgBuildersTransport.NewEndpoint(builderService, serviceSecretKey, logger)
	userBuilderEndpoint := pkgUserBuilderTransport.NewEndpoint(userBuilderService, serviceSecretKey, logger)
	codesEndpoint := pkgCodesTransport.NewEndpoint(codeService, serviceSecretKey, logger)
	fileStorageEndpoint := pkgFileStorageTransport.NewEndpoint(fileStorageService, serviceSecretKey, logger)
	leadBuildersEndpoint := pkgLeadsBuildersTransport.NewEndpoint(leadBuilderService, serviceSecretKey, logger)
	managersEndpoint := pkgManagersTransport.NewEndpoint(managersService, serviceSecretKey, logger)
	usersEndpoint := pkgUserTransport.NewEndpoint(usersService, serviceSecretKey, logger)
	userCottagesEndpoint := pkgUserCottagesTransport.NewEndpoints(userCottagesService, serviceSecretKey, logger)

	// Grpc server layer.
	//
	leadsCottageGrpcServer := pkgLeadCottageTransport.NewGRPCServer(leadsCottageEndpoint, logger)
	cottageGrpcServer := pkgCottageTransport.NewGRPCServer(cottageEndpoint, logger)
	newsGrpcServer := pkgNewsTransport.NewGRPCServer(newsEndpoint, logger)
	residenceGrpcServer := pkgResidencesTransport.NewGRPCServer(residenceEndpoint, logger)
	identityGrpcServer := pkgIdentityTransport.NewGRPCServer(identityEndpoint, logger)
	userResidenceGrpcServer := pkgUserResidencesTransport.NewGRPCServer(userResidenceEndpoint, logger)
	permissionsGrpcServer := pkgPermissionsTransport.NewGRPCServer(permissionEndpoint, logger)
	profileGrpcServer := pkgProfileTransport.NewGRPCServer(profileEndpoint, logger)
	leadGrpcServer := pkgLeadTransport.NewGRPCServer(leadResidenceEndpoint, logger)
	filtersGrpcServer := pkgFiltersTransport.NewGRPCServer(filtersEndpoint, logger)
	contactsGrpcServer := pkgContactsTransport.NewGRPCServer(contactsEndpoint, logger)
	buildersGrpcServer := pkgBuildersTransport.NewGRPCServer(buildersEndpoint, logger)
	userBuilderGrpcServer := pkgUserBuilderTransport.NewGRPCServer(userBuilderEndpoint, logger)
	codesGrpcServer := pkgCodesTransport.NewGRPCServer(codesEndpoint, logger)
	fileStorageGrpcServer := pkgFileStorageTransport.NewGRPCServerV1(fileStorageEndpoint, logger)
	leadBuildersGrpcServer := pkgLeadsBuildersTransport.NewGRPCServer(leadBuildersEndpoint, logger)
	managersGrpcServer := pkgManagersTransport.NewGRPCServer(managersEndpoint, logger)
	systemGrpcServer := pkgUserTransport.NewGRPCServer(usersEndpoint, logger)
	userCottagesGrpcServer := pkgUserCottagesTransport.NewGrpcServer(userCottagesEndpoint, logger)

	// Setup base grpc-server.
	//
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		kitgrpc.Interceptor,
	), helpers.ServerKeepaliveParams)
	residencesapiv1.RegisterLeadCottageServiceServer(grpcServer, leadsCottageGrpcServer)
	residencesapiv1.RegisterCottageServiceServer(grpcServer, cottageGrpcServer)
	residencesapiv1.RegisterNewsServiceServer(grpcServer, newsGrpcServer)
	residencesapiv1.RegisterResidenceServiceServer(grpcServer, residenceGrpcServer)
	identityapiv1.RegisterAuthenticationServiceServer(grpcServer, identityGrpcServer)
	residencesapiv1.RegisterUserResidenceServiceServer(grpcServer, userResidenceGrpcServer)
	identityapiv1.RegisterPermissionServiceServer(grpcServer, permissionsGrpcServer)
	identityapiv1.RegisterProfileServiceServer(grpcServer, profileGrpcServer)
	residencesapiv1.RegisterLeadResidencesServiceServer(grpcServer, leadGrpcServer)
	residencesapiv1.RegisterResidenceFilterServiceServer(grpcServer, filtersGrpcServer)
	residencesapiv1.RegisterContactDetailsServiceServer(grpcServer, contactsGrpcServer)
	residencesapiv1.RegisterBuilderServiceServer(grpcServer, buildersGrpcServer)
	residencesapiv1.RegisterSubscribersServiceServer(grpcServer, userBuilderGrpcServer)
	identityapiv1.RegisterCodeServiceServer(grpcServer, codesGrpcServer)
	residencesapiv1.RegisterFileStorageServiceServer(grpcServer, fileStorageGrpcServer)
	residencesapiv1.RegisterLeadBuildersServiceServer(grpcServer, leadBuildersGrpcServer)
	residencesapiv1.RegisterManagerServiceServer(grpcServer, managersGrpcServer)
	residencesapiv1.RegisterSystemServiceServer(grpcServer, systemGrpcServer)
	residencesapiv1.RegisterUserCottagesServiceServer(grpcServer, userCottagesGrpcServer)

	var g run.Group
	// Startup the gRPC listener
	{
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			logFatal(err)
		}

		g.Add(func() error {
			_ = logger.Log("transport", "gRPC", "addr", *grpcAddr)
			return grpcServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	// This function just sits and waits for ctrl-C.
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
