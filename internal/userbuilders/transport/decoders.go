package transport

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

func encodeBuildersV1(r []*domain.Builder) []*apiv1.BuilderRead {
	if r == nil {
		return nil
	}

	builders := make([]*apiv1.BuilderRead, len(r))
	for i, opt := range r {
		builders[i] = encodeBuilderV1(opt)
	}

	return builders
}

func encodeBuilderV1(r *domain.Builder) *apiv1.BuilderRead {
	if r == nil {
		return nil
	}

	return &apiv1.BuilderRead{
		Id:                      int64(r.ID),
		FirstName:               r.FirstName,
		LastName:                r.LastName,
		Email:                   r.Email,
		City:                    r.City,
		ConsultationPhoneNumber: r.ConsultationPhoneNumber,
		Image:                   r.Image,
		IsFavourite:             r.IsFavourite,
	}
}

func encodeSearchCriteria(apiCriteria *apiv1.ListSubscribedBuildersRequest_SearchCriteria) domain.UserBuilderSearchCriteria {
	if apiCriteria == nil {
		// default criteria.
		return domain.UserBuilderSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.UserBuilderSearchCriteria{
		Page:  helpers.DecodePageRequestV1(apiCriteria.Page),
		Sorts: helpers.DecodeSorts(apiCriteria.Sorts),
	}
}
