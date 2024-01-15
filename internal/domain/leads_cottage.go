package domain

import (
	"context"
	"gorm.io/plugin/soft_delete"
)

const (
	// LeadCottageTableName - name of the lead cottages in the postgres
	LeadCottageTableName = "cottage_leads"

	// LeadCottageAssociation - name of the lead cottages association for the gorm
	LeadCottageAssociation = "Cottage"

	// LeadCottageStatusAssociation - name of the association of the status for the gorm
	LeadCottageStatusAssociation = "Status"
)

// LeadCottage - service for the lead of the cottages, to highlight them on the list, almost same as lead Cottages
type LeadCottage struct {
	// ID - primary key of each lead cottage issue
	//
	ID int64 `json:"id" gorm:"column:id"`

	// CottageID - cottage id of each lead cottage issue
	//
	CottageID int64 `json:"cottage_id" gorm:"column:cottage_id"`

	// Cottage - cottage struct of each lead cottage issue
	//
	Cottage *Cottage `json:"cottage" gorm:"constraint:OnDelete:Cascade; onUpdate:Cascade"`

	// StatusID - id of the status of the lead cottage issue
	//
	StatusID int64 `json:"status_id" gorm:"column:status_id"`

	// Status - status of the lead cottage issue
	//
	Status *LeadStatus `json:"status" gorm:"constraint:OnDelete:Cascade; onUpdate:Cascade"`

	// IssuedAt - date when cottage lead was issued
	//
	IssuedAt int64 `json:"issued_at" gorm:"column:issued_at"`

	// ExpiresAt - date when cottage lead was issued
	//
	ExpiresAt int64 `json:"expires_at" gorm:"column:expires_at"`

	// DeletedAt - date when the issue was deleted
	//
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
}

// LeadCottageSearchCriteria - collection of criteria for the searching cottage-leads
type LeadCottageSearchCriteria struct {
	Page     PageRequest
	Name     string // filter by name
	StatusID int64  // filter by status
}

// LeadCottageRepository - repository for lead cottage
type LeadCottageRepository interface {
	// CreateLeadCottage - creates the lead instance of the cottage in the database
	//
	CreateLeadCottage(
		ctx context.Context,
		lead *LeadCottage,
	) (int64, error)

	// UpdateLeadCottage - updates the lead instance of the cottage in the database
	//
	UpdateLeadCottage(
		ctx context.Context,
		id int64,
		lead *LeadCottage,
	) (*LeadCottage, error)

	// DeleteLeadCottage - deletes the lead instance of the cottage in the database
	//
	DeleteLeadCottage(
		ctx context.Context,
		id int64,
	) error

	// GetLeadCottage - gets the lead instance of the cottage from the database
	//
	GetLeadCottage(
		ctx context.Context,
		id int64,
	) (*LeadCottage, error)

	// ListLeadCottage - lists the lead instances of the cottage from the database
	//
	ListLeadCottage(
		ctx context.Context,
		criteria LeadCottageSearchCriteria,
	) ([]*LeadCottage, Total, error)

	// IsLeadExistByDate - returns if lead exists between given timestamps
	//
	IsLeadExistByDate(
		ctx context.Context,
		cottageID int64,
		issuedAt int64,
		expiresAt int64,
	) (bool, error)

	// IsOtherLeadExist - returns if other lead exists between given timestamps
	//
	IsOtherLeadExist(
		ctx context.Context,
		leadID LeadID,
		cottageID int64,
		issuedAt int64,
		expiresAt int64,
	) (bool, error)

	// RevokeLeadCottage - revoke lead Cottage
	//
	RevokeLeadCottage(
		ctx context.Context,
		leadID LeadID,
	) error
}

// LeadCottageService - service for lead cottage
type LeadCottageService interface {
	// CreateLeadCottage - creates the lead instance of the cottage in the database
	//
	CreateLeadCottage(
		ctx context.Context,
		lead *LeadCottage,
		caller CallerID,
	) (int64, error)

	// UpdateLeadCottage - updates the lead instance of the cottage in the database
	//
	UpdateLeadCottage(
		ctx context.Context,
		id int64,
		lead *LeadCottage,
		caller CallerID,
	) (*LeadCottage, error)

	// DeleteLeadCottage - deletes the lead instance of the cottage in the database
	//
	DeleteLeadCottage(ctx context.Context,
		id int64,
		caller CallerID,
	) error

	// GetLeadCottage - gets the lead instance of the cottage from the database
	//
	GetLeadCottage(ctx context.Context,
		id int64,
		caller CallerID,
	) (*LeadCottage, error)

	// ListLeadCottage - lists the lead instances of the cottage from the database
	//
	ListLeadCottage(
		ctx context.Context,
		criteria LeadCottageSearchCriteria,
		caller CallerID,
	) ([]*LeadCottage, Total, error)

	// RevokeLeadCottage - revoke lead Cottage
	//
	RevokeLeadCottage(
		ctx context.Context,
		leadID LeadID,
		callerID CallerID,
	) error
}
