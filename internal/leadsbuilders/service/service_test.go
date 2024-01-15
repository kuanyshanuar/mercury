package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"github.com/golang/mock/gomock"
)

func TestService_CreateLeadBuilder(t *testing.T) {

	var (
		ctx      = context.Background()
		nowTime  = time.Now()
		callerID = domain.CallerID{}
		lead     = &domain.LeadBuilder{
			BuilderID: 1,
			StatusID:  1,
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mock repositories
	repoStub := mocks.NewMockLeadBuilderRepository(stubCtrl)

	repoStub.EXPECT().
		CreateLeadBuilder(ctx, lead).
		Return(nil).
		AnyTimes()

	service := newBasicService(repoStub)

	type arguments struct {
		lead *domain.LeadBuilder
	}

	// Define tests
	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: create lead",
			arguments: arguments{
				lead: lead,
			},
			expectError: false,
		},
		{
			name: "Fail: invalid builder id",
			arguments: arguments{
				lead: &domain.LeadBuilder{
					BuilderID: 0,
					StatusID:  1,
					IssuedAt:  1,
					ExpiresAt: 1,
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid status id",
			arguments: arguments{
				lead: &domain.LeadBuilder{
					BuilderID: 1,
					StatusID:  0,
					IssuedAt:  1,
					ExpiresAt: 1,
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid issued date",
			arguments: arguments{
				lead: &domain.LeadBuilder{
					BuilderID: 1,
					StatusID:  1,
					IssuedAt:  0,
					ExpiresAt: 1,
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid expire date",
			arguments: arguments{
				lead: &domain.LeadBuilder{
					BuilderID: 1,
					StatusID:  1,
					IssuedAt:  1,
					ExpiresAt: 0,
				},
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := service.CreateLeadBuilder(ctx, test.arguments.lead, callerID)
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

func TestService_UpdateLeadBuilder(t *testing.T) {

	var (
		ctx      = context.Background()
		nowTime  = time.Now()
		callerID = domain.CallerID{}
		leadID   = domain.LeadID(1)
		lead     = &domain.LeadBuilder{
			BuilderID: 1,
			StatusID:  1,
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: nowTime.Add(24 * time.Hour).Unix(),
		}
		leadID2 = domain.LeadID(2)
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mock repositories
	repoStub := mocks.NewMockLeadBuilderRepository(stubCtrl)

	repoStub.EXPECT().
		UpdateLeadBuilder(ctx, leadID, lead).
		Return(nil).
		AnyTimes()

	repoStub.EXPECT().
		UpdateLeadBuilder(ctx, leadID2, lead).
		Return(errors.New("record not found")).
		AnyTimes()

	service := newBasicService(repoStub)

	type arguments struct {
		leadID domain.LeadID
		lead   *domain.LeadBuilder
	}

	// Define tests
	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: updated lead",
			arguments: arguments{
				leadID: leadID,
				lead:   lead,
			},
			expectError: false,
		},
		{
			name: "Fail: invalid id",
			arguments: arguments{
				leadID: 0,
				lead:   lead,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid builder id",
			arguments: arguments{
				leadID: 1,
				lead: &domain.LeadBuilder{
					BuilderID: 0,
					StatusID:  1,
					IssuedAt:  1,
					ExpiresAt: 1,
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid status id",
			arguments: arguments{
				leadID: 1,
				lead: &domain.LeadBuilder{
					BuilderID: 1,
					StatusID:  0,
					IssuedAt:  1,
					ExpiresAt: 1,
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid issued date",
			arguments: arguments{
				leadID: 1,
				lead: &domain.LeadBuilder{
					BuilderID: 1,
					StatusID:  1,
					IssuedAt:  0,
					ExpiresAt: 1,
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid expiration date",
			arguments: arguments{
				leadID: 1,
				lead: &domain.LeadBuilder{
					BuilderID: 1,
					StatusID:  1,
					IssuedAt:  1,
					ExpiresAt: 0,
				},
			},
			expectError: true,
		},
		{
			name: "Fail: not found",
			arguments: arguments{
				leadID: leadID2,
				lead:   lead,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			err := service.UpdateLeadBuilder(ctx, args.leadID, args.lead, callerID)
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

func TestService_GetLeadBuilder(t *testing.T) {

}

func TestService_ListLeadBuilder(t *testing.T) {

}

func TestService_DeleteLeadBuilder(t *testing.T) {

}

func TestService_RevokeLeadBuilder(t *testing.T) {

}
