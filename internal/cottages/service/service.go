package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.CottageRepository
}

// NewService - creates a new service with necessary dependencies
func NewService(
	repository domain.CottageRepository,
	logger log.Logger,
) domain.CottageService {
	var service domain.CottageService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(repository domain.CottageRepository) domain.CottageService {
	return &service{
		repository: repository,
	}
}

func (s *service) UpdateHousePlan(
	ctx context.Context,
	housePlanID int64,
	cottagePlan *domain.HousePlan,
	callerID domain.CallerID,
) (*domain.HousePlan, error) {
	if housePlanID <= 0 {
		errors.NewErrInvalidArgument("cottage id required")
	}
	return s.repository.UpdateHousePlan(ctx, housePlanID, cottagePlan)
}

func (s *service) DeleteHousePlan(
	ctx context.Context,
	housePlanID int64,
	callerID domain.CallerID,
) error {
	if housePlanID <= 0 {
		return errors.NewErrInvalidArgument("cottage id required")
	}
	return s.repository.DeleteHousePlan(ctx, housePlanID)
}

func (s *service) CreateHousePlan(
	ctx context.Context,
	cottage *domain.HousePlan,
	callerID domain.CallerID,
) (*domain.HousePlan, error) {
	return s.repository.CreateHousePlan(ctx, cottage)
}

func (s *service) UpdateCottage(
	ctx context.Context,
	cottageID int64,
	cottage *domain.Cottage,
	allerID domain.CallerID,
) (*domain.Cottage, error) {

	if s.validateCottageInternal(cottage) != nil {
		return nil, errors.NewErrInvalidArgument("bad cottage input")
	}

	if cottageID <= 0 {
		return nil, errors.NewErrInvalidArgument("bad cottageID argument")
	}
	return s.repository.Update(ctx, cottageID, cottage)
}

func (s *service) DeleteCottage(
	ctx context.Context,
	cottageID int64,
	callerID domain.CallerID,
) error {
	if cottageID <= 0 {
		return errors.NewErrInvalidArgument("bad cottageID argument")
	}
	return s.repository.Delete(ctx, cottageID)
}

func (s *service) CreateCottage(
	ctx context.Context,
	cottage *domain.Cottage,
	callerID domain.CallerID,
) (*domain.Cottage, error) {
	if s.validateCottageInternal(cottage) != nil {
		return nil, errors.NewErrInvalidArgument("bad cottage input")
	}
	return s.repository.Create(ctx, cottage)
}

func (s *service) GetCottage(
	ctx context.Context,
	id int64,
	callerID domain.CallerID,
) (*domain.Cottage, error) {
	if id <= 0 {
		errors.NewErrInvalidArgument("invalid cottage id")
	}
	return s.repository.Get(ctx, id)
}

func (s *service) ListCottage(
	ctx context.Context,
	criteria domain.CottageSearchCriteria,
	callerID domain.CallerID,
) ([]*domain.Cottage, domain.Total, error) {
	return s.repository.List(ctx, criteria)
}

func (s *service) ListPopularCottages(
	ctx context.Context,
	criteria domain.CottageSearchCriteria,
	callerID domain.CallerID,
) ([]*domain.Cottage, domain.Total, error) {
	return s.repository.ListPopularCottages(ctx, criteria)

}

func (s *service) ListCottagesByIDs(
	ctx context.Context,
	ids []int64,
	callerID domain.CallerID,
) ([]*domain.Cottage, error) {
	return s.repository.ListCottagesByIDs(ctx, ids)
}

func (s *service) IsFavouriteCottage(
	ctx context.Context,
	cottageID int64,
	userID int64,
	callerID domain.CallerID,
) (bool, error) {
	return s.repository.IsFavouriteCottage(ctx, cottageID, domain.UserID(userID))
}

func (s *service) IsCottageExist(
	ctx context.Context,
	cottageID int64,
	callerID domain.CallerID,
) (bool, error) {
	res, err := s.repository.Get(ctx, cottageID)
	return res != nil, err
}

func (s *service) GetConsultationEmailByCottageID(
	ctx context.Context,
	cottageID int64,
	callerID domain.CallerID,
) (string, error) {
	return s.repository.GetConsultationEmailByCottageID(ctx, cottageID)
}

func (s *service) validateCottageInternal(cottage *domain.Cottage) error {

	if len(cottage.Title) == 0 {
		return errors.NewErrInvalidArgument("title required")
	}

	if cottage.StatusID <= 0 {
		return errors.NewErrInvalidArgument("status required")
	}

	if cottage.SaleStatusID <= 0 {
		return errors.NewErrInvalidArgument("sale_status required")
	}

	if cottage.CityID <= 0 {
		return errors.NewErrInvalidArgument("city required")
	}

	if cottage.BuildingArea < 0 {
		return errors.NewErrInvalidArgument("invalid building area")
	}

	if cottage.AreaMin < 0 {
		return errors.NewErrInvalidArgument("invalid min flat square")
	}

	if cottage.AreaMax < 0 {
		return errors.NewErrInvalidArgument("invalid max flat square")
	}

	if cottage.PricePerSquareMax < 0 {
		return errors.NewErrInvalidArgument("invalid max price")
	}

	if cottage.CeilingHeightMax < 0 {
		return errors.NewErrInvalidArgument("invalid ceiling height")
	}

	return nil
}
