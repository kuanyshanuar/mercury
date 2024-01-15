package domain

import "context"

// ResetMailPasswordEmailContent - mail struct
type ResetMailPasswordEmailContent struct {
	// Token - token
	//
	Token string `json:"token"`
}

// ResidenceContactDetailContent - mail struct
type ResidenceContactDetailContent struct {

	// Residence name
	//
	ResidenceName string `json:"residence_name"`

	// Slug - residence slug
	//
	Slug string `json:"slug"`

	// Full name
	//
	FullName string `json:"full_name"`

	// Phone
	//
	Phone string `json:"phone"`
}

// ContactDetailContent - content
type ContactDetailContent struct {
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

// MailService - provides access to business logic
type MailService interface {
	// SendResetPasswordEmail - sends reset password email.
	//
	SendResetPasswordEmail(
		ctx context.Context,
		email string,
		content ResetMailPasswordEmailContent,
	) error

	// SendResidenceContactDetailEmail - sends residence contact detail email
	//
	SendResidenceContactDetailEmail(
		ctx context.Context,
		email string,
		content ResidenceContactDetailContent,
	) error

	// SendContactDetailEmail - sends residence contact detail email
	//
	SendContactDetailEmail(
		ctx context.Context,
		email string,
		content ContactDetailContent,
	) error
}
