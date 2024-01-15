package repository

import (
	"context"
	"strings"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository creates the new repository for News Domain
func NewRepository(db *gorm.DB) domain.NewsRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(
	ctx context.Context,
	article *domain.Article,
) (int64, error) {
	err := r.db.WithContext(ctx).
		Table(domain.ArticlesTableName).
		Create(article).
		Error
	if err != nil {
		return 0, err
	}

	return article.ID, err
}

func (r *repository) Update(
	ctx context.Context,
	id int64,
	article *domain.Article,
) error {
	err := r.db.WithContext(ctx).
		Table(domain.ArticlesTableName).
		Where("id = ?", id).
		Save(&article).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the article for articles table
func (r *repository) Delete(
	ctx context.Context,
	id int64,
) error {
	article := new(domain.Article)
	err := r.db.WithContext(ctx).
		Table(domain.ArticlesTableName).
		Where("id = ?", id).
		Delete(&article).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) List(
	ctx context.Context,
	criteria *domain.NewsSearchCriteria,
) ([]*domain.Article, domain.Total, error) {
	var (
		db         = r.db
		article    = new([]*domain.Article)
		totalCount int64
	)

	if criteria.ID > 0 {
		db = db.Where("id = ?", criteria.ID)
	}
	if len(strings.ToLower(criteria.Title)) > 0 {
		db = db.Where("LOWER(title) LIKE LOWER (?)", criteria.Title)
	}

	order := helpers.PrepareOrder(criteria.Sorts)
	if len(order) == 0 {
		order = `"created_at" DESC`
	}
	db = db.Order(order)

	err := db.
		Table(domain.ArticlesTableName).
		Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.ArticlesTableName).
		Find(&article).Error
	return *article, domain.Total(totalCount), err

}

func (r *repository) Get(
	ctx context.Context,
	id int64,
) (*domain.Article, error) {
	var article *domain.Article
	err := r.db.WithContext(ctx).
		Table(domain.ArticlesTableName).
		Where("id = ?", id).
		First(&article).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

func (r *repository) AddLike(
	ctx context.Context,
	articleID int64,
	userID int64,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.LikesTableName).
		Create(&domain.Like{
			UserID:    userID,
			ArticleID: articleID,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteLike(
	ctx context.Context,
	articleID int64,
	userID int64,
) error {
	var like = new(domain.Like)
	err := r.db.
		WithContext(ctx).
		Table(domain.LikesTableName).
		Where("article_id = ? AND user_id = ?", articleID, userID).
		Delete(&like).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) AddDislike(
	ctx context.Context,
	articleID int64,
	userID int64,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.DislikesTableName).
		Create(&domain.Dislike{
			UserID:    userID,
			ArticleID: articleID,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteDislike(
	ctx context.Context,
	articleID int64,
	userID int64,
) error {
	var dislike = new(domain.Dislike)
	err := r.db.
		WithContext(ctx).
		Table(domain.DislikesTableName).
		Where("article_id = ? AND user_id = ?", articleID, userID).
		Delete(&dislike).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) HasLiked(
	ctx context.Context,
	articleID int64,
	userID int64,
) (bool, error) {
	var hasLiked bool
	err := r.db.
		WithContext(ctx).
		Table(domain.LikesTableName).
		Select("count(user_id) > 0").
		Where("article_id = ? AND user_id = ?", articleID, userID).
		Find(&hasLiked).Error
	if err != nil {
		return false, err
	}

	return hasLiked, nil
}

func (r *repository) HasDisliked(
	ctx context.Context,
	articleID int64,
	userID int64,
) (bool, error) {
	var hasDisliked bool
	err := r.db.
		WithContext(ctx).
		Table(domain.DislikesTableName).
		Select("count(user_id) > 0").
		Where("article_id = ? AND user_id = ?", articleID, userID).
		Find(&hasDisliked).Error
	if err != nil {
		return false, err
	}

	return hasDisliked, nil
}

func (r *repository) GetViewedOrNot(
	ctx context.Context,
	articleID int64,
	userID int64,
) (bool, error) {
	hasViewed := false
	err := r.db.
		WithContext(ctx).
		Table(domain.ArticleViewTableName).
		Select("count(user_id) > 0").
		Where("article_id = ? AND user_id = ?", articleID, userID).
		Find(&hasViewed).Error
	if err != nil {
		return false, err
	}
	return hasViewed, nil
}

func (r *repository) AddView(
	ctx context.Context,
	articleID int64,
	userID int64,
) error {
	viewInsert := &domain.ArticleViews{
		ArticleID: articleID,
		UserID:    userID,
	}
	err := r.db.
		WithContext(ctx).
		Table(domain.ArticleViewTableName).
		Create(viewInsert).
		Error
	if err != nil {
		return err
	}

	err = r.db.
		Exec("UPDATE articles SET views_count = views_count + 1 WHERE id = ?", articleID).
		Error
	if err != nil {
		return err
	}

	return nil
}
