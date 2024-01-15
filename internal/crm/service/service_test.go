package service

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/mocks"
)

func TestService_SendResidenceContactDetail(t *testing.T) {

	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}

		content = domain.ResidenceContactDetailContent{
			ResidenceName: "test",
			Slug:          "test",
			FullName:      "test",
			Phone:         "test",
		}
		baseURL = "https://motion.zoyd.space/webhook"
		url     = "https://motion.zoyd.space/webhook/test/shtab"
	)

	// Setup mocks.
	//
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks
	//
	httpClient := mocks.NewMockHTTPClientService(stubCtrl)

	// Mock repositories
	//
	httpClient.EXPECT().Post(ctx, url, content).Return(
		&http.Response{
			StatusCode: 200,
		}, nil).AnyTimes()

	service := newBasicService(
		baseURL,
		"staging",
		httpClient,
	)

	// Define tests
	//
	type arguments struct {
		content domain.ResidenceContactDetailContent
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: residence contact details sent",
			arguments: arguments{
				content: content,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			err := service.SendResidenceContactDetail(ctx, args.content, callerID)
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
