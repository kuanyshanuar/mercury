package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	repository      domain.ResidencesRepository
	redisRepository domain.ResidenceRedisRepository
}

// NewResidenceService creates a new service
func NewResidenceService(
	repository domain.ResidencesRepository,
	redisRepository domain.ResidenceRedisRepository,
	logger log.Logger,
) domain.ResidencesService {
	var service domain.ResidencesService
	{
		service = newBasicService(repository, redisRepository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

// Returns a naive, stateless implementation of service.
func newBasicService(
	repository domain.ResidencesRepository,
	redisRepository domain.ResidenceRedisRepository,
) domain.ResidencesService {
	return &service{
		repository:      repository,
		redisRepository: redisRepository,
	}
}

func (s *service) CreateResidence(
	ctx context.Context,
	residence *domain.Residence,
	callerID domain.CallerID,
) (*domain.Residence, error) {

	// Validate residence
	//
	if err := s.validateResidenceInternal(residence); err != nil {
		return nil, err
	}

	return s.repository.Create(ctx, residence)
}

func (s *service) ListResidences(
	ctx context.Context,
	criteria domain.ResidenceSearchCriteria,
	callerID domain.CallerID,
) ([]*domain.Residence, domain.Total, error) {
	return s.repository.List(ctx, criteria)
}

func (s *service) GetResidence(
	ctx context.Context,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) (*domain.Residence, error) {

	// Validate input
	//
	if residenceID == 0 {
		return nil, errors.NewErrInvalidArgument("required residence id")
	}

	// Get residence
	//
	residence, err := s.repository.Get(ctx, residenceID)
	if err != nil {
		return nil, err
	}

	// Get is favourite
	//
	if callerID.UserID > 0 {
		isFavourite, err := s.repository.IsFavouriteResidence(ctx, residenceID, domain.UserID(callerID.UserID))
		if err != nil {
			return nil, err
		}
		residence.IsFavourite = isFavourite
	}

	return residence, nil
}

func (s *service) UpdateResidence(
	ctx context.Context,
	residenceID domain.ResidenceID,
	residence *domain.Residence,
	callerID domain.CallerID,
) (*domain.Residence, error) {

	// Validate residence id
	//
	if residenceID == 0 {
		return nil, errors.NewErrInvalidArgument("invlid residence id")
	}

	//Validate residence
	//
	if err := s.validateResidenceInternal(residence); err != nil {
		return nil, err
	}

	return s.repository.Update(ctx, residenceID, residence)
}

func (s *service) DeleteResidence(
	ctx context.Context,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) error {

	// Validate residence id
	//
	if residenceID == 0 {
		return errors.NewErrInvalidArgument("required residence id")
	}

	return s.repository.Delete(ctx, residenceID)
}

func (s *service) validateResidenceInternal(residence *domain.Residence) error {

	if len(residence.Title) == 0 {
		return errors.NewErrInvalidArgument("title required")
	}

	if residence.StatusID <= 0 {
		return errors.NewErrInvalidArgument("status required")
	}

	if residence.SaleStatusID <= 0 {
		return errors.NewErrInvalidArgument("sale_status required")
	}

	if residence.CityID <= 0 {
		return errors.NewErrInvalidArgument("city required")
	}

	if residence.HousingClassID <= 0 {
		return errors.NewErrInvalidArgument("housing class required")
	}

	if residence.ConstructionTypeID <= 0 {
		return errors.NewErrInvalidArgument("construction type id required")
	}

	if residence.AreaMin < 0 {
		return errors.NewErrInvalidArgument("invalid min area")
	}

	if residence.AreaMax < 0 {
		return errors.NewErrInvalidArgument("invalid max area")
	}

	if residence.PriceMin < 0 {
		return errors.NewErrInvalidArgument("invalid min price")
	}

	if residence.PriceMax < 0 {
		return errors.NewErrInvalidArgument("invalid max price")
	}

	if residence.CeilingHeight <= 0 {
		return errors.NewErrInvalidArgument("invalid ceiling height")
	}

	if residence.DeadlineYear <= 0 {
		return errors.NewErrInvalidArgument("invalid deadline year")
	}

	if residence.DeadlineQuarter <= 0 || residence.DeadlineQuarter > 4 {
		return errors.NewErrInvalidArgument("invalid deadline quarter")
	}

	return nil
}

func (s *service) IsResidenceExist(
	ctx context.Context,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) (bool, error) {

	// Validate residence id.
	//
	if residenceID <= 0 {
		return false, errors.NewErrInvalidArgument("invalid residence id")
	}

	// Check if residence exists
	//
	residence, err := s.repository.Get(ctx, residenceID)
	if err != nil {
		return false, err
	}

	if err != nil {
		return false, errors.NewErrNotFound(
			fmt.Sprintf("residence not found by id: %v, details: %v", residenceID, err),
		)
	}
	if residence == nil {
		return false, errors.NewErrNotFound(
			fmt.Sprintf("residence not found by id: %v", residenceID),
		)
	}

	return true, nil
}

func (s *service) ListResidencesByIDs(
	ctx context.Context,
	residencesIDs []domain.ResidenceID,
	callerID domain.CallerID,
) ([]*domain.Residence, error) {
	return s.repository.ListResidencesByIDs(ctx, residencesIDs)
}

func (s *service) ListPopularResidences(
	ctx context.Context,
	criteria domain.ResidenceSearchCriteria,
	callerID domain.CallerID,
) ([]*domain.Residence, domain.Total, error) {
	return s.repository.ListPopularResidences(ctx, criteria)
}

func (s *service) GetConsultationEmailByResidenceID(
	ctx context.Context,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) (string, error) {

	// Validate inputs
	//
	if residenceID <= 0 {
		return "", errors.NewErrInvalidArgument("residence id required")
	}

	return s.repository.GetConsultationEmailByResidenceID(ctx, residenceID)
}

func (s *service) CreateFlatPlan(
	ctx context.Context,
	flatPlan *domain.FlatPlan,
	callerID domain.CallerID,
) (*domain.FlatPlan, error) {

	// Validate inputs
	//
	err := s.validateFlatPlanInternal(flatPlan)
	if err != nil {
		return nil, err
	}

	return s.repository.CreateFlatPlan(ctx, flatPlan)
}

func (s *service) UpdateFlatPlan(
	ctx context.Context,
	flatPlanID domain.FlatPlanID,
	flatPlan *domain.FlatPlan,
	callerID domain.CallerID,
) (*domain.FlatPlan, error) {

	// Validate inputs
	//
	if flatPlanID <= 0 {
		return nil, errors.NewErrInvalidArgument("flat plan id required")
	}
	err := s.validateFlatPlanInternal(flatPlan)
	if err != nil {
		return nil, err
	}

	return s.repository.UpdateFlatPlan(ctx, flatPlanID, flatPlan)
}

func (s *service) DeleteFlatPlan(
	ctx context.Context,
	flatPlanID domain.FlatPlanID,
	callerID domain.CallerID,
) error {
	// Validate inputs
	//
	if flatPlanID <= 0 {
		return errors.NewErrInvalidArgument("flat plan id required")
	}

	return s.repository.DeleteFlatPlan(ctx, flatPlanID)
}

func (s *service) GetSimilarResidences(
	ctx context.Context,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) ([]*domain.Residence, error) {
	residence, err := s.repository.Get(ctx, residenceID)
	if err != nil {
		return nil, err
	}
	criteria := domain.ResidenceSearchCriteria{
		CityID:     residence.CityID,
		DistrictID: int64(residence.DistrictID),
		PriceMin:   max(residence.PriceMin-5000000, 10000000),
		PriceMax:   residence.PriceMax + 5000000,
		AreaMax:    float64(residence.AreaMax + 50),
		AreaMin:    max(float64(residence.AreaMin)-25, 1),
		RoomsMin:   max(residence.RoomsMin-1, 1),
		RoomsMax:   residence.RoomsMax + 2,
		StatusID:   residence.StatusID,
		Sorts: []domain.Sort{
			{FieldName: "price_min"},
		},
	}
	residences, total, err := s.repository.List(ctx, criteria)
	if err != nil {
		return nil, err
	}
	if total < 6 {
		criteria = domain.ResidenceSearchCriteria{
			CityID:     residence.CityID,
			BuilderIDs: []int64{residence.UserID},
			StatusID:   residence.StatusID,
		}
		residences, total, err = s.repository.List(ctx, criteria)
		if err != nil {
			return nil, err
		}

		if total < 6 {
			criteria = domain.ResidenceSearchCriteria{
				CityID:   residence.CityID,
				StatusID: residence.StatusID,
			}
			residences, _, _ = s.repository.List(ctx, criteria)

		}
	}
	return residences, nil

}
func max[V int64 | float64](a1 V, a2 V) V {
	if a1 > a2 {
		return a1
	}
	return a2
}
func (s *service) validateFlatPlanInternal(plan *domain.FlatPlan) error {
	if plan == nil {
		return errors.NewErrInternal("flat plan required")
	}
	if plan.ResidenceID <= 0 {
		return errors.NewErrInternal("residence id required")
	}
	if plan.Area <= 0 {
		return errors.NewErrInternal("area required")
	}
	if plan.NumberOfRooms <= 0 {
		return errors.NewErrInternal("number of rooms required")
	}
	if plan.Area <= 0 {
		return errors.NewErrInternal("area required")
	}
	if plan.Price <= 0 {
		return errors.NewErrInternal("price required")
	}

	return nil
}
