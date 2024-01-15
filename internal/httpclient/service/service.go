package service

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
}

// NewService - creates a new service
func NewService(
	logger log.Logger,
) domain.HTTPClientService {
	var service domain.HTTPClientService
	{
		service = newBasicService()
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService() domain.HTTPClientService {
	return &service{}
}

func (s *service) Get(
	ctx context.Context,
	url string,
) (*http.Response, error) {

	// Validate input
	//
	if len(url) == 0 {
		return nil, errors.NewErrInvalidArgument("url required")
	}

	// Create request
	//
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	// Do request
	//
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// Send response
	//
	return httpResponse, nil
}

func (s *service) Post(
	ctx context.Context,
	url string,
	payload interface{},
) (*http.Response, error) {

	// Validate input
	//
	if len(url) == 0 {
		return nil, errors.NewErrInvalidArgument("url required")
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Create request
	//
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	// Do request
	//
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// Send response
	//
	return httpResponse, nil
}
