package domain

import "context"

const (
	// UserCottageTableName - database table name for favourite cottages
	UserCottageTableName = "cottages_bookmarks"
)

// UserCottage - unit of the bookmark
type UserCottage struct {
	// CottageID - is the id of the cottage in the user-cottage connection instance
	//
	CottageID int64 `json:"cottage_id" gorm:"column:cottage_id"`
	// UserID - is the id of the user in the user-cottage connection instance
	//
	UserID int64 `json:"user_id" gorm:"column:user_id"`
	// CreatedAt is the date of creation of the user-cottage connection instance
	//
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
}

// FavouriteCottagesSearchCriteria - search criteria
type FavouriteCottagesSearchCriteria struct {
	Page  PageRequest
	Sorts []Sort // sorting
}

// UserCottageRepository repository for favourite cottages
type UserCottageRepository interface {
	// AddFavouriteCottage - adds the cottage to the user's favourite list
	//
	AddFavouriteCottage(
		ctx context.Context,
		userID int64,
		cottageID int64,
	) error
	// IsCottageExists - returns if the cottage exists in the favourites of the user's choices
	//
	IsCottageExists(
		ctx context.Context,
		userID int64,
		cottageID int64,
	) (bool, error)
	// DeleteFavouriteCottage - deletes the cottage from the user's favourite list
	//
	DeleteFavouriteCottage(
		ctx context.Context,
		userID int64,
		cottageID int64,
	) error
	// ListFavouriteCottages - lists the cottages from the user's favourite list
	//
	ListFavouriteCottages(
		ctx context.Context,
		userID int64,
		criteria FavouriteCottagesSearchCriteria,
	) ([]int64, Total, error)
}

// UserCottageService repository for favourite cottages
type UserCottageService interface {
	// AddCottageToFavourites - adds Cottage to favourites.
	//
	AddCottageToFavourites(
		ctx context.Context,
		userID int64,
		CottageID int64,
		callerID CallerID,
	) error

	// DeleteCottageFromFavourites - deletes Cottage from favourites.
	//
	DeleteCottageFromFavourites(
		ctx context.Context,
		userID int64,
		CottageID int64,
		callerID CallerID,
	) error

	// ListFavouriteCottages - returns Cottage ids by user id.
	//
	ListFavouriteCottages(
		ctx context.Context,
		userID int64,
		criteria FavouriteCottagesSearchCriteria,
		callerID CallerID,
	) ([]int64, Total, error)
}
