package domain

// Roles
const (
	RoleClient  = 1
	RoleBuilder = 2
	RoleManager = 3
	RoleAdmin   = 4
)

// RoleID - id of role
type RoleID int64

// Role - represents roles struct
type Role struct {

	// ID - id of role
	//
	ID RoleID `json:"id"`

	// Name - name
	//
	Name string `json:"name"`
}
