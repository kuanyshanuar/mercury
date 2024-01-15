package service

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	smsService     domain.SmsService
	repository     domain.CodeRedisRepository
	userRepository domain.UserRepository
}

// NewService - creates a new service
func NewService(
	smsService domain.SmsService,
	repository domain.CodeRedisRepository,
	userRepository domain.UserRepository,
	logger log.Logger,
) domain.CodeService {
	var service domain.CodeService
	{
		service = newBasicService(
			smsService,
			repository,
			userRepository,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	smsService domain.SmsService,
	repository domain.CodeRedisRepository,
	userRepository domain.UserRepository,
) domain.CodeService {
	return &service{
		smsService:     smsService,
		repository:     repository,
		userRepository: userRepository,
	}
}

func (s *service) SendVerificationCode(
	ctx context.Context,
	phone string,
) error {
	// Validates input
	//
	if len(phone) == 0 {
		return errors.NewErrInvalidArgument("phone required")
	}

	// Generates code
	//
	generatedCode := s.assertGenerateCode()

	code := &domain.Code{
		Code:      generatedCode,
		Phone:     phone,
		ExpiresAt: time.Now().Add(time.Second * domain.RedisCodeDuration).Unix(),
	}

	// Store code in a storage
	//
	err := s.repository.Set(ctx, generatedCode, code, domain.RedisCodeDuration)
	if err != nil {
		return err
	}

	// Sends verification sms to phone
	//
	err = s.smsService.SendSms(ctx, &domain.Sms{
		Phone:     phone,
		Message:   generatedCode,
		ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SendCode(
	ctx context.Context,
	phone string,
) error {

	// Validates input
	//
	if len(phone) == 0 {
		return errors.NewErrInvalidArgument("phone required")
	}

	// Generates code
	//
	generatedCode := s.assertGenerateCode()
	code := &domain.Code{
		Code:      generatedCode,
		Phone:     phone,
		ExpiresAt: time.Now().Add(time.Second * domain.RedisCodeDuration).Unix(),
	}

	// Store code in a storage
	//
	err := s.repository.Set(ctx, generatedCode, code, domain.RedisCodeDuration)
	if err != nil {
		return err
	}

	// Sends verification sms to phone
	//
	err = s.smsService.SendSms(ctx, &domain.Sms{
		Phone:     phone,
		Message:   generatedCode,
		ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ValidateCode(
	ctx context.Context,
	code string,
) (string, error) {

	// Validate code
	//
	if len(code) == 0 {
		return "", errors.NewErrInvalidArgument("code required ")
	}

	// Get code from storage
	//
	storedCode, err := s.repository.Get(ctx, code)
	if err != nil {
		return "", errors.NewErrNotFound(
			fmt.Sprintf("code not found: %v", err),
		)
	}
	if storedCode == nil {
		return "", errors.NewErrNotFound("code not found")
	}

	// Check if entered code matches
	//
	if code != storedCode.Code {
		return "", errors.NewErrNotFound("code not found")
	}

	// Returns user phone
	//
	return storedCode.Phone, nil
}

func (s *service) assertGenerateCode() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}
