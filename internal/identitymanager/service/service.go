package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

type service struct {
	codesService    domain.CodeService
	mailService     domain.MailService
	repository      domain.UserRepository
	redisRepository domain.IdentityManagerRedisRepository
}

// NewService creates a new service
func NewService(
	codesService domain.CodeService,
	mailService domain.MailService,
	repository domain.UserRepository,
	redisRepository domain.IdentityManagerRedisRepository,
	logger log.Logger,
) domain.IdentityManagerService {
	var service domain.IdentityManagerService
	{
		service = newBasicService(
			codesService,
			mailService,
			repository,
			redisRepository,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	codesService domain.CodeService,
	mailService domain.MailService,
	repository domain.UserRepository,
	redisRepository domain.IdentityManagerRedisRepository,
) domain.IdentityManagerService {
	return &service{
		codesService:    codesService,
		mailService:     mailService,
		repository:      repository,
		redisRepository: redisRepository,
	}
}

func (s *service) CreateUser(
	ctx context.Context,
	user *domain.User,
) error {

	// Validate inputs
	//
	if err := s.validateInternalCustomer(ctx, user); err != nil {
		return err
	}

	// Store user in a redis storage
	//
	err := s.redisRepository.Set(ctx, user.Phone, user, domain.RedisUserDuration)
	if err != nil {
		return err
	}

	// Send verification code to user
	//
	err = s.codesService.SendVerificationCode(
		ctx,
		user.Phone,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ValidateUser(
	ctx context.Context,
	email string,
	password string,
) (*domain.User, error) {

	// Validate inputs
	//
	if len(email) == 0 {
		return nil, errors.NewErrInvalidArgument("email required")
	}

	if len(password) == 0 {
		return nil, errors.NewErrInvalidArgument("password required")
	}

	// Check if email exist in a storage
	//
	isEmailExist, err := s.repository.IsEmailExist(ctx, email)
	if err != nil {
		return nil, err
	}

	if !isEmailExist {
		return nil, errors.NewErrNotFound("wrong email")
	}

	// Get user by email and password
	//
	foundUser, err := s.repository.GetUserByEmailPassword(ctx, email, password)
	if err != nil {
		return nil, errors.NewErrNotFound("wrong password")
	}

	// Check if user is verified
	if !foundUser.IsVerified {
		return nil, errors.NewErrPermissionDenied("user is not verified")
	}

	return foundUser, nil
}

func (s *service) ValidateCode(
	ctx context.Context,
	code string,
) (domain.UserID, domain.RoleID, error) {

	// Validate inputs
	//
	if len(code) == 0 {
		return 0, 0, errors.NewErrInvalidArgument("code required")
	}

	// Validate code
	//
	phone, err := s.codesService.ValidateCode(ctx, code)
	if err != nil {
		return 0, 0, errors.NewErrNotFound(
			fmt.Sprintf("code not found: %v", err),
		)
	}

	// Get user from a cache by phone number
	//
	user, err := s.redisRepository.Get(ctx, phone)
	if err != nil {
		return 0, 0, err
	}

	// Create user in a storage
	//
	createdUser, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return 0, 0, err
	}

	// Make user verified
	//
	err = s.repository.MakeUserVerified(ctx, createdUser.ID)
	if err != nil {
		return 0, 0, err
	}

	return createdUser.ID, createdUser.RoleID, nil
}

func (s *service) SendResetPasswordToken(
	ctx context.Context,
	email string,
) error {
	// Validate inputs
	//
	if len(email) == 0 {
		return errors.NewErrInvalidArgument("email required")
	}

	// Check email existence
	//
	isExist, err := s.repository.IsEmailExist(ctx, email)
	if err != nil {
		return errors.NewErrNotFound(
			fmt.Sprintf("email not found:, details: %v", err),
		)
	}
	if !isExist {
		return errors.NewErrNotFound("email not found")
	}

	user, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return errors.NewErrNotFound(
			fmt.Sprintf("user by email %s not found", email),
		)
	}

	// Generate random token
	//
	token, err := uuid.NewRandom()
	if err != nil {
		return errors.NewErrInternal(
			fmt.Sprintf("token does not generated, details %v", err),
		)
	}

	// Store token in a storage
	//
	err = s.repository.CreateResetPasswordToken(ctx, &domain.ResetPasswordToken{
		UserID: user.ID,
		Token:  token.String(),
	})
	if err != nil {
		return errors.NewErrInternal(
			fmt.Sprintf("reset password token does not stored, details: %v", err),
		)
	}

	// Send email
	//
	err = s.mailService.SendResetPasswordEmail(
		ctx,
		email,
		domain.ResetMailPasswordEmailContent{
			Token: token.String(),
		},
	)
	if err != nil {
		return errors.NewErrInternal(
			fmt.Sprintf("email did not send to email %s, details: %v", email, err),
		)
	}

	return nil
}

func (s *service) ResetPassword(
	ctx context.Context,
	token string,
	newPassword string,
) error {

	// Validate inputs
	//
	if len(token) == 0 {
		return errors.NewErrInvalidArgument("token required")
	}
	if len(newPassword) == 0 {
		return errors.NewErrInvalidArgument("password required")
	}

	// Get reset password token
	//
	resetPasswordToken, err := s.repository.GetResetPasswordToken(ctx, token)
	if err != nil {
		return errors.NewErrNotFound(
			fmt.Sprintf("token %s not found", token),
		)
	}

	// Check reset password confirmation
	//
	if resetPasswordToken.Confirmed {
		return errors.NewErrNotFound(
			fmt.Sprintf("token %s is already used", token),
		)
	}

	// Update reset password token
	//
	err = s.repository.UpdateResetPasswordToken(ctx, &domain.ResetPasswordToken{
		UserID:    resetPasswordToken.UserID,
		Token:     token,
		Confirmed: true,
	})
	if err != nil {
		return err
	}

	// Reset password
	//
	err = s.repository.ResetPassword(
		ctx,
		resetPasswordToken.UserID,
		newPassword,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) validateInternalCustomer(ctx context.Context, user *domain.User) error {
	if user == nil {
		return fmt.Errorf("user required")
	}

	if user.RoleID != domain.RoleClient {
		return errors.NewErrInvalidArgument("invalid role id")
	}
	if len(user.FirstName) == 0 {
		return errors.NewErrInvalidArgument("first name required")
	}
	if len(user.LastName) == 0 {
		return errors.NewErrInvalidArgument("last name required")
	}
	if len(user.Phone) == 0 {
		return errors.NewErrInvalidArgument("phone required")
	}
	if len(user.Email) == 0 {
		return errors.NewErrInvalidArgument("email required")
	}
	if len(user.Password) == 0 {
		return errors.NewErrInvalidArgument("password required")
	}

	// Check if email exist in a storage
	//
	isEmailExist, err := s.repository.IsEmailExist(ctx, user.Email)
	if err != nil {
		return err
	}
	if isEmailExist {
		return errors.NewErrAlreadyExist("email exists in system")
	}

	// Check if phone number exist in a storage
	//
	isPhoneExist, err := s.repository.IsPhoneNumberExist(ctx, user.Phone)
	if err != nil {
		return err
	}
	if isPhoneExist {
		return errors.NewErrAlreadyExist("phone exist")
	}

	return nil
}
