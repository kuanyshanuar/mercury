package domain

import "context"

// CrmService - provides access to business logic
type CrmService interface {
	// SendResidenceContactDetail - sends residence contact details to crm
	//
	SendResidenceContactDetail(
		ctx context.Context,
		content ResidenceContactDetailContent,
		callerID CallerID,
	) error
}
