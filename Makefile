##
# Important! Before running any command make sure you have setup GOPATH:
# export GOPATH="$HOME/go"
# PATH="$GOPATH/bin:$PATH"

start:
	# Start the application with postgresql database
	./scripts/start.sh

generate-api:
	# Generate proto stubs.
	./scripts/generate-api.sh self

unittest:
	go test ./...

testmocks:
	mockgen \
		-package mocks -destination=internal/mocks/mock_residences_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain ResidencesRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_residences_redis_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain ResidenceRedisRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_residences_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain ResidencesService

	mockgen \
		-package mocks -destination=internal/mocks/mock_sms_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain SmsService

	mockgen \
		-package mocks -destination=internal/mocks/mock_code_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain CodeService

	mockgen \
		-package mocks -destination=internal/mocks/mock_code_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain CodeRedisRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_permissions_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain PermissionsRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_user_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain UserRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_identity_manager_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain IdentityManagerService

	mockgen \
		-package mocks -destination=internal/mocks/mock_mail_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain MailService

	mockgen \
		-package mocks -destination=internal/mocks/mock_filters_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain FiltersService

	mockgen \
		-package mocks -destination=internal/mocks/mock_filters_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain FiltersRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_identity_manager_redis_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain IdentityManagerRedisRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_crm_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain CrmService

	mockgen \
		-package mocks -destination=internal/mocks/mock_http_client_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain HTTPClientService

	mockgen \
		-package mocks -destination=internal/mocks/mock_lead_builders_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain LeadBuilderRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_lead_builders_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain LeadBuilderService

	mockgen \
		-package mocks -destination=internal/mocks/mock_lead_residences_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain LeadResidenceRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_lead_residences_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain LeadResidenceService

	mockgen \
		-package mocks -destination=internal/mocks/mock_news_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain NewsRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_сcontact_details_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain ContactDetailsRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_сcontact_details_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain ContactDetailsService

	mockgen \
		-package mocks -destination=internal/mocks/mock_amo_crm_service.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain AmoCrmService

	mockgen \
		-package mocks -destination=internal/mocks/mock_profile_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain ProfileRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_profile_redis_repository.go \
		-package mocks gitlab.com/zharzhanov/mercury/internal/domain ProfileRedisRepository