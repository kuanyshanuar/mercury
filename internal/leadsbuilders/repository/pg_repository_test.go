package repository

import (
	"context"
	"testing"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"

	"bou.ke/monkey"
	"github.com/go-test/deep"
)

func TestRepository_CreateLeadBuilder(t *testing.T) {

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
		nowTime = time.Now()
	)

	testRepository := NewRepository(db)

	tests := []struct {
		name        string
		lead        *domain.LeadBuilder
		expected    *domain.LeadBuilder
		expectError bool
	}{
		{
			name: "#1: success: create lead",
			lead: &domain.LeadBuilder{
				BuilderID: 1,
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
			expected: &domain.LeadBuilder{
				ID:        1,
				BuilderID: 1,
				StatusID:  1,
				Status: &domain.LeadStatus{
					ID:   1,
					Name: "Активный",
				},
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
			expectError: false,
		},
		{
			name: "#2: fail: invalid builder id",
			lead: &domain.LeadBuilder{
				BuilderID: 0,
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "#3: fail: builder id exists",
			lead: &domain.LeadBuilder{
				BuilderID: 1,
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			err := testRepository.CreateLeadBuilder(ctx, test.lead)
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

func TestRepository_GetLeadBuilder(t *testing.T) {

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
		nowTime = time.Now()
		lead    = &domain.LeadBuilder{
			BuilderID: 1,
			StatusID:  1,
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
		}
	)

	testRepository := NewRepository(db)

	// Create a new lead
	//
	err = testRepository.CreateLeadBuilder(ctx, lead)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	tests := []struct {
		name        string
		leadID      domain.LeadID
		expected    *domain.LeadBuilder
		expectError bool
	}{
		{
			name:        "Success: get lead",
			leadID:      1,
			expected:    lead,
			expectError: false,
		},
		{
			name:        "Fail: lead not exists",
			leadID:      2,
			expected:    nil,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := testRepository.GetLeadBuilder(ctx, test.leadID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == nil {
					t.Errorf("unexpected error nil result")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_UpdateLeadBuilder(t *testing.T) {

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
		nowTime = time.Now()
		lead    = &domain.LeadBuilder{
			BuilderID: 1,
			StatusID:  1,
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
		}
	)

	testRepository := NewRepository(db)

	// Create a new lead
	//
	err = testRepository.CreateLeadBuilder(ctx, lead)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	// Define tests
	//
	tests := []struct {
		name        string
		leadID      domain.LeadID
		lead        *domain.LeadBuilder
		expected    *domain.LeadBuilder
		expectError bool
	}{
		{
			name:   "#1: success: update lead",
			leadID: 1,
			lead: &domain.LeadBuilder{
				BuilderID: 1,
				StatusID:  1,
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
			expected: &domain.LeadBuilder{
				BuilderID: 1,
				StatusID:  1,
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
			expectError: false,
		},
		{
			name:   "#1: fail: lead does not exist",
			leadID: 2,
			lead: &domain.LeadBuilder{
				BuilderID: 1,
				StatusID:  1,
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
			expected:    &domain.LeadBuilder{},
			expectError: false,
		},
		{
			name:        "#1: fail: invalid id",
			leadID:      0,
			lead:        &domain.LeadBuilder{},
			expected:    &domain.LeadBuilder{},
			expectError: true,
		},
		{
			name:        "#1: fail: invalid builder id",
			leadID:      0,
			lead:        &domain.LeadBuilder{},
			expected:    &domain.LeadBuilder{},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.UpdateLeadBuilder(ctx, test.leadID, test.lead)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				result, err := testRepository.GetLeadBuilder(ctx, test.leadID)
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if diff := deep.Equal(result, test.expected); diff != nil {
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

func TestRepository_ListLeadBuilder(t *testing.T) {

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
		nowTime = time.Now()
		leads   = []*domain.LeadBuilder{
			{
				BuilderID: 1,
				StatusID:  1,
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
			{
				BuilderID: 2,
				StatusID:  1,
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
			{
				BuilderID: 3,
				StatusID:  1,
				IssuedAt:  nowTime.Unix(),
				ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
			},
		}
	)

	testRepository := NewRepository(db)

	// Create a new lead
	//
	for _, lead := range leads {
		err = testRepository.CreateLeadBuilder(ctx, lead)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
	}

	tests := []struct {
		name        string
		criteria    domain.LeadBuilderSearchCriteria
		expectError bool
	}{
		{
			name: "#1: list all leads",
			criteria: domain.LeadBuilderSearchCriteria{
				Page: domain.PageRequest{
					Offset: 0,
					Size:   3,
				},
			},
			expectError: false,
		},
		{
			name: "#1: list all leads",
			criteria: domain.LeadBuilderSearchCriteria{
				Page: domain.PageRequest{
					Offset: 3,
					Size:   3,
				},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, total, err := testRepository.ListLeadBuilders(ctx, test.criteria)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if total == 0 {
					t.Errorf("wrong total number: %d", total)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_DeleteLeadBuilder(t *testing.T) {

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
		nowTime = time.Now()
		lead    = &domain.LeadBuilder{
			BuilderID: 1,
			StatusID:  1,
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
		}
	)

	testRepository := NewRepository(db)

	// Create a new lead
	//
	err = testRepository.CreateLeadBuilder(ctx, lead)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	// Define tests
	//
	tests := []struct {
		name        string
		leadID      domain.LeadID
		expectError bool
	}{
		{
			name:        "#1: delete user",
			leadID:      1,
			expectError: false,
		},
		{
			name:        "#1: delete not existing user",
			leadID:      2,
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.DeleteLeadBuilder(ctx, test.leadID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				result, err := testRepository.GetLeadBuilder(ctx, test.leadID)
				if err == nil {
					t.Errorf("unexpected error %s", err)
				}
				if result != nil {
					t.Error("lead was not deleted")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_RevokeLeadBuilder(t *testing.T) {

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
		ctx   = context.Background()
		leads = []*domain.LeadBuilder{
			{
				BuilderID: 1,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
			{
				BuilderID: 2,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
		}
	)

	// Mock time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 06, 28, 0, 0, 0, 0, time.UTC)
	})

	testRepository := NewRepository(db)

	for _, lead := range leads {
		err := testRepository.CreateLeadBuilder(ctx, lead)
		if err != nil {
			t.Error(err)
		}
	}

	// Define tests
	//
	tests := []struct {
		name        string
		leadID      domain.LeadID
		expectError bool
	}{
		{
			name:        "Success: lead is revoked",
			leadID:      domain.LeadID(1),
			expectError: false,
		},
		{
			name:        "Success: lead is revoked",
			leadID:      domain.LeadID(2),
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.RevokeLeadBuilder(ctx, test.leadID)
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

func TestRepository_IsLeadExistByDateRange(t *testing.T) {
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
		ctx   = context.Background()
		leads = []*domain.LeadBuilder{
			{
				BuilderID: 1,
				StatusID:  1,
				IssuedAt:  1687284000,
				ExpiresAt: 1687975200,
			},
			{
				BuilderID: 2,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
		}
	)

	testRepository := NewRepository(db)

	// Create leads
	//
	for _, lead := range leads {
		err := testRepository.CreateLeadBuilder(ctx, lead)
		if err != nil {
			t.Error(err)
		}
	}

	// Define tests
	//
	type arguments struct {
		builderID domain.BuilderID
		issuedAt  int64
		expiresAt int64
	}

	type result struct {
		isExist bool
	}

	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: lead exists",
			arguments: arguments{
				builderID: 1,
				issuedAt:  1687543200,
				expiresAt: 1687629600,
			},
			expected: result{
				isExist: true,
			},
			expectError: false,
		},
		{
			name: "Success: lead exists",
			arguments: arguments{
				builderID: 1,
				issuedAt:  1690221600,
				expiresAt: 1690394400,
			},
			expected: result{
				isExist: false,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			isExist, err := testRepository.IsLeadExistByDateRange(ctx, args.builderID, args.issuedAt, args.expiresAt)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if diff := deep.Equal(test.expected.isExist, isExist); diff != nil {
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
