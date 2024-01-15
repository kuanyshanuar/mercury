package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.ManagerRepository
}

// NewService - creates a new service
func NewService(
	repository domain.ManagerRepository,
	logger log.Logger,
) domain.ManagerService {
	var service domain.ManagerService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}

	return service
}

func newBasicService(
	repository domain.ManagerRepository,
) domain.ManagerService {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateManager(
	ctx context.Context,
	manager *domain.Manager,
	callerID domain.CallerID,
) (domain.ManagerID, error) {
	// Assign role
	//
	{
		manager.RoleID = domain.RoleManager
		manager.IsVerified = true
	}

	// Validate inputs
	//
	if err := s.validateCreateManager(manager); err != nil {
		return 0, err
	}

	// Create manager
	//
	managerID, err := s.repository.Create(ctx, manager)
	if err != nil {
		return 0, err
	}

	return managerID, nil
}

func (s *service) GetManager(
	ctx context.Context,
	managerID domain.ManagerID,
	callerID domain.CallerID,
) (*domain.Manager, error) {
	// Validate inputs
	//
	if managerID <= 0 {
		return nil, errors.NewErrInvalidArgument("manager id required")
	}

	return s.repository.Get(ctx, managerID)
}

func (s *service) ListManagers(
	ctx context.Context,
	criteria domain.ManagerSearchCriteria,
	callerID domain.CallerID,
) ([]*domain.Manager, domain.Total, error) {
	return s.repository.List(ctx, criteria)
}

func (s *service) UpdateManager(
	ctx context.Context,
	managerID domain.ManagerID,
	manager *domain.Manager,
	callerID domain.CallerID,
) error {

	// Validate inputs
	//
	if managerID <= 0 {
		return errors.NewErrInvalidArgument("manager id required")
	}
	if manager == nil {
		return errors.NewErrInvalidArgument("manager required")
	}

	return s.repository.Update(ctx, managerID, manager)
}

func (s *service) DeleteManager(
	ctx context.Context,
	managerID domain.ManagerID,
	callerID domain.CallerID,
) error {
	// Validate inputs
	//
	if managerID <= 0 {
		return errors.NewErrInvalidArgument("manager id required")
	}

	return s.repository.Delete(ctx, managerID)
}

func (s *service) validateCreateManager(manager *domain.Manager) error {

	if manager == nil {
		return errors.NewErrInvalidArgument("manager required")
	}
	if manager.RoleID <= 0 {
		return errors.NewErrInvalidArgument("role id required")
	}
	if manager.RoleID != domain.RoleManager {
		return errors.NewErrInvalidArgument("role id does not match")
	}
	if len(manager.FirstName) == 0 {
		return errors.NewErrInvalidArgument("first name required")
	}
	if len(manager.LastName) == 0 {
		return errors.NewErrInvalidArgument("last name required")
	}
	if len(manager.Email) == 0 {
		return errors.NewErrInvalidArgument("email required")
	}
	if len(manager.Phone) == 0 {
		return errors.NewErrInvalidArgument("phone required")
	}
	if len(manager.Password) == 0 {
		return errors.NewErrInvalidArgument("password required")
	}

	return nil
}
