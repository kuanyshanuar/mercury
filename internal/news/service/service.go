package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.NewsRepository
}

// NewService - creates a new service with necessary dependencies
func NewService(
	repository domain.NewsRepository,
	logger log.Logger,
) domain.NewsService {
	var service domain.NewsService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(repository domain.NewsRepository) domain.NewsService {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateArticle(
	ctx context.Context,
	article *domain.Article,
	callerID domain.CallerID,
) (int64, error) {

	// Validates input
	//
	err := validateArticle(article)
	if err != nil {
		return 0, err
	}

	{
		article.CreatedBy = callerID.UserID
	}

	return s.repository.Create(ctx, article)
}

func (s *service) UpdateArticle(
	ctx context.Context,
	id int64,
	article *domain.Article,
	callerID domain.CallerID,
) error {

	// Validates input
	//
	err := validateArticle(article)
	if err != nil {
		return err
	}

	{
		article.UpdatedBy = callerID.UserID
	}

	return s.repository.Update(ctx, id, article)
}

func (s *service) DeleteArticle(
	ctx context.Context,
	id int64,
	callerID domain.CallerID,
) error {

	// Validates input
	//
	if id <= 0 {
		return errors.NewErrInvalidArgument("required id")
	}

	return s.repository.Delete(ctx, id)
}

func (s *service) ListArticles(
	ctx context.Context,
	criteria *domain.NewsSearchCriteria,
	callerID domain.CallerID,
) ([]*domain.Article, domain.Total, error) {

	// Validates input
	//
	if criteria == nil {
		return nil, 0, errors.NewErrInvalidArgument("required criteria")
	}

	return s.repository.List(ctx, criteria)
}

func (s *service) GetArticle(
	ctx context.Context,
	id int64,
	callerID domain.CallerID,
) (*domain.Article, error) {

	// Validates input
	//
	if id <= 0 {
		return nil, errors.NewErrInvalidArgument("required id")
	}

	// Gets the article from the database
	//
	article, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Adds the ViewCount to the database
	//
	err = s.AddView(ctx, id, callerID)
	if err != nil {
		return nil, err
	}

	// Checks whether user has liked or disliked and then assigns the true labels
	//
	if callerID.UserID > 0 {
		hasLiked, err := s.repository.HasLiked(ctx, id, callerID.UserID)
		if err != nil {
			return nil, err
		}

		hasDisliked, err := s.repository.HasDisliked(ctx, id, callerID.UserID)
		if err != nil {
			return nil, err
		}

		article.IsLikedByMe = hasLiked
		article.IsDislikedByMe = hasDisliked
	}

	return article, nil
}

func (s *service) AddLike(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) error {

	// Validates input
	//
	if articleID <= 0 {
		return errors.NewErrInvalidArgument("required article id")
	}
	if callerID.UserID <= 0 {
		return errors.NewErrInvalidArgument("user id required")
	}

	// Get article
	//
	article, err := s.repository.Get(ctx, articleID)
	if err != nil {
		return err
	}

	if article == nil {
		return errors.NewErrNotFound(
			fmt.Sprintf("article by id %d not found", articleID),
		)
	}

	// Check is user has disliked
	//
	hasDisliked, err := s.repository.HasDisliked(ctx, articleID, callerID.UserID)
	if err != nil {
		return err
	}

	// If the article it has disliked then remove the dislike
	//
	if hasDisliked {
		err = s.repository.DeleteDislike(ctx, articleID, callerID.UserID)
		if err != nil {
			return err
		}

		article.DecreaseDisLikes()
	}

	article.IncreaseLikes()

	err = s.UpdateArticle(ctx, articleID, article, callerID)
	if err != nil {
		return err
	}

	return s.repository.AddLike(ctx, articleID, callerID.UserID)
}

func (s *service) AddDislike(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) error {

	// Validates inputs
	//
	if articleID <= 0 {
		return errors.NewErrInvalidArgument("required article id")
	}
	if callerID.UserID <= 0 {
		return errors.NewErrInvalidArgument("user id required")
	}

	article, err := s.repository.Get(ctx, articleID)
	if err != nil {
		return err
	}
	if article == nil {
		return errors.NewErrNotFound(
			fmt.Sprintf("article by id %d not found", articleID),
		)
	}

	// Check is user has liked
	//
	hasLiked, err := s.repository.HasLiked(ctx, articleID, callerID.UserID)
	if err != nil {
		return err
	}

	fmt.Println(hasLiked)

	// If the article it has liked then remove the like
	//
	if hasLiked {
		// Delete like
		//
		err = s.repository.DeleteLike(ctx, articleID, callerID.UserID)
		if err != nil {
			return err
		}

		// Decrease likes counter
		//
		article.DecreaseLikes()
	}

	// Add like
	//
	article.IncreaseDisLikes()

	err = s.UpdateArticle(ctx, articleID, article, callerID)
	if err != nil {
		return err
	}

	return s.repository.AddDislike(ctx, articleID, callerID.UserID)
}

func (s *service) DeleteLike(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) error {

	// Validates inputs
	//
	if articleID <= 0 {
		return errors.NewErrInvalidArgument("required article id")
	}
	if callerID.UserID <= 0 {
		return errors.NewErrInvalidArgument("user id required")
	}

	// Check if article exists
	//
	article, err := s.repository.Get(ctx, articleID)
	if err != nil {
		return err
	}
	if article == nil {
		return errors.NewErrNotFound(
			fmt.Sprintf("article by id %d not found", articleID),
		)
	}

	// Decrease likes
	article.DecreaseLikes()

	// Updates article
	//
	err = s.UpdateArticle(ctx, articleID, article, callerID)
	if err != nil {
		return err
	}

	return s.repository.DeleteLike(ctx, articleID, callerID.UserID)
}

func (s *service) DeleteDislike(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) error {

	// Validates inputs
	//
	if articleID <= 0 {
		return errors.NewErrInvalidArgument("required article id")
	}
	if callerID.UserID <= 0 {
		return errors.NewErrInvalidArgument("user id required")
	}

	// Check if article exists
	//
	article, err := s.repository.Get(ctx, articleID)
	if err != nil {
		return err
	}
	if article == nil {
		return errors.NewErrNotFound(
			fmt.Sprintf("article by id %d not found", articleID),
		)
	}

	article.DecreaseDisLikes()

	// Updates article
	//
	err = s.UpdateArticle(ctx, articleID, article, callerID)
	if err != nil {
		return err
	}

	return s.repository.DeleteDislike(ctx, articleID, callerID.UserID)
}

func (s *service) AddView(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) error {
	hasViewed, err := s.repository.GetViewedOrNot(ctx, articleID, callerID.UserID)
	if !hasViewed {
		return s.repository.AddView(ctx, articleID, callerID.UserID)
	}

	return err
}

func validateArticle(article *domain.Article) error {
	if len(article.Content) == 0 {
		return errors.NewErrInvalidArgument("required content")
	}
	if len(article.Title) == 0 {
		return errors.NewErrInvalidArgument("required title")
	}

	return nil
}
