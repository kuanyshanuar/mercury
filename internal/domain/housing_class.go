package domain

import (
	"context"
	"fmt"
)

// HousingClassTableName - name of residence table in storage.
const (
	HousingClassTableName = "housing_classes"
)

// HousingClassID - id of the housing class id.
type HousingClassID int

// HousingClass - housing class.
type HousingClass struct {
	// ID - id of housing class id.
	//
	ID HousingClassID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of housing class.
	//
	Name string `json:"name" gorm:"not null;column:name"`
}

// IsValid - validates entire of struct
func (h *HousingClass) IsValid() error {
	if len(h.Name) == 0 {
		return fmt.Errorf("name required")
	}

	return nil
}

// HousingClassRepository - provides access to storage.
type HousingClassRepository interface {
	// CreateHousingClass - creates housing class in a storage.
	//
	CreateHousingClass(
		ctx context.Context,
		class *HousingClass,
	) (HousingClassID, error)

	// ListHousingClasses - returns a list housing classes from storage.
	//
	ListHousingClasses(
		ctx context.Context,
	) ([]*HousingClass, error)

	// DeleteHousingClass - deletes the housing class
	//
	DeleteHousingClass(
		ctx context.Context,
		classID HousingClassID,
	) error
}

// HousingClassService - provides access to business logic.
type HousingClassService interface {

	// CreateHousingClass - create housing class
	//
	CreateHousingClass(
		ctx context.Context,
		class *HousingClass,
		callerID CallerID,
	) (HousingClassID, error)

	// ListHousingClasses - returns a list housing classes.
	//
	ListHousingClasses(
		ctx context.Context,
		callerID CallerID,
	) ([]*HousingClass, error)

	// DeleteHousingClass - deletes the housing class
	//
	DeleteHousingClass(
		ctx context.Context,
		classID HousingClassID,
		callerID CallerID,
	) error
}
