package domain

import (
	"context"
	"fmt"
)

// HeatingTypesTableName - table name
const (
	HeatingTypesTableName = "heating_types"
)

// HeatingTypeID - id of heating type.
type HeatingTypeID int64

// HeatingType - represents heating type.
type HeatingType struct {

	// ID - id of the heating type.
	//
	ID HeatingTypeID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of heating type.
	//
	Name string `json:"name" gorm:"not null;column:name"`
}

// IsValid - validates entire of struct
func (t *HeatingType) IsValid() error {
	if len(t.Name) == 0 {
		return fmt.Errorf("name required")
	}

	return nil
}

// HeatingTypesRepository - provides access to storage.
type HeatingTypesRepository interface {
	// CreateHeatingType - creates a new heating type.
	//
	CreateHeatingType(
		ctx context.Context,
		heatingType *HeatingType,
	) (HeatingTypeID, error)

	// ListHeatingTypes - returns a list of heating types.
	//
	ListHeatingTypes(
		ctx context.Context,
	) ([]*HeatingType, error)

	// DeleteHeatingType - deletes a heating type.
	//
	DeleteHeatingType(
		ctx context.Context,
		heatingTypeID HeatingTypeID,
	) error
}

// HeatingTypesService - provides access to business logic
type HeatingTypesService interface {
	// CreateHeatingType - creates a new heating type.
	//
	CreateHeatingType(
		ctx context.Context,
		heatingType *HeatingType,
		callerID CallerID,
	) (HeatingTypeID, error)

	// ListHeatingTypes - returns a list of heating types.
	//
	ListHeatingTypes(
		ctx context.Context,
		callerID CallerID,
	) ([]*HeatingType, error)

	// DeleteHeatingType - deletes a heating type.
	//
	DeleteHeatingType(
		ctx context.Context,
		heatingTypeID HeatingTypeID,
		callerID CallerID,
	) error
}
