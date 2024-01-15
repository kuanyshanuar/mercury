package service

import (
	"context"
	"fmt"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"github.com/golang/mock/gomock"
)

func TestService_UpdateProfile(t *testing.T) {
	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}
		email    = "test@mail.com"
		phone    = "1234567"
		profile  = &domain.Profile{
			ID:        1,
			RoleID:    1,
			FirstName: "Test",
			LastName:  "Test",
			City:      "Test",
			Phone:     "1234567",
			Email:     "test@mail.com",
			Password:  "Test",
		}
	)

	// Setup mocks.
	//
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mock dependencies
	//
	codeService := mocks.NewMockCodeService(stubCtrl)
	repository := mocks.NewMockProfileRepository(stubCtrl)
	redis := mocks.NewMockProfileRedisRepository(stubCtrl)

	service := newBasicService(
		codeService,
		repository,
		redis,
	)

	repository.EXPECT().
		IsEmailExist(ctx, email).
		Return(false, nil).
		AnyTimes()

	repository.EXPECT().
		IsPhoneExist(ctx, phone).
		Return(false, nil).
		AnyTimes()

	repository.EXPECT().
		GetProfile(ctx, profile.ID).
		Return(profile, nil).
		AnyTimes()

	repository.EXPECT().
		UpdateProfile(ctx, profile.ID, profile).
		Return(nil).
		AnyTimes()

	// Define tests
	//
	type arguments struct {
		userID  domain.UserID
		profile *domain.Profile
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: profile updated",
			arguments: arguments{
				userID: 1,
				profile: &domain.Profile{
					ID:        1,
					RoleID:    1,
					FirstName: "Test",
					LastName:  "Test",
					City:      "Test",
					Phone:     "1234567",
					Email:     "test@mail.com",
					Password:  "Test",
				},
			},
			expectError: false,
		},
		{
			name: "Fail: invalid user id",
			arguments: arguments{
				userID: 0,
				profile: &domain.Profile{
					ID:        1,
					RoleID:    1,
					FirstName: "Test",
					LastName:  "Test",
					City:      "Test",
					Phone:     "1234567",
					Email:     "test@mail.com",
					Password:  "Test",
				},
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			err := service.UpdateProfile(ctx, args.userID, args.profile, callerID)
			if !test.expectError {
				if err != nil {
					t.Error(err)
				}
			} else {
				if err == nil {
					t.Error("unexpected error but got nothing")
				}
			}
		})
	}
}

func TestService_UpdateProfileWithError(t *testing.T) {
	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}
		email    = "test@mail.com"
		phone    = "1234567"
	)

	// Setup mocks.
	//
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mock dependencies
	//
	codeService := mocks.NewMockCodeService(stubCtrl)
	repository := mocks.NewMockProfileRepository(stubCtrl)
	redis := mocks.NewMockProfileRedisRepository(stubCtrl)

	service := newBasicService(
		codeService,
		repository,
		redis,
	)

	repository.EXPECT().
		IsEmailExist(ctx, email).
		Return(true, nil).
		AnyTimes()

	repository.EXPECT().
		IsPhoneExist(ctx, phone).
		Return(true, nil).
		AnyTimes()

	// Define tests
	//
	type arguments struct {
		userID  domain.UserID
		profile *domain.Profile
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: profile updated",
			arguments: arguments{
				userID: 1,
				profile: &domain.Profile{
					ID:        1,
					RoleID:    1,
					FirstName: "Test",
					LastName:  "Test",
					City:      "Test",
					Phone:     "1234567",
					Email:     "test@mail.com",
					Password:  "Test",
				},
			},
			expectError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			err := service.UpdateProfile(ctx, args.userID, args.profile, callerID)
			if !test.expectError {
				if err != nil {
					t.Error(err)
				}
			} else {
				if err == nil {
					t.Error("unexpected error but got nothing")
				}
			}
		})
	}
}

func TestService_UpdatePhone(t *testing.T) {
	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}
		email    = "test@mail.com"
		phone    = "12345678"
		profile  = &domain.Profile{
			ID:        1,
			RoleID:    1,
			FirstName: "Test",
			LastName:  "Test",
			City:      "Test",
			Phone:     "1234567",
			Email:     "test@mail.com",
			Password:  "Test",
		}
	)

	// Setup mocks.
	//
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mock dependencies
	//
	codeService := mocks.NewMockCodeService(stubCtrl)
	repository := mocks.NewMockProfileRepository(stubCtrl)
	redis := mocks.NewMockProfileRedisRepository(stubCtrl)

	service := newBasicService(
		codeService,
		repository,
		redis,
	)

	repository.EXPECT().
		IsEmailExist(ctx, email).
		Return(false, nil).
		AnyTimes()

	repository.EXPECT().
		IsPhoneExist(ctx, phone).
		Return(false, nil).
		AnyTimes()

	repository.EXPECT().
		GetProfile(ctx, profile.ID).
		Return(profile, nil).
		AnyTimes()

	codeService.EXPECT().
		SendCode(ctx, phone).
		Return(nil).
		AnyTimes()

	redis.EXPECT().Set(
		ctx,
		phone,
		&domain.Profile{
			ID:        1,
			RoleID:    1,
			FirstName: "Test",
			LastName:  "Test",
			City:      "Test",
			Phone:     "12345678",
			Email:     "test@mail.com",
			Password:  "Test",
		},
		domain.RedisUserDuration,
	).Return(nil).
		AnyTimes()

	// Define tests
	//
	type arguments struct {
		userID  domain.UserID
		profile *domain.Profile
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: profile updated",
			arguments: arguments{
				userID: 1,
				profile: &domain.Profile{
					ID:        1,
					RoleID:    1,
					FirstName: "Test",
					LastName:  "Test",
					City:      "Test",
					Phone:     "12345678",
					Email:     "test@mail.com",
					Password:  "Test",
				},
			},
			expectError: false,
		},
		{
			name: "Fail: invalid user id",
			arguments: arguments{
				userID: 0,
				profile: &domain.Profile{
					ID:        1,
					RoleID:    1,
					FirstName: "Test",
					LastName:  "Test",
					City:      "Test",
					Phone:     "12345679",
					Email:     "test@mail.com",
					Password:  "Test",
				},
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			err := service.UpdateProfile(ctx, args.userID, args.profile, callerID)
			if !test.expectError {
				if err != nil {
					t.Error(err)
				}
			} else {
				if err == nil {
					t.Error("unexpected error but got nothing")
				}
			}
		})
	}
}

func TestService_ValidatePhone(t *testing.T) {
	var (
		ctx         = context.Background()
		callerID    = domain.CallerID{}
		code        = "1234"
		invalidCode = "123456"
		phone       = "12345567"
		profile     = &domain.Profile{
			ID:        1,
			RoleID:    1,
			FirstName: "Test",
			LastName:  "Test",
			City:      "Test",
			Phone:     "1234567",
			Email:     "test@mail.com",
			Password:  "Test",
		}
	)

	// Setup mocks.
	//
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mock dependencies
	//
	codeService := mocks.NewMockCodeService(stubCtrl)
	repository := mocks.NewMockProfileRepository(stubCtrl)
	redis := mocks.NewMockProfileRedisRepository(stubCtrl)

	service := newBasicService(
		codeService,
		repository,
		redis,
	)

	codeService.EXPECT().
		ValidateCode(ctx, code).
		Return(phone, nil).
		AnyTimes()

	codeService.EXPECT().
		ValidateCode(ctx, invalidCode).
		Return("", fmt.Errorf("code is not exist")).
		AnyTimes()

	redis.EXPECT().
		Get(ctx, phone).
		Return(profile, nil).
		AnyTimes()

	repository.EXPECT().
		UpdateProfile(ctx, profile.ID, profile).
		Return(nil).
		AnyTimes()

	type arguments struct {
		code string
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: code is validated",
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
				code: invalidCode,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			err := service.ValidatePhone(ctx, args.code, callerID)
			if !test.expectError {
				if err != nil {
					t.Error(err)
				}
			} else {
				if err == nil {
					t.Error("unexpected error but nothing")
				}
			}

		})
	}
}
