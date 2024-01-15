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

func TestService_RevokeExpiredLeadResidences(t *testing.T) {

	// Mock time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 06, 28, 0, 0, 0, 0, time.UTC)
	})

	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}
		total    = domain.Total(2)
		leads    = []*domain.LeadResidence{
			{
				ID:          1,
				ResidenceID: 1,
				StatusID:    1,
				IssuedAt:    1687543200,
				ExpiresAt:   1690394400,
			},
			{
				ID:          2,
				ResidenceID: 2,
				StatusID:    1,
				IssuedAt:    1687543200,
				ExpiresAt:   1690653600,
			},
		}
		nowTime = time.Now().Unix()
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks
	leadCottageService := mocks.NewMockLeadCottageService(stubCtrl)
	leadResidencesService := mocks.NewMockLeadResidenceService(stubCtrl)
	leadBuildersService := mocks.NewMockLeadBuilderService(stubCtrl)

	service := newBasicService(
		leadCottageService,
		leadResidencesService,
		leadBuildersService,
	)

	leadResidencesService.EXPECT().ListLeadResidences(ctx, domain.LeadResidenceSearchCriteria{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   1,
		},
	}, callerID).
		Return(nil, total, nil).
		AnyTimes()

	for offset := 0; offset < int(total); offset = offset + domain.DefaultPageSize {
		leadResidencesService.EXPECT().ListLeadResidences(ctx, domain.LeadResidenceSearchCriteria{
			Page: domain.PageRequest{
				Offset: offset,
				Size:   domain.DefaultPageSize,
			},
		}, callerID).
			Return(leads, domain.Total(2), nil).
			AnyTimes()

		for _, lead := range leads {
			if nowTime >= lead.ExpiresAt {
				leadResidencesService.EXPECT().
					RevokeLeadResidence(ctx, lead.ID, callerID).
					Return(nil).
					AnyTimes()
			}
		}
	}

	err := service.RevokeExpiredLeadResidences(ctx, callerID)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}

func TestService_RevokeExpiredLeadCottages(t *testing.T) {

	// Mock time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 06, 28, 0, 0, 0, 0, time.UTC)
	})

	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}
		total    = domain.Total(2)
		leads    = []*domain.LeadCottage{
			{
				ID:        1,
				CottageID: 1,
				StatusID:  1,
				IssuedAt:  1687543200,
				ExpiresAt: 1690394400,
			},
			{
				ID:        2,
				CottageID: 2,
				StatusID:  1,
				IssuedAt:  1687543200,
				ExpiresAt: 1690653600,
			},
			{
				ID:        3,
				CottageID: 3,
				StatusID:  1,
				IssuedAt:  1687543200,
				ExpiresAt: 1687559822,
			},
		}
		nowTime = time.Now().Unix()
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks
	leadCottageService := mocks.NewMockLeadCottageService(stubCtrl)
	leadResidencesService := mocks.NewMockLeadResidenceService(stubCtrl)
	leadBuildersService := mocks.NewMockLeadBuilderService(stubCtrl)

	service := newBasicService(
		leadCottageService,
		leadResidencesService,
		leadBuildersService,
	)

	leadCottageService.EXPECT().ListLeadCottage(ctx, domain.LeadCottageSearchCriteria{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   1,
		},
	}, callerID).
		Return(nil, total, nil).
		AnyTimes()

	for offset := 0; offset < int(total); offset = offset + domain.DefaultPageSize {
		leadCottageService.EXPECT().ListLeadCottage(ctx, domain.LeadCottageSearchCriteria{
			Page: domain.PageRequest{
				Offset: offset,
				Size:   domain.DefaultPageSize,
			},
		}, callerID).
			Return(leads, domain.Total(3), nil).
			AnyTimes()

		for _, lead := range leads {
			if nowTime >= lead.ExpiresAt {
				leadCottageService.EXPECT().
					RevokeLeadCottage(ctx, domain.LeadID(lead.ID), callerID).
					Return(nil)
			}
		}
	}

	err := service.RevokeExpiredLeadCottages(ctx, callerID)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}

func TestService_RevokeExpiredLeadBuilders(t *testing.T) {
	// Mock time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 06, 28, 0, 0, 0, 0, time.UTC)
	})

	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}
		total    = domain.Total(2)
		leads    = []*domain.LeadBuilder{
			{
				ID:        1,
				BuilderID: 1,
				StatusID:  1,
				IssuedAt:  1687543200,
				ExpiresAt: 1690394400,
			},
			{
				ID:        2,
				BuilderID: 2,
				StatusID:  1,
				IssuedAt:  1687543200,
				ExpiresAt: 1690653600,
			},
		}
		nowTime = time.Now().Unix()
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks
	leadCottageService := mocks.NewMockLeadCottageService(stubCtrl)
	leadResidencesService := mocks.NewMockLeadResidenceService(stubCtrl)
	leadBuildersService := mocks.NewMockLeadBuilderService(stubCtrl)

	service := newBasicService(
		leadCottageService,
		leadResidencesService,
		leadBuildersService,
	)

	leadBuildersService.EXPECT().ListLeadBuilders(ctx, domain.LeadBuilderSearchCriteria{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   1,
		},
	}, callerID).
		Return(nil, total, nil).
		AnyTimes()

	for offset := 0; offset < int(total); offset = offset + domain.DefaultPageSize {
		leadBuildersService.EXPECT().ListLeadBuilders(ctx, domain.LeadBuilderSearchCriteria{
			Page: domain.PageRequest{
				Offset: offset,
				Size:   domain.DefaultPageSize,
			},
		}, callerID).
			Return(leads, domain.Total(2), nil).
			AnyTimes()

		for _, lead := range leads {
			if nowTime >= lead.ExpiresAt {
				leadBuildersService.EXPECT().
					RevokeLeadBuilder(ctx, lead.ID, callerID).
					Return(nil).
					AnyTimes()
			}
		}
	}

	err := service.RevokeExpiredLeadBuilders(ctx, callerID)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}
