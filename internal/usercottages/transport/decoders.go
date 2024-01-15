package transport

import (
	userresidenceapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

func decodeFavouriteCottageSearchCriteria(apiCriteria *userresidenceapiv1.FavouriteCottageSearchCriteria) domain.FavouriteCottagesSearchCriteria {

	if apiCriteria == nil {
		// default criteria.
		return domain.FavouriteCottagesSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.FavouriteCottagesSearchCriteria{
		Page:  helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		Sorts: helpers.DecodeSorts(apiCriteria.Sorts),
	}
}
