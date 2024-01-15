package domain

import (
	"context"

	"github.com/lib/pq"
)

// ArticlesTableName - table name
// ArticleViewTableName - table for the storage of the users who viewed the article
const (
	ArticlesTableName    = "articles"
	ArticleViewTableName = "article_views"
)

// ArticleID - id of the article
type ArticleID int64

// Article - article struct
type Article struct {
	// ID - id of the article
	//
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Title is a title of the article
	//
	Title string `json:"title" gorm:"column:title"`

	// ShortDescription is a short description at the beginning
	//
	ShortDescription string `json:"short_Description" gorm:"column:short_description"`

	// Slug is a slug of the article on the URL
	//
	Slug string `json:"slug" gorm:"column:slug"`

	// Content is a content of the article
	//
	Content string `json:"content" gorm:"column:content"`

	// ViewsCount is value of the viewers
	//
	ViewsCount int64 `json:"views_count" gorm:"column:views_count"`

	// SourceURL is for the case when the article was taken from other website
	//
	SourceURL string `json:"source_url" gorm:"column:source_url"`

	// AuthorName is the name of the article's author
	//
	AuthorName string `json:"author_name" gorm:"column:author_name"`

	// Images array of the URL of images
	//
	Images pq.StringArray `json:"images" gorm:"type:varchar[];column:images"`

	// IsLikedByMe stores if user has liked the article
	//
	IsLikedByMe bool `json:"liked_by_me" gorm:"-"`

	// IsDislikedByMe stores if user has disliked the article
	//
	IsDislikedByMe bool `json:"dislike_by_me" gorm:"-"`

	// Likes stores the number of likes per article
	//
	Likes int64 `json:"likes" gorm:"<-:update"`

	// Dislikes stores the number of dislikes per article
	//
	Dislikes int64 `json:"dislikes" gorm:"<-:update"`

	// CreatedAt - created timestamp.
	//
	CreatedAt int64 `json:"created_at" gorm:"<-:create;not null;autoCreateTime;column:created_at;"`

	// CreatedBy stores the ID of the moderator
	//
	CreatedBy int64 `json:"created_by" gorm:"column:created_by"`

	// UpdatedAt stores update date of the article
	//
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// UpdatedBy stores the ID of the moderator
	//
	UpdatedBy int64 `json:"updated_by" gorm:"column:updated_by"`

	// DeletedAt stores the deletion date
	//
	DeletedAt int64 `json:"deleted_at" gorm:"column:deleted_at"`

	// DeletedBy stores the ID of the moderator
	//
	DeletedBy int64 `json:"deleted_by" gorm:"column:deleted_by"`
}

// IncreaseLikes - increases likes by one
func (a *Article) IncreaseLikes() {
	a.Likes = a.Likes + 1
}

// DecreaseLikes - decreases likes by one
func (a *Article) DecreaseLikes() {
	if a.Likes == 0 {
		a.Likes = 0
	} else {
		a.Likes = a.Likes - 1
	}
}

// IncreaseDisLikes - increases dislikes by one
func (a *Article) IncreaseDisLikes() {
	a.Dislikes = a.Dislikes + 1
}

// DecreaseDisLikes - decreases dislikes by one
func (a *Article) DecreaseDisLikes() {
	if a.Dislikes == 0 {
		a.Dislikes = 0
	} else {
		a.Dislikes = a.Dislikes - 1
	}
}

// ArticleViews is a struct for views
type ArticleViews struct {
	ID        int64
	ArticleID int64
	UserID    int64
}

// NewsSearchCriteria - news search criteria
type NewsSearchCriteria struct {
	Page  PageRequest
	Sorts []Sort // sorting
	ID    int64  // filter by id
	Title string // filter by Title
	Short string // filter by short description
	Slug  string // filter by Slug
}

// NewsReadRepository - provides access to a read storage
type NewsReadRepository interface {
	// List - lists all the articles according to the search criteria
	//
	List(
		ctx context.Context,
		searchCriteria *NewsSearchCriteria,
	) ([]*Article, Total, error)

	// Get - returns the article given its id
	//
	Get(
		ctx context.Context,
		id int64,
	) (*Article, error)

	// HasDisliked - has disliked
	//
	HasDisliked(
		ctx context.Context,
		articleID int64,
		userID int64,
	) (bool, error)

	// HasLiked - has liked
	//
	HasLiked(
		ctx context.Context,
		articleID int64,
		userID int64,
	) (bool, error)

	// GetViewedOrNot - returns if article was viewed
	//
	GetViewedOrNot(
		ctx context.Context,
		article int64,
		userID int64,
	) (bool, error)
}

// NewsRepository - provides access to a storage
type NewsRepository interface {
	NewsReadRepository

	// Create - creates a new article
	//
	Create(
		ctx context.Context,
		news *Article,
	) (int64, error)

	// Update - updates the article given id
	//
	Update(ctx context.Context,
		id int64,
		news *Article,
	) error

	// Delete - deletes the article given id
	//
	Delete(
		ctx context.Context,
		id int64,
	) error

	// AddLike - adds the like
	//
	AddLike(
		ctx context.Context,
		articleID int64,
		userID int64,
	) error

	// AddDislike - adds the like chosen by user
	//
	AddDislike(
		ctx context.Context,
		articleID int64,
		userID int64,
	) error

	// DeleteLike - deletes the like
	//
	DeleteLike(
		ctx context.Context,
		articleID int64,
		userID int64,
	) error

	// DeleteDislike - deletes the given dislike
	//
	DeleteDislike(
		ctx context.Context,
		articleID int64,
		userID int64,
	) error

	// AddView - increases the view count
	//
	AddView(
		ctx context.Context,
		article int64,
		userID int64,
	) error
}

// NewsService - provides access to a business logic
type NewsService interface {

	// CreateArticle - creates a new article
	//
	CreateArticle(
		ctx context.Context,
		news *Article,
		callerID CallerID,
	) (int64, error)

	// UpdateArticle - updates the article by id
	//
	UpdateArticle(
		ctx context.Context,
		id int64,
		news *Article,
		callerID CallerID,
	) error

	// DeleteArticle - deletes the article by id
	//
	DeleteArticle(
		ctx context.Context,
		id int64,
		callerID CallerID,
	) error

	// ListArticles - returns a list of articles
	//
	ListArticles(
		ctx context.Context,
		searchCriteria *NewsSearchCriteria,
		callerID CallerID,
	) ([]*Article, Total, error)

	// GetArticle - returns article
	//
	GetArticle(
		ctx context.Context,
		id int64,
		callerID CallerID,
	) (*Article, error)

	// AddLike - adds the like
	//
	AddLike(
		ctx context.Context,
		articleID int64,
		callerID CallerID,
	) error

	// DeleteLike - deletes the like
	//
	DeleteLike(
		ctx context.Context,
		articleID int64,
		callerID CallerID,
	) error

	// AddDislike - adds the like
	//
	AddDislike(
		ctx context.Context,
		articleID int64,
		callerID CallerID,
	) error

	// DeleteDislike - deletes the dislike
	//
	DeleteDislike(
		ctx context.Context,
		articleID int64,
		callerID CallerID,
	) error
}
