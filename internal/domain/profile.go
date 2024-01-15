package domain

import "context"

// Profile - represents profile structs
type Profile struct {

	// ID - id of user.
	//
	ID UserID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// RoleID - if of the role
	//
	RoleID RoleID `json:"role_id" gorm:"column:role_id"`

	// FirstName - first name
	//
	FirstName string `json:"first_name" gorm:"column:first_name"`

	// LastName - last name
	//
	LastName string `json:"last_name" gorm:"column:last_name"`

	// City - city
	//
	City string `json:"city" gorm:"column:city"`

	// Phone - phone
	//
	Phone string `json:"phone" gorm:"column:phone"`

	// Email - email
	//
	Email string `json:"email" gorm:"column:email"`

	// Password - password
	//
	Password string `json:"password" gorm:"password"`

	// ConsultationPhoneNumber - phone for consultation
	// only for builders
	//
	ConsultationPhoneNumber string `json:"consultation_phone_number" gorm:"column:consultation_phone_number"`

	// CreatedAt - created timestamp
	//
	CreatedAt int64 `json:"created_at" gorm:"created_at"`
}

// ProfileRepository - provides access to storage.
type ProfileRepository interface {
	// GetProfile - returns profile by user id
	//
	GetProfile(
		ctx context.Context,
		userID UserID,
	) (*Profile, error)

	// UpdateProfile - updates profile
	//
	UpdateProfile(
		ctx context.Context,
		userID UserID,
		profile *Profile,
	) error

	// IsEmailExist - is email exists
	//
	IsEmailExist(
		ctx context.Context,
		email string,
	) (bool, error)

	// IsPhoneExist - is phone exists
	//
	IsPhoneExist(
		ctx context.Context,
		phone string,
	) (bool, error)
}

// ProfileRedisRepository - provides access to a storage
type ProfileRedisRepository interface {
	// Get - returns profile from cache by key.
	//
	Get(
		ctx context.Context,
		key string,
	) (*Profile, error)

	// Set - sets an profile to cache with specific key.
	//
	Set(
		ctx context.Context,
		key string,
		value *Profile,
		seconds int,
	) error

	// Delete - deletes an profile from cache by key.
	//
	Delete(
		ctx context.Context,
		key string,
	) error
}

// ProfileService - provides access to business logic
type ProfileService interface {
	// GetProfile - returns profile
	//
	GetProfile(
		ctx context.Context,
		userID UserID,
		callerID CallerID,
	) (*Profile, error)

	// UpdateProfile - updates profile
	//
	UpdateProfile(
		ctx context.Context,
		userID UserID,
		user *Profile,
		callerID CallerID,
	) error

	// ValidatePhone - validates phone
	//
	ValidatePhone(
		ctx context.Context,
		code string,
		callerID CallerID,
	) error
}
