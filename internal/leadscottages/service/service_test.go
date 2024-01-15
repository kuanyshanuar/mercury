package service

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
	mocks "gitlab.com/zharzhanov/mercury/internal/mocks"
	"reflect"
	"testing"
)

func TestService_CreateLeadCottage(t *testing.T) {
	var (
		ctx        = context.Background()
		validLeads = []*domain.LeadCottage{
			{
				CottageID: 1,
				IssuedAt:  1693471997,
				StatusID:  1,
				ExpiresAt: 1693571997,
				DeletedAt: 0,
			},
			{
				CottageID: 2,
				IssuedAt:  1693471997,
				StatusID:  1,
				ExpiresAt: 1693571997,
				DeletedAt: 0,
			},
			{
				CottageID: 2,
				IssuedAt:  1693473997,
				StatusID:  0,
				ExpiresAt: 1693571996,
				DeletedAt: 0,
			},
			{
				CottageID: 2,
				IssuedAt:  1693473997,
				StatusID:  1,
				ExpiresAt: 1693571996,
				DeletedAt: 0,
			},
		}
	)
	ctrl := gomock.NewController(t)
	mockLeadCottage := mocks.NewMockLeadCottageRepository(ctrl)
	mockLeadCottage.EXPECT().CreateLeadCottage(ctx, validLeads[0]).Return(int64(1), nil)
	mockLeadCottage.EXPECT().IsLeadExistByDate(ctx, validLeads[0].CottageID, validLeads[0].IssuedAt, validLeads[0].ExpiresAt).Return(false, nil)
	mockLeadCottage.EXPECT().CreateLeadCottage(ctx, validLeads[1]).Return(int64(2), nil)
	mockLeadCottage.EXPECT().IsLeadExistByDate(ctx, validLeads[1].CottageID, validLeads[1].IssuedAt, validLeads[1].ExpiresAt).Return(false, nil)
	mockLeadCottage.EXPECT().IsLeadExistByDate(ctx, validLeads[2].CottageID, validLeads[2].IssuedAt, validLeads[2].ExpiresAt).Return(true, nil)

	type test struct {
		name        string
		arg1        *domain.LeadCottage
		arg2        int64
		expectError bool
	}

	tests := []test{
		{
			name: "#1: CreateLeadCottage Success",
			arg1: &domain.LeadCottage{
				CottageID: 1,
				StatusID:  1,
				IssuedAt:  1693471997,
				ExpiresAt: 1693571997,
				DeletedAt: 0,
			},
			arg2:        1,
			expectError: false,
		},
		{
			name: "#2: CreateLeadCottage Success",
			arg1: &domain.LeadCottage{
				CottageID: 2,
				StatusID:  1,
				IssuedAt:  1693471997,
				ExpiresAt: 1693571997,
				DeletedAt: 0,
			},
			arg2:        2,
			expectError: false,
		},
		{
			name: "#3: CreateLeadCottage Fail",
			arg1: &domain.LeadCottage{
				CottageID: 2,
				StatusID:  1,
				IssuedAt:  1693473997,
				ExpiresAt: 1693571996,
				DeletedAt: 0,
			},
			arg2:        0,
			expectError: true,
		},
	}

	service := newBasicService(mockLeadCottage)
	//errors.NewErrInvalidArgument("lead exists between these dates")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !test.expectError {
				id, err := service.CreateLeadCottage(ctx, test.arg1, helpers.CallerID(ctx))
				if id != test.arg2 {
					t.Fatal("service test id isn't valid")
				}
				if err != nil {
					t.Fatal(err)
				}
			} else {
				_, err := service.CreateLeadCottage(ctx, test.arg1, helpers.CallerID(ctx))
				if err == nil {
					t.Fatal("expected error, but got nothing")
				}

			}

		})
	}
}
func TestService_UpdateLeadCottage(t *testing.T) {
	var (
		ctx        = context.Background()
		validLeads = []*domain.LeadCottage{
			{
				ID:        1,
				CottageID: 1,
				IssuedAt:  1693471997,
				StatusID:  1,
				ExpiresAt: 1693571997,
				DeletedAt: 0,
			},
			{
				ID:        1,
				CottageID: 2,
				IssuedAt:  1693471997,
				StatusID:  1,
				ExpiresAt: 1693571997,
				DeletedAt: 0,
			},
			{
				ID:        1,
				CottageID: 3,
				IssuedAt:  1693473997,
				StatusID:  0,
				ExpiresAt: 1693571996,
				DeletedAt: 0,
			},
			{
				ID:        1,
				CottageID: 4,
				IssuedAt:  1693473997,
				StatusID:  1,
				ExpiresAt: 1693571996,
				DeletedAt: 0,
			},
		}
	)

	ctrl := gomock.NewController(t)
	mockLeadCottage := mocks.NewMockLeadCottageRepository(ctrl)
	for i := range validLeads {

		mockLeadCottage.EXPECT().UpdateLeadCottage(ctx, int64(1), validLeads[i]).Return(validLeads[i], nil)
		mockLeadCottage.EXPECT().IsOtherLeadExist(ctx, domain.LeadID(1), validLeads[i].CottageID, validLeads[i].IssuedAt, validLeads[i].ExpiresAt).
			Return(false, nil)
	}

	type test struct {
		name        string
		arg1        *domain.LeadCottage
		expectError bool
	}

	tests := []test{
		{
			name: "#1: UpdateLeadCottage Success",
			arg1: &domain.LeadCottage{
				ID:        1,
				CottageID: 1,
				IssuedAt:  1693471997,
				StatusID:  1,
				ExpiresAt: 1693571997,
				DeletedAt: 0,
			},

			expectError: false,
		},
		{
			name: "#2: UpdateLeadCottage Success",
			arg1: &domain.LeadCottage{
				ID:        1,
				CottageID: 2,
				IssuedAt:  1693471997,
				StatusID:  1,
				ExpiresAt: 1693571997,
				DeletedAt: 0,
			},
			expectError: false,
		},
		{
			name: "#3: UpdateLeadCottage Success",
			arg1: &domain.LeadCottage{
				ID:        1,
				CottageID: 3,
				IssuedAt:  1693473997,
				StatusID:  0,
				ExpiresAt: 1693571996,
				DeletedAt: 0,
			},

			expectError: false,
		},
		{
			name: "#4: UpdateLeadCottage Success",
			arg1: &domain.LeadCottage{
				ID:        1,
				CottageID: 4,
				IssuedAt:  1693473997,
				StatusID:  1,
				ExpiresAt: 1693571996,
				DeletedAt: 0,
			},

			expectError: false,
		},
	}

	service := newBasicService(mockLeadCottage)
	//errors.NewErrInvalidArgument("lead exists between these dates")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !test.expectError {
				lead, err := service.UpdateLeadCottage(ctx, int64(1), test.arg1, helpers.CallerID(ctx))
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(lead, test.arg1) {
					fmt.Print(lead, test.arg1)
					t.Fatal("service update invalid")
				}
			}

		})
	}

}
