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

// NewCottageRepository creates a new repository.
func NewCottageRepository(db *gorm.DB) domain.CottageRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(
	ctx context.Context,
	cottage *domain.Cottage,
) (*domain.Cottage, error) {
	err := r.db.
		WithContext(ctx).
		Table(domain.CottageTableName).
		Omit("IsFavourite").
		Create(cottage).
		Error
	if err != nil {
		return nil, err
	}

	return cottage, nil
}

func (r *repository) List(
	ctx context.Context,
	criteria domain.CottageSearchCriteria,
) ([]*domain.Cottage, domain.Total, error) {
	var (
		db         = r.db
		cottages   = make([]*domain.Cottage, 0)
		totalCount int64
	)

	if len(criteria.Title) > 0 {
		db = db.Where(`UPPER(title) LIKE UPPER(?)`, "%"+criteria.Title+"%")
	}
	if len(criteria.BuilderIDs) > 0 {
		db = db.Where("user_id IN ?", criteria.BuilderIDs)
	}
	if criteria.CityID > 0 {
		db = db.Where("city_id = ?", criteria.CityID)
	}
	if criteria.DistrictID > 0 {
		db = db.Where("district_id = ?", criteria.DistrictID)
	}
	if criteria.HouseType > 0 {
		db = db.Where("housing_class_id = ?", criteria.HouseType)
	}
	if criteria.StatusID > 0 {
		db = db.Where("status_id = ?", criteria.StatusID)
	}
	if criteria.RoomsMin > 0 {
		db = db.Where("rooms_min <= ?", criteria.RoomsMin)
	}
	if criteria.RoomsMax > 0 {
		db = db.Where("rooms_min <= ?", criteria.RoomsMax)
	}
	if criteria.CeilingHeightMin > 0 {
		db = db.Where("ceiling_height_max >= ?", criteria.CeilingHeightMin)
	}
	if criteria.CeilingHeightMax > 0 {
		db = db.Where("ceiling_height_min <= ?", criteria.CeilingHeightMax)
	}

	if len(criteria.InteriorDecoration) > 0 {
		db = db.Joins("JOIN cottages_interior_decorations AS idt ON idt.cottage_id = cottages.id AND idt.interior_decoration_id IN ?", criteria.InteriorDecoration)
	}
	if len(criteria.HeatingTypes) > 0 {
		db = db.Joins("JOIN cottages_heating_types AS ht ON ht.cottage_id = cottages.id AND ht.heating_type_id IN ?", criteria.HeatingTypes)
	}
	if len(criteria.ElevatorTypes) > 0 {
		db = db.Joins("JOIN cottages_elevator_types AS et ON et.cottage_id = cottages.id AND et.elevator_type_id IN ?", criteria.ElevatorTypes)
	}
	if len(criteria.PurchaseMethods) > 0 {
		db = db.Joins("JOIN cottages_purchase_methods AS pm ON pm.cottage_id = cottages.id AND pm.purchase_method_id IN ?", criteria.PurchaseMethods)
	}
	if len(criteria.WallTypes) > 0 {
		db = db.Joins("JOIN cottages_wall_types AS wt ON wt.cottage_id = cottages.id AND wt.wall_type_id IN ?", criteria.WallTypes)
	}
	if len(criteria.ParkingTypes) > 0 {
		db = db.Joins("JOIN cottages_parking_types AS pt ON pt.cottage_id = cottages.id AND pt.parking_type_id IN ?", criteria.ParkingTypes)
	}

	if len(criteria.WarmingTypes) > 0 {
		db = db.Joins("JOIN cottages_warming_types AS wrm ON wrm.cottage_id = cottages.id AND wrm.warming_type_id IN ?", criteria.WarmingTypes)
	}
	if len(criteria.WindowTypes) > 0 {
		db = db.Joins("JOIN cottages_windows AS win ON win.cottage_id = cottages.id AND win.window_type_id IN ?", criteria.WindowTypes)
	}

	if criteria.AreaMin > 0 && criteria.AreaMax > 0 {
		db = db.Where("area_max >= ? AND area_min <= ?", criteria.AreaMin, criteria.AreaMax)
	} else if criteria.AreaMin > 0 {
		db = db.Where("area_max >= ?", criteria.AreaMin)
	} else if criteria.AreaMax > 0 {
		db = db.Where("area_min <= ?", criteria.AreaMax)
	}
	if criteria.PriceMin > 0 {
		db = db.Where("price_per_square_max >= ?", criteria.PriceMin)
	}
	if criteria.PriceMax > 0 {
		db = db.Where("price_per_square_min <= ?", criteria.PriceMax)
	}
	if criteria.FloorsMin > 0 {
		db = db.Where("floors_count >= ?", criteria.FloorsMin)
	}
	if criteria.FloorsMax > 0 {
		db = db.Where("floors_count <= ?", criteria.FloorsMax)
	}
	db = db.Select("*, EXISTS(SELECT * FROM cottages_bookmarks AS rb WHERE rb.cottage_id = cottages.id and rb.user_id = ?) as is_favourite", criteria.UserID)

	order := helpers.PrepareOrder(criteria.Sorts)
	if len(order) == 0 {
		order = `RANDOM() < 0.7`
	}
	db = db.Order(order)

	totalCount, err := r.assertGetTotalCottages(db)
	if err != nil {
		return nil, 0, err
	}

	err = db.
		WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Preload(domain.CityAssociation).
		Preload(domain.DistrictAssociation).
		Preload(domain.StatusAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.WallTypesAssociation).
		Preload(domain.WindowTypesAssociation).
		Preload(domain.WarmingTypesAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.InteriorDecorationsAssociation).
		Preload(domain.HeatingTypesAssociation).
		Preload(domain.ElevatorTypesAssociation).
		Preload(domain.PurchaseMethodAssociation).
		Preload(domain.ParkingTypesAssociation).
		Preload(domain.HousePlansAssociation).
		Preload(domain.UserAssociation, func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, email, consultation_phone_number, image")
		}).
		Table(domain.CottageTableName).
		Table(domain.CottageTableName).
		Find(&cottages).Error
	if err != nil {
		return nil, 0, err
	}

	return cottages, domain.Total(totalCount), nil
}

func (r *repository) ListCottagesByIDs(
	ctx context.Context,
	cottageIDs []int64,
) ([]*domain.Cottage, error) {
	var (
		db       = r.db
		cottages = make([]*domain.Cottage, 0)
	)

	err := db.
		WithContext(ctx).
		Table(domain.CottageTableName).
		Preload(domain.CityAssociation).
		Preload(domain.DistrictAssociation).
		Preload(domain.StatusAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.WallTypesAssociation).
		Preload(domain.WindowTypesAssociation).
		Preload(domain.WarmingTypesAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.InteriorDecorationsAssociation).
		Preload(domain.HeatingTypesAssociation).
		Preload(domain.ElevatorTypesAssociation).
		Preload(domain.PurchaseMethodAssociation).
		Preload(domain.ParkingTypesAssociation).
		Preload(domain.HousePlansAssociation).
		Preload(domain.UserAssociation, func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, email, consultation_phone_number, image")
		}).
		Table(domain.CottageTableName).
		Where("id IN ?", cottageIDs).
		Find(&cottages).Error

	if err != nil {
		return nil, err
	}

	return cottages, err
}

func (r *repository) ListPopularCottages(
	ctx context.Context,
	criteria domain.CottageSearchCriteria,
) ([]*domain.Cottage, domain.Total, error) {
	var (
		db         = r.db
		cottages   = make([]*domain.Cottage, 0)
		totalCount int64
	)

	db = db.Where("id IN (?)", db.Table(domain.LeadCottageTableName).Select("cottage_id").Where("status_id = 1 AND deleted_at=0"))

	if len(criteria.Title) > 0 {
		db = db.Where(`UPPER(title) LIKE UPPER(?)`, "%"+criteria.Title+"%")
	}
	if len(criteria.BuilderIDs) > 0 {
		db = db.Where("user_id IN ?", criteria.BuilderIDs)
	}
	if criteria.CityID > 0 {
		db = db.Where("city_id = ?", criteria.CityID)
	}
	if criteria.DistrictID > 0 {
		db = db.Where("district_id = ?", criteria.DistrictID)
	}
	if criteria.HouseType > 0 {
		db = db.Where("housing_class_id = ?", criteria.HouseType)
	}
	if criteria.StatusID > 0 {
		db = db.Where("status_id = ?", criteria.StatusID)
	}
	if criteria.RoomsMin > 0 {
		db = db.Where("rooms_min <= ?", criteria.RoomsMin)
	}
	if criteria.RoomsMax > 0 {
		db = db.Where("rooms_min <= ?", criteria.RoomsMax)
	}
	if criteria.CeilingHeightMin > 0 {
		db = db.Where("ceiling_height_max >= ?", criteria.CeilingHeightMin)
	}
	if criteria.CeilingHeightMax > 0 {
		db = db.Where("ceiling_height_min <= ?", criteria.CeilingHeightMax)
	}

	if len(criteria.InteriorDecoration) > 0 {
		db = db.Joins("JOIN cottages_interior_decorations AS idt ON idt.cottage_id = cottages.id AND idt.interior_decoration_id IN ?", criteria.InteriorDecoration)
	}
	if len(criteria.HeatingTypes) > 0 {
		db = db.Joins("JOIN cottages_heating_types AS ht ON ht.cottage_id = cottages.id AND ht.heating_type_id IN ?", criteria.HeatingTypes)
	}
	if len(criteria.ElevatorTypes) > 0 {
		db = db.Joins("JOIN cottages_elevator_types AS et ON et.cottage_id = cottages.id AND et.elevator_type_id IN ?", criteria.ElevatorTypes)
	}
	if len(criteria.PurchaseMethods) > 0 {
		db = db.Joins("JOIN cottages_purchase_methods AS pm ON pm.cottage_id = cottages.id AND pm.purchase_method_id IN ?", criteria.PurchaseMethods)
	}
	if len(criteria.WallTypes) > 0 {
		db = db.Joins("JOIN cottages_wall_types AS wt ON wt.cottage_id = cottages.id AND wt.wall_type_id IN ?", criteria.WallTypes)
	}
	if len(criteria.ParkingTypes) > 0 {
		db = db.Joins("JOIN cottages_parking_types AS pt ON pt.cottage_id = cottages.id AND pt.parking_type_id IN ?", criteria.ParkingTypes)
	}

	if len(criteria.WarmingTypes) > 0 {
		db = db.Joins("JOIN cottages_warming_types AS wrm ON wrm.cottage_id = cottages.id AND wrm.warming_type_id IN ?", criteria.WarmingTypes)
	}
	if len(criteria.WindowTypes) > 0 {
		db = db.Joins("JOIN cottages_windows AS win ON win.cottage_id = cottages.id AND win.window_type_id IN ?", criteria.WindowTypes)
	}

	if criteria.AreaMin > 0 && criteria.AreaMax > 0 {
		db = db.Where("area_min >= ? AND area_max <= ?", criteria.AreaMin, criteria.AreaMax)
	} else if criteria.AreaMin > 0 {
		db = db.Where("area_max >= ?", criteria.AreaMin)
	} else if criteria.AreaMax > 0 {
		db = db.Where("area_min <= ?", criteria.AreaMax)
	}
	if criteria.PriceMin > 0 {
		db = db.Where("price_per_square_max >= ?", criteria.PriceMin)
	}
	if criteria.PriceMax > 0 {
		db = db.Where("price_per_square_min <= ?", criteria.PriceMax)
	}
	if criteria.FloorsMin > 0 {
		db = db.Where("floors_count >= ?", criteria.FloorsMin)
	}
	if criteria.FloorsMax > 0 {
		db = db.Where("floors_count <= ?", criteria.FloorsMax)
	}
	db = db.Select("*, EXISTS(SELECT * FROM cottages_bookmarks AS rb WHERE rb.cottage_id = cottages.id and rb.user_id = ?) as is_favourite", criteria.UserID)

	order := helpers.PrepareOrder(criteria.Sorts)
	if len(order) == 0 {
		order = `RANDOM() < 0.7`
	}
	db = db.Order(order)

	totalCount, err := r.assertGetTotalCottages(db)
	if err != nil {
		return nil, 0, err
	}

	err = db.
		WithContext(ctx).
		Scopes(helpers.PaginateRandom(criteria.Page)).
		Preload(domain.CityAssociation).
		Preload(domain.DistrictAssociation).
		Preload(domain.StatusAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.WallTypesAssociation).
		Preload(domain.WindowTypesAssociation).
		Preload(domain.WarmingTypesAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.InteriorDecorationsAssociation).
		Preload(domain.HeatingTypesAssociation).
		Preload(domain.ElevatorTypesAssociation).
		Preload(domain.PurchaseMethodAssociation).
		Preload(domain.ParkingTypesAssociation).
		Preload(domain.HousePlansAssociation).
		Preload(domain.UserAssociation, func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, email, consultation_phone_number, image")
		}).
		Table(domain.CottageTableName).
		Find(&cottages).Error
	if err != nil {
		return nil, 0, err
	}

	return cottages, domain.Total(totalCount), nil
}

func (r *repository) Get(
	ctx context.Context,
	cottageID int64,
) (*domain.Cottage, error) {
	var cottage *domain.Cottage
	if err := r.db.
		WithContext(ctx).
		Table(domain.CottageTableName).
		Preload(domain.DistrictAssociation).
		Preload(domain.HousePlansAssociation).
		Preload(domain.CityAssociation).
		Preload(domain.StatusAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.WallTypesAssociation).
		Preload(domain.InteriorDecorationsAssociation).
		Preload(domain.WindowTypesAssociation).
		Preload(domain.WarmingTypesAssociation).
		Preload(domain.HeatingTypesAssociation).
		Preload(domain.ElevatorTypesAssociation).
		Preload(domain.PurchaseMethodAssociation).
		Preload(domain.ParkingTypesAssociation).
		Preload(domain.UserAssociation, func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, email, consultation_phone_number, image")
		}).
		Take(&cottage, cottageID).Error; err != nil {
		return nil, err
	}

	return cottage, nil
}

func (r *repository) Update(
	ctx context.Context,
	cottageID int64,
	cottage *domain.Cottage,
) (*domain.Cottage, error) {
	// Assign cottage id
	cottage.ID = cottageID

	// Clear many2many relations.
	wallTypes := cottage.WallTypes
	r.db.Model(&cottage).Association(domain.WallTypesAssociation).Clear()
	cottage.WallTypes = wallTypes

	warmingTypes := cottage.WarmingTypes
	r.db.Model(&cottage).Association(domain.WarmingTypesAssociation).Clear()
	cottage.WarmingTypes = warmingTypes

	windowTypes := cottage.WindowTypes
	r.db.Model(&cottage).Association(domain.WindowTypesAssociation).Clear()
	cottage.WindowTypes = windowTypes

	interiorDecorationTypes := cottage.InteriorDecorations
	r.db.Model(&cottage).Association(domain.InteriorDecorationsAssociation).Clear()
	cottage.InteriorDecorations = interiorDecorationTypes

	heatingTypes := cottage.HeatingTypes
	r.db.Model(&cottage).Association(domain.HeatingTypesAssociation).Clear()
	cottage.HeatingTypes = heatingTypes

	elevatorTypes := cottage.ElevatorTypes
	r.db.Model(&cottage).Association(domain.ElevatorTypesAssociation).Clear()
	cottage.ElevatorTypes = elevatorTypes

	housePlans := cottage.HousePlans
	r.db.Model(&cottage).Association(domain.HousePlansAssociation).Clear()
	cottage.HousePlans = housePlans

	// update cottage data
	if err := r.db.
		WithContext(ctx).
		Table(domain.CottageTableName).
		Where("id = ?", cottageID).
		Omit("IsFavourite").
		Updates(&cottage).Error; err != nil {
		return nil, err
	}

	return cottage, nil
}

func (r *repository) Delete(
	ctx context.Context,
	cottageID int64,
) error {
	cottage := new(domain.Cottage)
	if err := r.db.
		WithContext(ctx).
		Where("id = ?", cottageID).
		Table(domain.CottageTableName).
		Delete(&cottage).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetConsultationEmailByCottageID(
	ctx context.Context,
	cottageID int64,
) (string, error) {
	var email string
	err := r.db.
		WithContext(ctx).
		Table(domain.CottageTableName).
		Select("email").
		Where("id = ?", cottageID).
		First(&email).Error
	if err != nil {
		return "", err
	}

	return email, nil
}

func (r *repository) assertGetTotalCottages(db *gorm.DB) (int64, error) {
	var totalCount int64

	err := r.db.
		Table(domain.CottageTableName).
		Count(&totalCount).
		Error

	if err != nil {
		return 0, err
	}

	return totalCount, nil
}

func (r *repository) IsFavouriteCottage(
	ctx context.Context,
	cottageID int64,
	userID domain.UserID,
) (bool, error) {
	var isFavourite bool

	err := r.db.
		WithContext(ctx).
		Table(domain.FavouriteCottageTable).
		Select("count(*) > 0").
		Where("user_id = ? AND cottage_id = ?", userID, cottageID).
		Find(&isFavourite).Error
	if err != nil {
		return false, err
	}

	return isFavourite, nil
}

func (r *repository) CreateHousePlan(
	ctx context.Context,
	cottagePlan *domain.HousePlan,
) (*domain.HousePlan, error) {

	err := r.db.
		WithContext(ctx).
		Table(domain.HousePlansTable).
		Create(&cottagePlan).Error
	if err != nil {
		return nil, err
	}

	return cottagePlan, nil
}

func (r *repository) UpdateHousePlan(
	ctx context.Context,
	cottagePlanID int64,
	cottagePlan *domain.HousePlan,
) (*domain.HousePlan, error) {

	if err := r.db.
		WithContext(ctx).
		Table(domain.HousePlansTable).
		Where("id = ?", cottagePlanID).
		Omit("IsFavourite").
		Updates(cottagePlan).Error; err != nil {
		return nil, err
	}

	return cottagePlan, nil
}

func (r *repository) DeleteHousePlan(
	ctx context.Context,
	cottagePlanID int64,
) error {
	cottagePlan := new(domain.HousePlan)
	if err := r.db.
		WithContext(ctx).
		Table(domain.HousePlansTable).
		Where("id = ?", cottagePlanID).
		Delete(&cottagePlan).Error; err != nil {
		return err
	}

	return nil
}
