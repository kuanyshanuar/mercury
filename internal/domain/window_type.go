package domain

// WindowTypeTableName - table name
const (
	WindowTypeTableName = "window_types"
)

// WindowType - is the table for all the heating types available
type WindowType struct {
	// ID - id of the window type
	//
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of the window type
	//
	Name string `json:"name" gorm:"column:name"`
}
