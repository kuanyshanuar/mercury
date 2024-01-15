package domain

import "context"

// Table names
const (
	ContactDetailsTableName          = "contact_details"
	ResidenceContactDetailsTableName = "residences_contact_details"
)

// ContactDetails - represents contact details struct
type ContactDetails struct {
	// ID - logical id
	//
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// FullName - full name of user
	//
	FullName string `json:"full_name" gorm:"column:full_name"`

	// Phone - phone
	//
	Phone string `json:"phone" gorm:"column:phone"`

	// Message - message
	//
	Message string `json:"message" gorm:"column:message"`
}

// ResidenceContactDetails - residence contact details struct.
type ResidenceContactDetails struct {
	// ID - logical id
	//
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// ResidenceID - residence id
	//
	ResidenceID int64 `json:"residence_id" gorm:"column:residence_id"`

	// FullName - full name of user
	//
	FullName string `json:"full_name" gorm:"column:full_name"`

	// Phone - phone
	//
	Phone string `json:"phone" gorm:"column:phone"`

	// IsDelivered - is delivered
	//
	IsDelivered bool `json:"is_delivered" gorm:"column:is_delivered"`

	// CreatedAt - created timestamp.
	//
	CreatedAt int64 `json:"created_at" gorm:"<-:create;not null;autoCreateTime;column:created_at;"`
}

// ResidenceContactDetailsSearchCriteria - search criteria
type ResidenceContactDetailsSearchCriteria struct {
	Page        PageRequest
	IsDelivered int64 // filter by is_delivered 0 - all 1 - is delivered 2 - not delivered
	ResidenceID int64
	Phone       string
	FromTime    int64
	ToTime      int64
}

// ContactDetailsRepository - provides access to a storage.
type ContactDetailsRepository interface {
	// CreateContactDetails - creates contact details
	//
	CreateContactDetails(
		ctx context.Context,
		contactDetails *ContactDetails,
	) error

	// ListContactDetails - returns a list of contact details.
	//
	ListContactDetails(
		ctx context.Context,
	) ([]*ContactDetails, error)

	// CreateResidenceContactDetails - creates residence contact details
	//
	CreateResidenceContactDetails(
		ctx context.Context,
		contactDetails *ResidenceContactDetails,
	) error

	// ListResidenceContactDetails - returns a list of  residence contact details.
	//
	ListResidenceContactDetails(
		ctx context.Context,
		criteria ResidenceContactDetailsSearchCriteria,
	) ([]*ResidenceContactDetails, Total, error)

	// MarkAsDelivered - marks as delivered
	//
	MarkAsDelivered(
		ctx context.Context,
		contactID int64,
	) error
}

// ContactDetailsService - provides access to a business logic.
type ContactDetailsService interface {
	// CreateContactDetails - creates contact details
	//
	CreateContactDetails(
		ctx context.Context,
		contactDetails *ContactDetails,
		callerID CallerID,
	) error

	// CreateResidenceContactDetails - creates residence contact details
	//
	CreateResidenceContactDetails(
		ctx context.Context,
		contactDetails *ResidenceContactDetails,
		callerID CallerID,
	) error

	// ListContactDetails - returns a list of contact details.
	//
	ListContactDetails(
		ctx context.Context,
		callerID CallerID,
	) ([]*ContactDetails, error)

	// ListResidenceContactDetails - returns a list of  residence contact details.
	//
	ListResidenceContactDetails(
		ctx context.Context,
		criteria ResidenceContactDetailsSearchCriteria,
		callerID CallerID,
	) ([]*ResidenceContactDetails, Total, error)

	// MarkAsDelivered - marks as delivered
	//
	MarkAsDelivered(
		ctx context.Context,
		contactID int64,
		callerID CallerID,
	) error
}
