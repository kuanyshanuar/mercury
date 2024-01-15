package domain

// DislikesTableName - table name
const (
	DislikesTableName = "dislikes"
)

// Dislike is a table to store dislikes
type Dislike struct {
	// ArticleID - id of the article
	//
	ArticleID int64 `json:"article_id" gorm:"constraint:OnDelete:CASCADE;"`

	// UserID - id of the user
	//
	UserID int64 `json:"user_id" gorm:"constraint:OnDelete:CASCADE;"`
}
