package domain

import "context"

// Filter - represents filter struct
type Filter struct {
	// ID - id of housing class id.
	//
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of housing class.
	//
	Name string `json:"name" gorm:"not null;column:name"`
}

// Filters - represents filters
type Filters struct {
	Filters map[string]Filter `json:"filters"`
}

// FilterBuilder - represents builder in filter
type FilterBuilder struct {

	// ID - id of user.
	//
	ID int64 `json:"id" gorm:"column:id"`

	// FirstName - first name
	//
	FirstName string `json:"first_name" gorm:"column:first_name"`

	// LastName - last name
	//
	LastName string `json:"last_name" gorm:"column:last_name"`
}

// FiltersRepository - provides access to storage.
type FiltersRepository interface {

	// CreateCity - creates a city
	//
	CreateCity(
		ctx context.Context,
		city *City,
	) (CityID, error)

	// ListCities - returns a list cities.
	//
	ListCities(
		ctx context.Context,
		criteria CitySearchCriteria,
	) ([]*City, Total, error)

	// UpdateCity - updates city by id
	//
	UpdateCity(
		ctx context.Context,
		cityID CityID,
		city *City,
	) error

	// DeleteCity - deletes a city by id
	//
	DeleteCity(
		ctx context.Context,
		cityID CityID,
	) error

	CreateDistrict(
		ctx context.Context,
		district *District,
	) (DistrictID, error)

	// UpdateDistrict - updates district
	//
	UpdateDistrict(
		ctx context.Context,
		districtID DistrictID,
		district *District,
	) error

	// ListDistricts - returns a list of districts
	//
	ListDistricts(
		ctx context.Context,
		criteria DistrictSearchCriteria,
	) ([]*District, Total, error)

	// DeleteDistrict - deletes district
	//
	DeleteDistrict(
		ctx context.Context,
		districtID DistrictID,
	) error

	// ListFilterKeys - returns a list of filter keys
	//
	ListFilterKeys(
		ctx context.Context,
	) ([]*FilterKey, error)

	// ListFilters - returns a list of filters by given key
	//
	ListFilters(
		ctx context.Context,
		key string,
	) ([]*Filter, error)

	// ListBuilders - returns a list of builders
	//
	ListBuilders(
		ctx context.Context,
	) ([]*FilterBuilder, error)

	// CreateFilter - creates filter
	//
	CreateFilter(
		ctx context.Context,
		key string,
		filter *Filter,
	) (*Filter, error)

	// DeleteFilter - deletes filter
	//
	DeleteFilter(
		ctx context.Context,
		id int64,
		key string,
	) error
}

// FiltersService - provides access to business logic.
type FiltersService interface {

	// CreateCity - creates city
	//
	CreateCity(
		ctx context.Context,
		city *City,
		callerID CallerID,
	) (CityID, error)

	// UpdateCity - updates city by id
	//
	UpdateCity(
		ctx context.Context,
		cityID CityID,
		city *City,
		callerID CallerID,
	) error

	// ListCities - returns a list cities.
	//
	ListCities(
		ctx context.Context,
		criteria CitySearchCriteria,
		callerID CallerID,
	) ([]*City, Total, error)

	// DeleteCity - deletes city
	//
	DeleteCity(
		ctx context.Context,
		cityID CityID,
		callerID CallerID,
	) error

	CreateDistrict(
		ctx context.Context,
		district *District,
		callerID CallerID,
	) (DistrictID, error)

	// UpdateDistrict - updates district
	//
	UpdateDistrict(
		ctx context.Context,
		districtID DistrictID,
		district *District,
		callerID CallerID,
	) error

	// DeleteDistrict - deletes district
	//
	DeleteDistrict(
		ctx context.Context,
		districtID DistrictID,
		callerID CallerID,
	) error

	// ListDistricts - returns a list of districts
	//
	ListDistricts(
		ctx context.Context,
		criteria DistrictSearchCriteria,
		callerID CallerID,
	) ([]*District, Total, error)

	// ListFilters -  returns a list of filters
	//
	ListFilters(
		ctx context.Context,
		callerID CallerID,
	) (map[string][]*Filter, error)

	// ListBuilders - returns a list of builders
	//
	ListBuilders(
		ctx context.Context,
		callerID CallerID,
	) ([]*FilterBuilder, error)

	// CreateFilter - creates filter
	//
	CreateFilter(
		ctx context.Context,
		key string,
		filter *Filter,
		callerID CallerID,
	) (*Filter, error)

	// DeleteFilter - deletes filter
	//
	DeleteFilter(
		ctx context.Context,
		id int64,
		key string,
		callerID CallerID,
	) error
}
