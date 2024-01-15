package repository

import (
	"context"
	"github.com/go-test/deep"
	cottage "gitlab.com/zharzhanov/mercury/internal/cottages/repository"
	"testing"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"

	"bou.ke/monkey"
)

func TestRepository_CreateLeadCottage(t *testing.T) {
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
		leadID      int64
		lead        *domain.LeadCottage
		expectError bool
	}{
		{
			name:   "Success: create lead",
			leadID: 1,
			lead: &domain.LeadCottage{
				CottageID: 1,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
			expectError: false,
		},
		{
			name:   "Success: create lead",
			leadID: 2,
			lead: &domain.LeadCottage{
				CottageID: 2,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
			expectError: false,
		},
		{
			name:   "Success: create lead",
			leadID: 1,
			lead: &domain.LeadCottage{
				CottageID: 3,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := testRepository.CreateLeadCottage(ctx, test.lead)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if id <= 0 {
					t.Errorf("unexpected error %s", err)
				}
				lead, err := testRepository.GetLeadCottage(ctx, test.leadID)
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

func TestRepository_ListLeadCottages(t *testing.T) {
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
		criteria = domain.LeadCottageSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
		validCottages = []*domain.Cottage{
			{
				Title:             "Simple Fields #1",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10,
				Longitude:         50,
				Territory:         "excellent facade",
				CeilingHeightMin:  4,
				BuildingArea:      10,
				RoomsMin:          1,
				RoomsMax:          6,
				HouseAmount:       404,
				FloorsCount:       50,
				Facade:            "good facade",
				CanRePlan:         true,
				AreaMin:           100,
				AreaMax:           300,
				PricePerSquareMin: 1500,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
			},
			{
				Title:            "Simple Fields #2",
				Description:      "Complex of cottages in the higher district",
				Address:          "Uralsk, Kazanoe, 56",
				Latitude:         10000,
				Longitude:        50000,
				Territory:        "excellent facade",
				CeilingHeightMin: 4000000000,
				BuildingArea:     10000000,
				HouseAmount:      404000,
				RoomsMin:         1,
				RoomsMax:         6,
				FloorsCount:      50000,
				Facade:           "good facade",

				CanRePlan:         true,
				AreaMin:           100000000,
				AreaMax:           300000000,
				PricePerSquareMin: 150000000,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
			},
		}
	)

	testRepository := NewRepository(db)
	cottagesRepo := cottage.NewCottageRepository(db)
	for _, cottage := range validCottages {
		_, err := cottagesRepo.Create(ctx, cottage)
		if err != nil {
			t.Log(err)
		}
	}
	tests := []struct {
		name        string
		leadID      domain.LeadID
		lead        *domain.LeadCottage
		expectError bool
	}{
		{
			name:   "Success: create lead",
			leadID: 1,
			lead: &domain.LeadCottage{
				CottageID: 1,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
			expectError: false,
		},
		{
			name:   "Success: create lead",
			leadID: 2,
			lead: &domain.LeadCottage{
				CottageID: 2,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := testRepository.CreateLeadCottage(ctx, test.lead)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if id <= 0 {
					t.Errorf("lead id is negative or zero")
				}
				leads, total, err := testRepository.ListLeadCottage(ctx, criteria)
				t.Log(*leads[0])
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

func TestRepository_UpdateLeadCottage(t *testing.T) {
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
		leads = []*domain.LeadCottage{
			{
				CottageID: 1,
				StatusID:  1,
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Unix(),
			},
			{
				CottageID: 2,
				StatusID:  2,
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Unix(),
			},
		}
	)

	// Mock time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC)
	})

	testRepository := NewRepository(db)

	for _, lead := range leads {
		id, err := testRepository.CreateLeadCottage(ctx, lead)
		if err != nil {
			t.Error(err)
		}
		if id <= 0 {
			t.Error("lead id return isn't valid")
		}
	}

	// Define tests
	//
	tests := []struct {
		name        string
		lead        *domain.LeadCottage
		expectError bool
	}{
		{
			name: "Success: update lead Cottage",
			lead: &domain.LeadCottage{
				ID:        1,
				CottageID: 1,
				StatusID:  2,
				IssuedAt:  0,
				ExpiresAt: 0,
			},
			expectError: false,
		},
		{
			name: "Success: update lead Cottage",
			lead: &domain.LeadCottage{
				ID:        1,
				CottageID: 1,
				StatusID:  2,
				IssuedAt:  0,
				ExpiresAt: 0,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lead, err := testRepository.UpdateLeadCottage(ctx, 1, test.lead)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if lead == nil {
					t.Error("update error")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_DeleteLeadCottage(t *testing.T) {
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
		leads = []*domain.LeadCottage{
			{
				CottageID: 1,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
			{
				CottageID: 2,
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
		id, err := testRepository.CreateLeadCottage(ctx, lead)
		if err != nil {
			t.Error(err)
		}
		if id <= 0 {
			t.Error("create error")
		}
	}

	// Define tests
	//
	tests := []struct {
		name        string
		leadID      int64
		expectError bool
	}{
		{
			name:        "Success: disable lead Cottage",
			leadID:      1,
			expectError: false,
		},
		{
			name:        "Success: disable lead Cottage",
			leadID:      2,
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.DeleteLeadCottage(ctx, test.leadID)
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

func TestRepository_RevokeLeadCottage(t *testing.T) {
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
		leads = []*domain.LeadCottage{
			{
				CottageID: 1,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
			{
				CottageID: 2,
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
		_, err := testRepository.CreateLeadCottage(ctx, lead)
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
			name:        "Success: disable lead Cottage",
			leadID:      1,
			expectError: false,
		},
		{
			name:        "Success: disable lead Cottage",
			leadID:      2,
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.RevokeLeadCottage(ctx, test.leadID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				lead, err := testRepository.GetLeadCottage(ctx, int64(test.leadID))
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
		leads = []*domain.LeadCottage{
			{
				CottageID: 1,
				StatusID:  1,
				IssuedAt:  1687197600,
				ExpiresAt: 1687975200,
			},
			{
				CottageID: 2,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
		}
	)

	testRepository := NewRepository(db)

	for _, lead := range leads {
		id, err := testRepository.CreateLeadCottage(ctx, lead)
		if err != nil {
			t.Error(err)
		}
		if id <= 0 {
			t.Fatal("id is negative or zero")
		}
	}

	type arguments struct {
		CottageID int64
		issuedAt  int64
		expiresAt int64
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
				CottageID: 1,
				issuedAt:  1687370400,
				expiresAt: 1687543200,
			},
			result: result{
				isExist: true,
			},
			expectError: false,
		},
		{
			name: "Lead does not exists",
			arguments: arguments{
				CottageID: 2,
				issuedAt:  1685556000,
				expiresAt: 1685642400,
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
			isLeadExists, err := testRepository.IsLeadExistByDate(ctx, args.CottageID, args.issuedAt, args.expiresAt)
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
