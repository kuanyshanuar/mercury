package transport

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

func decodeArticleV1(article *apiv1.Article) *domain.Article {
	if article == nil {
		return nil
	}

	return &domain.Article{
		Title:            article.Title,
		ShortDescription: article.ShortDescription,
		Slug:             article.Slug,
		ViewsCount:       article.ViewsCount,
		SourceURL:        article.SourceUrl,
		AuthorName:       article.AuthorName,
		Images:           article.Images,
		CreatedAt:        article.CreatedAt,
		UpdatedAt:        article.UpdatedAt,
		UpdatedBy:        article.UpdatedBy,
		CreatedBy:        article.CreatedBy,
		DeletedAt:        article.DeletedAt,
		DeletedBy:        article.DeletedBy,
		IsLikedByMe:      article.HasLikedByMe,
		IsDislikedByMe:   article.HasDislikedByMe,
		Content:          article.Content,
		Likes:            article.Likes,
		Dislikes:         article.Dislikes,
	}

}

func encodeArticleV1(article *domain.Article) *apiv1.Article {
	if article == nil {
		return nil
	}

	return &apiv1.Article{
		ID:               article.ID,
		Title:            article.Title,
		ShortDescription: article.ShortDescription,
		Slug:             article.Slug,
		ViewsCount:       article.ViewsCount,
		SourceUrl:        article.SourceURL,
		AuthorName:       article.AuthorName,
		Images:           article.Images,
		CreatedAt:        article.CreatedAt,
		UpdatedAt:        article.UpdatedAt,
		UpdatedBy:        article.UpdatedBy,
		CreatedBy:        article.CreatedBy,
		DeletedAt:        article.DeletedAt,
		DeletedBy:        article.DeletedBy,
		HasLikedByMe:     article.IsLikedByMe,
		HasDislikedByMe:  article.IsDislikedByMe,
		Content:          article.Content,
		Likes:            article.Likes,
		Dislikes:         article.Dislikes,
	}
}

func encodeArticleListV1(r []*domain.Article) []*apiv1.Article {
	if r == nil {
		return nil
	}

	articles := make([]*apiv1.Article, len(r))
	for idx, article := range r {
		articles[idx] = encodeArticleV1(article)
	}

	return articles
}

func decodeCriteriaV1(apiCriteria *apiv1.NewsSearchCriteria) *domain.NewsSearchCriteria {
	if apiCriteria == nil {
		return nil
	}

	var pageReq domain.PageRequest
	if apiCriteria.PageRequest != nil {
		pageReq = domain.PageRequest{
			Offset: int(apiCriteria.PageRequest.Offset),
			Size:   int(apiCriteria.PageRequest.Size),
		}
	}

	return &domain.NewsSearchCriteria{
		Page:  pageReq,
		ID:    apiCriteria.Id,    // filter by id
		Title: apiCriteria.Title, // filter by title
		Short: apiCriteria.Short, // filter by short description
		Slug:  apiCriteria.Slug,  // filter by slug
		Sorts: helpers.DecodeSorts(apiCriteria.Sorts),
	}
}
