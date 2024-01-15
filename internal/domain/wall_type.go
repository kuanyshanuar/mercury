package domain

// WallTypeTableName - table name
const (
	WallTypeTableName = "wall_types"
)

// WallType is the table for all wall types for cottage
type WallType struct {
	// ID - id
	//
	ID int64 `json:"id" gorm:"column:id"`

	// Name - name
	//
	Name string `json:"name" gorm:"column:name"`
}
