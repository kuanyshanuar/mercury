package domain

import (
	"github.com/lib/pq"

	"gorm.io/plugin/soft_delete"
)

// HousePlansTable - house plan table in the database
const (
	HousePlansTable = "house_plans"
)

// HousePlan is the collection of plannings of the particular cottage city
type HousePlan struct {
	// ID is an id of the cottage
	//
	ID int64 `json:"id" gorm:"column:id"`

	// CottageID is an id of the cottage
	//
	CottageID int64 `json:"cottage_id" gorm:"not null;column:cottage_id"`

	// HousingClassID is id of the cottage's house type(elite or whatever)
	//
	HousingClassID int64 `json:"housing_class_id" gorm:"column:housing_class_id"`

	// Title is just title information about the plan
	//
	Title string `json:"title" gorm:"column:title"`

	// NumberOfRooms is number of rooms inside the plan
	//
	NumberOfRooms int64 `json:"number_of_rooms" gorm:"column:number_of_rooms"`

	// Area is an area of the particular plan
	//
	Area float64 `json:"area" gorm:"column:area"`

	// Longitude is a longitude of the cottage house
	//
	Longitude float64 `json:"longitude" gorm:"column:longitude"`

	// Territory is an id of the cottage
	//
	Territory float64 `json:"territory" gorm:"column:territory"`

	// CeilingHeight is the height of the ceiling of the cottage
	//
	CeilingHeight float64 `json:"ceiling_height" gorm:"column:ceiling_height"`

	// Price is a price of the cottage_plan
	//
	Price int64 `json:"price" gorm:"column:price"`

	// PricePerSquare is a price per square calculated by price/area of the cottage_plan
	//
	PricePerSquare float64 `json:"price_per_square" gorm:"column:price_per_square"`

	// PlanImages is images of the plan(schemes)
	//
	PlanImages pq.StringArray `json:"plan_images" gorm:"type:varchar[];column:plan_images"`

	// HouseImages is images of the cottage itself
	//
	HouseImages pq.StringArray `json:"house_images" gorm:"type:varchar[];column:house_images"`

	// CreatedAt - created time.
	//
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`

	// UpdatedAt - updated time.
	//
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// DeletedAt - deleted timestamp.
	//
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
