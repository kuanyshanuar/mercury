package domain

// DistrictsTableName - name of table
const (
	DistrictsTableName = "districts"
)

// DistrictID id of district
type DistrictID int64

// District - district struct
type District struct {

	// ID - id of district
	//
	ID DistrictID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// CityID - id of city
	//
	CityID int64 `json:"city_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// City - city
	// only read
	//
	City *City `json:"city" gorm:""`

	// Name - name of district
	//
	Name string `json:"name" gorm:"column:name"`
}

// DistrictSearchCriteria - search criteria
type DistrictSearchCriteria struct {
	Page   PageRequest
	CityID int64
}
