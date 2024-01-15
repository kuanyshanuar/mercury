package domain

import "context"

// CronContactsService - provides access to a business logic
type CronContactsService interface {

	// SendContactsToCRM - sends contact details to crm
	//
	SendContactsToCRM(
		ctx context.Context,
		callerID CallerID,
	) error
}
