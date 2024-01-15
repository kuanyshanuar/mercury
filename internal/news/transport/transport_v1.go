package transport

import (
	"context"

	newserviceapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

type grpcServer struct {
	create        grpctransport.Handler
	delete        grpctransport.Handler
	update        grpctransport.Handler
	list          grpctransport.Handler
	get           grpctransport.Handler
	addLike       grpctransport.Handler
	deleteLike    grpctransport.Handler
	addDislike    grpctransport.Handler
	deleteDislike grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoint Endpoints, logger log.Logger) newserviceapiv1.NewsServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoint.CreateArticleEndpoint,
			decodeCreateArticleRequestV1,
			encodeCreateArticleResponseV1,
			options...,
		),
		delete: grpctransport.NewServer(
			endpoint.DeleteArticleEndpoint,
			decodeDeleteArticleRequestV1,
			encodeDeleteArticleResponseV1,
			options...,
		),
		update: grpctransport.NewServer(
			endpoint.UpdateArticleEndpoint,
			decodeUpdateArticleRequestV1,
			encodeUpdateArticleResponseV1,
			options...,
		),
		get: grpctransport.NewServer(
			endpoint.GetArticleEndpoint,
			decodeGetArticleRequestV1,
			encodeGetArticleResponseV1,
			options...,
		),
		list: grpctransport.NewServer(
			endpoint.ListArticlesEndpoint,
			decodeListArticleRequestV1,
			encodeListArticleResponseV1,
			options...,
		),
		addLike: grpctransport.NewServer(
			endpoint.AddLikeEndpoint,
			decodeAddLikeRequestV1,
			encodeAddLikeResponseV1,
			options...,
		),
		deleteLike: grpctransport.NewServer(
			endpoint.DeleteLikeEndpoint,
			decodeDeleteLikeRequestV1,
			encodeDeleteLikeResponseV1,
			options...,
		),
		addDislike: grpctransport.NewServer(
			endpoint.AddDislikeEndpoint,
			decodeAddDislikeRequestV1,
			encodeAddDislikeResponseV1,
			options...,
		),
		deleteDislike: grpctransport.NewServer(
			endpoint.DeleteDislikeEndpoint,
			decodeDeleteDislikeRequestV1,
			encodeDeleteDislikeResponseV1,
			options...,
		),
	}
}

func decodeCreateArticleRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*newserviceapiv1.CreateNewsRequest)

	return createArticleRequest{
		Article: decodeArticleV1(req.Article),
	}, nil
}

func encodeCreateArticleResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(createArticleResponse)
	if resp.Err != nil {
		return &newserviceapiv1.CreateNewsResponse{}, resp.Err
	}
	return &newserviceapiv1.CreateNewsResponse{
		Id: resp.ID,
	}, nil
}

func (s *grpcServer) Create(ctx context.Context, request *newserviceapiv1.CreateNewsRequest) (*newserviceapiv1.CreateNewsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.create)
	if err != nil {
		return nil, err
	}
	return rep.(*newserviceapiv1.CreateNewsResponse), nil
}

func decodeDeleteArticleRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*newserviceapiv1.DeleteNewsRequest)
	return deleteArticleRequest{
		ID: req.Id,
	}, nil
}

func encodeDeleteArticleResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteArticleResponse)
	if resp.Err != nil {
		return &newserviceapiv1.DeleteNewsResponse{}, resp.Err
	}
	return &newserviceapiv1.DeleteNewsResponse{}, nil
}

func (s *grpcServer) Delete(ctx context.Context, request *newserviceapiv1.DeleteNewsRequest) (*newserviceapiv1.DeleteNewsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.delete)
	if err != nil {
		return nil, err
	}
	return rep.(*newserviceapiv1.DeleteNewsResponse), nil
}

func decodeUpdateArticleRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*newserviceapiv1.UpdateNewsRequest)
	return updateArticleRequest{
		ID:      req.Id,
		Article: decodeArticleV1(req.Article),
	}, nil
}

func encodeUpdateArticleResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateArticleResponse)
	if resp.Err != nil {
		return &newserviceapiv1.UpdateNewsResponse{}, resp.Err
	}
	return &newserviceapiv1.UpdateNewsResponse{}, nil
}

func (s *grpcServer) Update(ctx context.Context, request *newserviceapiv1.UpdateNewsRequest) (*newserviceapiv1.UpdateNewsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.update)
	if err != nil {
		return nil, err
	}
	return rep.(*newserviceapiv1.UpdateNewsResponse), nil
}

func decodeGetArticleRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*newserviceapiv1.GetNewsRequest)

	return getArticleRequest{
		ID: req.Id,
	}, nil
}

func encodeGetArticleResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(getArticleResponse)
	if resp.Err != nil {
		return &newserviceapiv1.GetNewsResponse{}, resp.Err
	}
	return &newserviceapiv1.GetNewsResponse{
		Article: encodeArticleV1(resp.Article),
	}, nil
}

func (s *grpcServer) Get(ctx context.Context, request *newserviceapiv1.GetNewsRequest) (*newserviceapiv1.GetNewsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.get)
	if err != nil {
		return nil, err
	}
	return rep.(*newserviceapiv1.GetNewsResponse), nil
}

func decodeListArticleRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*newserviceapiv1.ListNewsRequest)

	return listArticleRequest{
		decodeCriteriaV1(req.Criteria),
	}, nil
}

func encodeListArticleResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listArticleResponse)
	if resp.Err != nil {
		return &newserviceapiv1.ListNewsResponse{}, resp.Err
	}
	return &newserviceapiv1.ListNewsResponse{
		Articles:   encodeArticleListV1(resp.Articles),
		TotalCount: int64(resp.TotalCount),
	}, nil
}

func (s *grpcServer) List(ctx context.Context, request *newserviceapiv1.ListNewsRequest) (*newserviceapiv1.ListNewsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.list)
	if err != nil {
		return nil, err
	}
	return rep.(*newserviceapiv1.ListNewsResponse), nil
}

func (s *grpcServer) Like(ctx context.Context, request *newserviceapiv1.LikeRequest) (*newserviceapiv1.LikeResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.addLike)
	if err != nil {
		return nil, err
	}
	return rep.(*newserviceapiv1.LikeResponse), nil
}

func decodeAddLikeRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*newserviceapiv1.LikeRequest)

	return addLikeRequest{
		ArticleID: req.ArticleId,
	}, nil
}

func encodeAddLikeResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addLikeResponse)
	if resp.Err != nil {
		return &newserviceapiv1.LikeResponse{}, resp.Err
	}
	return &newserviceapiv1.LikeResponse{}, nil
}

func (s *grpcServer) DeleteLike(ctx context.Context, request *newserviceapiv1.DeleteLikeRequest) (*newserviceapiv1.DeleteLikeResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteLike)
	if err != nil {
		return nil, err
	}
	return rep.(*newserviceapiv1.DeleteLikeResponse), nil
}

func decodeDeleteLikeRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*newserviceapiv1.DeleteLikeRequest)

	return deleteLikeRequest{
		ArticleID: req.ArticleId,
	}, nil
}

func encodeDeleteLikeResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteLikeResponse)
	if resp.Err != nil {
		return &newserviceapiv1.DeleteLikeResponse{}, resp.Err
	}
	return &newserviceapiv1.DeleteLikeResponse{}, nil
}

func (s *grpcServer) Dislike(ctx context.Context, request *newserviceapiv1.DislikeRequest) (*newserviceapiv1.DislikeResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.addDislike)
	if err != nil {
		return nil, err
	}
	return rep.(*newserviceapiv1.DislikeResponse), nil
}

func decodeAddDislikeRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*newserviceapiv1.DislikeRequest)

	return addDislikeRequest{
		ArticleID: req.ArticleId,
	}, nil
}

func encodeAddDislikeResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addDislikeResponse)
	if resp.Err != nil {
		return &newserviceapiv1.DislikeResponse{}, resp.Err
	}
	return &newserviceapiv1.DislikeResponse{}, nil
}

func (s *grpcServer) DeleteDislike(ctx context.Context, request *newserviceapiv1.DeleteDislikeRequest) (*newserviceapiv1.DeleteDislikeResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteDislike)
	if err != nil {
		return nil, err
	}
	return rep.(*newserviceapiv1.DeleteDislikeResponse), nil
}

func decodeDeleteDislikeRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*newserviceapiv1.DeleteDislikeRequest)

	return deleteDislikeRequest{
		ArticleID: req.ArticleId,
	}, nil
}

func encodeDeleteDislikeResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteDislikeResponse)
	if resp.Err != nil {
		return &newserviceapiv1.DeleteDislikeResponse{}, resp.Err
	}
	return &newserviceapiv1.DeleteDislikeResponse{}, nil
}
