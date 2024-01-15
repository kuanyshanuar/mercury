package domain

// LikesTableName - table name
const (
	LikesTableName = "likes"
)

// Like - like struct
type Like struct {
	// ArticleID - id of the Article
	//
	ArticleID int64 `json:"article_id" gorm:"constraint:OnDelete:CASCADE;"`

	// UserID - id of the user liked the article
	//
	UserID int64 `json:"user_id" gorm:"constraint:OnDelete:CASCADE;"`
}
