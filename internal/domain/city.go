package domain

// CitiesTableName - cities table name
const (
	CitiesTableName = "cities"
)

// CityID - id of city.
type CityID int64

// City - city struct.
type City struct {

	// ID - id of city
	//
	ID CityID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name of city
	//
	Name string `json:"name" gorm:"column:name"`
}

// CitySearchCriteria - search criteria
type CitySearchCriteria struct {
	Page PageRequest
}
