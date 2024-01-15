package domain

// Logging fields for tracing.
const (
	LogFieldTraceID = "traceId"
	LogFieldSpanID  = "spanId"
)

// ServiceSecretKey - represents a service secret key.
type ServiceSecretKey string

// CallerID - the endpoint caller id.
type CallerID struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}

// CanonicalEnvironments - represents environment modes
var (
	CanonicalEnvironments = map[string]bool{
		"dev":        true,
		"staging":    true,
		"production": true,
	}
)

// AppConfig - represents application configuration
type AppConfig struct {
	Environment            string `env:"ENVIRONMENT" validate:"required"`
	Port                   string `env:"PORT" validate:"required"`
	GRPCDefaultBearerToken string `env:"GRPC_DEFAULT_BEARER_TOKEN" validate:"required"`
	DSN                    string `env:"DSN" validate:"required"`

	RedisHost     string `env:"REDIS_HOST" validate:"required"`
	RedisPort     string `env:"REDIS_PORT" validate:"required"`
	RedisPassword string `env:"REDIS_PASSWORD" validate:"required"`
	RedisPoolSize int    `env:"REDIS_POOLSIZE" validate:"required"`

	PrometheusURL string `env:"PROMETHEUS_URL" validate:"required"`
	ServiceName   string `env:"SERVICE_NAME" validate:"required"`

	MinioEndpoint        string `env:"MINIO_ENDPOINT" validate:"required"`
	MinioAccessKeyID     string `env:"MINIO_ACCESS_KEY_ID" validate:"required"`
	MinioSecretAccessKey string `env:"MINIO_SECRET_ACCESS_KEY" validate:"required"`
	BucketName           string `env:"BUCKET_NAME" validate:"required"`
	MinioBaseURL         string `env:"MINIO_BASE_URL" validate:"required"`
	//MinioUseSSL          bool   `env:"MINIO_USE_SSL" validate:"required"`

	//JaegerLogSpans bool   `env:"LOG_SPANS" validate:"required"`
	JaegerHost string `env:"JAEGER_HOST" validate:"required"`

	OpenTelemetryCollectorURL string `env:"OPEN_TELEMETRY_COLLECTOR_URL" validate:"required"`

	SmsServiceURL string `env:"SMS_SERVICE_URL" validate:"required"`
	SmsAlphaName  string `env:"SMS_ALPHA_NAME" validate:"required"`
	SmsAPIKey     string `env:"SMS_SERVICE_API_KEY" validate:"required"`

	CrmBaseURL string `env:"CRM_BASE_URL" validate:"required"`
}

// EmailConfig - represents email configuration
type EmailConfig struct {
	Email    string `env:"EMAIL" validate:"required"`
	Password string `env:"EMAIL_PASSWORD" validate:"required"`
	Host     string `env:"EMAIL_HOST" validate:"required"`
	Port     string `env:"EMAIL_PORT" validate:"required"`
}

// EmailTemplate - email templates
type EmailTemplate struct {
	ResetPasswordTemplate          string `env:"RESET_PASSWORD_TEMPLATE" validate:"required"`
	ContactDetailTemplate          string `env:"CONTACT_DETAIL_TEMPLATE" validate:"required"`
	ResidenceContactDetailTemplate string `env:"RESIDENCE_CONTACT_DETAIL_TEMPLATE" validate:"required"`
}
