package service

import (
	"context"
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/mocks"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
)

func TestService_Allow(t *testing.T) {
	var (
		ctx           = context.Background()
		roleID1       = domain.RoleID(1)
		roleID2       = domain.RoleID(2)
		permissionKey = "Residences.CreateResidence"
		permission    = &domain.Permission{
			ID:           1,
			EndpointName: "Residences.CreateResidence",
			Action:       "CreateResidence",
			CrudType:     "C",
			IsActive:     true,
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	repoStub := mocks.NewMockPermissionsRepository(stubCtrl)
	service := newBasicService(repoStub)

	repoStub.EXPECT().
		GetPermissionByEndpoint(ctx, permissionKey).
		Return(permission, nil).
		AnyTimes()

	repoStub.EXPECT().
		IsPermissionAllowed(ctx, roleID1, permission.ID).
		Return(true, nil).
		AnyTimes()

	repoStub.EXPECT().
		IsPermissionAllowed(ctx, roleID2, permission.ID).
		Return(false, nil).
		AnyTimes()

	// Arguments
	type arguments struct {
		callerID      domain.CallerID
		permissionKey string
	}

	// Result
	type result struct {
		isAllowed bool
	}

	// Test cases
	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: is allowed",
			arguments: arguments{
				callerID: domain.CallerID{
					UserID: 1,
					RoleID: 1,
				},
				permissionKey: permissionKey,
			},
			expected: result{
				isAllowed: true,
			},
			expectError: false,
		},
		{
			name: "Success: not allowed",
			arguments: arguments{
				callerID: domain.CallerID{
					UserID: 2,
					RoleID: 2,
				},
				permissionKey: permissionKey,
			},
			expected: result{
				isAllowed: false,
			},
			expectError: false,
		},
		{
			name: "Fail: invalid user id",
			arguments: arguments{
				callerID: domain.CallerID{
					UserID: 0,
					RoleID: 2,
				},
			},
			expected: result{
				isAllowed: false,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid role id",
			arguments: arguments{
				callerID: domain.CallerID{
					UserID: 1,
					RoleID: 0,
				},
			},
			expected: result{
				isAllowed: false,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			expected := test.expected

			isAllowed, err := service.Allow(ctx, args.permissionKey, domain.UserID(args.callerID.UserID), domain.RoleID(args.callerID.RoleID), args.callerID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					isAllowed: isAllowed,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
					t.Error(diff)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}
