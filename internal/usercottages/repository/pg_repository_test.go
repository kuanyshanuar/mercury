package repository

import (
	"context"
	"testing"

	cottage "gitlab.com/zharzhanov/mercury/internal/cottages/repository"
	"gitlab.com/zharzhanov/mercury/internal/domain"

	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
)

func TestRepository_AddFavouriteCottage(t *testing.T) {
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

	var (
		ctx                   = context.Background()
		validFavouriteCottage = []*domain.UserCottage{
			{
				CottageID: 1,
				UserID:    1,
			},
			{
				CottageID: 2,
				UserID:    1,
			},
			{
				CottageID: 1,
				UserID:    2,
			},
			{
				CottageID: 1,
				UserID:    2,
			},
		}
		validCottages = []*domain.Cottage{{
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
			CanRePlan:         false,
			AreaMin:           30,
			AreaMax:           180,
			PricePerSquareMin: 15000,
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
			CeilingHeightMin:  2.4,
			BuildingArea:      1000,
			HouseAmount:       6,
			FloorsCount:       2,
			Facade:            "open facade",
			WindowTypes:       nil,
			CanRePlan:         true,
			AreaMin:           100,
			AreaMax:           1000,
			PricePerSquareMin: 0,
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
			PricePerSquareMin: 15000,
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
				PricePerSquareMin: 1500,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
			},
		}
	)

	userCottageRepo := NewRepository(db)
	cottageRepo := cottage.NewCottageRepository(db)
	type userCottageTest struct {
		name        string
		arg1        interface{}
		arg2        interface{}
		expectError bool
	}
	tests := []userCottageTest{
		{
			name:        "UserCottage Create success #1",
			arg1:        validFavouriteCottage[0],
			arg2:        validCottages,
			expectError: false,
		},
		{
			name:        "UserCottage Create success #2",
			arg1:        validFavouriteCottage[1],
			arg2:        validCottages,
			expectError: false,
		},
		{
			name:        "UserCottage Create success #3",
			arg1:        validFavouriteCottage[2],
			arg2:        validCottages,
			expectError: false,
		},
		{
			name:        "UserCottage Create Fail #4(copy value)",
			arg1:        validFavouriteCottage[3],
			arg2:        validCottages,
			expectError: false,
		},
	}
	for _, ctg := range validCottages {
		_, err = cottageRepo.Create(ctx, ctg)
		if err != nil {
			t.Fatal(err, "couldn't create cottage")
		}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !test.expectError {
				userCottage := test.arg1.(*domain.UserCottage)
				err = userCottageRepo.AddFavouriteCottage(ctx, userCottage.UserID, userCottage.CottageID)
				if err != nil {
					t.Fatalf("couldn't add favourite cottage %s", err)
				}
				xists, err := userCottageRepo.IsCottageExists(ctx, userCottage.UserID, userCottage.CottageID)
				if err != nil {
					t.Fatalf("couldn't verify favourite cottage %s", err)
				}
				if !xists {
					t.Fatal("favourite cottage wasn't created or verified")
				}

			}
		})
	}
}

func TestRepository_DeleteFavouriteCottage(t *testing.T) {
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

	var (
		ctx                   = context.Background()
		validFavouriteCottage = []*domain.UserCottage{
			{
				CottageID: 1,
				UserID:    1,
			},
			{
				CottageID: 2,
				UserID:    1,
			},
			{
				CottageID: 1,
				UserID:    2,
			},
			{
				CottageID: 1,
				UserID:    2,
			},
		}
		validCottages = []*domain.Cottage{{
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
		}, {
			Title:             "ZK-vostochnoe",
			Description:       "Vostochnaya oblast'",
			Address:           "Semey, District Valikhanova",
			Latitude:          0.5,
			Longitude:         2.6,
			Territory:         "open territory",
			CeilingHeightMin:  2.4,
			BuildingArea:      1000,
			HouseAmount:       6,
			FloorsCount:       2,
			Facade:            "open facade",
			WindowTypes:       nil,
			CanRePlan:         true,
			AreaMin:           100,
			AreaMax:           1000,
			PricePerSquareMin: 0,
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
			PricePerSquareMin: 15000,
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
				PricePerSquareMin: 1500,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
			},
		}
	)

	userCottageRepo := NewRepository(db)
	cottageRepo := cottage.NewCottageRepository(db)
	type userCottageTest struct {
		name        string
		arg1        interface{}
		arg2        interface{}
		expectError bool
	}
	tests := []userCottageTest{
		{
			name:        "UserCottage Delete success #1",
			arg1:        validFavouriteCottage[0],
			arg2:        validCottages,
			expectError: false,
		},
		{
			name:        "UserCottage Delete success #2",
			arg1:        validFavouriteCottage[1],
			arg2:        validCottages,
			expectError: false,
		},
		{
			name:        "UserCottage Delete success #3",
			arg1:        validFavouriteCottage[2],
			arg2:        validCottages,
			expectError: false,
		},
		{
			name:        "UserCottage Delete Fail #4(copy value)",
			arg1:        validFavouriteCottage[3],
			arg2:        validCottages,
			expectError: false,
		},
	}
	for _, ctg := range validCottages {
		_, err = cottageRepo.Create(ctx, ctg)
		if err != nil {
			t.Fatal(err, "couldn't create cottage")
		}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !test.expectError {
				userCottage := test.arg1.(*domain.UserCottage)
				err = userCottageRepo.AddFavouriteCottage(ctx, userCottage.UserID, userCottage.CottageID)
				if err != nil {
					t.Fatalf("couldn't add favourite cottage %s", err)
				}

				err = userCottageRepo.DeleteFavouriteCottage(ctx, userCottage.UserID, userCottage.CottageID)
				if err != nil {
					t.Fatal("couldn't delete favourite cottage")
				}
				xists, err := userCottageRepo.IsCottageExists(ctx, userCottage.UserID, userCottage.CottageID)
				if err != nil {
					t.Fatalf("couldn't verify favourite cottage %s", err)
				}
				if xists {
					t.Fatal("favourite cottage wasn't deleted")
				}

			}
		})
	}
}

func TestRepository_ListFavouriteCottages(t *testing.T) {
	var (
		ctx                   = context.Background()
		validFavouriteCottage = []*domain.UserCottage{
			{
				CottageID: 1,
				UserID:    1,
			},
			{
				CottageID: 2,
				UserID:    1,
			},
			{
				CottageID: 1,
				UserID:    2,
			},
			{
				CottageID: 1,
				UserID:    2,
			},
		}
		validCottages = []*domain.Cottage{{
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
		}, {
			Title:             "ZK-vostochnoe",
			Description:       "Vostochnaya oblast'",
			Address:           "Semey, District Valikhanova",
			Latitude:          0.5,
			Longitude:         2.6,
			Territory:         "open territory",
			CeilingHeightMin:  2.4,
			BuildingArea:      1000,
			HouseAmount:       6,
			FloorsCount:       2,
			Facade:            "open facade",
			WindowTypes:       nil,
			CanRePlan:         true,
			AreaMin:           100,
			AreaMax:           1000,
			PricePerSquareMin: 0,
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
			PricePerSquareMin: 15000,
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
				PricePerSquareMin: 1500,
				CityID:            3,
				DistrictID:        2,
				StatusID:          1,
				HousingClassID:    1,
				UserID:            1,
			},
		}
	)
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	defer func() {
		_ = sqlDB.Close()
	}()
	userCottageRepo := NewRepository(db)
	cottageRepo := cottage.NewCottageRepository(db)
	type userCottageTest struct {
		name        string
		arg1        interface{}
		arg2        interface{}
		arg3        domain.FavouriteCottagesSearchCriteria
		expectError bool
	}
	tests := []userCottageTest{
		{
			name: "UserCottage List success #1",
			arg1: validFavouriteCottage[0:2],
			arg2: validCottages,
			arg3: domain.FavouriteCottagesSearchCriteria{
				Page: domain.PageRequest{
					Offset: 0,
					Size:   2,
				},
				Sorts: []domain.Sort{
					{
						FieldName: "created_at",
						Order:     "ASC",
					},
				},
			},
			expectError: false,
		},
		{
			name: "UserCottage List success #2",
			arg1: validFavouriteCottage[1:3],
			arg2: validCottages,
			arg3: domain.FavouriteCottagesSearchCriteria{
				Page: domain.PageRequest{
					Offset: 0,
					Size:   2,
				},
				Sorts: []domain.Sort{
					{
						FieldName: "created_at",
						Order:     "ASC",
					},
				},
			},
			expectError: false,
		},
		{
			name: "UserCottage List success #3",
			arg1: validFavouriteCottage[1:2],
			arg2: validCottages,
			arg3: domain.FavouriteCottagesSearchCriteria{
				Page: domain.PageRequest{
					Offset: 0,
					Size:   2,
				},
				Sorts: []domain.Sort{
					{
						FieldName: "created_at",
						Order:     "ASC",
					},
				},
			},
			expectError: false,
		},
		{
			name: "UserCottage List Success #4",
			arg1: validFavouriteCottage[0:4],
			arg2: validCottages,
			arg3: domain.FavouriteCottagesSearchCriteria{
				Page: domain.PageRequest{
					Offset: 0,
					Size:   4,
				},
				Sorts: []domain.Sort{
					{
						FieldName: "created_at",
						Order:     "ASC",
					},
				},
			},
			expectError: false,
		},
	}
	for _, ctg := range validCottages {
		_, err = cottageRepo.Create(ctx, ctg)
		if err != nil {
			t.Fatal(err, "couldn't create cottage")
		}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !test.expectError {
				userCottages := test.arg1.([]*domain.UserCottage)
				for _, userCtg := range userCottages {
					err = userCottageRepo.AddFavouriteCottage(ctx, userCtg.UserID, userCtg.CottageID)
					if err != nil {
						t.Fatalf("couldn't add favourite cottage %s", err)
					}
				}
				list, total, err := userCottageRepo.ListFavouriteCottages(ctx, 1, test.arg3)
				t.Logf("list is %v, total is %v", list, total)
				if total == 0 {
					t.Fatal("expected more than 0 favourite cottages")
				}
				if list == nil {
					t.Fatal("list of favourite cottages is nil")
				}
				if err != nil {
					t.Fatalf("unexpected error, %s", err.Error())
				}
			}
		})
	}
}
