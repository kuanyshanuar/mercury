package repository

import (
	"context"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
	"testing"
)

func TestRepository_Create(t *testing.T) {

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

	// Define tests
	//
	tests := []struct {
		name        string
		manager     *domain.Manager
		expectError bool
	}{
		{
			name: "Success: manager created",
			manager: &domain.Manager{
				ID:         100,
				RoleID:     domain.RoleManager,
				FirstName:  "test",
				LastName:   "test",
				Email:      "test",
				Phone:      "test",
				Image:      "test",
				Password:   "test",
				IsVerified: false,
				IsBanned:   helpers.PointerBool(false),
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			managerID, err := testRepository.Create(ctx, test.manager)
			if err != nil {
				t.Error(err)
			}
			if managerID <= 0 {
				t.Errorf("invalid manager id: %d", managerID)
			}
		})
	}
}

func TestRepository_Update(t *testing.T) {

}

func TestRepository_Delete(t *testing.T) {

}
