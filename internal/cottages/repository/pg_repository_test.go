package repository

import (
	"context"
	"fmt"
	leadscottages "gitlab.com/zharzhanov/mercury/internal/leadscottages/repository"
	"reflect"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
)

func TestGetListCottages(t *testing.T) {
	var (
		ctx = context.Background()
	)
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal("Database is invalid")
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	cottageRepository := NewCottageRepository(db)
	criteriaValid := []*domain.CottageSearchCriteria{{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   10,
		},
		Sorts:              nil,
		Title:              "",
		BuilderIDs:         []int64{1, 2},
		RoomsMin:           2,
		RoomsMax:           3,
		CeilingHeightMin:   0,
		CeilingHeightMax:   0,
		AreaMin:            0,
		AreaMax:            0,
		PriceMin:           0,
		PriceMax:           0,
		InteriorDecoration: nil,
		HeatingTypes:       nil,
		PurchaseMethods:    nil,
		ElevatorTypes:      nil,
		HouseType:          0,
		FloorsMin:          0,
		FloorsMax:          0,
	},
		{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   10,
			},
			Title:            "ZHK",
			RoomsMin:         1,
			RoomsMax:         5,
			CeilingHeightMin: 0,
			CeilingHeightMax: 6,
			AreaMin:          10,
			AreaMax:          1000,
			PriceMin:         10,
			PriceMax:         50000,
		},
		{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   100,
			},
			Title:              "ZHK-borisovka",
			BuilderIDs:         nil,
			RoomsMin:           1,
			RoomsMax:           5,
			AreaMin:            10,
			AreaMax:            1000,
			PriceMin:           10,
			PriceMax:           50000,
			InteriorDecoration: nil,
			HeatingTypes:       nil,
			PurchaseMethods:    nil,
			ElevatorTypes:      nil,
			HouseType:          0,
			FloorsMin:          0,
			FloorsMax:          0,
		},
		{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   10,
			},
			Title:            "",
			BuilderIDs:       nil,
			RoomsMin:         1,
			RoomsMax:         5,
			CeilingHeightMin: 0,
			CeilingHeightMax: 5,
			AreaMin:          12,
			AreaMax:          1000,
			PriceMin:         10,
			PriceMax:         50000,
		},
	}
	// Valid cottages that suit all constraints. USE THESE FOR THE FUTURE TESTS!
	cottagesValid := []*domain.Cottage{{
		Title:             "ZHK-borisovka",
		Description:       "This is the ZK in the Almaty, in the village of borisovka",
		Address:           "Almaty-Auezova-46",
		Latitude:          10,
		Longitude:         10,
		Territory:         "open territory",
		CeilingHeightMin:  0,
		CeilingHeightMax:  20,
		BuildingArea:      100,
		HouseAmount:       4,
		FloorsCount:       5,
		Facade:            "excellent facade",
		WindowTypes:       nil,
		CanRePlan:         false,
		AreaMin:           30,
		AreaMax:           180,
		PricePerSquareMin: 10,
		PricePerSquareMax: 25000,
		CityID:            4,
		DistrictID:        1,
		StatusID:          1,
		HousingClassID:    1,
		UserID:            1,
	}, {
		Title:             "ZK-vostochnoe",
		Description:       "Vostochnaya oblast'",
		Address:           "Semey, District Valikhanova",
		Latitude:          0.5,
		Longitude:         2.6,
		Territory:         "open territory",
		CeilingHeightMin:  0,
		CeilingHeightMax:  20,
		BuildingArea:      1000,
		HouseAmount:       6,
		FloorsCount:       2,
		Facade:            "open facade",
		WindowTypes:       nil,
		CanRePlan:         true,
		AreaMin:           100,
		AreaMax:           1000,
		PricePerSquareMin: 10,
		PricePerSquareMax: 25000,
		CityID:            2,
		DistrictID:        2,
		StatusID:          1,
		HousingClassID:    1,
		UserID:            1,
	}, {
		Title:             "Kottedzhnyy gorodok",
		Description:       "Cottage city complex in Astana",
		Address:           "Esil district, 46",
		Latitude:          10,
		Longitude:         100,
		Territory:         "closed territory",
		CeilingHeightMin:  47,
		BuildingArea:      100,
		HouseAmount:       40,
		FloorsCount:       50,
		Facade:            "good facade",
		WindowTypes:       nil,
		CanRePlan:         false,
		AreaMin:           30,
		AreaMax:           180,
		PricePerSquareMin: 10,
		PricePerSquareMax: 25000,
		CityID:            1,
		DistrictID:        1,
		StatusID:          1,
		HousingClassID:    1,
		UserID:            1,
	},
		{
			Title:             "Maxi-Cottage",
			Description:       "Complex of cottages in the higher district",
			Address:           "Uralsk, Kazanoe, 56",
			Latitude:          10,
			Longitude:         50,
			Territory:         "excellent facade",
			CeilingHeightMin:  4,
			BuildingArea:      10,
			HouseAmount:       404,
			FloorsCount:       50,
			Facade:            "good facade",
			WindowTypes:       nil,
			CanRePlan:         true,
			AreaMin:           100,
			AreaMax:           300,
			PricePerSquareMin: 10,
			PricePerSquareMax: 25000,
			CityID:            3,
			DistrictID:        2,
			StatusID:          1,
			HousingClassID:    1,
			UserID:            1,
		},
	}

	type ListTests struct {
		name          string
		arg2          *domain.CottageSearchCriteria
		arg1          []*domain.Cottage
		expectedError bool
	}
	tests := []ListTests{
		{
			name:          "ListOfCottages success #1",
			arg1:          cottagesValid,
			arg2:          criteriaValid[0],
			expectedError: false,
		},
		{
			name:          "ListOfCottages success #2",
			arg1:          cottagesValid,
			arg2:          criteriaValid[1],
			expectedError: false,
		},
		{
			name:          "ListOfCottages success #3",
			arg1:          cottagesValid,
			arg2:          criteriaValid[2],
			expectedError: false,
		},
		{
			name:          "ListOfCottages success #4",
			arg1:          cottagesValid,
			arg2:          criteriaValid[3],
			expectedError: false,
		},
	}
	for _, val := range tests[0].arg1 {
		cottage, err := cottageRepository.Create(ctx, val)
		if err != nil {
			t.Fatal(err, "failed create repository")
		}

		if cottage.ID <= 0 {
			t.Fatal(err, "failed create repository, invalid ID")
		}

		if cottage.Title != val.Title {
			t.Fatal(err, "failed create repository, invalid title")
		}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			cottages, count, err := cottageRepository.List(ctx, *test.arg2)

			if count <= 0 {
				t.Fatal(err, "count return is invalid")
			}
			if len(cottages) <= 0 {
				t.Fatal(err, "cottages len is invalid")
			}
		})
	}
}

func TestDeleteCottage(t *testing.T) {

	var (
		ctx = context.Background()
	)
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal("Database is invalid")
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = sqlDB.Close()
	}()

	cottageRepository := NewCottageRepository(db)

	type testDelete struct {
		name string
		arg1 *domain.Cottage
	}
	tests := []testDelete{{
		name: "Success delete test #1",
		arg1: &domain.Cottage{
			Title:             "ZHK-borisovka",
			Description:       "This is the ZK in the Almaty, in the village of borisovka",
			Address:           "Almaty-Auezova-46",
			Latitude:          10,
			Longitude:         10,
			Territory:         "open territory",
			CeilingHeightMin:  47,
			BuildingArea:      100,
			HouseAmount:       4,
			FloorsCount:       5,
			Facade:            "excellent facade",
			WindowTypes:       nil,
			CanRePlan:         false,
			AreaMin:           30,
			AreaMax:           180,
			PricePerSquareMin: 15000,
			CityID:            4,
			DistrictID:        1,
			StatusID:          1,
			HousingClassID:    1,
			UserID:            1,
		},
	},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cottage, err := cottageRepository.Create(ctx, test.arg1)
			if err != nil {
				t.Fatal(err)
			}
			if cottage == nil {
				t.Fatal("invalid cottage create return")
			}
			err = cottageRepository.Delete(ctx, cottage.ID)
			if err != nil {
				t.Fatal(err)
			}
			cottage, err = cottageRepository.Get(ctx, cottage.ID)
			if err == nil {
				if cottage != nil {
					t.Fatal("cottage wasn't deleted")
				}
				t.Fatal("expected error, but got nothing")
			}
		})
	}

}

func TestUpdateCottage(t *testing.T) {
	var (
		ctx = context.Background()
	)
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal("Database is invalid")
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = sqlDB.Close()
	}()

	type test struct {
		name          string
		arg1          *domain.Cottage // original cottage creation
		arg2          *domain.Cottage // modified cottage
		expectedError bool
	}
	cottageRepository := NewCottageRepository(db)
	tests := []test{
		{
			name: "test update cottage: change basic fields",
			arg1: &domain.Cottage{
				Title:             "Maxi-Cottage",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10,
				Longitude:         50,
				Territory:         "excellent facade",
				CeilingHeightMin:  4,
				BuildingArea:      10,
				HouseAmount:       404,
				FloorsCount:       50,
				Facade:            "good facade",
				WindowTypes:       nil,
				CanRePlan:         true,
				AreaMin:           100,
				AreaMax:           300,
				PricePerSquareMin: 1500,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
				//HousePlans: []*domain.HousePlan{
				//	{
				//		CottageID:      1,
				//		Title:          "",
				//		NumberOfRooms:  0,
				//		Area:           0,
				//		Longitude:      0,
				//		Territory:      0,
				//		CeilingHeightMin:  0,
				//		Price:          0,
				//		PricePerSquareMin: 0,
				//		PlanImages:     nil,
				//		HouseImages:    nil,
				//		HousingClassID: 0,
				//		CreatedAt:      0,
				//		UpdatedAt:      0,
				//		DeletedAt:      0,
				//	},
				//},
			},
			arg2: &domain.Cottage{
				Title:             "Maxi-Cottage",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10,
				Longitude:         50,
				Territory:         "Good Territory",
				CeilingHeightMin:  4.3,
				BuildingArea:      100,
				HouseAmount:       404,
				FloorsCount:       500,
				Facade:            "facade is a little bit different",
				WindowTypes:       nil,
				CanRePlan:         true,
				AreaMin:           100,
				AreaMax:           300,
				PricePerSquareMin: 1500,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
				//HousePlans: []*domain.HousePlan{
				//	{
				//		CottageID:      1,
				//		Title:          "House_plan_id",
				//		NumberOfRooms:  0,
				//		Area:           0,
				//		Longitude:      0,
				//		Territory:      0,
				//		CeilingHeightMin:  0,
				//		Price:          0,
				//		PricePerSquareMin: 0,
				//		PlanImages:     nil,
				//		HouseImages:    nil,
				//		HousingClassID: 0,
				//		CreatedAt:      0,
				//		UpdatedAt:      0,
				//		DeletedAt:      0,
				//	},
				//},
			},
			expectedError: false,
		},
		{
			name: "test update cottage: change fields that are connected through associations",
			arg1: &domain.Cottage{
				Title:             "Maxi-Cottage",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10,
				Longitude:         50,
				Territory:         "excellent facade",
				CeilingHeightMin:  4,
				BuildingArea:      10,
				HouseAmount:       404,
				FloorsCount:       50,
				Facade:            "good facade",
				WindowTypes:       nil,
				CanRePlan:         true,
				AreaMin:           100,
				AreaMax:           300,
				PricePerSquareMin: 1500,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
				//HousePlans: []*domain.HousePlan{
				//	{
				//		CottageID:      1,
				//		Title:          "House_plan_id",
				//		NumberOfRooms:  0,
				//		Area:           0,
				//		Longitude:      0,
				//		Territory:      0,
				//		CeilingHeightMin:  0,
				//		Price:          0,
				//		PricePerSquareMin: 0,
				//		PlanImages:     nil,
				//		HouseImages:    nil,
				//		HousingClassID: 0,
				//		CreatedAt:      0,
				//		UpdatedAt:      0,
				//		DeletedAt:      0,
				//	},
				//},
			},
			arg2: &domain.Cottage{
				Title:             "Maxi-Cottage",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10,
				Longitude:         50,
				Territory:         "Good Territory",
				CeilingHeightMin:  4.3,
				BuildingArea:      100,
				HouseAmount:       404,
				FloorsCount:       500,
				Facade:            "facade is a little bit different",
				WindowTypes:       nil,
				CanRePlan:         true,
				AreaMin:           100,
				AreaMax:           300,
				PricePerSquareMin: 1500,
				CityID:            1,
				DistrictID:        1,
				StatusID:          2,
				HousingClassID:    2,
				UserID:            3,
				//HousePlans: []*domain.HousePlan{
				//	{
				//		CottageID:      1,
				//		Title:          "other_house_plan",
				//		NumberOfRooms:  0,
				//		Area:           0,
				//		Longitude:      0,
				//		Territory:      0,
				//		CeilingHeightMin:  0,
				//		Price:          0,
				//		PricePerSquareMin: 0,
				//		PlanImages:     nil,
				//		HouseImages:    nil,
				//		HousingClassID: 0,
				//	},
				//},
			},
			expectedError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cottage1, err := cottageRepository.Create(ctx, test.arg1)
			fmt.Println(cottage1.City, cottage1.Status, cottage1.User, cottage1.District)
			if err != nil {
				t.Fatal("couldn't create cottage")
			}
			cottage2, err := cottageRepository.Update(ctx, cottage1.ID, test.arg2)
			if err != nil {
				t.Fatal("couldn't update cottage")
			}
			if !reflect.DeepEqual(cottage2, test.arg2) {
				t.Fatal("cottage update was false")
			}

		})
	}
}

func TestListCottagesByID(t *testing.T) {
}

func TestListPopularCottages(t *testing.T) {

	var (
		ctx           = context.Background()
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
			{
				Title:             "Simple Fields #3(ceilingHegihtMinMax, areaMinMax, PriceMinMax)",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10000,
				Longitude:         50000,
				Territory:         "excellent facade",
				CeilingHeightMin:  1.2,
				CeilingHeightMax:  5.0,
				BuildingArea:      100,
				HouseAmount:       404000,
				RoomsMin:          1000,
				RoomsMax:          6000,
				FloorsCount:       50000,
				Facade:            "good facade",
				CanRePlan:         true,
				AreaMin:           60,
				AreaMax:           150,
				PricePerSquareMin: 13000,
				PricePerSquareMax: 19000,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
			},
			{
				/*
								CeilingHeightMin: 1.5,
					CeilingHeightMax: 4.8,
					AreaMin:          60,
					AreaMax:          150,
					PriceMin:         10000,
					PriceMax:         22500,

				*/
				Title:             "Simple Fields #4(ceilingHegihtMinMax, areaMinMax, PriceMinMax)",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10000,
				Longitude:         50000,
				Territory:         "excellent facade",
				CeilingHeightMin:  1.2,
				CeilingHeightMax:  2.6,
				BuildingArea:      100,
				HouseAmount:       404000,
				RoomsMin:          1000,
				RoomsMax:          6000,
				FloorsCount:       50000,
				Facade:            "good facade",
				CanRePlan:         true,
				AreaMin:           100000000,
				AreaMax:           300000000,
				PricePerSquareMin: 12000,
				PricePerSquareMax: 20000,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
			},
			{
				/*
					HouseType:   2,
						FloorsMin:   100,
						FloorsMax:   120,
						UserID:      2,
						CanRePlan:   true,
						HouseAmount: 45,
				*/
				Title:             "Simple Fields #5(HouseType, FloorsMinMax, UserID, CanRePlan, HouseAmount)",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10000,
				Longitude:         50000,
				Territory:         "excellent facade",
				CeilingHeightMin:  1.9,
				BuildingArea:      100,
				HouseAmount:       45,
				RoomsMin:          1000,
				RoomsMax:          6000,
				FloorsCount:       115,
				Facade:            "good facade",
				CanRePlan:         true,
				AreaMin:           100000000,
				AreaMax:           300000000,
				PricePerSquareMin: 12000,
				InteriorDecorations: []*domain.InteriorDecoration{
					{ID: 2},
				},
				HeatingTypes: []*domain.HeatingType{
					{ID: 2},
				},
				PurchaseMethods: []*domain.PurchaseMethod{
					{ID: 2},
				},
				ElevatorTypes: []*domain.ElevatorType{
					{ID: 2},
				},
				PricePerSquareMax: 20000,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    2,
				UserID:            2,
			},
			{
				Title:             "Simple Fields #6(ceilingHegihtMinMax, areaMinMax, PriceMinMax)",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10000,
				Longitude:         50000,
				Territory:         "excellent facade",
				CeilingHeightMin:  1.9,
				BuildingArea:      100,
				HouseAmount:       404000,
				RoomsMin:          1000,
				RoomsMax:          6000,
				FloorsCount:       50000,
				Facade:            "good facade",
				CanRePlan:         true,
				AreaMin:           100000000,
				AreaMax:           300000000,
				PricePerSquareMin: 12000,
				WallTypes: []*domain.WallType{
					{ID: 2},
				},
				WarmingTypes: []*domain.WarmingType{
					{ID: 2},
				},
				WindowTypes: []*domain.WindowType{
					{ID: 2},
				},
				PricePerSquareMax: 20000,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
				HousePlans: []*domain.HousePlan{
					{
						CottageID:      6,
						HousingClassID: 2,
						Title:          "House Plan #1",
						NumberOfRooms:  0,
						Area:           0,
						Longitude:      0,
						Territory:      0,
						CeilingHeight:  0,
						Price:          0,
						PricePerSquare: 0,
						PlanImages:     nil,
						HouseImages:    nil,
					},
				},
			},
		}
		validLeads = []*domain.LeadCottage{
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

			{
				CottageID: 3,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},

			{
				CottageID: 4,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},

			{
				CottageID: 5,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},

			{
				CottageID: 6,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},

			{
				CottageID: 7,
				StatusID:  1,
				IssuedAt:  1687802400,
				ExpiresAt: 1687975200,
			},
		}
	)
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = sqlDB.Close()
	}()
	cottageRepo := NewCottageRepository(db)
	leadsCottageRepo := leadscottages.NewRepository(db)
	type test struct {
		name           string
		searchCriteria domain.CottageSearchCriteria
		expectedTotal  int64
		expectError    bool
	}

	tests := []test{
		{
			name:           "#1 Success Empty test",
			searchCriteria: domain.CottageSearchCriteria{},
			expectedTotal:  6,
			expectError:    false,
		},
		{
			name: "#2 Success simple group test(city, district, status, rooms_min, rooms_max)",
			searchCriteria: domain.CottageSearchCriteria{
				CityID:     3,
				DistrictID: 2,
				StatusID:   1,
				RoomsMin:   2,
				RoomsMax:   5,
			},
			expectedTotal: 2,
			expectError:   false,
		},
		{
			name: "#3 Success simple group test(ceilingHegihtMinMax, areaMinMax, PriceMinMax)",
			searchCriteria: domain.CottageSearchCriteria{
				CeilingHeightMin: 1.5,
				CeilingHeightMax: 4.8,
				AreaMin:          60,
				AreaMax:          150,
				PriceMin:         10000,
				PriceMax:         22500,
			},
			expectedTotal: 1,
			expectError:   false,
		},
		{
			name: "#4 Success simple group test(HouseType, FloorsMinMax, UserID, CanRePlan, HouseAmount)",
			searchCriteria: domain.CottageSearchCriteria{
				HouseType:   2,
				FloorsMin:   100,
				FloorsMax:   120,
				UserID:      2,
				CanRePlan:   true,
				HouseAmount: 45,
			},
			expectedTotal: 1,
			expectError:   false,
		},
		{
			name: "#5 Success array group test(InteriorDecoration, HeatingTypes, PurchaseMethods, ElevatorTypes)",
			searchCriteria: domain.CottageSearchCriteria{

				InteriorDecoration: []int64{2},
				HeatingTypes:       []int64{2},
				PurchaseMethods:    []int64{2},
				ElevatorTypes:      []int64{2},
			},
			expectedTotal: 1,
			expectError:   false,
		},
		{
			name: "#6 Success array group test(WallTypes, WindowTypes, WarmingTypes)",
			searchCriteria: domain.CottageSearchCriteria{
				WallTypes:    []int64{2},
				WindowTypes:  []int64{2},
				WarmingTypes: []int64{2},
			},
			expectedTotal: 1,
			expectError:   false,
		},
	}

	// adding window types
	db.Exec("INSERT INTO window_types(name) VALUES ('plastic'), ('old');")

	// adding wall types
	db.Exec("INSERT INTO wall_types(name) VALUES ('white'), ('black');")

	// adding warming types
	db.Exec("INSERT INTO warming_types(name) VALUES ('cool'), ('uncool');")

	for i, cottage := range validCottages {
		_, err := cottageRepo.Create(ctx, cottage)

		if err != nil {
			t.Fatalf("couldn't create cottage %s", err)
		}
		_, err = leadsCottageRepo.CreateLeadCottage(ctx, validLeads[i])
		if err != nil {
			t.Fatalf("couldn't create leads cottage %s", err)
		}
	}

	for _, unit := range tests {
		t.Run(unit.name, func(t *testing.T) {

			list, total, err := cottageRepo.ListPopularCottages(ctx, unit.searchCriteria)
			if int64(len(list)) != unit.expectedTotal {
				t.Fatalf("expected length of popular lists didn't meet requirements %d, %v", len(list), list)
			}
			if err != nil {
				t.Fatalf("couldn't execute listPopularCottages, %s", err)
			}
			if total <= 0 {
				t.Fatalf("total popular cottages length is invalid")
			}
			// Checking associations
			t.Logf("criteria: %v, search result: %v, \n %v \n %v", unit.searchCriteria, *list[0].Status, *list[0].City, *list[0].User)
		})

	}
}

func TestDeleteHousePlan(t *testing.T) {
	var (
		ctx = context.Background()
	)
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = sqlDB.Close()
	}()
	cottageRepository := NewCottageRepository(db)
	type testDeletePlan struct {
		name        string
		arg1        *domain.HousePlan
		arg2        *domain.Cottage
		expectError bool
	}
	tests := []testDeletePlan{
		{
			name: "Success Delete House Plan #1",
			arg1: &domain.HousePlan{
				CottageID:      1,
				Title:          "House Plan #1",
				NumberOfRooms:  2,
				Area:           100,
				Longitude:      23,
				Territory:      0,
				CeilingHeight:  0,
				Price:          0,
				PricePerSquare: 0,
				PlanImages:     nil,
				HouseImages:    nil,
				HousingClassID: 2,
			},
			arg2: &domain.Cottage{
				Title:             "Maxi-Cottage",
				Description:       "Complex of cottages in the higher district",
				Address:           "Uralsk, Kazanoe, 56",
				Latitude:          10,
				Longitude:         50,
				Territory:         "Good Territory",
				CeilingHeightMin:  4.3,
				BuildingArea:      100,
				HouseAmount:       404,
				FloorsCount:       500,
				Facade:            "facade is a little bit different",
				WindowTypes:       nil,
				CanRePlan:         true,
				AreaMin:           100,
				AreaMax:           300,
				PricePerSquareMin: 1500,
				CityID:            1,
				DistrictID:        1,
				StatusID:          1,
				HousingClassID:    2,
				UserID:            3,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cottage, err := cottageRepository.Create(ctx, test.arg2)
			if err != nil || cottage == nil {
				t.Fatalf("couldn't create cottage %s %v", err, cottage)
			}
			housePlan, err := cottageRepository.CreateHousePlan(ctx, test.arg1)
			if housePlan != nil {
				t.Log(housePlan)
			} else if err != nil {
				t.Fatalf("couldn't create housePlan %v", housePlan)
			}
			err = cottageRepository.DeleteHousePlan(ctx, 1)
			if err != nil {
				t.Fatal("couldn't delete house plan")
			}
		})
	}
}
