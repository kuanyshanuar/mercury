package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	baseURL    string
	httpClient domain.HTTPClientService
}

// NewService - creates a new service
func NewService(
	httpClient domain.HTTPClientService,
	logger log.Logger,
) domain.AmoCrmService {
	var service domain.AmoCrmService
	{
		service = newBasicService(httpClient)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	httpClient domain.HTTPClientService,
) domain.AmoCrmService {
	return &service{
		baseURL:    "https://1kz.site/bitrix/shtab-kvartir/services/webform/index.php",
		httpClient: httpClient,
	}
}

func (s *service) CreateApplication(
	ctx context.Context,
	application *domain.AmoApplication,
) error {

	// Send content to crm
	//
	httpResponse, err := s.httpClient.Post(ctx, s.baseURL, application)
	if err != nil {
		return err
	}

	// Check http response status
	//
	if httpResponse.StatusCode >= 400 && httpResponse.StatusCode <= 500 {
		return errors.NewErrInternal(
			fmt.Sprintf("response status: %d", httpResponse.StatusCode),
		)
	}

	return nil
}
