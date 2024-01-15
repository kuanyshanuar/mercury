package domain

// ParkingTypesTableName - table name
const (
	ParkingTypesTableName = "parking_types"
)

// ParkingTypeID - id of parking type.
type ParkingTypeID int64

// ParkingType - represents parking type.
type ParkingType struct {

	// ID - id of parking type.
	//
	ID ParkingTypeID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of parking type.
	//
	Name string `json:"name" gorm:"not null;column:name"`
}
