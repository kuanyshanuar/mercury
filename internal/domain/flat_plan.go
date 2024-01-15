package domain

import (
	"github.com/lib/pq"
	"gorm.io/plugin/soft_delete"
)

// FlatPlansTableName - table name
const (
	FlatPlansTableName = "flat_plans"
)

// FlatPlanID - id of flat plan
type FlatPlanID int64

// FlatPlan - represents flat plan
type FlatPlan struct {
	// ID - id of the plan.
	//
	ID FlatPlanID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// ResidenceID - id of the residence.
	//
	ResidenceID ResidenceID `json:"residence_id" gorm:"not null;column:residence_id"`

	// NumberOfRooms - number of rooms.
	//
	NumberOfRooms int `json:"number_of_rooms" gorm:"not null;column:number_of_rooms"`

	// Area - area of flat.
	//
	Area float32 `json:"area" gorm:"column:area"`

	// Price - price of the flat.
	//
	Price int `json:"price" gorm:"column:price"`

	// Images - images of the flat plans.
	//
	Images pq.StringArray `json:"images" gorm:"type:varchar[];column:images"`

	// CreatedAt - created time.
	//
	CreatedAt int64 `json:"created_at" gorm:"not null;autoCreateTime;column:created_at"`

	// UpdatedAt - updated time
	//
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// DeletedAt - deleted timestamp.
	//
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
