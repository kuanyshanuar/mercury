package domain

import (
	"context"

	"github.com/lib/pq"
	"gorm.io/plugin/soft_delete"
)

// ResidencesTableName - name of residence table in storage.
const (
	ResidencesTableName = "residences"
)

// ResidenceID - id of the residence
type ResidenceID int64

// Residence - represents residence struct
type Residence struct {
	// ID - id of the residence.
	//
	ID ResidenceID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// UserID - user id.
	//
	UserID int64 `json:"user_id" gorm:"column:user_id"`

	// User - user
	//
	User *User `json:"user" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// StatusID - status of the residence.
	//
	StatusID int64 `json:"status_id" gorm:"column:status_id"`

	// Status - represents status.
	// only read
	Status Status `json:"status" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// CityID - id of the city.
	//
	CityID int64 `json:"city_id" gorm:"column:city_id"`

	// City - city reference
	//
	City *City `json:"city" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// HousingClassID- housing class id.
	//
	HousingClassID HousingClassID `json:"housing_class_id" gorm:"not null;column:housing_class_id"`

	// HousingClass - housing class id.
	//
	HousingClass HousingClass `json:"housing_class" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// ConstructionTypeID - construction type id.
	//
	ConstructionTypeID int64 `json:"construction_type_id" gorm:"not null;column:construction_type_id"`

	// ConstructionType  - construction type.
	//
	ConstructionType ConstructionType `json:"construction_type" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Title - title of residence
	//
	Title string `json:"title" gorm:"column:title"`

	// Description - description of the residence.
	//
	Description string `json:"description" gorm:"column:description"`

	// Address - address of residence.
	//
	Address string `json:"address" gorm:"column:address"`

	// DistrictID - id of district
	//
	DistrictID DistrictID `json:"district_id" gorm:"not null;column:district_id"`

	// District - name of district where residence located.
	//
	District District `json:"district" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Latitude - latitude of the residence.
	//
	Latitude float64 `json:"latitude" gorm:"column:latitude"`

	// Longitude - longitude of the residence.
	//
	Longitude float64 `json:"longitude" gorm:"column:longitude"`

	// DeadlineYear - deadline year of the construction.
	//
	DeadlineYear int64 `json:"deadline_year" gorm:"column:deadline_year"`

	// DeadlineQuarter - deadline quarter of the construction.
	//
	DeadlineQuarter int64 `json:"deadline_quarter" gorm:"column:deadline_quarter"`

	// FlatsCount - the amount of flats in residence.
	//
	FlatsCount int64 `json:"flats_count" gorm:"column:flats_count"`

	// FloorsMax - the maximum number of floors in the residence.
	//
	FloorsMax int64 `json:"floors_max" gorm:"column:floors_max"`

	// RoomsMin - the minimum number of rooms in the flat.
	//
	RoomsMin int64 `json:"rooms_min" gorm:"column:rooms_min"`

	// RoomsMax - the maximum amount of rooms in the flat.
	//
	RoomsMax int64 `json:"rooms_max" gorm:"column:rooms_max"`

	// CeilingHeight - the height of the ceiling.
	//
	CeilingHeight float32 `json:"ceiling_height" gorm:"column:ceiling_height"`

	// HasHGF - has guarantee from Kazakhstan housing
	//
	HasHGF bool `json:"has_hgf" gorm:"column:has_hgf"`

	// PricePerSquareMin - the minimum price per square.
	//
	PricePerSquareMin int64 `json:"price_per_square_min" gorm:"column:price_per_square_min"`

	// PriceMin - the maximum price of the flat.
	//
	PriceMin int64 `json:"price_min" gorm:"column:price_min"`

	// PriceMax - the maximum price of the flat.
	//
	PriceMax int64 `json:"price_max" gorm:"column:price_max"`

	// AreaMin - the minimum area of the flat.
	//
	AreaMin float32 `json:"area_min" gorm:"column:area_min"`

	// AreaMax - the maximum area.
	//
	AreaMax float32 `json:"area_max" gorm:"column:area_max"`

	// TitleImage - the title image.
	//
	TitleImage string `json:"title_image" gorm:"column:title_image"`

	// Images - the images of residence.
	//
	Images pq.StringArray `json:"images" gorm:"type:varchar[];column:images"`

	// Views - the amount of total views of the residence.
	//
	Views int64 `json:"views" gorm:"column:views"`

	// Likes - the amount of likes.
	//
	Likes int64 `json:"likes" gorm:"column:likes"`

	// FlatPlans
	//
	FlatPlans []*FlatPlan `json:"flat_plans" gorm:"foreignKey:residence_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// ParkingTypes - parking types.
	//
	ParkingTypes []*ParkingType `json:"parking_types" gorm:"many2many:residences_parking_types;"`

	// InteriorDecoration - interior decoration of the flat.
	//
	InteriorDecorations []*InteriorDecoration `json:"interior_decorations" gorm:"many2many:residences_interior_decorations;"`

	// HeatingTypes - heating types.
	//
	HeatingTypes []*HeatingType `json:"heating_types" gorm:"many2many:residences_heating_types;"`

	// ElevatorTypes - heating types.
	//
	ElevatorTypes []*ElevatorType `json:"elevator_types" gorm:"many2many:residences_elevator_types;"`

	// PurchaseMethods - purchase methods
	//
	PurchaseMethods []*PurchaseMethod `json:"purchase_methods" gorm:"many2many:residences_purchase_methods;"`

	// ConsultationEmail - email
	//
	ConsultationEmail string `json:"consultation_email" gorm:"column:consultation_email"`

	// IsFavourite - is favourite
	// read-only
	IsFavourite bool `json:"is_favourite" gorm:"column:is_favourite"`

	// Slug - slug used for crm
	//
	Slug string `json:"slug"`

	// CreatedAt - created timestamp.
	//
	CreatedAt int64 `json:"created_at" gorm:"<-:create;not null;autoCreateTime;column:created_at;"`

	// UpdatedAt - updated timestamp.
	//
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// DeletedAt - deleted timestamp.
	//
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`

	// SaleStatusID - sale status.
	//
	SaleStatusID int64 `json:"sale_status_id" gorm:"column:sale_status_id"`

	// Status - represents status.
	// only read
	SaleStatus SaleStatus `json:"sale_status" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// ResidenceSearchCriteria - represents search criteria.
type ResidenceSearchCriteria struct {
	Page                  PageRequest
	Sorts                 []Sort  // sorting
	Title                 string  // filter by title
	BuilderIDs            []int64 // filter by builders
	CityID                int64   // filter by city id
	DistrictID            int64   // filter by district id
	StatusID              int64   // filter by status
	RoomsMin              int64   // filter by the amount of rooms
	RoomsMax              int64   // filter by the amount of rooms
	CeilingHeightMin      float32 // filter by ceiling height
	CeilingHeightMax      float32 // filter by ceiling height
	HasHGF                *bool   // filter by has_hgf
	AreaMin               float64 // filter by area
	AreaMax               float64 // filter by area
	PriceMin              int64   // filter by price
	PriceMax              int64   // filter by price
	ConstructionTypesIDs  []int64 // filter by construction types
	ParkingTypesIds       []int64 // filter by parking types
	InteriorDecorationIDs []int64 // filter by interior decoration
	HeatingTypesIDs       []int64 // filter by heating types
	PurchaseMethodsIDs    []int64 // filter by purchase methods
	ElevatorTypesIDs      []int64 // filter by elevator types
	HousingClassID        int64   // filter by housing class
	FloorsMin             int64   // filter by floors
	FloorsMax             int64   // filter by floors
	UserID                int64   // filter favourites
}

// ResidencesReadRepository - provides read access to a storage.
type ResidencesReadRepository interface {

	// List - returns list of residences by criteria.
	//
	List(
		ctx context.Context,
		criteria ResidenceSearchCriteria,
	) ([]*Residence, Total, error)

	// ListResidencesByIDs returns list all the residences by given ids
	//
	ListResidencesByIDs(
		ctx context.Context,
		residencesIDs []ResidenceID,
	) ([]*Residence, error)

	// Get - returns residence by id.
	//
	Get(
		ctx context.Context,
		residenceID ResidenceID,
	) (*Residence, error)

	// ListPopularResidences returns list all the residences
	//
	ListPopularResidences(
		ctx context.Context,
		criteria ResidenceSearchCriteria,
	) ([]*Residence, Total, error)

	// GetConsultationEmailByResidenceID - returns a consultation email of the residence
	//
	GetConsultationEmailByResidenceID(
		ctx context.Context,
		residenceID ResidenceID,
	) (string, error)

	// IsFavouriteResidence - returns residence by id.
	//
	IsFavouriteResidence(
		ctx context.Context,
		residenceID ResidenceID,
		userID UserID,
	) (bool, error)
}

// ResidencesRepository - provides access to a storage.
type ResidencesRepository interface {
	ResidencesReadRepository
	ResidenceFlatPlansRepository

	// Create - creates residence in storage.
	//
	Create(
		ctx context.Context,
		residence *Residence,
	) (*Residence, error)

	// Update - updates residence by id.
	//
	Update(
		ctx context.Context,
		residenceID ResidenceID,
		residence *Residence,
	) (*Residence, error)

	// Delete - deletes residence by id.
	//
	Delete(
		ctx context.Context,
		residenceID ResidenceID,
	) error
}

// ResidenceFlatPlansRepository - provides access to a flat plans storage.
type ResidenceFlatPlansRepository interface {
	// CreateFlatPlan - creates flat plan
	//
	CreateFlatPlan(
		ctx context.Context,
		flatPlan *FlatPlan,
	) (*FlatPlan, error)

	// UpdateFlatPlan - updates flat plan
	//
	UpdateFlatPlan(
		ctx context.Context,
		flatPlanID FlatPlanID,
		flatPlan *FlatPlan,
	) (*FlatPlan, error)

	// DeleteFlatPlan - deletes flat plan
	//
	DeleteFlatPlan(
		ctx context.Context,
		flatPlanID FlatPlanID,
	) error
}

// ResidenceRedisRepository - provides access to a redis storage.
type ResidenceRedisRepository interface {

	// Get - returns residence from cache by key.
	//
	Get(ctx context.Context, key string) (*Residence, error)

	// Set - sets a residence to cache with specific key.
	//
	Set(ctx context.Context, key string, seconds int, residence *Residence) error

	// Delete - deletes a residence from cache by key.
	Delete(ctx context.Context, key string) error
}

// ResidenceFlatPlansService - provides access to a business logic of flat plans.
type ResidenceFlatPlansService interface {
	// CreateFlatPlan - creates flat plan
	//
	CreateFlatPlan(
		ctx context.Context,
		flatPlan *FlatPlan,
		callerID CallerID,
	) (*FlatPlan, error)

	// UpdateFlatPlan - updates flat plan
	//
	UpdateFlatPlan(
		ctx context.Context,
		flatPlanID FlatPlanID,
		flatPlan *FlatPlan,
		callerID CallerID,
	) (*FlatPlan, error)

	// DeleteFlatPlan - deletes flat plan
	//
	DeleteFlatPlan(
		ctx context.Context,
		flatPlanID FlatPlanID,
		callerID CallerID,
	) error
}

// ResidencesService - provides access to a business logic.
type ResidencesService interface {
	ResidenceFlatPlansService

	// CreateResidence creates a residence
	//
	CreateResidence(
		ctx context.Context,
		residence *Residence,
		callerID CallerID,
	) (*Residence, error)

	// ListResidences returns list all the residences
	//
	ListResidences(
		ctx context.Context,
		criteria ResidenceSearchCriteria,
		callerID CallerID,
	) ([]*Residence, Total, error)

	// ListPopularResidences returns list all the residences
	//
	ListPopularResidences(
		ctx context.Context,
		criteria ResidenceSearchCriteria,
		callerID CallerID,
	) ([]*Residence, Total, error)

	// ListResidencesByIDs returns list all the residences by given ids
	//
	ListResidencesByIDs(
		ctx context.Context,
		residencesIDs []ResidenceID,
		callerID CallerID,
	) ([]*Residence, error)

	// GetResidence returns residence by id
	//
	GetResidence(
		ctx context.Context,
		residenceID ResidenceID,
		callerID CallerID,
	) (*Residence, error)

	// UpdateResidence updates residence
	//
	UpdateResidence(
		ctx context.Context,
		residenceID ResidenceID,
		residence *Residence,
		callerID CallerID,
	) (*Residence, error)

	// DeleteResidence deletes residence
	//
	DeleteResidence(
		ctx context.Context,
		residenceID ResidenceID,
		callerID CallerID,
	) error

	// IsResidenceExist - returns if residence exists.
	//
	IsResidenceExist(
		ctx context.Context,
		residenceID ResidenceID,
		callerID CallerID,
	) (bool, error)

	// GetConsultationEmailByResidenceID - returns a consultation email of the residence
	//
	GetConsultationEmailByResidenceID(
		ctx context.Context,
		residenceID ResidenceID,
		callerID CallerID,
	) (string, error)
}
