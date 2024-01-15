package service

import (
	"context"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/log"
)

type service struct {
	leadCottageService   domain.LeadCottageService
	leadResidenceService domain.LeadResidenceService
	leadBuilderService   domain.LeadBuilderService
}

// NewService - creates a new service
func NewService(
	leadCottageService domain.LeadCottageService,
	leadResidenceService domain.LeadResidenceService,
	leadBuilderService domain.LeadBuilderService,
	logger log.Logger,
) domain.CronLeadService {
	var service domain.CronLeadService
	{
		service = newBasicService(
			leadCottageService,
			leadResidenceService,
			leadBuilderService,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	leadCottageService domain.LeadCottageService,
	leadResidenceService domain.LeadResidenceService,
	leadBuilderService domain.LeadBuilderService,
) domain.CronLeadService {
	return &service{
		leadCottageService:   leadCottageService,
		leadResidenceService: leadResidenceService,
		leadBuilderService:   leadBuilderService,
	}
}

func (s *service) RevokeExpiredLeadResidences(
	ctx context.Context,
	callerID domain.CallerID,
) (err error) {

	var (
		nowTime = time.Now().Unix()
	)

	// Do 1 request to calculate total rows
	// Fetch all leads
	_, total, err := s.leadResidenceService.ListLeadResidences(ctx, domain.LeadResidenceSearchCriteria{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   1,
		},
	}, helpers.CallerID(ctx))
	if err != nil {
		return nil
	}

	for offset := 0; offset < int(total); offset = offset + domain.DefaultPageSize {
		// List leads
		leads, _, err := s.leadResidenceService.ListLeadResidences(ctx, domain.LeadResidenceSearchCriteria{
			Page: domain.PageRequest{
				Offset: offset,
				Size:   domain.DefaultPageSize,
			},
		}, helpers.CallerID(ctx))
		if err != nil {
			return nil
		}

		for _, lead := range leads {
			// Check expiration time
			if nowTime >= lead.ExpiresAt && lead.StatusID == domain.StatusActive {
				// Disable lead by id
				err = s.leadResidenceService.RevokeLeadResidence(ctx, lead.ID, callerID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *service) RevokeExpiredLeadCottages(
	ctx context.Context,
	callerID domain.CallerID,
) (err error) {

	var (
		nowTime = time.Now().Unix()
	)

	// Do 1 request to calculate total rows
	// Fetch all leads
	_, total, err := s.leadCottageService.ListLeadCottage(ctx, domain.LeadCottageSearchCriteria{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   1,
		},
	}, helpers.CallerID(ctx))
	if err != nil {
		return nil
	}

	for offset := 0; offset < int(total); offset = offset + domain.DefaultPageSize {
		// List leads
		leads, _, err := s.leadCottageService.ListLeadCottage(ctx, domain.LeadCottageSearchCriteria{
			Page: domain.PageRequest{
				Offset: offset,
				Size:   domain.DefaultPageSize,
			},
		}, helpers.CallerID(ctx))
		if err != nil {
			return nil
		}

		for _, lead := range leads {
			// Check expiration time
			if nowTime >= lead.ExpiresAt && lead.StatusID == domain.StatusActive {
				// Disable lead by id
				err = s.leadCottageService.RevokeLeadCottage(ctx, domain.LeadID(lead.ID), callerID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *service) RevokeExpiredLeadBuilders(
	ctx context.Context,
	callerID domain.CallerID,
) error {
	var (
		nowTime = time.Now().Unix()
	)

	// Do 1 request to calculate total rows
	// Fetch all leads
	_, total, err := s.leadBuilderService.ListLeadBuilders(ctx, domain.LeadBuilderSearchCriteria{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   1,
		},
	}, helpers.CallerID(ctx))
	if err != nil {
		return nil
	}

	for offset := 0; offset < int(total); offset = offset + domain.DefaultPageSize {
		// List leads
		leads, _, err := s.leadBuilderService.ListLeadBuilders(ctx, domain.LeadBuilderSearchCriteria{
			Page: domain.PageRequest{
				Offset: offset,
				Size:   domain.DefaultPageSize,
			},
		}, helpers.CallerID(ctx))
		if err != nil {
			return nil
		}

		for _, lead := range leads {
			// Check expiration time
			if nowTime >= lead.ExpiresAt && lead.StatusID == domain.StatusActive {
				// Disable lead by id
				err = s.leadBuilderService.RevokeLeadBuilder(ctx, lead.ID, callerID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
