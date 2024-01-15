package repository

import (
	"context"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"

	"github.com/go-test/deep"
)

func TestRepository_CreatePermission(t *testing.T) {

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
		permission  *domain.Permission
		expectError bool
	}{
		{
			name: "Permission #700 created",
			permission: &domain.Permission{
				ID:           700,
				EndpointName: "System.CreateResidence",
				Action:       "CreateResidence",
				CrudType:     "C",
				IsActive:     true,
			},
			expectError: false,
		},
		{
			name: "Permission #2 created",
			permission: &domain.Permission{
				ID:           800,
				EndpointName: "System.UpdateResidence",
				Action:       "UpdateResidence",
				CrudType:     "U",
				IsActive:     true,
			},
			expectError: false,
		},
		{
			name: "Permission #3 created",
			permission: &domain.Permission{
				ID:           900,
				EndpointName: "System.DeleteResidence",
				Action:       "DeleteResidence",
				CrudType:     "D",
				IsActive:     true,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := testRepository.CreatePermission(ctx, test.permission)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == 0 {
					t.Error("permission has not been created")
					return
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_GetPermissionByEndpoint(t *testing.T) {
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
		permissions = []*domain.Permission{
			{
				ID:           700,
				EndpointName: "System.CreateResidence",
				Action:       "CreateResidence",
				CrudType:     "C",
				IsActive:     true,
			},
			{
				ID:           800,
				EndpointName: "System.UpdateResidence",
				Action:       "UpdateResidence",
				CrudType:     "U",
				IsActive:     true,
			},
			{
				ID:           900,
				EndpointName: "System.DeleteResidence",
				Action:       "DeleteResidence",
				CrudType:     "D",
				IsActive:     true,
			},
		}
	)

	testRepository := NewRepository(db)

	for _, permission := range permissions {
		resultID, err := testRepository.CreatePermission(ctx, permission)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if resultID == 0 {
			t.Fatal("residence has not been created")
		}
	}

	tests := []struct {
		name         string
		endpointName string
		expectError  bool
	}{
		{
			name:         "Permission #1 found",
			endpointName: "System.CreateResidence",
			expectError:  false,
		},
		{
			name:         "Permission #2 found",
			endpointName: "System.UpdateResidence",
			expectError:  false,
		},
		{
			name:         "Permission #3 found",
			endpointName: "System.DeleteResidence",
			expectError:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := testRepository.GetPermissionByEndpoint(ctx, test.endpointName)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == nil {
					t.Error("permission has not been created")
					return
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_IsPermissionAllowed(t *testing.T) {

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
		name         string
		roleID       domain.RoleID
		permissionID domain.PermissionID
		expected     bool
		expectError  bool
	}{
		{
			name:         "RoleID: 2 PermissionID: 1 - Allowed",
			roleID:       2,
			permissionID: 1,
			expected:     true,
			expectError:  false,
		},
		{
			name:         "RoleID: 2 PermissionID: 2 - Allowed",
			roleID:       2,
			permissionID: 2,
			expected:     true,
			expectError:  false,
		},
		{
			name:         "RoleID: 3 PermissionID: 1 - Not Allowed",
			roleID:       3,
			permissionID: 1,
			expected:     false,
			expectError:  false,
		},
		{
			name:         "RoleID: 2 PermissionID: 5 - Not Allowed",
			roleID:       2,
			permissionID: 5,
			expected:     false,
			expectError:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := testRepository.IsPermissionAllowed(ctx, test.roleID, test.permissionID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if diff := deep.Equal(test.expected, actual); diff != nil {
					t.Error(diff)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}
