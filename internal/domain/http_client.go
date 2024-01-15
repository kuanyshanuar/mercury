package domain

import (
	"context"
	"net/http"
)

// HTTPClientService - http client service
type HTTPClientService interface {
	// Get - creates GET http request
	//
	Get(ctx context.Context, url string) (*http.Response, error)

	// Post - creates POST http request
	Post(ctx context.Context, url string, payload interface{}) (*http.Response, error)
}
