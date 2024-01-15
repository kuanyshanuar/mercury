package domain

// ConstructionTypesTableName - table name
const (
	ConstructionTypesTableName = "construction_types"
)

// ConstructionTypeID - id of construction
type ConstructionTypeID int64

// ConstructionType - represents construction
type ConstructionType struct {

	// ID - id of the construction type.
	//
	ID ConstructionTypeID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of construction type.
	//
	Name string `json:"name" gorm:"not null;column:name"`
}
