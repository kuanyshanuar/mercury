package domain

// ElevatorTypesTableName - table name
const (
	ElevatorTypesTableName = "elevator_types"
)

// ElevatorTypeID - id of elevator type.
type ElevatorTypeID int64

// ElevatorType - represents elevator type.
type ElevatorType struct {
	// ID
	//
	ID ElevatorTypeID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name
	//
	Name string `json:"name" gorm:"not null:column:name"`
}
