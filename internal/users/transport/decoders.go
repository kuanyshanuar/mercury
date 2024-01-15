package transport

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

func decodeUserSearchCriteriaV1(apiCriteria *apiv1.UserSearchCriteria) domain.UserSearchCriteria {
	if apiCriteria == nil {
		// default criteria.
		return domain.UserSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.UserSearchCriteria{
		Page:  helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		ID:    apiCriteria.Id,
		Name:  apiCriteria.Name,
		Email: apiCriteria.Email,
		Phone: apiCriteria.Phone,
	}
}

func encodeUsersV1(r []*domain.User) []*apiv1.UserRead {
	if r == nil {
		return nil
	}

	users := make([]*apiv1.UserRead, len(r))
	for i, opt := range r {
		users[i] = encodeUserV1(opt)
	}

	return users
}

func encodeUserV1(r *domain.User) *apiv1.UserRead {
	if r == nil {
		return nil
	}

	return &apiv1.UserRead{
		Id:        int64(r.ID),
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Phone:     r.Phone,
		IsBanned:  r.IsBanned != nil && *r.IsBanned,
	}
}

func decodeUserV1(w *apiv1.UserWrite) *domain.User {
	if w == nil {
		return nil
	}

	return &domain.User{
		FirstName: w.FirstName,
		LastName:  w.LastName,
		Email:     w.Email,
		City:      w.City,
		Phone:     w.Phone,
		IsBanned:  &w.IsBanned,
	}
}
