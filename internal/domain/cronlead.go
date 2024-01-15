package domain

import "context"

// CronLeadService - provides access to a business logic
type CronLeadService interface {

	// RevokeExpiredLeadResidences - revoke expired lead residences
	//
	RevokeExpiredLeadResidences(
		ctx context.Context,
		callerID CallerID,
	) error

	// RevokeExpiredLeadCottages - revoke expired lead residences
	//
	RevokeExpiredLeadCottages(
		ctx context.Context,
		callerID CallerID,
	) error

	// RevokeExpiredLeadBuilders - revoke expired lead builders
	//
	RevokeExpiredLeadBuilders(
		ctx context.Context,
		callerID CallerID,
	) error
}
