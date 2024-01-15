package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	"github.com/hashicorp/go-multierror"
)

type service struct {
	contactDetailsService domain.ContactDetailsService
	crmService            domain.CrmService
	amoCrmService         domain.AmoCrmService
	residenceService      domain.ResidencesService
	logger                log.Logger
}

// NewService - create a new service with necessary dependencies
func NewService(
	contactDetailsService domain.ContactDetailsService,
	crmService domain.CrmService,
	amoCrmService domain.AmoCrmService,
	residenceService domain.ResidencesService,
	logger log.Logger,
) domain.CronContactsService {
	var service domain.CronContactsService
	{
		service = newBasicService(
			contactDetailsService,
			crmService,
			amoCrmService,
			residenceService,
			logger,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	contactDetailsService domain.ContactDetailsService,
	crmService domain.CrmService,
	amoCrmService domain.AmoCrmService,
	residenceService domain.ResidencesService,
	logger log.Logger,
) domain.CronContactsService {
	return &service{
		contactDetailsService: contactDetailsService,
		crmService:            crmService,
		amoCrmService:         amoCrmService,
		residenceService:      residenceService,
		logger:                logger,
	}
}

func (s *service) SendContactsToCRM(
	ctx context.Context,
	callerID domain.CallerID,
) error {

	var (
		errs error
	)
	// List all the contacts
	_, total, err := s.contactDetailsService.ListResidenceContactDetails(ctx, domain.ResidenceContactDetailsSearchCriteria{
		Page: domain.PageRequest{
			Offset: 0,
			Size:   1,
		},
		IsDelivered: 2,
	}, callerID)
	if err != nil {
		return err
	}

	for offset := 0; offset < int(total); offset = offset + domain.DefaultPageSize {
		// List contacts
		contactDetails, _, err := s.contactDetailsService.ListResidenceContactDetails(
			ctx,
			domain.ResidenceContactDetailsSearchCriteria{
				Page: domain.PageRequest{
					Offset: offset,
					Size:   domain.DefaultPageSize,
				},
				IsDelivered: 2,
			}, callerID)
		if err != nil {
			return err
		}

		for _, contactDetail := range contactDetails {
			var (
				residence = &domain.Residence{}
			)

			// Check if contact details contains residence id
			if contactDetail.ResidenceID > 0 {
				residence, err = s.residenceService.GetResidence(
					ctx,
					domain.ResidenceID(contactDetail.ResidenceID),
					callerID,
				)
				if err != nil {
					errs = multierror.Append(errs, err)
					_ = s.logger.Log("err",
						fmt.Sprintf("residence not found by id %d, details %v", contactDetail.ResidenceID, err),
					)
					continue
				}
			}

			// Send to third party crm (builders)
			//if len(residence.Slug) > 0 {
			//	//err = s.crmService.SendResidenceContactDetail(
			//	//	ctx,
			//	//	domain.ResidenceContactDetailContent{
			//	//		ResidenceName: residence.Title,
			//	//		Slug:          residence.Slug,
			//	//		FullName:      contactDetail.FullName,
			//	//		Phone:         contactDetail.Phone,
			//	//	},
			//	//	callerID,
			//	//)
			//	//if err != nil {
			//	//	errs = multierror.Append(errs, err)
			//	//	_ = s.logger.Log("err",
			//	//		fmt.Sprintf("not sent to third party crm, details %v", err),
			//	//	)
			//	//	continue
			//	//}
			//}

			// Create application in internal crm service (bitrix)
			err = s.amoCrmService.CreateApplication(
				ctx,
				&domain.AmoApplication{
					FullName:         contactDetail.FullName,
					Phone:            contactDetail.Phone,
					ResidenceDetails: residence,
				},
			)
			if err != nil {
				errs = multierror.Append(errs, err)
				_ = s.logger.Log("err",
					fmt.Sprintf("not sent to bitrix crm, details %v", err),
				)
				continue
			}

			err = s.contactDetailsService.MarkAsDelivered(
				ctx,
				contactDetail.ID,
				callerID,
			)
			if err != nil {
				errs = multierror.Append(errs, err)
				_ = s.logger.Log("err",
					fmt.Sprintf("can not mar as delivered, details %v", err),
				)
				continue
			}
		}
	}

	return errs
}
