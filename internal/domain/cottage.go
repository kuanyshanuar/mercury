package domain

import (
	"context"
	"github.com/lib/pq"

	"gorm.io/plugin/soft_delete"
)

// Cottages - constants
const (
	CottageTableName      = "cottages"
	FavouriteCottageTable = "cottages_bookmarks"
)

// Cottage is the struct for one unit of a cottage
type Cottage struct {
	// ID is an id of the cottage
	//
	ID int64 `json:"id" gorm:"column:id"`

	// CityID - id of a city in the database table
	//
	CityID int64 `json:"city_id" gorm:"column:city_id"`

	// City - city information detailed
	//
	City *City `json:"city" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	// UserID - user id in the database table
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
	Status *Status `json:"status" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// SaleStatusID - sale status of the residence.
	//
	SaleStatusID int64 `json:"sale_status_id" gorm:"column:sale_status_id"`

	// SaleStatus - represents sale status.
	// only read
	SaleStatus *SaleStatus `json:"sale_status" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// DistrictID - id of the district from the database table
	//
	DistrictID int64 `json:"district_id" gorm:"column:district_id"`

	// District
	//
	District *District `json:"district" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// HousingCLassID - id of the housing class
	//
	HousingClassID HousingClassID `json:"housing_class_id" gorm:"column:housing_class_id"`

	// HousingCLass - class of the house('elite', 'regular', see migrations)
	//
	HousingClass HousingClass `json:"housing_class" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Title is a short title of the cottage, name of the organization
	//
	Title string `json:"title" gorm:"column:title"`

	// Description is a description of the cottage - like about, benefits, additional info
	//
	Description string `json:"description" gorm:"column:description"`

	// Address is an address of the cottage
	//
	Address string `json:"address" gorm:"column:address"`

	// Latitude is an latitude(ширина участка) of the cottage
	//
	Latitude float64 `json:"latitude" gorm:"column:latitude"`

	// Longitude is longitude(длина участка) of the cottage
	//
	Longitude float64 `json:"longitude" gorm:"column:longitude"`

	// Territory is a description of the territory
	//
	Territory string `json:"territory" gorm:"column:territory"`

	// CeilingHeightMin is the minimal height of the ceiling among all the houses inside the cottage city
	//
	CeilingHeightMin float64 `json:"ceiling_height_min" gorm:"column:ceiling_height_min"`

	// CeilingHeightMax is height of the ceiling among all the houses inside the cottage city
	//
	CeilingHeightMax float64 `json:"ceiling_height_max" gorm:"column:ceiling_height_max"`

	// BuildingArea is an area of the total building
	//
	BuildingArea float64 `json:"building_area" gorm:"column:building_area"`

	// AreaMin - minimal area among the houses in the cottage city
	//
	AreaMin float64 `json:"area_min" gorm:"column:area_min"`
	// AreaMax - maximal area among the houses in the cottage city
	//
	AreaMax float64 `json:"area_max" gorm:"column:area_max"`
	// HouseAmount shows how many houses are around here
	//
	HouseAmount int64 `json:"house_amount" gorm:"column:house_amount"`

	// FloorsCount shows how many floors there can be in the cottage
	//
	FloorsCount int64 `json:"floors_count" gorm:"column:floors_count"`

	// Facade is a facade type of the cottage
	//
	Facade string `json:"facade" gorm:"column:facade"`

	// CanRePlan shows whether the territory can be planned differently, rebuilt
	//
	CanRePlan bool `json:"can_replan" gorm:"column:can_replan"`

	// RoomsMin - minimal number of rooms among the cottage house plans
	//
	RoomsMin int64 `json:"rooms_min" gorm:"column:rooms_min"`

	// RoomsMax - maximal number of rooms among the cottage house plans
	//
	RoomsMax int64 `json:"rooms_max" gorm:"column:rooms_max"`

	// PricePerSquareMin is a minimal price per square for the cottage planning
	//
	PricePerSquareMin float64 `json:"price_per_square_min" gorm:"column:price_per_square_min"`

	// PricePerSquare is a maximal price per square for the cottage planning
	//
	PricePerSquareMax float64 `json:"price_per_square_max" gorm:"column:price_per_square_max"`

	// IsFavourite - indicates whether the cottage (in return value for the transport) - was bookmarked by user or not...
	//
	IsFavourite bool `json:"is_favourite" gorm:"column:is_favourite"`

	// HousePlan - house plans of the cottage city
	//
	HousePlans []*HousePlan `json:"house_plans" gorm:"foreignKey:cottage_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// WindowTypes is a window type of the cottage
	//
	WindowTypes []*WindowType `json:"windows" gorm:"many2many:cottages_windows"`

	// WallTypes
	//
	WallTypes []*WallType `json:"wall_types" gorm:"many2many:cottages_wall_types"`

	// ElevatorTypes
	//
	ElevatorTypes []*ElevatorType `json:"elevator_types" gorm:"many2many:cottages_elevator_types"`

	// WarmingTypeIDs
	//
	WarmingTypes []*WarmingType `json:"warming_types" gorm:"many2many:cottages_warming_types"`

	// InteriorDecorationIDs
	//
	InteriorDecorations []*InteriorDecoration `json:"interior_decorations" gorm:"many2many:cottages_interior_decorations;"`

	// PurchaseMethods - available purchase methods
	//
	PurchaseMethods []*PurchaseMethod `json:"purchase_methods" gorm:"many2many:cottages_purchase_methods"`

	// HeatingTypes - available heating types for the cottage
	//
	HeatingTypes []*HeatingType `json:"heating_types" gorm:"many2many:cottages_heating_types"`

	// ParkingTypes - available parking types
	//
	ParkingTypes []*ParkingType `json:"parking_types" gorm:"many2many:cottages_parking_types"`

	// Images - introductory images of the cottage
	//
	Images pq.StringArray `json:"images" gorm:"type:varchar[];column:images"`

	// CreatedAt - created time.
	//
	CreatedAt int64 `json:"created_at" gorm:"<-:create;not null;autoCreateTime;column:created_at"`

	// UpdatedAt - updated time.
	//
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// DeletedAt - deleted timestamp.
	//
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

// CottageSearchCriteria is a collection of criteria to search for the cottage
type CottageSearchCriteria struct {
	Page               PageRequest
	Sorts              []Sort  // sorting
	Title              string  // filter by title
	BuilderIDs         []int64 // filter by builders
	CityID             int64   // filter by city id
	DistrictID         int64   // filter by district id
	StatusID           int64   // filter by status
	RoomsMin           int64   // filter by the amount of rooms
	RoomsMax           int64   // filter by the amount of rooms
	CeilingHeightMin   float64 // filter by ceiling height
	CeilingHeightMax   float64 // filter by ceiling height
	AreaMin            float64 // filter by area
	AreaMax            float64 // filter by area
	PriceMin           int64   // filter by price
	PriceMax           int64   // filter by price
	InteriorDecoration []int64 // filter by interior decoration
	HeatingTypes       []int64 // filter by heating types
	PurchaseMethods    []int64 // filter by purchase methods
	ElevatorTypes      []int64 // filter by elevator types
	WallTypes          []int64 // filter by wall types
	WindowTypes        []int64 // filter by window types
	WarmingTypes       []int64 // filter by warming types
	ParkingTypes       []int64 // filter by parking types
	HouseType          int64   // filter by housing class
	FloorsMin          int64   // filter by floors
	FloorsMax          int64   // filter by floors
	UserID             int64   // filter favourites
	CanRePlan          bool    // filter by can replan
	HouseAmount        int64   // filter by house amount in the cottage city
}

// CottageRepository is a repository for cottages
type CottageRepository interface {
	CottageHousePlanRepository

	// Get allows to retrieve the cottage information given its id in the database
	//
	Get(
		ctx context.Context,
		id int64,
	) (*Cottage, error)

	// List allows to retrieve the cottage information given search criterias
	//
	List(
		ctx context.Context,
		criteria CottageSearchCriteria,
	) ([]*Cottage, Total, error)

	// ListPopularCottages allows to retrieve the popular cottages, see
	//
	ListPopularCottages(
		ctx context.Context,
		criteria CottageSearchCriteria,
	) ([]*Cottage, Total, error)

	// ListCottagesByIDs allows to retrieve cottage samples by ID array
	//
	ListCottagesByIDs(
		ctx context.Context,
		ids []int64,
	) ([]*Cottage, error)

	// IsFavouriteCottage allows to know whether the cottage was bookmarked by the user
	//
	IsFavouriteCottage(ctx context.Context,
		cottageID int64,
		userID UserID,
	) (bool, error)

	// Update allows to update the cottage by its id and modified information
	//
	Update(
		ctx context.Context,
		cottageID int64,
		cottage *Cottage,
	) (*Cottage, error)

	// Delete allows to delete the cottage by its id
	//
	Delete(
		ctx context.Context,
		cottageID int64,
	) error

	// Create allows to create the cottage given enough information,
	// see validate cottage method in service
	Create(
		ctx context.Context,
		cottage *Cottage,
	) (*Cottage, error)
}

// CottageHousePlanRepository is a repository for manipulating cottage flat plans
type CottageHousePlanRepository interface {
	// UpdateHousePlan allows to update the cottage house plan by its id and modified information
	//
	UpdateHousePlan(
		ctx context.Context,
		housePlanID int64,
		cottage *HousePlan,
	) (*HousePlan, error)

	// DeleteHousePlan allows to delete the cottage house plan by its id
	//
	DeleteHousePlan(
		ctx context.Context,
		housePlanID int64,
	) error

	// CreateHousePlan allows to create the cottage given enough information,
	//
	CreateHousePlan(
		ctx context.Context,
		cottagePlan *HousePlan,
	) (*HousePlan, error)

	// GetConsultationEmailByCottageID - returns consultation email by cottage id
	//
	GetConsultationEmailByCottageID(
		ctx context.Context,
		cottageID int64,
	) (string, error)
}

// CottageHousePlanService is a service for cottage flat plans
type CottageHousePlanService interface {
	// UpdateHousePlan allows to update the cottage house plan by its id and modified information
	//
	UpdateHousePlan(
		ctx context.Context,
		housePlanID int64,
		housePlan *HousePlan,
		callerID CallerID,
	) (*HousePlan, error)

	// DeleteHousePlan allows to delete the cottage house plan by its id
	//
	DeleteHousePlan(
		ctx context.Context,
		housePlanID int64,
		callerID CallerID,
	) error

	// CreateHousePlan allows to create the cottage given enough information,
	// see validate cottage method in service
	CreateHousePlan(
		ctx context.Context,
		housePlan *HousePlan,
		callerID CallerID,
	) (*HousePlan, error)
}

// CottageService is a service for cottages
type CottageService interface {
	CottageHousePlanService

	// UpdateCottage allows to update the cottage by its id and modified information
	//
	UpdateCottage(ctx context.Context,
		cottageID int64,
		cottage *Cottage,
		callerID CallerID,
	) (*Cottage, error)

	// DeleteCottage allows to delete the cottage by its id
	//
	DeleteCottage(
		ctx context.Context,
		cottageID int64,
		callerID CallerID,
	) error

	// CreateCottage allows to create the cottage given enough information,
	//
	CreateCottage(
		ctx context.Context,
		cottage *Cottage,
		callerID CallerID,
	) (*Cottage, error)

	// GetCottage allows to retrieve the cottage information given its id in the database
	//
	GetCottage(
		ctx context.Context,
		id int64,
		callerID CallerID,
	) (*Cottage, error)

	// ListCottage allows to retrieve the cottage information given search criteria
	//
	ListCottage(
		ctx context.Context,
		criteria CottageSearchCriteria,
		callerID CallerID,
	) ([]*Cottage, Total, error)

	// ListPopularCottages allows to retrieve the popular cottages, see
	// repository implementation for more details
	ListPopularCottages(
		ctx context.Context,
		criteria CottageSearchCriteria,
		callerID CallerID,
	) ([]*Cottage, Total, error)

	// ListCottagesByIDs allows to retrieve cottage samples by ID array
	//
	ListCottagesByIDs(
		ctx context.Context,
		id []int64,
		callerID CallerID,
	) ([]*Cottage, error)

	// IsFavouriteCottage allows to know whether the cottage was bookmarked by the user
	//
	IsFavouriteCottage(ctx context.Context,
		cottageID int64,
		userID int64,
		callerID CallerID,
	) (bool, error)

	// IsCottageExist - returns if cottage exists.
	//
	IsCottageExist(
		ctx context.Context,
		cottageID int64,
		callerID CallerID,
	) (bool, error)

	// GetConsultationEmailByCottageID - returns a consultation email of the residence
	//
	GetConsultationEmailByCottageID(
		ctx context.Context,
		cottageID int64,
		callerID CallerID,
	) (string, error)
}
