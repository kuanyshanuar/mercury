package service

import (
	"context"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"github.com/golang/mock/gomock"
)

type dummyLogger struct{}

func (l dummyLogger) Log(_ ...interface{}) error {
	return nil
}

func TestService_SendContactsToCRM(t *testing.T) {

	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}
		total    = domain.Total(2)
		contacts = []*domain.ResidenceContactDetails{
			{
				ID:          1,
				ResidenceID: 1,
				FullName:    "Test",
				Phone:       "Test",
				IsDelivered: false,
			},
			{
				ID:          2,
				ResidenceID: 2,
				FullName:    "Test",
				Phone:       "Test",
				IsDelivered: false,
			},
		}
		residence = &domain.Residence{}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mock dependencies
	contactDetailsService := mocks.NewMockContactDetailsService(stubCtrl)
	crmService := mocks.NewMockCrmService(stubCtrl)
	amoCrmService := mocks.NewMockAmoCrmService(stubCtrl)
	residenceService := mocks.NewMockResidencesService(stubCtrl)

	service := newBasicService(
		contactDetailsService,
		crmService,
		amoCrmService,
		residenceService,
		dummyLogger{},
	)

	contactDetailsService.EXPECT().ListResidenceContactDetails(ctx, domain.ResidenceContactDetailsSearchCriteria{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   1,
		},
		IsDelivered: 2,
	}, callerID).
		Return(nil, total, nil).
		AnyTimes()

	for offset := 0; offset < int(total); offset = offset + domain.DefaultPageSize {
		contactDetailsService.EXPECT().ListResidenceContactDetails(ctx, domain.ResidenceContactDetailsSearchCriteria{
			Page: domain.PageRequest{
				Offset: offset,
				Size:   domain.DefaultPageSize,
			},
			IsDelivered: 2,
		}, callerID).
			Return(contacts, total, nil).
			AnyTimes()

		for _, contact := range contacts {

			residenceService.EXPECT().
				GetResidence(ctx, domain.ResidenceID(contact.ResidenceID), callerID).
				Return(residence, nil).
				AnyTimes()

			amoCrmService.EXPECT().CreateApplication(
				ctx,
				&domain.AmoApplication{
					FullName:         contact.FullName,
					Phone:            contact.Phone,
					ResidenceDetails: residence,
				},
			).Return(nil).
				AnyTimes()

			contactDetailsService.EXPECT().MarkAsDelivered(
				ctx,
				contact.ID,
				callerID,
			).
				Return(nil).
				AnyTimes()
		}
	}

	err := service.SendContactsToCRM(ctx, callerID)
	if err != nil {
		t.Error(err)
	}
}
