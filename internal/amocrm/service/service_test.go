package service

import (
	"context"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"github.com/golang/mock/gomock"
)

func TestService_CreateApplication(t *testing.T) {

	var (
		ctx = context.Background()
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	httpClient := mocks.NewMockHTTPClientService(stubCtrl)

	service := newBasicService(httpClient)

	type arguments struct {
		application *domain.AmoApplication
	}

	// Define tests
	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: application created",
			arguments: arguments{
				application: &domain.AmoApplication{
					FullName:         "",
					Phone:            "",
					ResidenceDetails: &domain.Residence{},
				},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			err := service.CreateApplication(ctx, args.application)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}
