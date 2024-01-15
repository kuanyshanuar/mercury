package repository

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/go-test/deep"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
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

	testResidencesRepository := NewResidenceRepository(db)

	tests := []struct {
		name        string
		residence   *domain.Residence
		expectError bool
	}{
		{
			name: "Success: residence 1 create",
			residence: &domain.Residence{
				StatusID:           1,
				DistrictID:         1,
				UserID:             1,
				CityID:             1,
				HousingClassID:     1,
				ConstructionTypeID: 1,
				Title:              "test",
				Description:        "test",
				FlatPlans: []*domain.FlatPlan{
					{
						NumberOfRooms: 1,
						Area:          10,
						Price:         1000,
						Images:        []string{"test_image"},
					},
				},
			},
			expectError: false,
		},
		{
			name: "Success: residence 2 create",
			residence: &domain.Residence{
				StatusID:           2,
				DistrictID:         1,
				UserID:             1,
				CityID:             2,
				HousingClassID:     2,
				ConstructionTypeID: 2,
				Title:              "test2",
				Description:        "test2",
			},
			expectError: false,
		},
		{
			name: "Success: residence 3 create",
			residence: &domain.Residence{
				StatusID:           2,
				DistrictID:         1,
				UserID:             1,
				CityID:             3,
				HousingClassID:     3,
				ConstructionTypeID: 3,
				Title:              "test3",
				Description:        "test3",
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			residence := test.residence
			result, err := testResidencesRepository.Create(ctx, residence)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == nil {
					t.Error("residence has not been created")
					return
				}
				if !strings.EqualFold(residence.Title, result.Title) {
					t.Errorf("expected residence description is %s but actual is %s", residence.Title, result.Title)
				}
				if result.ID <= 0 {
					t.Error("bad residence id")
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

func TestRepository_List(t *testing.T) {
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
		ctx        = context.Background()
		residences = []*domain.Residence{
			{
				StatusID:           1,
				UserID:             1,
				CityID:             1,
				DistrictID:         1,
				HousingClassID:     1,
				ConstructionTypeID: 1,
				Title:              "test",
				Description:        "test",
				FloorsMax:          5,
				HasHGF:             false,
				AreaMin:            40,
				AreaMax:            90,
				RoomsMin:           1,
				RoomsMax:           5,
				PriceMin:           40,
				PriceMax:           60,
				PricePerSquareMin:  50,
				CeilingHeight:      2.5,
				FlatPlans: []*domain.FlatPlan{
					{
						NumberOfRooms: 1,
						Area:          10,
					},
					{
						NumberOfRooms: 2,
						Area:          20,
					},
				},
			},
			{
				StatusID:           1,
				UserID:             1,
				CityID:             1,
				DistrictID:         1,
				HousingClassID:     1,
				ConstructionTypeID: 2,
				Title:              "test2",
				Description:        "test2",
				FloorsMax:          10,
				HasHGF:             false,
				AreaMin:            50,
				AreaMax:            100,
				RoomsMin:           1,
				RoomsMax:           3,
				PriceMin:           50,
				PriceMax:           90,
				PricePerSquareMin:  90,
				CeilingHeight:      2.5,
				FlatPlans: []*domain.FlatPlan{
					{
						NumberOfRooms: 3,
						Area:          30,
					},
					{
						NumberOfRooms: 4,
						Area:          40,
					},
				},
			},
			{
				StatusID:           2,
				UserID:             1,
				CityID:             2,
				DistrictID:         1,
				HousingClassID:     2,
				ConstructionTypeID: 3,
				Title:              "test3",
				Description:        "test3",
				FloorsMax:          10,
				HasHGF:             true,
				AreaMin:            50,
				AreaMax:            150,
				RoomsMin:           3,
				RoomsMax:           4,
				PriceMin:           40,
				PriceMax:           90,
				PricePerSquareMin:  100,
				CeilingHeight:      3.5,
				FlatPlans: []*domain.FlatPlan{
					{
						NumberOfRooms: 3,
						Area:          50,
					},
					{
						NumberOfRooms: 4,
						Area:          60,
					},
				},
			},
			{
				StatusID:           2,
				UserID:             1,
				CityID:             2,
				DistrictID:         1,
				HousingClassID:     2,
				ConstructionTypeID: 3,
				Title:              "test4",
				Description:        "test4",
				FloorsMax:          4,
				HasHGF:             true,
				AreaMin:            50,
				AreaMax:            150,
				RoomsMin:           3,
				RoomsMax:           4,
				PriceMin:           100,
				PriceMax:           200,
				PricePerSquareMin:  200,
				CeilingHeight:      3.5,
				FlatPlans: []*domain.FlatPlan{
					{
						NumberOfRooms: 3,
						Area:          70,
					},
					{
						NumberOfRooms: 4,
						Area:          80,
					},
				},
			},
		}
	)

	testResidencesRepository := NewResidenceRepository(db)

	for _, residence := range residences {
		createdResidence, err := testResidencesRepository.Create(ctx, residence)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if createdResidence == nil {
			t.Fatal("residence has not been created")
		}
	}

	type expected struct {
		amount int
	}

	tests := []struct {
		name        string
		criteria    domain.ResidenceSearchCriteria
		expected    expected
		expectError bool
	}{
		{
			name: "Success: filter by title",
			criteria: domain.ResidenceSearchCriteria{
				Title: "test",
			},
			expected: expected{
				amount: 4,
			},
			expectError: false,
		},
		{
			name: "Success: filter by city",
			criteria: domain.ResidenceSearchCriteria{
				CityID: 1,
			},
			expected: expected{
				amount: 2,
			},
			expectError: false,
		},
		{
			name: "Success: filter by builders",
			criteria: domain.ResidenceSearchCriteria{
				BuilderIDs: []int64{1},
			},
			expected: expected{
				amount: 4,
			},
			expectError: false,
		},
		{
			name: "Success: filter by district",
			criteria: domain.ResidenceSearchCriteria{
				DistrictID: 1,
			},
			expected: expected{
				amount: 4,
			},
			expectError: false,
		},
		{
			name: "Success: filter by not existsing district",
			criteria: domain.ResidenceSearchCriteria{
				DistrictID: 4,
			},
			expected: expected{
				amount: 0,
			},
			expectError: false,
		},
		{
			name: "Success: filter by status",
			criteria: domain.ResidenceSearchCriteria{
				StatusID: 1,
			},
			expected: expected{
				amount: 2,
			},
			expectError: false,
		},
		{
			name: "Success: filter by rooms",
			criteria: domain.ResidenceSearchCriteria{
				RoomsMin: 1,
				RoomsMax: 2,
			},
			expected: expected{
				amount: 2,
			},
			expectError: false,
		},
		{
			name: "Success: filter by ceiling height",
			criteria: domain.ResidenceSearchCriteria{
				CeilingHeightMin: 2.5,
				CeilingHeightMax: 3,
			},
			expected: expected{
				amount: 2,
			},
			expectError: false,
		},
		{
			name: "Success: filter by min area",
			criteria: domain.ResidenceSearchCriteria{
				AreaMin: 50,
			},
			expected: expected{
				amount: 3,
			},
			expectError: false,
		},
		{
			name: "Success: filter by max area",
			criteria: domain.ResidenceSearchCriteria{
				AreaMax: 50,
			},
			expected: expected{
				amount: 4,
			},
			expectError: false,
		},
		{
			name: "Success: filter by area",
			criteria: domain.ResidenceSearchCriteria{
				AreaMin: 40,
				AreaMax: 150,
			},
			expected: expected{
				amount: 4,
			},
			expectError: false,
		},
		{
			name: "Success: filter by min price",
			criteria: domain.ResidenceSearchCriteria{
				PriceMin: 90,
				PriceMax: 300,
			},
			expected: expected{
				amount: 3,
			},
			expectError: false,
		},
		{
			name: "Success: filter by floors",
			criteria: domain.ResidenceSearchCriteria{
				FloorsMin: 3,
				FloorsMax: 5,
			},
			expected: expected{
				amount: 2,
			},
			expectError: false,
		},
		{
			name: "Success: filter by housing class",
			criteria: domain.ResidenceSearchCriteria{
				HousingClassID: 1,
			},
			expected: expected{
				amount: 2,
			},
			expectError: false,
		},
		{
			name: "Success: filter by construction types",
			criteria: domain.ResidenceSearchCriteria{
				ConstructionTypesIDs: []int64{1},
			},
			expected: expected{
				amount: 1,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			amount := test.expected.amount

			foundResidences, total, err := testResidencesRepository.List(ctx, test.criteria)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if foundResidences == nil {
					t.Error("residence has not been created")
					return
				}
				if total != domain.Total(amount) {
					t.Errorf("found residences: %d !=  %d", total, amount)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_ListPopularResidences(t *testing.T) {

}

func TestRepository_Get(t *testing.T) {

}

func TestRepository_Update(t *testing.T) {

}

func TestRepository_Delete(t *testing.T) {

}

func TestRepository_CreateFlatPlan(t *testing.T) {
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
		ctx       = context.Background()
		residence = &domain.Residence{
			StatusID:           1,
			UserID:             1,
			CityID:             1,
			DistrictID:         1,
			HousingClassID:     1,
			ConstructionTypeID: 1,
			Title:              "test",
			Description:        "test",
			FloorsMax:          5,
			HasHGF:             false,
			AreaMin:            40,
			AreaMax:            90,
			RoomsMin:           1,
			RoomsMax:           5,
			PriceMin:           40,
			PriceMax:           60,
			PricePerSquareMin:  50,
			CeilingHeight:      2.5,
		}
	)

	testResidencesRepository := NewResidenceRepository(db)

	// Create residence
	createdResidence, err := testResidencesRepository.Create(ctx, residence)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if createdResidence == nil {
		t.Fatal("residence has not been created")
	}

	// Define tests
	tests := []struct {
		name        string
		flatPlan    *domain.FlatPlan
		expectError bool
	}{
		{
			name: "#1 Flat plan",
			flatPlan: &domain.FlatPlan{
				ResidenceID:   1,
				NumberOfRooms: 1,
				Area:          10.0,
				Price:         1000,
				Images:        []string{"test", "test2"},
			},
			expectError: false,
		},
		{
			name: "#2 Flat plan",
			flatPlan: &domain.FlatPlan{
				ResidenceID:   1,
				NumberOfRooms: 2,
				Area:          20.0,
				Price:         2000,
				Images:        []string{"test", "test2"},
			},
			expectError: false,
		},
		{
			name: "#3 Flat plan",
			flatPlan: &domain.FlatPlan{
				ResidenceID:   1,
				NumberOfRooms: 3,
				Area:          30.0,
				Price:         3000,
				Images:        []string{"test", "test2"},
			},
			expectError: false,
		},
		{
			name: "#4 Flat plan",
			flatPlan: &domain.FlatPlan{
				ResidenceID:   1,
				NumberOfRooms: 4,
				Area:          40.0,
				Price:         4000,
				Images:        []string{"test", "test2"},
			},
			expectError: false,
		},
		{
			name: "#5 Flat plan",
			flatPlan: &domain.FlatPlan{
				ResidenceID:   1,
				NumberOfRooms: 5,
				Area:          50.0,
				Price:         5000,
				Images:        []string{"test", "test2"},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flatPlan := test.flatPlan
			result, err := testResidencesRepository.CreateFlatPlan(ctx, flatPlan)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == nil {
					t.Error("flat plan has not been created")
					return
				}
				if result.ID <= 0 {
					t.Error("bad flat plan id")
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

func TestRepository_UpdateFlatPlan(t *testing.T) {

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
		ctx       = context.Background()
		nowTime   = time.Now().Unix()
		residence = &domain.Residence{
			StatusID:           1,
			UserID:             1,
			CityID:             1,
			DistrictID:         1,
			HousingClassID:     1,
			ConstructionTypeID: 1,
			Title:              "test",
			Description:        "test",
			FloorsMax:          5,
			HasHGF:             false,
			AreaMin:            40,
			AreaMax:            90,
			RoomsMin:           1,
			RoomsMax:           5,
			PriceMin:           40,
			PriceMax:           60,
			PricePerSquareMin:  50,
			CeilingHeight:      2.5,
		}
		flatPlans = []*domain.FlatPlan{
			{
				ResidenceID:   1,
				NumberOfRooms: 1,
				Area:          10.0,
				Price:         1000,
				Images:        []string{"test", "test2"},
			},
			{
				ResidenceID:   1,
				NumberOfRooms: 2,
				Area:          20.0,
				Price:         2000,
				Images:        []string{"test", "test2"},
			},
			{
				ResidenceID:   1,
				NumberOfRooms: 3,
				Area:          30.0,
				Price:         3000,
				Images:        []string{"test", "test2"},
			},
			{
				ResidenceID:   1,
				NumberOfRooms: 4,
				Area:          40.0,
				Price:         4000,
				Images:        []string{"test", "test2"},
			},
			{
				ResidenceID:   1,
				NumberOfRooms: 5,
				Area:          50.0,
				Price:         5000,
				Images:        []string{"test", "test2"},
			},
		}
	)

	testResidencesRepository := NewResidenceRepository(db)

	// Create residence
	//
	createdResidence, err := testResidencesRepository.Create(ctx, residence)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if createdResidence == nil {
		t.Fatal("residence has not been created")
	}

	// Create flat plans
	//
	for _, flatPlan := range flatPlans {
		createdFlatPlan, err := testResidencesRepository.CreateFlatPlan(ctx, flatPlan)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if createdFlatPlan == nil {
			t.Fatal("residence has not been created")
		}
	}

	// Define tests
	//
	tests := []struct {
		name        string
		flatPlanID  domain.FlatPlanID
		argument    *domain.FlatPlan
		expected    *domain.FlatPlan
		expectError bool
	}{
		{
			name:       "#1: updated flat plan",
			flatPlanID: 1,
			argument: &domain.FlatPlan{
				ResidenceID:   1,
				NumberOfRooms: 1,
				Area:          20.0,
				Price:         1000,
				Images:        []string{"test", "test2", "test3"},
				UpdatedAt:     nowTime,
			},
			expected: &domain.FlatPlan{
				ResidenceID:   1,
				NumberOfRooms: 1,
				Area:          20.0,
				Price:         1000,
				Images:        []string{"test", "test2", "test3"},
				UpdatedAt:     nowTime,
			},
			expectError: false,
		},
		{
			name:       "#2: updated flat plan",
			flatPlanID: 1,
			argument: &domain.FlatPlan{
				ResidenceID:   1,
				NumberOfRooms: 2,
				Area:          40.0,
				Price:         2000,
				Images:        []string{"test"},
				UpdatedAt:     nowTime,
			},
			expected: &domain.FlatPlan{
				ResidenceID:   1,
				NumberOfRooms: 2,
				Area:          40.0,
				Price:         2000,
				Images:        []string{"test"},
				UpdatedAt:     nowTime,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := testResidencesRepository.UpdateFlatPlan(ctx, test.flatPlanID, test.argument)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if result == nil {
					t.Error("flat plan has not been created")
					return
				}
				if diff := deep.Equal(test.expected, result); diff != nil {
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

func TestRepository_DeleteFlatPlan(t *testing.T) {
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
		ctx       = context.Background()
		residence = &domain.Residence{
			StatusID:           1,
			UserID:             1,
			CityID:             1,
			DistrictID:         1,
			HousingClassID:     1,
			ConstructionTypeID: 1,
			Title:              "test",
			Description:        "test",
			FloorsMax:          5,
			HasHGF:             false,
			AreaMin:            40,
			AreaMax:            90,
			RoomsMin:           1,
			RoomsMax:           5,
			PriceMin:           40,
			PriceMax:           60,
			PricePerSquareMin:  50,
			CeilingHeight:      2.5,
		}
		flatPlans = []*domain.FlatPlan{
			{
				ResidenceID:   1,
				NumberOfRooms: 1,
				Area:          10.0,
				Price:         1000,
				Images:        []string{"test", "test2"},
			},
			{
				ResidenceID:   1,
				NumberOfRooms: 2,
				Area:          20.0,
				Price:         2000,
				Images:        []string{"test", "test2"},
			},
			{
				ResidenceID:   1,
				NumberOfRooms: 3,
				Area:          30.0,
				Price:         3000,
				Images:        []string{"test", "test2"},
			},
			{
				ResidenceID:   1,
				NumberOfRooms: 4,
				Area:          40.0,
				Price:         4000,
				Images:        []string{"test", "test2"},
			},
			{
				ResidenceID:   1,
				NumberOfRooms: 5,
				Area:          50.0,
				Price:         5000,
				Images:        []string{"test", "test2"},
			},
		}
	)

	testResidencesRepository := NewResidenceRepository(db)

	// Create residence
	//
	createdResidence, err := testResidencesRepository.Create(ctx, residence)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if createdResidence == nil {
		t.Fatal("residence has not been created")
	}

	// Create flat plans
	//
	for _, flatPlan := range flatPlans {
		createdFlatPlan, err := testResidencesRepository.CreateFlatPlan(ctx, flatPlan)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if createdFlatPlan == nil {
			t.Fatal("residence has not been created")
		}
	}

	// Define tests
	//
	tests := []struct {
		name        string
		flatPlanID  domain.FlatPlanID
		expectError bool
	}{
		{
			name:        "#1: delete flat plan",
			flatPlanID:  1,
			expectError: false,
		},
		{
			name:        "#1: delete flat plan",
			flatPlanID:  2,
			expectError: false,
		},
		{
			name:        "#1: delete flat plan",
			flatPlanID:  3,
			expectError: false,
		},
		{
			name:        "#1: delete flat plan",
			flatPlanID:  4,
			expectError: false,
		},
		{
			name:        "#1: delete flat plan",
			flatPlanID:  5,
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testResidencesRepository.DeleteFlatPlan(ctx, test.flatPlanID)
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

	// Check residence flat plans
	//
	existedResidence, err := testResidencesRepository.Get(ctx, domain.ResidenceID(1))
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if len(existedResidence.FlatPlans) != 0 {
		t.Errorf("flat plans not deleted:  %d != %d", len(existedResidence.FlatPlans), 0)
	}
}
