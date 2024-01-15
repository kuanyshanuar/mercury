package domain

import "context"

// UserBuildersTableName - table name
const (
	UserBuildersTableName = "subscribers"
)

// UserBuilder - represents user builders
type UserBuilder struct {

	// BuilderID - builder id
	//
	BuilderID int64 `json:"builder_id" gorm:"column:builder_id"`

	// SubscriberID - subscriber id
	//
	SubscriberID int64 `json:"subscriber_id" gorm:"column:subscriber_id"`
}

// UserBuilderSearchCriteria - search criteria
type UserBuilderSearchCriteria struct {
	Page  PageRequest
	Sorts []Sort // sorting
}

// UserBuilderRepository - provides access to a storage
type UserBuilderRepository interface {

	// Subscribe - subscribe builder
	//
	Subscribe(
		ctx context.Context,
		subscriberID int64,
		builderID int64,
	) (err error)

	// Unsubscribe - unsubscribe builder
	//
	Unsubscribe(
		ctx context.Context,
		subscriberID int64,
		builderID int64,
	) (err error)

	// ListBuilders - list subscribers
	//
	ListBuilders(
		ctx context.Context,
		subscriberID int64,
		criteria UserBuilderSearchCriteria,
	) (builders []*Builder, total Total, err error)
}

// UserBuilderService - provides access to a business logic
type UserBuilderService interface {

	// Subscribe - subscribe builder
	//
	Subscribe(
		ctx context.Context,
		subscriberID int64,
		builderID int64,
		callerID CallerID,
	) (err error)

	// Unsubscribe - unsubscribe builder
	//
	Unsubscribe(
		ctx context.Context,
		subscriberID int64,
		builderID int64,
		callerID CallerID,
	) (err error)

	// ListBuilders - list subscribers
	//
	ListBuilders(
		ctx context.Context,
		subscriberID int64,
		criteria UserBuilderSearchCriteria,
		callerID CallerID,
	) (builders []*Builder, total Total, err error)
}
