package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	baseURL     string
	environment string
	httpClient  domain.HTTPClientService
}

// NewService creates a new service
func NewService(
	baseURL string,
	environment string,
	httpClient domain.HTTPClientService,
	logger log.Logger,
) domain.CrmService {
	var service domain.CrmService
	{
		service = newBasicService(
			baseURL,
			environment,
			httpClient,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	baseURL string,
	environment string,
	httpClient domain.HTTPClientService,
) domain.CrmService {
	return &service{
		baseURL:     baseURL,
		environment: environment,
		httpClient:  httpClient,
	}
}

func (s *service) SendResidenceContactDetail(
	ctx context.Context,
	content domain.ResidenceContactDetailContent,
	callerID domain.CallerID,
) error {

	// Create url
	//
	url := fmt.Sprintf("%s/%s/shtab", s.baseURL, content.Slug)
	if s.environment == "staging" {
		url = fmt.Sprintf("%s/%s/website", s.baseURL, content.Slug)
	}

	// Send content to crm
	//
	httpResponse, err := s.httpClient.Post(ctx, url, content)
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
