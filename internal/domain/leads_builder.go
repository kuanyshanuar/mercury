package domain

import (
	"context"
	"gorm.io/plugin/soft_delete"
)

// LeadBuildersTableName - table name
const (
	LeadBuildersTableName = "builder_leads"
)

// LeadBuildersAssociation - association
const (
	LeadBuildersAssociation = "Builder"
)

// LeadBuilder - represents a lead builder
type LeadBuilder struct {
	// ID - id of the lead
	//
	ID LeadID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// BuilderID - id of the builder
	//
	BuilderID BuilderID `json:"builder_id" gorm:"column:builder_id"`

	// Builder - builder
	//
	Builder *User `json:"builder" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// StatusID - id of the status
	//
	StatusID int64 `json:"status_id" gorm:"default:1;column:status_id"`

	// Status - status reference
	//
	Status *LeadStatus `json:"status" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// IssuedAt - issued datetime
	//
	IssuedAt int64 `json:"issued_at" gorm:"not null;column:issued_at"`

	// ExpiresAt - expiration time
	//
	ExpiresAt int64 `json:"expires_at" gorm:"not null;column:expires_at"`

	// DeletedAt - deleted timestamp.
	//
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

// LeadBuilderSearchCriteria - represents search criteria.
type LeadBuilderSearchCriteria struct {
	Page     PageRequest
	Name     string // filter by name
	StatusID int64  // filter by status
}

// LeadBuilderRepository - provides access to a storage.
type LeadBuilderRepository interface {
	// CreateLeadBuilder - creates a lead builder
	//
	CreateLeadBuilder(
		ctx context.Context,
		lead *LeadBuilder,
	) error

	// GetLeadBuilder - returns lead by id
	//
	GetLeadBuilder(
		ctx context.Context,
		leadID LeadID,
	) (*LeadBuilder, error)

	// UpdateLeadBuilder - updates a lead builder
	//
	UpdateLeadBuilder(
		ctx context.Context,
		leadID LeadID,
		lead *LeadBuilder,
	) error

	// ListLeadBuilders - list builders in lead
	//
	ListLeadBuilders(
		ctx context.Context,
		criteria LeadBuilderSearchCriteria,
	) ([]*LeadBuilder, Total, error)

	// DeleteLeadBuilder - deletes lead builder by id
	//
	DeleteLeadBuilder(
		ctx context.Context,
		leadID LeadID,
	) error

	// RevokeLeadBuilder - revoke lead builder
	//
	RevokeLeadBuilder(
		ctx context.Context,
		leadID LeadID,
	) error

	// IsLeadExistByDateRange - returns if lead exists between given timestamps
	//
	IsLeadExistByDateRange(
		ctx context.Context,
		builderID BuilderID,
		issuedAt int64,
		expiresAt int64,
	) (bool, error)

	// IsOtherLeadExist - returns if other lead exists between given timestamps
	//
	IsOtherLeadExist(
		ctx context.Context,
		leadID LeadID,
		builderID BuilderID,
		issuedAt int64,
		expiresAt int64,
	) (bool, error)
}

// LeadBuilderService - provides access to a business logic.
type LeadBuilderService interface {
	// CreateLeadBuilder - creates a lead builder
	//
	CreateLeadBuilder(
		ctx context.Context,
		lead *LeadBuilder,
		callerID CallerID,
	) error

	// GetLeadBuilder - returns lead by id
	//
	GetLeadBuilder(
		ctx context.Context,
		leadID LeadID,
		callerID CallerID,
	) (*LeadBuilder, error)

	// UpdateLeadBuilder - updates lead builder by id
	//
	UpdateLeadBuilder(
		ctx context.Context,
		leadID LeadID,
		lead *LeadBuilder,
		callerID CallerID,
	) error

	// ListLeadBuilders - list builders in lead
	//
	ListLeadBuilders(
		ctx context.Context,
		criteria LeadBuilderSearchCriteria,
		callerID CallerID,
	) ([]*LeadBuilder, Total, error)

	// DeleteLeadBuilder - deletes lead builder by id
	//
	DeleteLeadBuilder(
		ctx context.Context,
		leadID LeadID,
		callerID CallerID,
	) error

	// RevokeLeadBuilder - revoke lead builder
	//
	RevokeLeadBuilder(
		ctx context.Context,
		leadID LeadID,
		callerID CallerID,
	) error
}
