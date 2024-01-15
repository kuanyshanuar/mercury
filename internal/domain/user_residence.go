package domain

import "context"

// FavouriteResidenceTable - name of table in a storage.
const (
	FavouriteResidenceTable = "residences_bookmarks"
)

// UserResidence - represents user residence struct.
type UserResidence struct {

	// UserID - id of user
	//
	UserID int64 `json:"user_id" gorm:"not null;column:user_id"`

	// ResidenceID - id of residence
	//
	ResidenceID int64 `json:"residence_id" gorm:"not null;column:residence_id"`

	// CreatedAt - created timestamp.
	//
	CreatedAt int64 `json:"created_at" gorm:"not null;autoCreateTime;column:created_at"`
}

// FavouriteResidencesSearchCriteria - search criteria
type FavouriteResidencesSearchCriteria struct {
	Page  PageRequest
	Sorts []Sort // sorting
}

// UserResidenceRepository - provides access to storage.
type UserResidenceRepository interface {

	// AddResidenceToFavourites - adds residence to favourites.
	//
	AddResidenceToFavourites(
		ctx context.Context,
		userID UserID,
		residenceID ResidenceID,
	) error

	// IsResidenceExist - returns if residence exist in favourites.
	//
	IsResidenceExist(
		ctx context.Context,
		userID UserID,
		residenceID ResidenceID,
	) (bool, error)

	// DeleteResidenceFromFavourites - deletes residence from favourites.
	//
	DeleteResidenceFromFavourites(
		ctx context.Context,
		userID UserID,
		residenceID ResidenceID,
	) error

	// ListFavouriteResidences - returns residence ids by user id.
	//
	ListFavouriteResidences(
		ctx context.Context,
		userID UserID,
		criteria FavouriteResidencesSearchCriteria,
	) ([]int64, Total, error)
}

// UserResidenceService - provides access to business logic.
type UserResidenceService interface {

	// AddResidenceToFavourites - adds residence to favourites.
	//
	AddResidenceToFavourites(
		ctx context.Context,
		userID UserID,
		residenceID ResidenceID,
		callerID CallerID,
	) error

	// DeleteResidenceFromFavourites - deletes residence from favourites.
	//
	DeleteResidenceFromFavourites(
		ctx context.Context,
		userID UserID,
		residenceID ResidenceID,
		callerID CallerID,
	) error

	// ListFavouriteResidences - returns residence ids by user id.
	//
	ListFavouriteResidences(
		ctx context.Context,
		userID UserID,
		criteria FavouriteResidencesSearchCriteria,
		callerID CallerID,
	) ([]int64, Total, error)
}
