package service

import (
	"context"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"github.com/golang/mock/gomock"
)

func TestService_CreateFilter(t *testing.T) {
	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}
		key      = "parking_types"
		filter   = &domain.Filter{
			Name: "air",
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repositories
	filtersRepoStub := mocks.NewMockFiltersRepository(stubCtrl)

	filtersRepoStub.EXPECT().
		CreateFilter(ctx, key, filter).
		Return(nil, nil).
		AnyTimes()

	// Service
	service := newBasicService(filtersRepoStub)

	type arguments struct {
		key    string
		filter *domain.Filter
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: create filter",
			arguments: arguments{
				key:    key,
				filter: filter,
			},
			expectError: false,
		},
		{
			name: "Fail: invalid key",
			arguments: arguments{
				key: "",
				filter: &domain.Filter{
					Name: "air",
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid filter name",
			arguments: arguments{
				key: key,
				filter: &domain.Filter{
					Name: "",
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid filter",
			arguments: arguments{
				key:    key,
				filter: nil,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			_, err := service.CreateFilter(ctx, args.key, args.filter, callerID)
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

func TestService_DeleteFilter(t *testing.T) {
	var (
		ctx      = context.Background()
		callerID = domain.CallerID{}
		key      = "parking_types"
		id       = int64(1)
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repositories
	filtersRepoStub := mocks.NewMockFiltersRepository(stubCtrl)

	filtersRepoStub.EXPECT().
		DeleteFilter(ctx, id, key).
		Return(nil).
		AnyTimes()

	// Service
	service := newBasicService(filtersRepoStub)

	type arguments struct {
		key string
		id  int64
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Success: create filter",
			arguments: arguments{
				key: key,
				id:  id,
			},
			expectError: false,
		},
		{
			name: "Fail: invalid key",
			arguments: arguments{
				key: "",
				id:  id,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid id",
			arguments: arguments{
				key: key,
				id:  0,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			err := service.DeleteFilter(ctx, args.id, args.key, callerID)
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
