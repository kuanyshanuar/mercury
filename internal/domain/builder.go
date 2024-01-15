package domain

import (
	"context"

	"gorm.io/plugin/soft_delete"
)

// BuilderID - id of the builder
type BuilderID int64

// Builder - represents builder struct
type Builder struct {

	// ID - id of user.
	//
	ID BuilderID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

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
	IsBanned *bool `json:"is_banned" gorm:"column:is_banned"`

	// IsFavourite - is favourite
	//
	IsFavourite bool `json:"is_favourite" gorm:"column:is_favourite"`

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

// BuilderSearchCriteria - builder search criteria
type BuilderSearchCriteria struct {
	Page   PageRequest
	UserID int64
	ID     int64  // filter by id
	Name   string // filter by name
	Email  string // filter by email
	Phone  string // filter by phone
}

// BuilderRepository - provides access to a storage.
type BuilderRepository interface {
	Create(
		ctx context.Context,
		builder *Builder,
	) (BuilderID, error)

	Update(
		ctx context.Context,
		builderID BuilderID,
		builder *Builder,
	) error

	Get(
		ctx context.Context,
		builderID BuilderID,
	) (*Builder, error)

	List(
		ctx context.Context,
		criteria BuilderSearchCriteria,
	) ([]*Builder, Total, error)

	Delete(
		ctx context.Context,
		builderID BuilderID,
	) error

	// IsFavouriteBuilder - returns is favourite builder
	//
	IsFavouriteBuilder(
		ctx context.Context,
		builder BuilderID,
		userID UserID,
	) (bool, error)
}

// BuilderService - provides access to a business logic.
type BuilderService interface {

	// CreateBuilder - creates builder.
	//
	CreateBuilder(
		ctx context.Context,
		user *Builder,
		callerID CallerID,
	) (BuilderID, error)

	// UpdateBuilder - updates builder
	//
	UpdateBuilder(
		ctx context.Context,
		builderID BuilderID,
		builder *Builder,
		callerID CallerID,
	) error

	// GetBuilder - returns the builder by id.
	//
	GetBuilder(
		ctx context.Context,
		builderID BuilderID,
		callerID CallerID,
	) (*Builder, error)

	// ListBuilders - returns a list of builders by criteria.
	//
	ListBuilders(
		ctx context.Context,
		criteria BuilderSearchCriteria,
		callerID CallerID,
	) ([]*Builder, Total, error)

	// DeleteBuilder - deletes a builder by id
	//
	DeleteBuilder(
		ctx context.Context,
		builderID BuilderID,
		callerID CallerID,
	) error
}
