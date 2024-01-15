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

// NewResidenceRepository creates a new repository.
func NewResidenceRepository(db *gorm.DB) domain.ResidencesRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(
	ctx context.Context,
	residence *domain.Residence,
) (*domain.Residence, error) {
	err := r.db.
		WithContext(ctx).
		Table(domain.ResidencesTableName).
		Omit("IsFavourite").
		Create(residence).Error
	if err != nil {
		return nil, err
	}

	return residence, nil
}

func (r *repository) List(
	ctx context.Context,
	criteria domain.ResidenceSearchCriteria,
) ([]*domain.Residence, domain.Total, error) {
	var (
		db         = r.db
		residences = make([]*domain.Residence, 0)
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
	if criteria.HousingClassID > 0 {
		db = db.Where("housing_class_id = ?", criteria.HousingClassID)
	}
	if criteria.StatusID > 0 {
		db = db.Where("status_id = ?", criteria.StatusID)
	}
	if criteria.RoomsMin > 0 {
		db = db.Where("rooms_min <= ?", criteria.RoomsMin)
	}
	if criteria.RoomsMax > 0 {
		db = db.Where("rooms_max >= ?", criteria.RoomsMax)
	}
	if criteria.CeilingHeightMin > 0 {
		db = db.Where("ceiling_height >= ?", criteria.CeilingHeightMin)
	}
	if criteria.CeilingHeightMax > 0 {
		db = db.Where("ceiling_height <= ?", criteria.CeilingHeightMax)
	}
	if len(criteria.ConstructionTypesIDs) > 0 {
		db = db.Where("construction_type_id IN ?", criteria.ConstructionTypesIDs)
	}
	if len(criteria.ParkingTypesIds) > 0 {
		db = db.Joins("JOIN residences_parking_types AS pt ON pt.residence_id = residences.id AND pt.parking_type_id IN ?", criteria.ParkingTypesIds)
	}
	if len(criteria.InteriorDecorationIDs) > 0 {
		db = db.Joins("JOIN residences_interior_decorations AS idt ON idt.residence_id = residences.id AND idt.interior_decoration_id IN ?", criteria.InteriorDecorationIDs)
	}
	if len(criteria.HeatingTypesIDs) > 0 {
		db = db.Joins("JOIN residences_heating_types AS ht ON ht.residence_id = residences.id AND ht.heating_type_id IN ?", criteria.HeatingTypesIDs)
	}
	if len(criteria.ElevatorTypesIDs) > 0 {
		db = db.Joins("JOIN residences_elevator_types AS et ON et.residence_id = residences.id AND et.elevator_type_id IN ?", criteria.ElevatorTypesIDs)
	}
	if len(criteria.PurchaseMethodsIDs) > 0 {
		db = db.Joins("JOIN residences_purchase_methods AS pm ON pm.residence_id = residences.id AND pm.purchase_method_id IN ?", criteria.PurchaseMethodsIDs)
	}
	if criteria.HasHGF != nil {
		if *criteria.HasHGF {
			db = db.Where("has_hgf = ?", *criteria.HasHGF)
		}
	}
	if criteria.AreaMin > 0 && criteria.AreaMax > 0 {
		db = db.Where("area_min >= ? AND area_min <= ?", criteria.AreaMin, criteria.AreaMax)
	} else if criteria.AreaMin > 0 {
		db = db.Where("area_min >= ?", criteria.AreaMin)
	} else if criteria.AreaMax > 0 {
		db = db.Where("area_min <= ?", criteria.AreaMax)
	}
	if criteria.PriceMin > 0 {
		db = db.Where("price_per_square_min >= ?", criteria.PriceMin)
	}
	if criteria.PriceMax > 0 {
		db = db.Where("price_per_square_min <= ?", criteria.PriceMax)
	}
	if criteria.FloorsMin > 0 {
		db = db.Where("floors_max >= ?", criteria.FloorsMin)
	}
	if criteria.FloorsMax > 0 {
		db = db.Where("floors_max <= ?", criteria.FloorsMax)
	}
	if criteria.UserID > 0 {
		db = db.Select("*, EXISTS(SELECT * FROM residences_bookmarks rb WHERE rb.residence_id = residences.id and rb.user_id = ?) as is_favourite", criteria.UserID)
	} else {
		db = db.Omit("IsFavourite")
	}

	order := helpers.PrepareOrder(criteria.Sorts)
	if len(order) == 0 {
		order = `"created_at" DESC`
	}
	db = db.Order(order)

	err := db.
		Model(&domain.Residence{}).
		Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.
		WithContext(ctx).
		Debug().
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.ResidencesTableName).
		Preload(domain.CityAssociation).
		Preload(domain.DistrictAssociation).
		Preload(domain.StatusAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.ConstructionTypeAssociation).
		Preload(domain.ParkingTypesAssociation).
		Preload(domain.InteriorDecorationsAssociation).
		Preload(domain.HeatingTypesAssociation).
		Preload(domain.ElevatorTypesAssociation).
		Preload(domain.PurchaseMethodAssociation).
		Preload(domain.FlatPlansAssociation, func(db *gorm.DB) *gorm.DB {
			return db.Order("number_of_rooms ASC")
		}).
		Preload(domain.UserAssociation, func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, email, consultation_phone_number, image")
		}).
		Find(&residences).Error
	if err != nil {
		return nil, 0, err
	}

	return residences, domain.Total(totalCount), nil
}

func (r *repository) ListResidencesByIDs(
	ctx context.Context,
	residencesIDs []domain.ResidenceID,
) ([]*domain.Residence, error) {
	var (
		db         = r.db
		residences = make([]*domain.Residence, 0)
	)

	err := db.
		WithContext(ctx).
		Table(domain.ResidencesTableName).
		Preload(domain.CityAssociation).
		Preload(domain.DistrictAssociation).
		Table(domain.ResidencesTableName).
		Preload(domain.DistrictAssociation).
		Preload(domain.StatusAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.ConstructionTypeAssociation).
		Preload(domain.PurchaseMethodAssociation).
		Preload(domain.FlatPlansAssociation).
		Preload(domain.UserAssociation, func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, email, consultation_phone_number, image")
		}).
		Where("id IN ?", residencesIDs).
		Find(&residences).Error
	if err != nil {
		return nil, err
	}

	return residences, err
}

func (r *repository) ListPopularResidences(
	ctx context.Context,
	criteria domain.ResidenceSearchCriteria,
) ([]*domain.Residence, domain.Total, error) {
	var (
		db         = r.db
		residences = make([]*domain.Residence, 0)
		totalCount int64
	)

	db = db.Where("id IN (?)", db.Table(domain.LeadResidenceTableName).Select("residence_id").Where("status_id = 1 AND deleted_at=0"))

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
	if criteria.HousingClassID > 0 {
		db = db.Where("housing_class_id = ?", criteria.HousingClassID)
	}
	if criteria.StatusID > 0 {
		db = db.Where("status_id = ?", criteria.StatusID)
	}
	if criteria.RoomsMin > 0 {
		db = db.Where("rooms_min <= ?", criteria.RoomsMin)
	}
	if criteria.RoomsMax > 0 {
		db = db.Where("rooms_max >= ?", criteria.RoomsMax)
	}
	if criteria.CeilingHeightMin > 0 {
		db = db.Where("ceiling_height >= ?", criteria.CeilingHeightMin)
	}
	if criteria.CeilingHeightMax > 0 {
		db = db.Where("ceiling_height <= ?", criteria.CeilingHeightMax)
	}
	if len(criteria.ConstructionTypesIDs) > 0 {
		db = db.Where("construction_type_id IN ?", criteria.ConstructionTypesIDs)
	}
	if len(criteria.ParkingTypesIds) > 0 {
		db = db.Joins("JOIN residences_parking_types AS pt ON pt.residence_id = residences.id AND pt.parking_type_id IN ?", criteria.ParkingTypesIds)
	}
	if len(criteria.InteriorDecorationIDs) > 0 {
		db = db.Joins("JOIN residences_interior_decorations AS idt ON idt.residence_id = residences.id AND idt.interior_decoration_id IN ?", criteria.InteriorDecorationIDs)
	}
	if len(criteria.HeatingTypesIDs) > 0 {
		db = db.Joins("JOIN residences_heating_types AS ht ON ht.residence_id = residences.id AND ht.heating_type_id IN ?", criteria.HeatingTypesIDs)
	}
	if len(criteria.ElevatorTypesIDs) > 0 {
		db = db.Joins("JOIN residences_elevator_types AS et ON et.residence_id = residences.id AND et.elevator_type_id IN ?", criteria.ElevatorTypesIDs)
	}
	if len(criteria.PurchaseMethodsIDs) > 0 {
		db = db.Joins("JOIN residences_purchase_methods AS pm ON pm.residence_id = residences.id AND pm.purchase_method_id IN ?", criteria.PurchaseMethodsIDs)
	}
	if criteria.HasHGF != nil {
		if *criteria.HasHGF {
			db = db.Where("has_hgf = ?", *criteria.HasHGF)
		}
	}
	if criteria.AreaMin > 0 && criteria.AreaMax > 0 {
		db = db.Where("area_min >= ? AND area_min <= ?", criteria.AreaMin, criteria.AreaMax)
	} else if criteria.AreaMin > 0 {
		db = db.Where("area_min >= ?", criteria.AreaMin)
	} else if criteria.AreaMax > 0 {
		db = db.Where("area_min <= ?", criteria.AreaMax)
	}
	if criteria.PriceMin > 0 {
		db = db.Where("price_per_square_min >= ?", criteria.PriceMin)
	}
	if criteria.PriceMax > 0 {
		db = db.Where("price_per_square_min <= ?", criteria.PriceMax)
	}
	if criteria.FloorsMin > 0 {
		db = db.Where("floors_max >= ?", criteria.FloorsMin)
	}
	if criteria.FloorsMax > 0 {
		db = db.Where("floors_max <= ?", criteria.FloorsMax)
	}

	db = db.Select("*, EXISTS(SELECT * FROM residences_bookmarks rb WHERE rb.residence_id = residences.id and rb.user_id = ?) as is_favourite", criteria.UserID)

	order := helpers.PrepareOrder(criteria.Sorts)
	if len(order) == 0 {
		order = `RANDOM() < 0.7`
	}
	db = db.Order(order)

	err := db.
		Model(&domain.Residence{}).
		Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.
		WithContext(ctx).
		Scopes(helpers.PaginateRandom(criteria.Page)).
		Table(domain.ResidencesTableName).
		Preload(domain.CityAssociation).
		Preload(domain.DistrictAssociation).
		Preload(domain.StatusAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.ConstructionTypeAssociation).
		Preload(domain.PurchaseMethodAssociation).
		Preload(domain.FlatPlansAssociation).
		Preload(domain.ParkingTypesAssociation).
		Preload(domain.InteriorDecorationsAssociation).
		Preload(domain.HeatingTypesAssociation).
		Preload(domain.ElevatorTypesAssociation).
		Preload(domain.PurchaseMethodAssociation).
		Preload(domain.UserAssociation, func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, email, consultation_phone_number, image")
		}).
		Find(&residences).Error
	if err != nil {
		return nil, 0, err
	}

	return residences, domain.Total(totalCount), nil
}

func (r *repository) Get(
	ctx context.Context,
	residenceID domain.ResidenceID,
) (*domain.Residence, error) {
	var residence *domain.Residence
	if err := r.db.
		WithContext(ctx).
		Table(domain.ResidencesTableName).
		Preload(domain.CityAssociation).
		Preload(domain.DistrictAssociation).
		Preload(domain.StatusAssociation).
		Preload(domain.HousingClassAssociation).
		Preload(domain.ConstructionTypeAssociation).
		Preload(domain.ParkingTypesAssociation).
		Preload(domain.InteriorDecorationsAssociation).
		Preload(domain.HeatingTypesAssociation).
		Preload(domain.ElevatorTypesAssociation).
		Preload(domain.FlatPlansAssociation).
		Preload(domain.PurchaseMethodAssociation).
		Preload(domain.UserAssociation, func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, email, consultation_phone_number, image")
		}).
		Take(&residence, residenceID).Error; err != nil {
		return nil, err
	}

	return residence, nil
}

func (r *repository) Update(
	ctx context.Context,
	residenceID domain.ResidenceID,
	residence *domain.Residence,
) (*domain.Residence, error) {
	// Assign residence id
	residence.ID = residenceID

	// Clear many2many relations.
	parkingTypes := residence.ParkingTypes
	r.db.Model(&residence).Association(domain.ParkingTypesAssociation).Clear()
	residence.ParkingTypes = parkingTypes

	interiorDecorationTypes := residence.InteriorDecorations
	r.db.Model(&residence).Association(domain.InteriorDecorationsAssociation).Clear()
	residence.InteriorDecorations = interiorDecorationTypes

	heatingTypes := residence.HeatingTypes
	r.db.Model(&residence).Association(domain.HeatingTypesAssociation).Clear()
	residence.HeatingTypes = heatingTypes

	elevatorTypes := residence.ElevatorTypes
	r.db.Model(&residence).Association(domain.ElevatorTypesAssociation).Clear()
	residence.ElevatorTypes = elevatorTypes

	flatPlans := residence.FlatPlans
	r.db.Model(&residence).Association(domain.FlatPlansAssociation).Clear()
	residence.FlatPlans = flatPlans

	// update residence data
	if err := r.db.
		WithContext(ctx).
		Model(&residence).
		Where("id = ?", residenceID).
		Omit("IsFavourite").
		Save(&residence).Error; err != nil {
		return nil, err
	}

	return residence, nil
}

func (r *repository) Delete(
	ctx context.Context,
	residenceID domain.ResidenceID,
) error {
	residence := new(domain.Residence)
	if err := r.db.
		WithContext(ctx).
		Where("id = ?", residenceID).
		Table(domain.ResidencesTableName).
		Delete(&residence).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetConsultationEmailByResidenceID(
	ctx context.Context,
	residenceID domain.ResidenceID,
) (string, error) {
	var email string
	err := r.db.
		WithContext(ctx).
		Table(domain.ResidencesTableName).
		Select("email").
		Where("id = ?", residenceID).
		First(&email).Error
	if err != nil {
		return "", err
	}

	return email, nil
}

func (r *repository) IsFavouriteResidence(
	ctx context.Context,
	residenceID domain.ResidenceID,
	userID domain.UserID,
) (bool, error) {
	var isFavourite bool

	err := r.db.
		WithContext(ctx).
		Table(domain.FavouriteResidenceTable).
		Select("count(*) > 0").
		Where("user_id = ? AND residence_id = ?", userID, residenceID).
		Find(&isFavourite).Error
	if err != nil {
		return false, err
	}

	return isFavourite, nil
}

func (r *repository) CreateFlatPlan(
	ctx context.Context,
	flatPlan *domain.FlatPlan,
) (*domain.FlatPlan, error) {

	err := r.db.
		WithContext(ctx).
		Table(domain.FlatPlansTableName).
		Create(&flatPlan).Error
	if err != nil {
		return nil, err
	}

	return flatPlan, nil
}

func (r *repository) UpdateFlatPlan(
	ctx context.Context,
	flatPlanID domain.FlatPlanID,
	flatPlan *domain.FlatPlan,
) (*domain.FlatPlan, error) {

	if err := r.db.
		WithContext(ctx).
		Table(domain.FlatPlansTableName).
		Where("id = ?", flatPlanID).
		Omit("IsFavourite").
		Updates(flatPlan).Error; err != nil {
		return nil, err
	}

	return flatPlan, nil
}

func (r *repository) DeleteFlatPlan(
	ctx context.Context,
	flatPlanID domain.FlatPlanID,
) error {
	residence := new(domain.FlatPlan)
	if err := r.db.
		WithContext(ctx).
		Table(domain.FlatPlansTableName).
		Where("id = ?", flatPlanID).
		Delete(&residence).Error; err != nil {
		return err
	}

	return nil
}
