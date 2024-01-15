package domain

import "context"

// AmoApplication - amo application
type AmoApplication struct {
	// FullName - full name
	//
	FullName string `json:"full_name"`

	// Phone - phone
	//
	Phone string `json:"phone"`

	// ResidenceDetails - residence details
	//
	ResidenceDetails *Residence `json:"residence"`
}

// AmoCrmService - provides access to business logic
type AmoCrmService interface {
	// CreateApplication - create aaplication in crm service
	//
	CreateApplication(
		ctx context.Context,
		application *AmoApplication,
	) error
}
