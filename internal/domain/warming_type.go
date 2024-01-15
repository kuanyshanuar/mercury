package domain

// WarmingTableName - table name
const (
	WarmingTableName = "warming_types"
)

// WarmingType is the table for all warming types for cottage
type WarmingType struct {
	// ID - id
	//
	ID int64 `json:"id" gorm:"column:id"`

	// Name - name
	//
	Name string `json:"name" gorm:"column:name"`
}
