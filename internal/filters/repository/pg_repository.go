package repository

import (
	"context"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository - creates a new service
func NewRepository(db *gorm.DB) domain.FiltersRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateCity(
	ctx context.Context,
	city *domain.City,
) (domain.CityID, error) {
	err := r.db.
		WithContext(ctx).
		Table(domain.CitiesTableName).
		Create(&city).Error
	if err != nil {
		return 0, err
	}

	return city.ID, nil
}

func (r *repository) UpdateCity(
	ctx context.Context,
	cityID domain.CityID,
	city *domain.City,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.CitiesTableName).
		Where("id = ?", cityID).
		Updates(&city).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteCity(
	ctx context.Context,
	cityID domain.CityID,
) error {
	city := new(domain.City)
	err := r.db.
		WithContext(ctx).
		Table(domain.CitiesTableName).
		Where("id = ?", cityID).
		Delete(&city).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListCities(
	ctx context.Context,
	criteria domain.CitySearchCriteria,
) ([]*domain.City, domain.Total, error) {
	var (
		db         = r.db
		list       []*domain.City
		totalCount int64
	)

	// Get total count
	err := db.
		Table(domain.CitiesTableName).
		Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.
		WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.CitiesTableName).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, domain.Total(totalCount), nil
}

func (r *repository) ListDistricts(
	ctx context.Context,
	criteria domain.DistrictSearchCriteria,
) ([]*domain.District, domain.Total, error) {
	var (
		db         = r.db
		list       []*domain.District
		totalCount int64
	)

	if criteria.CityID > 0 {
		db = db.Where("city_id = ?", criteria.CityID)
	}

	err := db.
		Table(domain.DistrictsTableName).
		Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.
		WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.DistrictsTableName).
		Preload(domain.CityAssociation).
		Order("city_id").
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, domain.Total(totalCount), nil
}

func (r *repository) ListFilters(
	ctx context.Context,
	key string,
) ([]*domain.Filter, error) {
	var filters []*domain.Filter
	err := r.db.
		WithContext(ctx).
		Table(key).
		Find(&filters).Error
	if err != nil {
		return nil, err
	}

	return filters, nil
}

func (r *repository) ListFilterKeys(
	ctx context.Context,
) ([]*domain.FilterKey, error) {
	var keys []*domain.FilterKey
	err := r.db.
		WithContext(ctx).
		Table(domain.FilterKeysTableName).
		Find(&keys).Error
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (r *repository) ListBuilders(
	ctx context.Context,
) ([]*domain.FilterBuilder, error) {
	var builders []*domain.FilterBuilder
	err := r.db.
		WithContext(ctx).
		Select("id, first_name, last_name").
		Where("role_id = ? and deleted_at = 0", domain.RoleBuilder).
		Table(domain.UsersTableName).
		Order("first_name ASC").
		Find(&builders).Error
	if err != nil {
		return nil, err
	}

	return builders, nil
}

func (r *repository) CreateFilter(
	ctx context.Context,
	key string,
	filter *domain.Filter,
) (*domain.Filter, error) {
	err := r.db.
		WithContext(ctx).
		Table(key).
		Create(&filter).Error
	if err != nil {
		return nil, err
	}

	return filter, nil
}

func (r *repository) DeleteFilter(
	ctx context.Context,
	id int64,
	key string,
) error {
	filter := new(domain.Filter)
	err := r.db.
		WithContext(ctx).
		Table(key).
		Where("id = ?", id).
		Delete(&filter).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateDistrict(
	ctx context.Context,
	district *domain.District,
) (domain.DistrictID, error) {
	err := r.db.
		WithContext(ctx).
		Table(domain.DistrictsTableName).
		Create(&district).Error
	if err != nil {
		return 0, err
	}

	return district.ID, nil
}

func (r *repository) UpdateDistrict(
	ctx context.Context,
	districtID domain.DistrictID,
	district *domain.District,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.DistrictsTableName).
		Where("id = ?", districtID).
		Updates(&district).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteDistrict(
	ctx context.Context,
	districtID domain.DistrictID,
) error {
	district := new(domain.District)
	err := r.db.
		WithContext(ctx).
		Table(domain.DistrictsTableName).
		Where("id = ?", districtID).
		Delete(&district).Error
	if err != nil {
		return err
	}

	return nil
}
