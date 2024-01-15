package transport

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

func decodeManagerV1(w *apiv1.ManagerWrite) *domain.Manager {
	if w == nil {
		return nil
	}

	return &domain.Manager{
		RoleID:    domain.RoleManager,
		FirstName: w.FirstName,
		LastName:  w.LastName,
		Email:     w.Email,
		Phone:     w.Phone,
		Image:     w.Image,
		Password:  w.Password,
		IsBanned:  &w.IsBanned,
	}
}

func encodeManagerV1(r *domain.Manager) *apiv1.ManagerRead {
	if r == nil {
		return nil
	}

	return &apiv1.ManagerRead{
		Id:        int64(r.ID),
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Phone:     r.Phone,
		Image:     r.Image,
		IsBanned:  r.IsBanned != nil && *r.IsBanned,
	}
}

func encodeManagersV1(r []*domain.Manager) []*apiv1.ManagerRead {
	if r == nil {
		return nil
	}

	managers := make([]*apiv1.ManagerRead, len(r))
	for i, manager := range r {
		managers[i] = encodeManagerV1(manager)
	}

	return managers
}

func decodeManagerSearchCriteriaV1(apiCriteria *apiv1.ManagerSearchCriteria) domain.ManagerSearchCriteria {
	if apiCriteria == nil {
		// default criteria.
		return domain.ManagerSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.ManagerSearchCriteria{
		Page:  helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		ID:    apiCriteria.Id,
		Name:  apiCriteria.Name,
		Email: apiCriteria.Email,
		Phone: apiCriteria.Phone,
	}
}
