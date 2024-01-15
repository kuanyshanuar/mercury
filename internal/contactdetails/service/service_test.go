package service

import (
	"context"
	"testing"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
)

func TestService_CreateContactDetails(t *testing.T) {

}

func TestService_CreateResidenceContactDetails(t *testing.T) {
	var (
		ctx            = context.Background()
		callerID       = domain.CallerID{}
		contactDetails = &domain.ResidenceContactDetails{
			ResidenceID: 1,
			FullName:    "Test",
			Phone:       "12345678",
		}
	)

	// Mock time.Now()
	patch := monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 06, 28, 12, 0, 0, 0, time.UTC)
	})
	defer patch.Unpatch()

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mock dependencies
	var (
		mailService              = mocks.NewMockMailService(stubCtrl)
		contactDetailsRepository = mocks.NewMockContactDetailsRepository(stubCtrl)
	)

	contactDetailsRepository.EXPECT().
		ListResidenceContactDetails(ctx, domain.ResidenceContactDetailsSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   1,
			},
			ResidenceID: contactDetails.ResidenceID,
			Phone:       contactDetails.Phone,
			FromTime:    1687996800,
			ToTime:      1687910400,
		}).
		Return([]*domain.ResidenceContactDetails{}, domain.Total(0), nil).
		AnyTimes()

	contactDetailsRepository.EXPECT().
		CreateResidenceContactDetails(ctx, contactDetails).
		Return(nil).
		AnyTimes()

	// Run tests
	service := newBasicService(
		mailService,
		contactDetailsRepository,
	)

	type arguments struct {
		contactDetails *domain.ResidenceContactDetails
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: contact created",
			arguments: arguments{
				contactDetails: contactDetails,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			err := service.CreateResidenceContactDetails(ctx, args.contactDetails, callerID)
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

func TestService_ListContactDetails(t *testing.T) {

}

func TestService_ListResidenceContactDetails(t *testing.T) {

}
