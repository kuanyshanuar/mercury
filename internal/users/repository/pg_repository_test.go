package repository

import (
	"context"
	"strings"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
)

func TestRepository_CreateUser(t *testing.T) {

	// Setup database.
	//
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	var (
		ctx = context.Background()
	)

	testRepository := NewRepository(db)

	tests := []struct {
		name        string
		user        *domain.User
		expected    *domain.User
		expectError bool
	}{
		{
			name: "",
			user: &domain.User{
				ID:                      1001,
				RoleID:                  domain.RoleClient,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test@mail.com",
				City:                    "test",
				Phone:                   "123",
				ConsultationPhoneNumber: "123",
				Image:                   "test",
				Password:                "test",
			},
			expected: &domain.User{
				ID:                      1001,
				RoleID:                  domain.RoleClient,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test@mail.com",
				City:                    "test",
				Phone:                   "123",
				ConsultationPhoneNumber: "123",
				Image:                   "test",
				Password:                "test",
			},
			expectError: false,
		},
		{
			name: "",
			user: &domain.User{
				ID:                      1002,
				RoleID:                  domain.RoleClient,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test2@mail.com",
				City:                    "test",
				Phone:                   "1234",
				ConsultationPhoneNumber: "1234",
				Image:                   "test",
				Password:                "test",
			},
			expected: &domain.User{
				ID:                      1002,
				RoleID:                  domain.RoleClient,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test2@mail.com",
				City:                    "test",
				Phone:                   "1234",
				ConsultationPhoneNumber: "1234",
				Image:                   "test",
				Password:                "test",
			},
			expectError: false,
		},
		{
			name: "",
			user: &domain.User{
				ID:                      1003,
				RoleID:                  domain.RoleClient,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test3@mail.com",
				City:                    "test",
				Phone:                   "12345",
				ConsultationPhoneNumber: "12345",
				Image:                   "test",
				Password:                "test",
			},
			expected: &domain.User{
				ID:                      1003,
				RoleID:                  domain.RoleClient,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test3@mail.com",
				City:                    "test",
				Phone:                   "12345",
				ConsultationPhoneNumber: "12345",
				Image:                   "test",
				Password:                "test",
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := test.expected

			result, err := testRepository.CreateUser(ctx, test.user)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == nil {
					t.Error("residence has not been created")
					return
				}
				if result.ID <= 0 {
					t.Error("bad residence id")
				}
				if !strings.EqualFold(expected.Email, result.Email) {
					t.Errorf("expected residence description is %s but actual is %s", expected.Email, result.Email)
				}
				if result.CreatedAt <= 0 {
					t.Error("creation time not set")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_GetUser(t *testing.T) {

}

func TestRepository_GetUserByEmail(t *testing.T) {

}

func TestRepository_GetUserByEmailPassword(t *testing.T) {

}

func TestRepository_GetUserByPhoneNumber(t *testing.T) {

}

func TestRepository_ListUsers(t *testing.T) {

}

func TestRepository_UpdateUser(t *testing.T) {

}

func TestRepository_DeleteUser(t *testing.T) {

}

func TestRepository_IsEmailExist(t *testing.T) {

}

func TestRepository_IsPhoneNumberExist(t *testing.T) {

}

func TestRepository_ValidateUser(t *testing.T) {

}

func TestRepository_BanUser(t *testing.T) {

}

func TestRepository_CreateResetPasswordToken(t *testing.T) {

	// Setup database.
	//
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	var (
		ctx = context.Background()
	)

	// Create test repository
	//
	testRepository := NewRepository(db)

	tests := []struct {
		name        string
		token       *domain.ResetPasswordToken
		expectError bool
	}{
		{
			name: "#1 Token",
			token: &domain.ResetPasswordToken{
				UserID: 1,
				Token:  "test1",
			},
		},
		{
			name: "#2 Token",
			token: &domain.ResetPasswordToken{
				UserID: 2,
				Token:  "test2",
			},
		},
		{
			name: "#3 Token",
			token: &domain.ResetPasswordToken{
				UserID: 3,
				Token:  "test3",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.CreateResetPasswordToken(ctx, test.token)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				result, err := testRepository.GetResetPasswordToken(ctx, test.token.Token)
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == nil {
					t.Error("reset password token has not been created")
					return
				}
				if result.Confirmed == true {
					t.Error("confirmed not set")
				}
				if len(result.Token) == 0 {
					t.Error("token not set")
				}
				if result.CreatedAt <= 0 {
					t.Error("creation time not set")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_UpdateResetPasswordToken(t *testing.T) {
	// Setup database.
	//
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	var (
		ctx         = context.Background()
		resetTokens = []*domain.ResetPasswordToken{
			{
				UserID: 1,
				Token:  "test1",
			},
			{
				UserID: 2,
				Token:  "test2",
			},
			{
				UserID: 3,
				Token:  "test3",
			},
		}
	)

	// Create test repository
	//
	testRepository := NewRepository(db)

	// Create reset password tokens
	//
	for _, opt := range resetTokens {
		err = testRepository.CreateResetPasswordToken(ctx, opt)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
	}

	// Create tests
	//
	tests := []struct {
		name        string
		token       *domain.ResetPasswordToken
		expectError bool
	}{
		{
			name: "#1 Update Token",
			token: &domain.ResetPasswordToken{
				UserID:    1,
				Token:     "test1",
				Confirmed: true,
			},
			expectError: false,
		},
		{
			name: "#2 Update Token",
			token: &domain.ResetPasswordToken{
				UserID:    2,
				Token:     "test2",
				Confirmed: true,
			},
			expectError: false,
		},
		{
			name: "#3 Update Token",
			token: &domain.ResetPasswordToken{
				UserID:    3,
				Token:     "test3",
				Confirmed: true,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.UpdateResetPasswordToken(ctx, test.token)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				result, err := testRepository.GetResetPasswordToken(ctx, test.token.Token)
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == nil {
					t.Error("reset password token has not been created")
					return
				}
				if result.Confirmed == false {
					t.Error("token is not confirmed")
				}
				if result.CreatedAt <= 0 {
					t.Error("creation time not set")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}
