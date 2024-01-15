package transport

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

func decodeCityV1(w *apiv1.City) *domain.City {
	if w == nil {
		return nil
	}

	return &domain.City{
		ID:   domain.CityID(w.Id),
		Name: w.Name,
	}
}

func encodeCities(r []*domain.City) []*apiv1.City {
	if r == nil {
		return nil
	}

	cities := make([]*apiv1.City, len(r))
	for i, opt := range r {
		cities[i] = encodeCity(opt)
	}

	return cities
}

func encodeCity(r *domain.City) *apiv1.City {
	if r == nil {
		return nil
	}

	return &apiv1.City{
		Id:   int64(r.ID),
		Name: r.Name,
	}
}

func decodeDistrictV1(w *apiv1.DistrictWrite) *domain.District {
	if w == nil {
		return nil
	}

	return &domain.District{
		CityID: w.CityId,
		Name:   w.Name,
	}
}

func encodeDistricts(r []*domain.District) []*apiv1.DistrictRead {
	if r == nil {
		return nil
	}

	districts := make([]*apiv1.DistrictRead, len(r))
	for i, opt := range r {
		districts[i] = encodeDistrict(opt)
	}

	return districts
}

func encodeDistrict(r *domain.District) *apiv1.DistrictRead {
	if r == nil {
		return nil
	}

	return &apiv1.DistrictRead{
		Id:     int64(r.ID),
		Name:   r.Name,
		CityId: r.CityID,
		City:   encodeCity(r.City),
	}
}

func encodeFilters(filter map[string][]*domain.Filter) map[string]*apiv1.Filters {
	var apiFilter = make(map[string]*apiv1.Filters)

	for key, value := range filter {
		apiFilter[key] = encodeFilter(value)
	}

	return apiFilter
}

func encodeFilter(r []*domain.Filter) *apiv1.Filters {
	if r == nil {
		return nil
	}

	filters := make([]*apiv1.Filter, len(r))
	for i, opt := range r {
		filters[i] = &apiv1.Filter{
			Id:   opt.ID,
			Name: opt.Name,
		}
	}

	return &apiv1.Filters{
		Filters: filters,
	}
}

func encodeFiltersV2(filters map[string][]*domain.Filter) []*apiv1.FiltersV2 {

	apiFilters := make([]*apiv1.FiltersV2, 0)

	for key, value := range filters {
		filter := &apiv1.FiltersV2{
			Key:     key,
			Filters: encodeFilterV2(value),
		}

		apiFilters = append(apiFilters, filter)
	}

	return apiFilters
}

func encodeFilterV2(r []*domain.Filter) []*apiv1.Filter {
	if r == nil {
		return nil
	}

	filters := make([]*apiv1.Filter, len(r))
	for i, opt := range r {
		filters[i] = &apiv1.Filter{
			Id:   opt.ID,
			Name: opt.Name,
		}
	}

	return filters
}

func encodeFilterBuilders(r []*domain.FilterBuilder) []*apiv1.FilterBuilder {
	if r == nil {
		return nil
	}

	filters := make([]*apiv1.FilterBuilder, len(r))
	for i, opt := range r {
		filters[i] = &apiv1.FilterBuilder{
			Id:       opt.ID,
			FullName: opt.FirstName + opt.LastName,
		}
	}

	return filters
}

func decodeFilterV1(w *apiv1.Filter) *domain.Filter {
	if w == nil {
		return nil
	}

	return &domain.Filter{
		ID:   w.Id,
		Name: w.Name,
	}
}

func encodeFilterV1(r *domain.Filter) *apiv1.Filter {
	if r == nil {
		return nil
	}

	return &apiv1.Filter{
		Id:   r.ID,
		Name: r.Name,
	}
}

func decodeCitySearchCriteria(apiCriteria *apiv1.CitySearchCriteria) domain.CitySearchCriteria {

	if apiCriteria == nil {
		// default criteria.
		return domain.CitySearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.CitySearchCriteria{
		Page: helpers.DecodePageRequestV1(apiCriteria.PageRequest),
	}
}

func decodeDistrictSearchCriteria(apiCriteria *apiv1.DistrictSearchCriteria) domain.DistrictSearchCriteria {

	if apiCriteria == nil {
		// default criteria.
		return domain.DistrictSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.DistrictSearchCriteria{
		Page:   helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		CityID: apiCriteria.GetCityId(),
	}
}
