package service

import (
	"context"
	"fmt"
	errors "gitlab.com/zharzhanov/mercury/internal/error"
	"net/http"
	"net/url"
	"strings"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.SmsRepository
	baseURL    string
	apiKey     string
	alphaName  string
}

// NewService - creates a new service.
func NewService(
	repository domain.SmsRepository,
	baseURL string,
	apiKey string,
	alphaName string,
	logger log.Logger,
) domain.SmsService {
	var service domain.SmsService
	{
		service = newBasicService(
			repository,
			baseURL,
			apiKey,
			alphaName,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	repository domain.SmsRepository,
	baseURL string,
	apiKey string,
	alphaName string,
) domain.SmsService {
	return &service{
		repository: repository,
		baseURL:    baseURL,
		apiKey:     apiKey,
		alphaName:  alphaName,
	}
}

func (s *service) SendSms(
	ctx context.Context,
	sms *domain.Sms,
) error {

	// Validate input
	//
	if err := s.validateInternalSms(sms); err != nil {
		return err
	}

	// Create url params
	//
	params := url.Values{}
	params.Add("recipient", strings.Trim(sms.Phone, "+"))
	params.Add("text", sms.GetMessage())
	params.Add("apiKey", s.apiKey)

	// Sms api url
	//
	smsURL := fmt.Sprintf("%s?%s", s.baseURL, params.Encode())

	// Send request
	//
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, smsURL, nil)
	if err != nil {
		return err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return fmt.Errorf("can not send sms")
	}

	return nil
}

func (s *service) validateInternalSms(sms *domain.Sms) error {
	if sms == nil {
		return errors.NewErrInvalidArgument("sms required")
	}
	if len(sms.Phone) == 0 {
		return fmt.Errorf("phone required")
	}
	if len(sms.Message) == 0 {
		return fmt.Errorf("message required")
	}

	return nil
}
