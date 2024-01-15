package domain

import (
	"context"
)

// ResetPasswordTokenTableName - table name
const (
	ResetPasswordTokenTableName = "reset_password_tokens"
)

// ResetPasswordToken - represents reset password token
type ResetPasswordToken struct {
	// UserID - id of the user
	//
	UserID UserID `json:"user_id" gorm:"column:user_id"`

	// Token - token
	//
	Token string `json:"token" gorm:"column:token"`

	// Confirmed - confirmed
	//
	Confirmed bool `json:"confirmed" gorm:"column:confirmed"`

	// CreatedAt - created timestamp.
	//
	CreatedAt int64 `json:"created_at" gorm:"not null;autoCreateTime;column:created_at"`

	// UpdatedAt - updated timestamp.
	//
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}

// IdentityManagerRedisRepository - provides access to redis storage.
type IdentityManagerRedisRepository interface {

	// Get - returns user from cache by key.
	//
	Get(
		ctx context.Context,
		key string,
	) (*User, error)

	// Set - sets an user to cache with specific key.
	//
	Set(
		ctx context.Context,
		key string,
		value *User,
		seconds int,
	) error

	// Delete - deletes an user from cache by key.
	//
	Delete(
		ctx context.Context,
		key string,
	) error
}

// IdentityManagerService - provides access to a business logic.
type IdentityManagerService interface {

	// CreateUser - creates user.
	//
	CreateUser(
		ctx context.Context,
		user *User,
	) error

	// ValidateUser - validates user.
	//
	ValidateUser(
		ctx context.Context,
		email string,
		password string,
	) (*User, error)

	// ValidateCode - validates sms code.
	//
	ValidateCode(
		ctx context.Context,
		code string,
	) (UserID, RoleID, error)

	// SendResetPasswordToken - sends reset password token
	//
	SendResetPasswordToken(
		ctx context.Context,
		email string,
	) error

	// ResetPassword - resets password
	//
	ResetPassword(
		ctx context.Context,
		token string,
		newPassword string,
	) error
}
