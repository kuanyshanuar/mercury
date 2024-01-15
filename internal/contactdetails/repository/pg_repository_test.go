package repository

import (
	"context"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
)

func TestRepository_CreateResidenceContactDetails(t *testing.T) {
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
	tests := []struct {
		name           string
		contactDetails *domain.ResidenceContactDetails
	}{
		{
			name: "Success: contacts #1",
			contactDetails: &domain.ResidenceContactDetails{
				ResidenceID: 1,
				FullName:    "Test",
				Phone:       "77051234567",
			},
		},
		{
			name: "Success: contacts #2",
			contactDetails: &domain.ResidenceContactDetails{
				ResidenceID: 2,
				FullName:    "Test2",
				Phone:       "77051234567",
			},
		},
		{
			name: "Success: contacts #3",
			contactDetails: &domain.ResidenceContactDetails{
				ResidenceID: 3,
				FullName:    "Test3",
				Phone:       "77051234567",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.CreateResidenceContactDetails(ctx, test.contactDetails)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRepository_ListResidenceContactDetails(t *testing.T) {
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

	testRepository := NewRepository(db)

	var (
		ctx      = context.Background()
		contacts = []*domain.ResidenceContactDetails{
			{
				ID:          0,
				ResidenceID: 1,
				FullName:    "Test",
				Phone:       "77051234567",
				IsDelivered: false,
				CreatedAt:   1,
			},
			{
				ID:          0,
				ResidenceID: 2,
				FullName:    "Test2",
				Phone:       "77051234567",
				IsDelivered: true,
				CreatedAt:   2,
			},
			{
				ID:          0,
				ResidenceID: 3,
				FullName:    "Test3",
				Phone:       "77051234567",
				IsDelivered: false,
				CreatedAt:   3,
			},
		}
	)

	for _, contact := range contacts {
		err = testRepository.CreateResidenceContactDetails(ctx, contact)
		if err != nil {
			t.Error(err)
		}
	}

	// Define tests
	tests := []struct {
		name          string
		criteria      domain.ResidenceContactDetailsSearchCriteria
		expectedList  int
		expectedTotal int
	}{
		{
			name: "Success: list all",
			criteria: domain.ResidenceContactDetailsSearchCriteria{
				Page: domain.PageRequest{
					Offset: 0,
					Size:   10,
				},
				IsDelivered: 0,
			},
			expectedList:  3,
			expectedTotal: 3,
		},
		{
			name: "Success: list delivered",
			criteria: domain.ResidenceContactDetailsSearchCriteria{
				Page: domain.PageRequest{
					Offset: 0,
					Size:   10,
				},
				IsDelivered: 1,
			},
			expectedList:  1,
			expectedTotal: 1,
		},
		{
			name: "Success: list not delivered",
			criteria: domain.ResidenceContactDetailsSearchCriteria{
				Page: domain.PageRequest{
					Offset: 0,
					Size:   10,
				},
				IsDelivered: 2,
			},
			expectedList:  2,
			expectedTotal: 2,
		},
		{
			name: "Success: list not delivered",
			criteria: domain.ResidenceContactDetailsSearchCriteria{
				Page: domain.PageRequest{
					Offset: 0,
					Size:   10,
				},
				ResidenceID: 1,
				Phone:       "77051234567",
				FromTime:    1,
				ToTime:      2,
			},
			expectedList:  1,
			expectedTotal: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			createdContacts, total, err := testRepository.ListResidenceContactDetails(ctx, test.criteria)
			if err != nil {
				t.Error(err)
			}
			if len(createdContacts) != test.expectedList {
				t.Errorf("list number is not matched %d != %d", len(createdContacts), test.expectedList)
			}
			if int(total) != test.expectedTotal {
				t.Errorf("total is not matched %d != %d", total, test.expectedTotal)
			}
		})
	}
}

func TestRepository_MarkAsDelivered(t *testing.T) {
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

	testRepository := NewRepository(db)

	var (
		ctx      = context.Background()
		contacts = []*domain.ResidenceContactDetails{
			{
				ID:          0,
				ResidenceID: 1,
				FullName:    "Test",
				Phone:       "77051234567",
				IsDelivered: false,
			},
			{
				ID:          0,
				ResidenceID: 2,
				FullName:    "Test2",
				Phone:       "77051234567",
				IsDelivered: false,
			},
			{
				ID:          0,
				ResidenceID: 3,
				FullName:    "Test3",
				Phone:       "77051234567",
				IsDelivered: false,
			},
		}
	)

	for _, contact := range contacts {
		err = testRepository.CreateResidenceContactDetails(ctx, contact)
		if err != nil {
			t.Error(err)
		}
	}

	// Define tests
	tests := []struct {
		name      string
		contactID int64
	}{
		{
			name:      "Success: #1",
			contactID: 1,
		},
		{
			name:      "Success: #2",
			contactID: 2,
		},
		{
			name:      "Success: #3",
			contactID: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err = testRepository.MarkAsDelivered(ctx, test.contactID)
			if err != nil {
				t.Error(err)
			}
		})
	}

	list, total, err := testRepository.ListResidenceContactDetails(ctx, domain.ResidenceContactDetailsSearchCriteria{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   10,
		},
		IsDelivered: 2,
	})
	if err != nil {
		t.Error(err)
	}

	if total != 0 {
		t.Errorf("total is unmatched")
	}
	if len(list) != 0 {
		t.Errorf("length of list is unmatched")
	}
}
