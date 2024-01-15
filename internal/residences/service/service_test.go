package service

import (
	"context"
	"fmt"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
)

func TestService_CreateResidence(t *testing.T) {
	var (
		ctx            = context.Background()
		validResidence = &domain.Residence{
			UserID:             1,
			StatusID:           1,
			CityID:             1,
			HousingClassID:     1,
			ConstructionTypeID: 1,
			Title:              "test",
			CeilingHeight:      3,
			DeadlineYear:       2022,
			DeadlineQuarter:    1,
		}
		callerID = domain.CallerID{}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mocks.NewMockResidencesRepository(stubCtrl)
	redisStub := mocks.NewMockResidenceRedisRepository(stubCtrl)
	service := newBasicService(repoStub, redisStub)

	// mock residences
	repoStub.EXPECT().
		Create(ctx, validResidence).
		Return(validResidence, nil).
		AnyTimes()

	// Define tests.
	type arguments struct {
		residence *domain.Residence
	}

	type result struct {
		residence *domain.Residence
	}

	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: residence created",
			arguments: arguments{
				residence: validResidence,
			},
			expected: result{
				residence: validResidence,
			},
			expectError: false,
		},
		{
			name: "Fail: invalid title",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid status id",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           0,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid city id",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             0,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid housing class id",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     0,
					ConstructionTypeID: 1,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid construction type id",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 0,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid ceiling height",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      0,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid deadline year",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      1,
					DeadlineYear:       0,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid deadline year",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      1,
					DeadlineYear:       2022,
					DeadlineQuarter:    5,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			expected := test.expected

			residence, err := service.CreateResidence(ctx, args.residence, callerID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					residence: residence,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
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

func TestService_UpdateResidence(t *testing.T) {

	var (
		ctx         = context.Background()
		residenceID = domain.ResidenceID(1)
		residence   = &domain.Residence{
			UserID:             1,
			StatusID:           1,
			CityID:             1,
			HousingClassID:     1,
			ConstructionTypeID: 1,
			Title:              "test",
			CeilingHeight:      3,
			DeadlineYear:       2022,
			DeadlineQuarter:    1,
		}
		updatedResidence = &domain.Residence{
			ID:                 residenceID,
			UserID:             1,
			StatusID:           2,
			CityID:             1,
			HousingClassID:     1,
			ConstructionTypeID: 1,
			Title:              "new title",
			CeilingHeight:      3.5,
			DeadlineYear:       2023,
			DeadlineQuarter:    2,
		}
		callerID = domain.CallerID{}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mocks.NewMockResidencesRepository(stubCtrl)
	redisStub := mocks.NewMockResidenceRedisRepository(stubCtrl)
	service := newBasicService(repoStub, redisStub)

	// mock residences
	repoStub.EXPECT().
		Update(ctx, residenceID, residence).
		Return(residence, nil).
		AnyTimes()

	// Define tests.
	type arguments struct {
		residenceID domain.ResidenceID
		residence   *domain.Residence
	}

	type result struct {
		residence *domain.Residence
	}

	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Succes:: updated residence",
			arguments: arguments{
				residenceID: domain.ResidenceID(1),
				residence:   residence,
			},
			expected: result{
				residence: updatedResidence,
			},
			expectError: false,
		},
		{
			name: "Fail: invalid title",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid status id",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           0,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid city id",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             0,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid housing class id",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     0,
					ConstructionTypeID: 1,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid construction type id",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 0,
					CeilingHeight:      3,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid ceiling height",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      0,
					DeadlineYear:       2022,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid deadline year",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      1,
					DeadlineYear:       0,
					DeadlineQuarter:    1,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid deadline year",
			arguments: arguments{
				residence: &domain.Residence{
					Title:              "",
					StatusID:           1,
					CityID:             1,
					HousingClassID:     1,
					ConstructionTypeID: 1,
					CeilingHeight:      1,
					DeadlineYear:       2022,
					DeadlineQuarter:    5,
				},
			},
			expected: result{
				residence: nil,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			expected := test.expected

			residence, err := service.UpdateResidence(ctx, args.residenceID, args.residence, callerID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					residence: residence,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
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

func TestService_DeleteResidence(t *testing.T) {

	var (
		ctx         = context.Background()
		callerID    = domain.CallerID{}
		residenceID = domain.ResidenceID(1)
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mocks.NewMockResidencesRepository(stubCtrl)
	redisStub := mocks.NewMockResidenceRedisRepository(stubCtrl)
	service := newBasicService(repoStub, redisStub)

	// mock residences
	repoStub.EXPECT().
		Delete(ctx, residenceID).
		Return(nil).
		AnyTimes()

	// Define tests.
	type arguments struct {
		residenceID domain.ResidenceID
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: residence deleted",
			arguments: arguments{
				residenceID: residenceID,
			},
			expectError: false,
		},
		{
			name: "Fail: invalid residence id",
			arguments: arguments{
				residenceID: 0,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			err := service.DeleteResidence(ctx, args.residenceID, callerID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestService_ListResidence(t *testing.T) {

	var (
		ctx        = context.Background()
		callerID   = domain.CallerID{}
		criteria   = domain.ResidenceSearchCriteria{}
		total      = domain.Total(2)
		residences = []*domain.Residence{
			{
				ID:                 1,
				UserID:             1,
				StatusID:           1,
				CityID:             1,
				HousingClassID:     1,
				ConstructionTypeID: 1,
				Title:              "test",
				CeilingHeight:      3,
				DeadlineYear:       2022,
				DeadlineQuarter:    1,
			},
			{
				ID:                 2,
				UserID:             1,
				StatusID:           1,
				CityID:             1,
				HousingClassID:     1,
				ConstructionTypeID: 1,
				Title:              "test2",
				CeilingHeight:      3,
				DeadlineYear:       2022,
				DeadlineQuarter:    2,
			},
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mocks.NewMockResidencesRepository(stubCtrl)
	redisStub := mocks.NewMockResidenceRedisRepository(stubCtrl)
	service := newBasicService(repoStub, redisStub)

	// mock residences
	repoStub.EXPECT().
		List(ctx, criteria).
		Return(residences, total, nil).
		AnyTimes()

	type result struct {
		total domain.Total
	}

	tests := []struct {
		name        string
		result      result
		expectError bool
	}{
		{
			name: "Success: list residences",
			result: result{
				total: total,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			residences, total, err := service.ListResidences(ctx, criteria, callerID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				if int(total) != len(residences) {
					t.Errorf("unexpected error: total not equal to length of residences")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestService_GetResidence(t *testing.T) {
	var (
		ctx          = context.Background()
		callerID     = domain.CallerID{}
		residenceID  = domain.ResidenceID(1)
		residenceID2 = domain.ResidenceID(2)
		residence    = &domain.Residence{
			ID:                 1,
			UserID:             1,
			StatusID:           1,
			CityID:             1,
			HousingClassID:     1,
			ConstructionTypeID: 1,
			Title:              "test",
			CeilingHeight:      3,
			DeadlineYear:       2022,
			DeadlineQuarter:    1,
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mocks.NewMockResidencesRepository(stubCtrl)
	redisStub := mocks.NewMockResidenceRedisRepository(stubCtrl)
	service := newBasicService(repoStub, redisStub)

	// mock residences
	repoStub.EXPECT().
		Get(ctx, residenceID).
		Return(residence, nil).
		AnyTimes()

	repoStub.EXPECT().
		Get(ctx, residenceID2).
		Return(nil, fmt.Errorf("not found")).
		AnyTimes()

	type result struct {
		residenceID domain.ResidenceID
	}

	tests := []struct {
		name        string
		result      result
		expectError bool
	}{
		{
			name: "Success: found residence",
			result: result{
				residenceID: residenceID,
			},
			expectError: false,
		},
		{
			name: "Success: not found",
			result: result{
				residenceID: residenceID2,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid residence",
			result: result{
				residenceID: domain.ResidenceID(0),
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			residence, err := service.GetResidence(ctx, test.result.residenceID, callerID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				if residence == nil {
					t.Errorf("unexpected error: nil residence")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestService_IsResidenceExist(t *testing.T) {

	var (
		ctx          = context.Background()
		callerID     = domain.CallerID{}
		residenceID1 = domain.ResidenceID(1)
		residenceID2 = domain.ResidenceID(2)
		residence1   = &domain.Residence{}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mocks.NewMockResidencesRepository(stubCtrl)
	redisStub := mocks.NewMockResidenceRedisRepository(stubCtrl)
	service := newBasicService(repoStub, redisStub)

	// mock residences
	repoStub.EXPECT().
		Get(ctx, residenceID1).
		Return(residence1, nil).
		AnyTimes()

	repoStub.EXPECT().
		Get(ctx, residenceID2).
		Return(nil, errors.NewErrInvalidArgument("not found")).
		AnyTimes()

	// arguments
	type arguments struct {
		residenceID domain.ResidenceID
	}

	// result
	type result struct {
		exist bool
	}

	// Test cases
	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: residence is exist",
			arguments: arguments{
				residenceID: 1,
			},
			expected: result{
				exist: true,
			},
			expectError: false,
		},
		{
			name: "Success: residence is not exist",
			arguments: arguments{
				residenceID: 2,
			},
			expected: result{
				exist: false,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid residence id",
			arguments: arguments{
				residenceID: 0,
			},
			expected: result{
				exist: false,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			expected := test.expected

			exist, err := service.IsResidenceExist(ctx, args.residenceID, callerID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					exist: exist,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
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
