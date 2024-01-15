package domain

// Statuses table name
const (
	StatusesTableName = "statuses"
)

// StatusID - status id
type StatusID int64

// Status - represents status
type Status struct {

	// ID - id of status.
	//
	ID StatusID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of status.
	//
	Name string `json:"name" gorm:"column:name"`
}
