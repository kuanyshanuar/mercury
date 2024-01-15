package domain

import (
	"context"

	"gorm.io/plugin/soft_delete"
)

// LeadResidenceTableName table name
const (
	LeadResidenceTableName   = "residence_leads"
	LeadStatusAssociation    = "Status"
	LeadResidenceAssociation = "Residence"
)

// LeadID - id of the lead
type LeadID int64

// LeadResidence - represents a lead residence
type LeadResidence struct {

	// ID - id of the lead
	//
	ID LeadID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// ResidenceID - id of the residence
	//
	ResidenceID ResidenceID `json:"residence_id" gorm:"not null;column:residence_id"`

	// Residence
	//
	Residence *Residence `json:"residence" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// StatusID - id of the status
	//
	StatusID int64 `json:"status_id" gorm:"column:status_id"`

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

// LeadResidenceSearchCriteria - represents search criteria.
type LeadResidenceSearchCriteria struct {
	Page     PageRequest
	Name     string // filter by name
	StatusID int64  // filter by status
}

// LeadResidenceReadRepository - provides access to a read storage
type LeadResidenceReadRepository interface {
	// ListLeadResidences - list residences in lead
	//
	ListLeadResidences(
		ctx context.Context,
		criteria LeadResidenceSearchCriteria,
	) ([]*LeadResidence, Total, error)

	// GetLeadResidence - returns lead by id
	//
	GetLeadResidence(
		ctx context.Context,
		leadID LeadID,
	) (*LeadResidence, error)

	// IsLeadExistByDate - returns if lead exists between given timestamps
	//
	IsLeadExistByDate(
		ctx context.Context,
		residenceID ResidenceID,
		issuedAt int64,
		expiresAt int64,
	) (bool, error)

	// IsOtherLeadExist - returns if other lead exists between given timestamps
	//
	IsOtherLeadExist(
		ctx context.Context,
		leadID LeadID,
		residenceID ResidenceID,
		issuedAt int64,
		expiresAt int64,
	) (bool, error)
}

// LeadResidenceRepository - provides access to a storage
type LeadResidenceRepository interface {
	LeadResidenceReadRepository

	// CreateLeadResidence - creates a lead residence
	//
	CreateLeadResidence(
		ctx context.Context,
		lead *LeadResidence,
	) error

	// UpdateLeadResidence - updates a lead residence
	//
	UpdateLeadResidence(
		ctx context.Context,
		lead *LeadResidence,
	) error

	// DeleteLeadResidence - deletes lead residence by id
	//
	DeleteLeadResidence(
		ctx context.Context,
		leadID LeadID,
	) error

	// RevokeLeadResidence - revoke lead residence
	//
	RevokeLeadResidence(
		ctx context.Context,
		leadID LeadID,
	) error
}

// LeadResidenceService - provides access to business logic
type LeadResidenceService interface {
	// CreateLeadResidence - creates a lead residence
	//
	CreateLeadResidence(
		ctx context.Context,
		lead *LeadResidence,
		callerID CallerID,
	) error

	// UpdateLeadResidence - updates lead residence by id
	//
	UpdateLeadResidence(
		ctx context.Context,
		leadID LeadID,
		lead *LeadResidence,
		callerID CallerID,
	) error

	// ListLeadResidences - list residences in lead
	//
	ListLeadResidences(
		ctx context.Context,
		criteria LeadResidenceSearchCriteria,
		callerID CallerID,
	) ([]*LeadResidence, Total, error)

	// DeleteLeadResidence - deletes lead residence by id
	//
	DeleteLeadResidence(
		ctx context.Context,
		leadID LeadID,
		callerID CallerID,
	) error

	// RevokeLeadResidence - revoke lead residence
	//
	RevokeLeadResidence(
		ctx context.Context,
		leadID LeadID,
		callerID CallerID,
	) error

	// GetLeadResidence - returns lead by id
	//
	GetLeadResidence(
		ctx context.Context,
		leadID LeadID,
		callerID CallerID,
	) (*LeadResidence, error)
}
