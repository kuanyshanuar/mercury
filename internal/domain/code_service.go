package domain

import "context"

// RedisCodeDuration - duration of storing sms code
const (
	RedisCodeDuration = 1080
	RedisUserDuration = 1080
)

// Code - represents code struct
type Code struct {
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	ExpiresAt int64  `json:"expires_at"`
}

// CodeRedisRepository - provides access to storage.
type CodeRedisRepository interface {
	// Get - returns residence from cache by key.
	//
	Get(
		ctx context.Context,
		key string,
	) (*Code, error)

	// Set - sets a residence to cache with specific key.
	//
	Set(
		ctx context.Context,
		key string,
		value *Code,
		seconds int,
	) error

	// Delete - deletes a residence from cache by key.
	//
	Delete(
		ctx context.Context,
		key string,
	) error
}

// CodeService -  provides access to business logic.
type CodeService interface {
	// SendVerificationCode - sends a verification code.
	//
	SendVerificationCode(
		ctx context.Context,
		phone string,
	) error

	// SendCode - sends a code.
	//
	SendCode(
		ctx context.Context,
		phone string,
	) error

	// ValidateCode - validates code if exists in redis storage.
	//
	ValidateCode(
		ctx context.Context,
		code string,
	) (string, error)
}
