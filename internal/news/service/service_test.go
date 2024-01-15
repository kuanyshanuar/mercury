package service

import (
	"context"
	"strings"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"github.com/golang/mock/gomock"
)

func TestService_CreateArticle(t *testing.T) {
	var (
		ctx          = context.Background()
		callerID     = domain.CallerID{}
		validArticle = &domain.Article{
			ID:               1,
			Title:            "Mara",
			ShortDescription: "Marat",
			Slug:             "mara",
			Content:          "3000 years ago",
			CreatedAt:        0,
			UpdatedAt:        0,
			DeletedAt:        0,
			CreatedBy:        0,
			DeletedBy:        0,
			UpdatedBy:        0,
			ViewsCount:       0,
			SourceURL:        "",
			AuthorName:       "10000",
			Images:           nil,
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	var (
		repository = mocks.NewMockNewsRepository(stubCtrl)
	)

	repository.EXPECT().
		Create(ctx, validArticle).
		Return(int64(1), nil).
		AnyTimes()
	repository.EXPECT().
		Get(ctx, int64(1)).
		Return(validArticle, nil).
		AnyTimes()

	newsService := newBasicService(repository)

	tests := []struct {
		name        string
		article     *domain.Article
		expectError bool
	}{
		{
			name: "",
			article: &domain.Article{
				ID:               1,
				Title:            "Mara",
				ShortDescription: "Marat",
				Slug:             "mara",
				Content:          "3000 years ago",
				CreatedAt:        0,
				UpdatedAt:        0,
				DeletedAt:        0,
				CreatedBy:        0,
				DeletedBy:        0,
				UpdatedBy:        0,
				ViewsCount:       0,
				SourceURL:        "",
				AuthorName:       "10000",
				Images:           nil,
			},
			expectError: false,
		},
	}

	// Define tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := newsService.CreateArticle(ctx, test.article, callerID)
			if err != nil {
				t.Errorf("unknown error %s", err)
			}
			if id <= 0 {
				t.Errorf("id is negative %s", err)
			}

			article, err := newsService.GetArticle(ctx, id, callerID)
			t.Log(id, article, "this is logging")

			if err != nil {
				t.Errorf("getArticle method error service %s", err)

			}
			if !strings.EqualFold(article.Title, test.article.Title) {
				t.Errorf("titles are not equal %s", err)
			}

			if !strings.EqualFold(article.Slug, test.article.Slug) {
				t.Errorf("slug are not equal %s", err)
			}

			if !strings.EqualFold(article.ShortDescription, test.article.ShortDescription) {
				t.Errorf("shortDescription are not equal %s", err)
			}
		})
	}
}

func TestService_GetArticle(t *testing.T) {
	var (
		ctx          = context.Background()
		callerID     = domain.CallerID{}
		validArticle = &domain.Article{
			ID:               1,
			Title:            "Mara",
			ShortDescription: "Marat",
			Slug:             "mara",
			Content:          "3000 years ago",
			CreatedAt:        0,
			UpdatedAt:        0,
			DeletedAt:        0,
			CreatedBy:        0,
			DeletedBy:        0,
			UpdatedBy:        0,
			ViewsCount:       0,
			SourceURL:        "",
			AuthorName:       "10000",
			Images:           nil,
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	var (
		repository = mocks.NewMockNewsRepository(stubCtrl)
	)

	newsService := newBasicService(repository)

	tests := []struct {
		name        string
		article     *domain.Article
		expectError bool
	}{
		{
			name: "",
			article: &domain.Article{
				ID:               1,
				Title:            "Mara",
				ShortDescription: "Marat",
				Slug:             "mara",
				Content:          "3000 years ago",
				CreatedAt:        0,
				UpdatedAt:        0,
				DeletedAt:        0,
				CreatedBy:        0,
				DeletedBy:        0,
				UpdatedBy:        0,
				ViewsCount:       0,
				SourceURL:        "",
				AuthorName:       "10000",
				Images:           nil,
			},
			expectError: false,
		},
	}
	repository.EXPECT().
		Create(ctx, validArticle).
		Return(int64(1), nil).
		AnyTimes()
	repository.EXPECT().
		Get(ctx, int64(1)).
		Return(validArticle, nil).
		AnyTimes()

	// Define tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := newsService.CreateArticle(ctx, test.article, callerID)
			if err != nil {
				t.Errorf("unknown error %s", err)
			}
			if id < 0 {
				t.Errorf("id is negative %s", err)
			}
			article, err := newsService.GetArticle(ctx, id, callerID)
			t.Log(id, article, "this is logging")

			if err != nil {
				t.Errorf("getArticle method error service %s", err)

			}
			if !strings.EqualFold(article.Title, test.article.Title) {
				t.Errorf("titles are not equal %s", err)
			}

			if !strings.EqualFold(article.Slug, test.article.Slug) {
				t.Errorf("slug are not equal %s", err)
			}

			if !strings.EqualFold(article.ShortDescription, test.article.ShortDescription) {
				t.Errorf("shortDescription are not equal %s", err)
			}
		})
	}
}

func TestService_ListArticles(t *testing.T) {
	var (
		ctx           = context.Background()
		callerID      = domain.CallerID{}
		validArticles = []*domain.Article{{
			ID:               1,
			Title:            "New-ZK",
			ShortDescription: "Shtab-kvartir is a service for finding the dream house",
			Slug:             "shtab-kvartir-news",
			Content:          "shtab-kvartir was found 3000 years ago",
			CreatedAt:        0,
			UpdatedAt:        0,
			DeletedAt:        0,
			CreatedBy:        0,
			DeletedBy:        0,
			UpdatedBy:        0,
			ViewsCount:       0,
			SourceURL:        "",
			AuthorName:       "10000",
			Images:           nil},

			{ID: 2,
				Title:            "New-ZK",
				ShortDescription: "New-ZK has arrived into our service",
				Slug:             "zk-almaty",
				Content:          "new-ZK has been found in our service",
				CreatedAt:        0,
				UpdatedAt:        0,
				DeletedAt:        0,
				CreatedBy:        0,
				DeletedBy:        0,
				UpdatedBy:        0,
				ViewsCount:       0,
				SourceURL:        "",
				AuthorName:       "10000",
				Images:           nil},
		}
		validCriterias = []*domain.NewsSearchCriteria{{
			ID:    -12,
			Title: "New-ZK",
			Slug:  "mara"},
			{ID: 2,
				Title: "",
				Slug:  "mara"},
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	var (
		repository = mocks.NewMockNewsRepository(stubCtrl)
	)

	newsService := newBasicService(repository)

	tests := []struct {
		name        string
		criteria    *domain.NewsSearchCriteria
		expectError bool
	}{
		{
			name: "",
			criteria: &domain.NewsSearchCriteria{
				ID:    -12,
				Title: "New-ZK",
				Slug:  "mara",
			},
			expectError: false,
		},
		{
			name: "",
			criteria: &domain.NewsSearchCriteria{
				ID:    2,
				Title: "",
				Slug:  "mara",
			},
			expectError: false,
		},
	}
	for i := range validArticles {
		repository.EXPECT().
			List(ctx, validCriterias[i]).
			Return(validArticles, domain.Total(2), nil)
	}

	// Define tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			article, totalCount, err := newsService.ListArticles(ctx, test.criteria, callerID)
			if totalCount != 2 {
				t.Errorf("TotalCounts isn't equal")
			}

			if err != nil {
				t.Errorf("getArticle method error service %s", err)

			}
			if len(article) != 2 {
				t.Errorf("article length isn't valid %s", err)

			}
		})
	}
}

func TestService_AddLike(t *testing.T) {

}

func TestService_DeleteLike(t *testing.T) {

}

func TestService_AddDislike(t *testing.T) {

}

func TestService_DeleteDislike(t *testing.T) {

}
