package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	codeService domain.CodeService
	repository  domain.ProfileRepository
	redis       domain.ProfileRedisRepository
}

// NewService creates a new service
func NewService(
	codeService domain.CodeService,
	repository domain.ProfileRepository,
	redis domain.ProfileRedisRepository,
	logger log.Logger,
) domain.ProfileService {
	var service domain.ProfileService
	{
		service = newBasicService(
			codeService,
			repository,
			redis,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	codeService domain.CodeService,
	repository domain.ProfileRepository,
	redis domain.ProfileRedisRepository,
) domain.ProfileService {
	return &service{
		codeService: codeService,
		repository:  repository,
		redis:       redis,
	}
}

func (s *service) GetProfile(
	ctx context.Context,
	userID domain.UserID,
	_ domain.CallerID,
) (*domain.Profile, error) {

	// Validate user id
	//
	if userID <= 0 {
		return nil, errors.NewErrInvalidArgument("user id required")
	}

	return s.repository.GetProfile(ctx, userID)
}

func (s *service) UpdateProfile(
	ctx context.Context,
	userID domain.UserID,
	profile *domain.Profile,
	_ domain.CallerID,
) error {

	// Validates inputs
	//
	if userID <= 0 {
		return errors.NewErrInvalidArgument("user id required")
	}

	// Assign id
	{
		profile.ID = userID
	}

	err := s.validateProfileInternal(ctx, profile)
	if err != nil {
		return err
	}

	// Check if user has changed the phone number
	hasPhoneChanged, err := s.hasUserChangedPhone(ctx, userID, profile)
	if err != nil {
		return err
	}

	if hasPhoneChanged {
		// Check if phone exists in a storage
		//
		isPhoneExist, err := s.repository.IsPhoneExist(ctx, profile.Phone)
		if err != nil {
			return err
		}
		if isPhoneExist {
			return errors.NewErrAlreadyExist(
				fmt.Sprintf("phone %s already exists in a storage", profile.Phone),
			)
		}

		// Send code
		//
		err = s.codeService.SendCode(ctx, profile.Phone)
		if err != nil {
			return err
		}

		// Store temporary user data in a redis
		//
		err = s.redis.Set(ctx, profile.Phone, profile, domain.RedisUserDuration)
		if err != nil {
			return err
		}

		return nil
	}

	return s.repository.UpdateProfile(ctx, userID, profile)
}

func (s *service) validateProfileInternal(
	ctx context.Context,
	profile *domain.Profile,
) error {

	if len(profile.Email) > 0 {
		isEmailExist, err := s.repository.IsEmailExist(ctx, profile.Email)
		if err != nil {
			return err
		}
		if isEmailExist {
			return errors.NewErrAlreadyExist("email exists")
		}
	}

	if len(profile.Phone) > 0 {
		isPhoneExist, err := s.repository.IsPhoneExist(ctx, profile.Phone)
		if err != nil {
			return err
		}
		if isPhoneExist {
			return errors.NewErrAlreadyExist("phone exists")
		}
	}

	return nil
}

func (s *service) hasUserChangedPhone(
	ctx context.Context,
	userID domain.UserID,
	profile *domain.Profile,
) (bool, error) {
	user, err := s.repository.GetProfile(ctx, userID)
	if err != nil {
		return false, err
	}

	if len(profile.Phone) == 0 {
		return false, nil
	}

	if user.Phone != profile.Phone {
		return true, nil
	}

	return false, nil
}

func (s *service) ValidatePhone(
	ctx context.Context,
	code string,
	callerID domain.CallerID,
) error {

	// Validate inputs
	//
	if len(code) == 0 {
		return errors.NewErrInvalidArgument("code required")
	}

	// Check is code valid
	//
	phone, err := s.codeService.ValidateCode(ctx, code)
	if err != nil {
		return err
	}

	// Get profile by code
	//
	profile, err := s.redis.Get(ctx, phone)
	if err != nil {
		return err
	}

	// Update profile
	//
	return s.repository.UpdateProfile(ctx, profile.ID, profile)
}
