package repository

import (
	"context"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
)

func TestRepository_CreateFilter(t *testing.T) {
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
		key         string
		filter      *domain.Filter
		expectError bool
	}{
		{
			name: "#1: Filter created",
			key:  "parking_types",
			filter: &domain.Filter{
				Name: "underground",
			},
			expectError: false,
		},
		{
			name: "#2: Filter created",
			key:  "parking_types",
			filter: &domain.Filter{
				Name: "ground",
			},
			expectError: false,
		},
		{
			name: "#2: Filter created",
			key:  "parking_types",
			filter: &domain.Filter{
				Name: "air",
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := testRepository.CreateFilter(ctx, test.key, test.filter)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == nil {
					t.Error("filter not created")
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

func TestRepository_DeleteFilter(t *testing.T) {
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
		ctx     = context.Background()
		key     = "parking_types"
		filters = []*domain.Filter{
			{
				ID:   1000,
				Name: "underground",
			},
			{
				ID:   1001,
				Name: "ground",
			},
			{
				ID:   1002,
				Name: "air",
			},
		}
	)

	testRepository := NewRepository(db)

	for _, filter := range filters {
		created, err := testRepository.CreateFilter(ctx, key, filter)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if created == nil {
			t.Fatal("filter has not been created")
		}
	}

	// Define tests
	//
	tests := []struct {
		name        string
		key         string
		id          int64
		expectError bool
	}{
		{
			name:        "#1: Filter deleted",
			key:         key,
			id:          1000,
			expectError: false,
		},
		{
			name:        "#2: Filter deleted",
			key:         key,
			id:          1001,
			expectError: false,
		},
		{
			name:        "#2: Filter deleted",
			key:         key,
			id:          1002,
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.DeleteFilter(ctx, test.id, test.key)
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
