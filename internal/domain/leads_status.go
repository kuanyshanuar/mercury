package domain

// Statuses
const (
	StatusActive   = 1
	StatusInactive = 2
)

// LeadStatus - status of leads
type LeadStatus struct {
	// ID - id.
	//
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name.
	//
	Name string `json:"name" gorm:"not null;column:name"`
}
