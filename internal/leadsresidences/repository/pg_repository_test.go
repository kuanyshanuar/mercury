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

func TestRepository_CreateLeadResidence(t *testing.T) {
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
		leadID      domain.LeadID
		lead        *domain.LeadResidence
		expectError bool
	}{
		{
			name:   "Success: create lead",
			leadID: domain.LeadID(1),
			lead: &domain.LeadResidence{
				ResidenceID: 1,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
			expectError: false,
		},
		{
			name:   "Success: create lead",
			leadID: domain.LeadID(2),
			lead: &domain.LeadResidence{
				ResidenceID: 2,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
			expectError: false,
		},
		{
			name:   "Success: create lead",
			leadID: domain.LeadID(1),
			lead: &domain.LeadResidence{
				ResidenceID: 1,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.CreateLeadResidence(ctx, test.lead)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				lead, err := testRepository.GetLeadResidence(ctx, test.leadID)
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				if lead.StatusID != domain.StatusActive {
					t.Errorf("invalid status: %d", lead.StatusID)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_ListLeadResidences(t *testing.T) {
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
		ctx      = context.Background()
		criteria = domain.LeadResidenceSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	)

	testRepository := NewRepository(db)

	tests := []struct {
		name        string
		leadID      domain.LeadID
		lead        *domain.LeadResidence
		expectError bool
	}{
		{
			name:   "Success: create lead",
			leadID: domain.LeadID(1),
			lead: &domain.LeadResidence{
				ResidenceID: 1,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
			expectError: false,
		},
		{
			name:   "Success: create lead",
			leadID: domain.LeadID(2),
			lead: &domain.LeadResidence{
				ResidenceID: 2,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.CreateLeadResidence(ctx, test.lead)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				leads, total, err := testRepository.ListLeadResidences(ctx, criteria)
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				if len(leads) != int(total) {
					t.Errorf("leads and total mismatched: %d != %d", len(leads), total)
				}

			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_UpdateLeadResidence(t *testing.T) {
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
		leads = []*domain.LeadResidence{
			{
				ResidenceID: 1,
				StatusID:    1,
				IssuedAt:    time.Now().Unix(),
				ExpiresAt:   time.Now().Unix(),
			},
			{
				ResidenceID: 2,
				StatusID:    2,
				IssuedAt:    time.Now().Unix(),
				ExpiresAt:   time.Now().Unix(),
			},
		}
	)

	// Mock time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 06, 28, 0, 0, 0, 0, time.UTC)
	})

	testRepository := NewRepository(db)

	for _, lead := range leads {
		err := testRepository.CreateLeadResidence(ctx, lead)
		if err != nil {
			t.Error(err)
		}
	}

	// Define tests
	//
	tests := []struct {
		name        string
		lead        *domain.LeadResidence
		expectError bool
	}{
		{
			name: "Success: update lead residence",
			lead: &domain.LeadResidence{
				ID:          1,
				ResidenceID: 1,
				StatusID:    2,
				IssuedAt:    0,
				ExpiresAt:   0,
			},
			expectError: false,
		},
		{
			name: "Success: update lead residence",
			lead: &domain.LeadResidence{
				ID:          1,
				ResidenceID: 1,
				StatusID:    2,
				IssuedAt:    0,
				ExpiresAt:   0,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.UpdateLeadResidence(ctx, test.lead)
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

func TestRepository_DeleteLeadResidence(t *testing.T) {
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
		leads = []*domain.LeadResidence{
			{
				ResidenceID: 1,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
			{
				ResidenceID: 2,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
		}
	)

	// Mock time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 06, 28, 0, 0, 0, 0, time.UTC)
	})

	testRepository := NewRepository(db)

	for _, lead := range leads {
		err := testRepository.CreateLeadResidence(ctx, lead)
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
			name:        "Success: disable lead residence",
			leadID:      domain.LeadID(1),
			expectError: false,
		},
		{
			name:        "Success: disable lead residence",
			leadID:      domain.LeadID(2),
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.DeleteLeadResidence(ctx, test.leadID)
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

func TestRepository_RevokeLeadResidence(t *testing.T) {
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
		leads = []*domain.LeadResidence{
			{
				ResidenceID: 1,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
			{
				ResidenceID: 2,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
		}
	)

	// Mock time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 06, 28, 0, 0, 0, 0, time.UTC)
	})

	testRepository := NewRepository(db)

	for _, lead := range leads {
		err := testRepository.CreateLeadResidence(ctx, lead)
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
			name:        "Success: disable lead residence",
			leadID:      domain.LeadID(1),
			expectError: false,
		},
		{
			name:        "Success: disable lead residence",
			leadID:      domain.LeadID(2),
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.RevokeLeadResidence(ctx, test.leadID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				lead, err := testRepository.GetLeadResidence(ctx, test.leadID)
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				if lead.StatusID != domain.StatusInactive {
					t.Errorf("invalid status: %d", lead.StatusID)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_IsLeadExistByDate(t *testing.T) {
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
		leads = []*domain.LeadResidence{
			{
				ResidenceID: 1,
				StatusID:    1,
				IssuedAt:    1687197600,
				ExpiresAt:   1687975200,
			},
			{
				ResidenceID: 2,
				StatusID:    1,
				IssuedAt:    1687802400,
				ExpiresAt:   1687975200,
			},
		}
	)

	testRepository := NewRepository(db)

	for _, lead := range leads {
		err := testRepository.CreateLeadResidence(ctx, lead)
		if err != nil {
			t.Error(err)
		}
	}

	type arguments struct {
		residenceID domain.ResidenceID
		issuedAt    int64
		expiresAt   int64
	}

	type result struct {
		isExist bool
	}

	// Define tests
	//
	tests := []struct {
		name        string
		arguments   arguments
		result      result
		expectError bool
	}{
		{
			name: "Lead exists",
			arguments: arguments{
				residenceID: 1,
				issuedAt:    1687370400,
				expiresAt:   1687543200,
			},
			result: result{
				isExist: true,
			},
			expectError: false,
		},
		{
			name: "Lead does not exists",
			arguments: arguments{
				residenceID: 2,
				issuedAt:    1685556000,
				expiresAt:   1685642400,
			},
			result: result{
				isExist: false,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			isLeadExists, err := testRepository.IsLeadExistByDate(ctx, args.residenceID, args.issuedAt, args.expiresAt)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				actual := result{
					isExist: isLeadExists,
				}
				if diff := deep.Equal(actual.isExist, test.result.isExist); diff != nil {
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
