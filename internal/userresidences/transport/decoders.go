package transport

import (
	userresidenceapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

func decodeFavouriteResidenceSearchCriteria(apiCriteria *userresidenceapiv1.FavouriteResidencesSearchCriteria) domain.FavouriteResidencesSearchCriteria {

	if apiCriteria == nil {
		// default criteria.
		return domain.FavouriteResidencesSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.FavouriteResidencesSearchCriteria{
		Page:  helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		Sorts: helpers.DecodeSorts(apiCriteria.Sorts),
	}
}
