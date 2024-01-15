package domain

import "context"

// CommentsTableName - comments table name
const (
	CommentsTableName = "news_comments"
)

// NewsComments - new comments struct
type NewsComments struct {
	//ID is an id of the comment
	//
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	//UserID is id of a user that is author of the comment
	//
	UserID int64 `json:"user_id" gorm:"foreignKey:user_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	//ArticleID is id of the article of the comment
	//
	ArticleID int64 `json:"article_id" gorm:"foreignKey:article_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	//Content is the content of the message in the comment
	//
	Content string `json:"content" gorm:"column:content"`

	//CreatedAt is the date of the creation of the comment
	//
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`

	//UpdatedAt is the date of the update of the comment
	//
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	//DeletedAt is the date of the deletion of the comment
	//
	DeletedAt int64 `json:"deleted_at" gorm:"column:deleted_at"`
}

// NewsCommentsRepo - provides access to a storage
type NewsCommentsRepo interface {

	// ListComments lists the comments given the article id
	//
	ListComments(ctx context.Context, articleID int64) ([]*NewsComments, int64, error)

	// GetUserComments lists the comments given the user id
	//
	GetUserComments(ctx context.Context, userID int64) ([]*NewsComments, int64, error)

	// AddComment writes the comment to the database
	//
	AddComment(ctx context.Context, comment *NewsComments) (int64, error)

	// DeleteComment deletes the comment from the database
	//
	DeleteComment(ctx context.Context, commentID int64) error

	// UpdateComment updates the comment inside the database given the commentID , comment content
	//
	UpdateComment(ctx context.Context, commentID int64, comment *NewsComments) error

	// GetComment gets the comment from the database given the commentID
	//
	GetComment(ctx context.Context, commentID int64) (*NewsComments, error)
}

// NewsCommentsService - provides access to a business logic
type NewsCommentsService interface {
	// ListComments lists the comments given the article id
	//
	ListComments(ctx context.Context, articleID int64, callerID CallerID) ([]*NewsComments, int64, error)

	// GetUserComments lists the comments given the user id
	//
	GetUserComments(ctx context.Context, userID int64, callerID CallerID) ([]*NewsComments, int64, error)

	// AddComment writes the comment to the database
	//
	AddComment(ctx context.Context, comment *NewsComments, callerID CallerID) (int64, error)

	// DeleteComment deletes the comment from the database
	//
	DeleteComment(ctx context.Context, commentID int64, callerID CallerID) error

	// UpdateComment updates the comment inside the database given the commentID , comment content
	//
	UpdateComment(ctx context.Context, commentID int64, comment *NewsComments, callerID CallerID) error

	// GetComment gets the comment from the database given the commentID
	//
	GetComment(ctx context.Context, commentID int64, callerID CallerID) (*NewsComments, error)
}
