package repository

import (
	"context"
	"strings"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	pkgPostgres "gitlab.com/zharzhanov/mercury/pkg/database/postgres"
)

func TestRepository_Create(t *testing.T) {

	// Setup database.
	//
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	var (
		ctx = context.Background()
	)

	testNewsRepo := NewRepository(db)

	tests := []struct {
		name        string
		article     *domain.Article
		expectError bool
	}{
		{
			name: "Success: residence 1 create",
			article: &domain.Article{
				Title:            "Kapster",
				ShortDescription: "Shtab-kvartir: vars",
				Slug:             "kapster",
				CreatedAt:        12089991,
				AuthorName:       "Jupiter team",
				CreatedBy:        1,
			},
			expectError: false,
		}, {
			name: "Success: residence 2 create",
			article: &domain.Article{
				Title:            "HomeLander",
				ShortDescription: "Destroyed USA",
				Slug:             "homelander",
				CreatedAt:        12089991,
				AuthorName:       "Jupiter team",
				CreatedBy:        1,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			article := test.article
			id, err := testNewsRepo.Create(ctx, article)
			if err != nil {
				t.Errorf("unexpected error %s", err)
			}
			if id <= 0 {
				t.Error("article has not been created")
				return
			}

			result, err := testNewsRepo.Get(ctx, id)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if !strings.EqualFold(article.Title, result.Title) {
					t.Errorf("expected article title is %s but actual is %s", article.Title, result.Title)
				}
				if result.ID <= 0 {
					t.Error("bad residence id")
				}
				if result.CreatedAt <= 0 {
					t.Error("creation time not set")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_Update(t *testing.T) {

	// Setup database.
	//

	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	var (
		ctx = context.Background()
	)

	testNewsRepo := NewRepository(db)

	tests := []struct {
		name            string
		article         *domain.Article
		modifiedArticle *domain.Article
		id              int64
		expectError     bool
	}{
		{
			name: "Success: residence 1 update",
			article: &domain.Article{
				Title:            "Kapster",
				ShortDescription: "Shtab-kvartir: vars",
				Slug:             "kapster",
				CreatedAt:        12089991,
				AuthorName:       "Jupiter team",
				CreatedBy:        1,
			},
			modifiedArticle: &domain.Article{
				Title:            "KapsterRank",
				ShortDescription: "Shtab-kvartir: vars",
				Slug:             "kapster_113",
				CreatedAt:        12089991,
				AuthorName:       "Jupiter team",
				CreatedBy:        1,
			},
			expectError: false,
			id:          1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newArticleID, err := testNewsRepo.Create(ctx, test.modifiedArticle)
			if err != nil {
				t.Errorf("unexpected error %s", err)
			}
			if newArticleID <= 0 {
				t.Errorf("%s", err)
			}

			err = testNewsRepo.Update(ctx, newArticleID, test.modifiedArticle)
			if err != nil {
				t.Errorf("unexpected error %s", err)
			}

			result, err := testNewsRepo.Get(ctx, newArticleID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if !strings.EqualFold(result.Title, test.modifiedArticle.Title) {
					t.Errorf("not equal error %s", err)
				}
				if !strings.EqualFold(result.Slug, test.modifiedArticle.Slug) {
					t.Errorf("not equal error %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_Delete(t *testing.T) {

	// Setup database.
	//
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	var (
		ctx = context.Background()
	)

	testNewsRepo := NewRepository(db)

	tests := []struct {
		name        string
		id          int64
		expectError bool
	}{
		{
			name:        "Success: article 1 delete",
			id:          1,
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id := test.id
			err := testNewsRepo.Delete(ctx, id)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				_, err = testNewsRepo.Get(ctx, id)
				if err == nil {
					t.Error("expected error but got nothing")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_List(t *testing.T) {

	// Setup database.
	//
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	var (
		ctx     = context.Background()
		article = &domain.Article{
			Title:            "Kapster",
			ShortDescription: "Shtab-kvartir: vars",
			Slug:             "kapster",
			CreatedAt:        12089991,
			AuthorName:       "Jupiter team",
			CreatedBy:        1,
		}
	)

	testNewsRepo := NewRepository(db)

	tests := []struct {
		criteria    domain.NewsSearchCriteria
		name        string
		expectError bool
	}{
		{
			name: "List News test #1 ",
			criteria: domain.NewsSearchCriteria{
				Page: domain.PageRequest{
					Offset: -1,
					Size:   2,
				},
				Title: "kapsteR",
				Short: "",
				Slug:  "",
			}, expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := testNewsRepo.Create(ctx, article)
			if err != nil {
				t.Errorf("unexpected error %s", err)
			}

			list, totalArticles, err := testNewsRepo.List(ctx, &test.criteria)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				if totalArticles <= 0 {
					t.Errorf("total article is 0 %s", err)
				}
				if len(list) == 0 {
					t.Error(err, id, list)
				}

			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_Get(t *testing.T) {

	// Setup database.
	//
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	var (
		ctx      = context.Background()
		articles = []*domain.Article{
			{
				Title:            "test1",
				ShortDescription: "test1",
				Slug:             "test1",
				Content:          "test1",
				SourceURL:        "test1",
				AuthorName:       "test1",
				Images:           []string{"test1"},
			},
			{
				Title:            "test2",
				ShortDescription: "test2",
				Slug:             "test2",
				Content:          "test2",
				SourceURL:        "test2",
				AuthorName:       "test2",
				Images:           []string{"test2"},
			},
		}
	)

	testNewsRepo := NewRepository(db)

	for _, article := range articles {
		createdID, err := testNewsRepo.Create(ctx, article)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if createdID == 0 {
			t.Error("unexpected error article not created")
		}
	}

	tests := []struct {
		name        string
		id          int64
		expectError bool
	}{
		{
			name:        "Get the article 1",
			id:          1,
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			article, err := testNewsRepo.Get(ctx, test.id)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if len(article.Title) == 0 {
					t.Errorf("article title is empty : %s", err)
				}
				if len(article.ShortDescription) == 0 {
					t.Errorf("article short is empty : %s", err)
				}
				if len(article.Content) == 0 {
					t.Errorf("article content is empty : %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRepository_AddLike(t *testing.T) {

	// Setup database.
	//
	db, err := pkgPostgres.SetupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = sqlDB.Close() }()

	var (
		ctx      = context.Background()
		articles = []*domain.Article{
			{
				Title:            "test1",
				ShortDescription: "test1",
				Slug:             "test1",
				Content:          "test1",
				SourceURL:        "test1",
				AuthorName:       "test1",
				Images:           []string{"test1"},
			},
			{
				Title:            "test2",
				ShortDescription: "test2",
				Slug:             "test2",
				Content:          "test2",
				SourceURL:        "test2",
				AuthorName:       "test2",
				Images:           []string{"test2"},
			},
		}
	)

	testNewsRepo := NewRepository(db)

	for _, article := range articles {
		createdID, err := testNewsRepo.Create(ctx, article)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if createdID == 0 {
			t.Error("unexpected error article not created")
		}
	}

	tests := []struct {
		name        string
		articleID   int64
		userID      int64
		expectError bool
	}{
		{
			name:        "Like article 1",
			articleID:   1,
			userID:      1,
			expectError: false,
		},
		{
			name:        "Like article 2",
			articleID:   2,
			userID:      1,
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testNewsRepo.AddLike(ctx, test.articleID, test.userID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}
