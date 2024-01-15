package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.FiltersRepository
}

// NewService - creates a new service.
func NewService(
	repository domain.FiltersRepository,
	logger log.Logger,
) domain.FiltersService {
	var service domain.FiltersService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	repository domain.FiltersRepository,
) domain.FiltersService {
	return &service{
		repository: repository,
	}
}

func (s *service) ListCities(
	ctx context.Context,
	criteria domain.CitySearchCriteria,
	_ domain.CallerID,
) ([]*domain.City, domain.Total, error) {
	return s.repository.ListCities(ctx, criteria)
}

func (s *service) ListDistricts(
	ctx context.Context,
	criteria domain.DistrictSearchCriteria,
	_ domain.CallerID,
) ([]*domain.District, domain.Total, error) {
	return s.repository.ListDistricts(ctx, criteria)
}

func (s *service) ListFilters(
	ctx context.Context,
	_ domain.CallerID,
) (map[string][]*domain.Filter, error) {
	var (
		filters = make(map[string][]*domain.Filter)
	)

	// List filter keys
	//
	filterKeys, err := s.repository.ListFilterKeys(ctx)
	if err != nil {
		return nil, err
	}

	for _, filterKey := range filterKeys {
		filter, err := s.repository.ListFilters(ctx, filterKey.Key)
		if err != nil {
			return nil, err
		}

		filters[filterKey.Key] = filter
	}

	return filters, nil
}

func (s *service) ListBuilders(
	ctx context.Context,
	_ domain.CallerID,
) ([]*domain.FilterBuilder, error) {
	return s.repository.ListBuilders(ctx)
}

func (s *service) CreateFilter(
	ctx context.Context,
	key string,
	filter *domain.Filter,
	_ domain.CallerID,
) (*domain.Filter, error) {

	// Validate inputs
	//
	if err := s.validateInternalCreateFilter(key, filter); err != nil {
		return nil, err
	}

	return s.repository.CreateFilter(ctx, key, filter)
}

func (s *service) DeleteFilter(
	ctx context.Context,
	id int64,
	key string,
	_ domain.CallerID,
) error {

	// Validate inputs
	//
	if id <= 0 {
		return errors.NewErrInvalidArgument("id required")
	}
	if len(key) == 0 {
		return errors.NewErrInvalidArgument("key required")
	}

	return s.repository.DeleteFilter(ctx, id, key)
}

func (s *service) validateInternalCreateFilter(key string, filter *domain.Filter) error {
	if len(key) == 0 {
		return errors.NewErrInvalidArgument("key required")
	}
	if filter == nil {
		return errors.NewErrInvalidArgument("filter required")
	}
	if len(filter.Name) == 0 {
		return errors.NewErrInvalidArgument("filter name required")
	}

	return nil
}

func (s *service) CreateCity(
	ctx context.Context,
	city *domain.City,
	callerID domain.CallerID,
) (domain.CityID, error) {

	// Validate input
	//
	if city == nil {
		return 0, errors.NewErrInvalidArgument("city required")
	}
	if len(city.Name) == 0 {
		return 0, errors.NewErrInvalidArgument("city name required")
	}

	return s.repository.CreateCity(ctx, city)
}

func (s *service) UpdateCity(
	ctx context.Context,
	cityID domain.CityID,
	city *domain.City,
	callerID domain.CallerID,
) error {
	// Validate input
	//
	if cityID <= 0 {
		return errors.NewErrInvalidArgument("city id required")
	}
	if city == nil {
		return errors.NewErrInvalidArgument("city required")
	}
	if len(city.Name) == 0 {
		return errors.NewErrInvalidArgument("city name required")
	}

	return s.repository.UpdateCity(ctx, cityID, city)
}

func (s *service) DeleteCity(
	ctx context.Context,
	cityID domain.CityID,
	callerID domain.CallerID,
) error {
	// Validate input
	//
	if cityID <= 0 {
		return errors.NewErrInvalidArgument("city id required")
	}

	return s.repository.DeleteCity(ctx, cityID)
}

func (s *service) CreateDistrict(
	ctx context.Context,
	district *domain.District,
	callerID domain.CallerID,
) (domain.DistrictID, error) {
	// Validate input
	//
	if district == nil {
		return 0, errors.NewErrInvalidArgument("district required")
	}
	if district.CityID <= 0 {
		return 0, errors.NewErrInvalidArgument("city id required")
	}
	if len(district.Name) == 0 {
		return 0, errors.NewErrInvalidArgument("district name required")
	}

	return s.repository.CreateDistrict(ctx, district)
}

func (s *service) UpdateDistrict(
	ctx context.Context,
	districtID domain.DistrictID,
	district *domain.District,
	callerID domain.CallerID,
) error {
	// Validate inputs
	//
	if districtID <= 0 {
		return errors.NewErrInvalidArgument("district id required")
	}
	if district == nil {
		return errors.NewErrInvalidArgument("district required")
	}
	if district.CityID <= 0 {
		return errors.NewErrInvalidArgument("city id required")
	}
	if len(district.Name) == 0 {
		return errors.NewErrInvalidArgument("district name required")
	}

	return s.repository.UpdateDistrict(ctx, districtID, district)
}

func (s *service) DeleteDistrict(
	ctx context.Context,
	districtID domain.DistrictID,
	callerID domain.CallerID,
) error {
	// Validate input
	//
	if districtID <= 0 {
		return errors.NewErrInvalidArgument("district id required")
	}

	return s.repository.DeleteDistrict(ctx, districtID)
}
