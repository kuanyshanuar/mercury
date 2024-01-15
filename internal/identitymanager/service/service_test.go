package service

import (
	"context"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"github.com/golang/mock/gomock"
)

func TestService_CreateUser(t *testing.T) {

	var (
		ctx  = context.Background()
		user = &domain.User{
			RoleID:     1,
			FirstName:  "test",
			LastName:   "test",
			Email:      "test",
			City:       "test",
			Phone:      "+77771112233",
			Image:      "test",
			Password:   "test",
			IsVerified: false,
			IsBanned:   helpers.PointerBool(false),
		}
		user2 = &domain.User{
			RoleID:     1,
			FirstName:  "test2",
			LastName:   "test2",
			Email:      "test2",
			City:       "test2",
			Phone:      "+77771112233",
			Image:      "test2",
			Password:   "test2",
			IsVerified: false,
			IsBanned:   helpers.PointerBool(false),
		}
	)

	// Setup mocks.
	//
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks
	//
	codeService := mocks.NewMockCodeService(stubCtrl)
	mailService := mocks.NewMockMailService(stubCtrl)
	repository := mocks.NewMockUserRepository(stubCtrl)
	redisRepository := mocks.NewMockIdentityManagerRedisRepository(stubCtrl)

	repository.EXPECT().
		IsEmailExist(ctx, user.Email).
		Return(false, nil).
		AnyTimes()

	repository.EXPECT().
		IsEmailExist(ctx, user2.Email).
		Return(true, errors.NewErrAlreadyExist("email exists")).
		AnyTimes()

	repository.EXPECT().
		IsPhoneNumberExist(ctx, user.Phone).
		Return(false, nil).
		AnyTimes()

	redisRepository.EXPECT().
		Set(ctx, user.Phone, user, domain.RedisUserDuration).
		Return(nil).
		AnyTimes()

	codeService.EXPECT().
		SendVerificationCode(ctx, user.Phone).
		Return(nil).
		AnyTimes()

	service := newBasicService(
		codeService,
		mailService,
		repository,
		redisRepository,
	)

	// Define tests
	//
	type arguments struct {
		user *domain.User
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: user is created",
			arguments: arguments{
				user: user,
			},
			expectError: false,
		},
		{
			name: "Fail: nil user",
			arguments: arguments{
				user: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid email",
			arguments: arguments{
				user: &domain.User{
					RoleID:                  1,
					FirstName:               "test",
					LastName:                "test",
					Email:                   "",
					City:                    "test",
					Phone:                   "+77771112233",
					ConsultationPhoneNumber: "+77771112233",
					Image:                   "test",
					Password:                "test",
					IsVerified:              false,
					IsBanned:                helpers.PointerBool(false),
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid phone",
			arguments: arguments{
				user: &domain.User{
					RoleID:     1,
					FirstName:  "test",
					LastName:   "test",
					Email:      "test",
					City:       "test",
					Phone:      "",
					Image:      "test",
					Password:   "test",
					IsVerified: false,
					IsBanned:   helpers.PointerBool(false),
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid password",
			arguments: arguments{
				user: &domain.User{
					RoleID:     1,
					FirstName:  "test",
					LastName:   "test",
					Email:      "test",
					City:       "test",
					Phone:      "test",
					Image:      "test",
					Password:   "",
					IsVerified: false,
					IsBanned:   helpers.PointerBool(false),
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid first name",
			arguments: arguments{
				user: &domain.User{
					RoleID:     1,
					FirstName:  "",
					LastName:   "test",
					Email:      "test",
					City:       "test",
					Phone:      "test",
					Image:      "test",
					Password:   "",
					IsVerified: false,
					IsBanned:   helpers.PointerBool(false),
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid first name",
			arguments: arguments{
				user: &domain.User{
					RoleID:     1,
					FirstName:  "",
					LastName:   "test",
					Email:      "test",
					City:       "test",
					Phone:      "test",
					Image:      "test",
					Password:   "",
					IsVerified: false,
					IsBanned:   helpers.PointerBool(false),
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid role id",
			arguments: arguments{
				user: &domain.User{
					RoleID:     domain.RoleBuilder,
					FirstName:  "test",
					LastName:   "test",
					Email:      "test",
					City:       "test",
					Phone:      "test",
					Image:      "test",
					Password:   "",
					IsVerified: false,
					IsBanned:   helpers.PointerBool(false),
				},
			},
			expectError: true,
		},
		{
			name: "Fail: email exists",
			arguments: arguments{
				user: user2,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			err := service.CreateUser(ctx, args.user)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestService_ValidateCode(t *testing.T) {

	var (
		ctx   = context.Background()
		code  = "test"
		code2 = "test2"
		phone = "+77771112233"
		user  = &domain.User{
			ID:         1,
			RoleID:     1,
			FirstName:  "test",
			LastName:   "test",
			Email:      "test",
			City:       "test",
			Phone:      "+77771112233",
			Image:      "test",
			Password:   "test",
			IsVerified: false,
			IsBanned:   helpers.PointerBool(false),
		}
	)

	// Setup mocks.
	//
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks
	//
	codeService := mocks.NewMockCodeService(stubCtrl)
	mailService := mocks.NewMockMailService(stubCtrl)
	repository := mocks.NewMockUserRepository(stubCtrl)
	redisRepository := mocks.NewMockIdentityManagerRedisRepository(stubCtrl)

	service := newBasicService(
		codeService,
		mailService,
		repository,
		redisRepository,
	)

	codeService.EXPECT().
		ValidateCode(ctx, code).
		Return(phone, nil).
		AnyTimes()

	codeService.EXPECT().
		ValidateCode(ctx, code2).
		Return("", errors.NewErrNotFound("code not found")).
		AnyTimes()

	redisRepository.EXPECT().
		Get(ctx, phone).
		Return(user, nil).
		AnyTimes()

	repository.EXPECT().
		CreateUser(ctx, user).
		Return(user, nil).
		AnyTimes()

	repository.EXPECT().
		MakeUserVerified(ctx, user.ID).
		Return(nil).
		AnyTimes()

	// Define tests
	//
	type arguments struct {
		code string
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: code validated",
			arguments: arguments{
				code: code,
			},
			expectError: false,
		},
		{
			name: "Fail: invalid code",
			arguments: arguments{
				code: "",
			},
			expectError: true,
		},
		{
			name: "Fail: invalid code",
			arguments: arguments{
				code: code2,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			userID, roleID, err := service.ValidateCode(ctx, args.code)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if userID == 0 {
					t.Error("user id equals to 0")
				}
				if roleID == 0 {
					t.Error("user id equals to 0")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestService_ValidateUser(t *testing.T) {

}

func TestService_GetProfile(t *testing.T) {

}

func TestService_UpdateProfile(t *testing.T) {

}

func TestService_SendResetPasswordToken(t *testing.T) {

}

func TestService_ResetPassword(t *testing.T) {

}
