package transport

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

// Endpoints collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Endpoints struct {
	CreateArticleEndpoint endpoint.Endpoint
	ListArticlesEndpoint  endpoint.Endpoint
	DeleteArticleEndpoint endpoint.Endpoint
	GetArticleEndpoint    endpoint.Endpoint
	UpdateArticleEndpoint endpoint.Endpoint
	AddLikeEndpoint       endpoint.Endpoint
	DeleteLikeEndpoint    endpoint.Endpoint
	AddDislikeEndpoint    endpoint.Endpoint
	DeleteDislikeEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.NewsService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(service domain.NewsService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "News", logKey)
	}

	return Endpoints{
		CreateArticleEndpoint: factory(MakeCreateArticleEndpoint, "CreateArticle"),
		ListArticlesEndpoint:  factory(MakeListArticlesEndpoint, "ListArticles"),
		DeleteArticleEndpoint: factory(MakeDeleteArticleEndpoint, "DeleteArticle"),
		GetArticleEndpoint:    factory(MakeGetArticleEndpoint, "GetArticle"),
		UpdateArticleEndpoint: factory(MakeUpdateArticleEndpoint, "UpdateArticle"),
		AddLikeEndpoint:       factory(MakeAddLikeEndpoint, "AddLike"),
		DeleteLikeEndpoint:    factory(MakeDeleteLikeEndpoint, "DeleteLike"),
		AddDislikeEndpoint:    factory(MakeAddDislikeEndpoint, "AddDislike"),
		DeleteDislikeEndpoint: factory(MakeDeleteDislikeEndpoint, "DeleteDislike"),
	}
}

type createArticleRequest struct {
	Article *domain.Article
}
type createArticleResponse struct {
	ID  int64
	Err error
}

// MakeCreateArticleEndpoint - Impl.
func MakeCreateArticleEndpoint(service domain.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createArticleRequest)

		resp, err := service.CreateArticle(ctx, req.Article, helpers.CallerID(ctx))
		return createArticleResponse{
			ID:  resp,
			Err: err,
		}, nil
	}
}

type listArticleRequest struct {
	Criteria *domain.NewsSearchCriteria
}
type listArticleResponse struct {
	Articles   []*domain.Article
	TotalCount domain.Total
	Err        error
}

// MakeListArticlesEndpoint - Impl.
func MakeListArticlesEndpoint(service domain.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listArticleRequest)

		resp, totalCount, err := service.ListArticles(ctx, req.Criteria, helpers.CallerID(ctx))
		if err != nil {
			return nil, err
		}
		return listArticleResponse{
			Articles:   resp,
			TotalCount: totalCount,
			Err:        err,
		}, nil
	}
}

type getArticleRequest struct {
	ID int64
}
type getArticleResponse struct {
	Article *domain.Article
	Err     error
}

// MakeGetArticleEndpoint - Impl.
func MakeGetArticleEndpoint(service domain.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getArticleRequest)

		resp, err := service.GetArticle(ctx, req.ID, helpers.CallerID(ctx))
		return getArticleResponse{
			Article: resp,
			Err:     err,
		}, nil
	}
}

type deleteArticleRequest struct {
	ID int64
}
type deleteArticleResponse struct {
	Err error
}

// MakeDeleteArticleEndpoint - Impl.
func MakeDeleteArticleEndpoint(service domain.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteArticleRequest)

		err = service.DeleteArticle(ctx, req.ID, helpers.CallerID(ctx))
		return deleteArticleResponse{
			Err: err,
		}, nil
	}
}

type updateArticleRequest struct {
	ID      int64
	Article *domain.Article
}
type updateArticleResponse struct {
	Err error
}

// MakeUpdateArticleEndpoint - Impl.
func MakeUpdateArticleEndpoint(service domain.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateArticleRequest)

		err = service.UpdateArticle(ctx, req.ID, req.Article, helpers.CallerID(ctx))
		return updateArticleResponse{
			Err: err,
		}, nil
	}
}

// MakeAddLikeEndpoint Impl.
func MakeAddLikeEndpoint(service domain.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(addLikeRequest)

		err = service.AddLike(ctx, req.ArticleID, helpers.CallerID(ctx))
		return addLikeResponse{
			Err: err,
		}, nil
	}
}

type addLikeRequest struct {
	ArticleID int64
}

type addLikeResponse struct {
	Err error
}

// MakeDeleteLikeEndpoint Impl.
func MakeDeleteLikeEndpoint(service domain.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteLikeRequest)

		err = service.DeleteLike(ctx, req.ArticleID, helpers.CallerID(ctx))
		return deleteLikeResponse{
			Err: err,
		}, nil
	}
}

type deleteLikeRequest struct {
	ArticleID int64
}

type deleteLikeResponse struct {
	Err error
}

// MakeAddDislikeEndpoint Impl.
func MakeAddDislikeEndpoint(service domain.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(addDislikeRequest)

		err = service.AddDislike(ctx, req.ArticleID, helpers.CallerID(ctx))
		return addDislikeResponse{
			Err: err,
		}, nil
	}
}

type addDislikeRequest struct {
	ArticleID int64
}

type addDislikeResponse struct {
	Err error
}

// MakeDeleteDislikeEndpoint Impl.
func MakeDeleteDislikeEndpoint(service domain.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteDislikeRequest)

		err = service.DeleteDislike(ctx, req.ArticleID, helpers.CallerID(ctx))
		return deleteDislikeResponse{
			Err: err,
		}, nil
	}
}

type deleteDislikeRequest struct {
	ArticleID int64
}

type deleteDislikeResponse struct {
	Err error
}
