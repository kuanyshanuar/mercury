package domain

import (
	"context"

	"gorm.io/plugin/soft_delete"
)

// UsersTableName - users table
const (
	UsersTableName = "users"
)

// UserID - id of user
type UserID int64

// User - represents user
type User struct {

	// ID - id of user.
	//
	ID UserID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// RoleID - id of role.
	//
	RoleID RoleID `json:"role_id" gorm:"role_id"`

	// FirstName - first name
	//
	FirstName string `json:"first_name" gorm:"column:first_name"`

	// LastName - last name
	//
	LastName string `json:"last_name" gorm:"column:last_name"`

	// Email - email
	//
	Email string `json:"email" gorm:"column:email"`

	// City - city
	//
	City string `json:"city" gorm:"column:city"`

	// Phone - phone
	//
	Phone string `json:"phone" gorm:"column:phone"`

	// ConsultationPhoneNumber - phone for consultation
	// only for builders
	//
	ConsultationPhoneNumber string `json:"consultation_phone_number" gorm:"column:consultation_phone_number"`

	// Image - image
	// only for builders
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

// UserSearchCriteria - user search criteria
type UserSearchCriteria struct {
	Page  PageRequest
	ID    int64  // search by id
	Name  string // search by name
	Email string // search by email
	Phone string // search by phone
}

// UserReadRepository - provides access to read storage.
type UserReadRepository interface {

	// GetUser - returns user
	//
	GetUser(
		ctx context.Context,
		userID UserID,
	) (*User, error)

	// ListUsers - returns a list of users.
	//
	ListUsers(
		ctx context.Context,
		criteria UserSearchCriteria,
	) ([]*User, Total, error)

	// GetUserByEmailPassword - returns a user with email and password.
	//
	GetUserByEmailPassword(
		ctx context.Context,
		email string,
		password string,
	) (*User, error)

	// GetUserByPhoneNumber - returns user
	//
	GetUserByPhoneNumber(
		ctx context.Context,
		phone string,
	) (*User, error)

	// IsEmailExist - returns is email exist
	//
	IsEmailExist(
		ctx context.Context,
		email string,
	) (bool, error)

	// IsPhoneNumberExist - returns is phone number exist
	//
	IsPhoneNumberExist(
		ctx context.Context,
		phoneNumber string,
	) (bool, error)

	// GetUserByEmail - returns user by the provided email
	//
	GetUserByEmail(
		ctx context.Context,
		email string,
	) (*User, error)
}

// UserRepository - provides access to storage.
type UserRepository interface {
	UserReadRepository

	// CreateUser - creates user.
	//
	CreateUser(
		ctx context.Context,
		user *User,
	) (*User, error)

	// UpdateUser - updates user
	//
	UpdateUser(
		ctx context.Context,
		userID UserID,
		user *User,
	) error

	// DeleteUser - delets user
	//
	DeleteUser(
		ctx context.Context,
		userID UserID,
	) error

	// BanUser - bans user
	//
	BanUser(
		ctx context.Context,
		userID UserID,
	) error

	// MakeUserVerified - make user verified
	//
	MakeUserVerified(
		ctx context.Context,
		userID UserID,
	) error

	// CreateResetPasswordToken - creates reset password token
	//
	CreateResetPasswordToken(
		ctx context.Context,
		resetPasswordToken *ResetPasswordToken,
	) error

	// GetResetPasswordToken - returns reset password token
	//
	GetResetPasswordToken(
		ctx context.Context,
		token string,
	) (*ResetPasswordToken, error)

	// UpdateResetPasswordToken - updates reset password token
	//
	UpdateResetPasswordToken(
		ctx context.Context,
		resetPasswordToken *ResetPasswordToken,
	) error

	// ResetPassword - resets password
	//
	ResetPassword(
		ctx context.Context,
		userID UserID,
		password string,
	) error
}

// UserService - provides access to a business logic
type UserService interface {

	// ListUsers - returns a list of users by criteria
	//
	ListUsers(
		ctx context.Context,
		criteria UserSearchCriteria,
		callerID CallerID,
	) ([]*User, Total, error)

	// UpdateUser - updates user
	//
	UpdateUser(
		ctx context.Context,
		userID UserID,
		user *User,
		callerID CallerID,
	) error
}
