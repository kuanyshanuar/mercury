package domain

import (
	"context"
	"gorm.io/plugin/soft_delete"
)

// ManagerID - id of the manager
type ManagerID int64

// Manager - manager struct
type Manager struct {
	// ID - id of user.
	//
	ID ManagerID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// RoleID - id of role.
	//
	RoleID RoleID `json:"role_id" gorm:"column:role_id"`

	// FirstName - first name
	//
	FirstName string `json:"first_name" gorm:"column:first_name"`

	// LastName - last name
	//
	LastName string `json:"last_name" gorm:"column:last_name"`

	// Email - email
	//
	Email string `json:"email" gorm:"column:email"`

	// Phone - phone
	//
	Phone string `json:"phone" gorm:"column:phone"`

	// Image - image
	//
	Image string `json:"image" gorm:"column:image"`

	// Password - password
	//
	Password string `json:"password" gorm:"column:password"`

	// IsVerified - is verified
	//
	IsVerified bool `json:"is_verified" gorm:"column:is_verified"`

	// IsBanned - is banned
	//
	IsBanned *bool `json:"is_banned" gorm:"default:false;column:is_banned"`

	// CreatedAt - created time
	//
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`

	// UpdatedAt - updated time
	//
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// DeletedAt - deleted time
	//
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

// ManagerSearchCriteria - manager search criteria
type ManagerSearchCriteria struct {
	Page  PageRequest
	ID    int64  // filter by id
	Name  string // filter by name
	Email string // filter by email
	Phone string // filter by phone
}

// ManagerRepository - provides access to a storage.
type ManagerRepository interface {
	Create(
		ctx context.Context,
		manager *Manager,
	) (ManagerID, error)

	Update(
		ctx context.Context,
		managerID ManagerID,
		manager *Manager,
	) error

	Get(
		ctx context.Context,
		managerID ManagerID,
	) (*Manager, error)

	List(
		ctx context.Context,
		criteria ManagerSearchCriteria,
	) ([]*Manager, Total, error)

	Delete(
		ctx context.Context,
		managerID ManagerID,
	) error
}

// ManagerService - provides access to a business logic.
type ManagerService interface {
	// CreateManager - creates a new manager
	//
	CreateManager(
		ctx context.Context,
		manager *Manager,
		callerID CallerID,
	) (ManagerID, error)

	// GetManager - returns manager by id
	//
	GetManager(
		ctx context.Context,
		managerID ManagerID,
		callerID CallerID,
	) (*Manager, error)

	// ListManagers - returns a list of managers
	//
	ListManagers(
		ctx context.Context,
		criteria ManagerSearchCriteria,
		callerID CallerID,
	) ([]*Manager, Total, error)

	// UpdateManager - updates manager
	//
	UpdateManager(
		ctx context.Context,
		managerID ManagerID,
		manager *Manager,
		callerID CallerID,
	) error

	// DeleteManager - deletes manager by id
	//
	DeleteManager(
		ctx context.Context,
		managerID ManagerID,
		callerID CallerID,
	) error
}
