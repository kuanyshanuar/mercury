package domain

import "context"

// Names of table
const (
	PermissionsTable        = "permissions"
	RolesToPermissionsTable = "roles_to_permissions"
)

// PermissionID - id of the permission
type PermissionID int64

// Permission - permission struct
type Permission struct {
	// ID - id of permission
	//
	ID PermissionID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// EndpointName - name of endpoint
	//
	EndpointName string `json:"endpoint_name" gorm:"column:endpoint_name"`

	// Action - action
	//
	Action string `json:"action" gorm:"column:action"`

	// CrudType - crud type
	//
	CrudType string `json:"crud_type" gorm:"column:crud_type"`

	// IsActive - is active
	//
	IsActive bool `json:"is_active" gorm:"column:is_active"`

	// CreatedAt - created timestamp.
	//
	CreatedAt int64 `json:"created_at" gorm:"not null;autoCreateTime;column:created_at"`

	// UpdatedAt - updated timestamp.
	//
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}

// PermissionsReadRepository provides access to read storage.
type PermissionsReadRepository interface {

	// ListPermissionsByRole - returns user's permissions by role id.
	//
	ListPermissionsByRole(
		ctx context.Context,
		roleID int64,
	) ([]*Permission, error)

	// GetPermissionByEndpoint - returns permission id by endpoint name.
	//
	GetPermissionByEndpoint(
		ctx context.Context,
		endpointName string,
	) (*Permission, error)

	// IsPermissionAllowed - returns if permission allowed to role id.
	//
	IsPermissionAllowed(
		ctx context.Context,
		roleID RoleID,
		permissionID PermissionID,
	) (bool, error)
}

// PermissionsRepository provides access to read storage.
type PermissionsRepository interface {
	PermissionsReadRepository

	// CreatePermission - creates a new permission in a storage.
	//
	CreatePermission(
		ctx context.Context,
		permission *Permission,
	) (PermissionID, error)
}

// PermissionsService provides access to business logic.
type PermissionsService interface {

	// CreatePermission - creates a new permission.
	//
	CreatePermission(
		ctx context.Context,
		permission *Permission,
		callerID CallerID,
	) (PermissionID, error)

	// List - returns user's permissions.
	//
	List(
		ctx context.Context,
		callerID CallerID,
	) ([]*Permission, error)

	// Allow - allows user.
	//
	Allow(
		ctx context.Context,
		permissionKey string,
		userID UserID,
		roleID RoleID,
		callerID CallerID,
	) (bool, error)
}
