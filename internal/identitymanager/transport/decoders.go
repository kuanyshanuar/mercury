package transport

import (
	imsapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/identityserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

func decodeUserV1(w *imsapiv1.UserWrite) *domain.User {
	if w == nil {
		return nil
	}

	return &domain.User{
		RoleID:                  domain.RoleID(w.RoleId),
		FirstName:               w.FirstName,
		LastName:                w.LastName,
		ConsultationPhoneNumber: w.ConsultationPhoneNumber,
		City:                    w.City,
		Phone:                   w.Phone,
		Email:                   w.Email,
		Password:                w.Password,
	}
}

func encodeUserV1(r *domain.User) *imsapiv1.UserRead {
	if r == nil {
		return nil
	}

	return &imsapiv1.UserRead{
		Id:                      int64(r.ID),
		RoleId:                  int64(r.RoleID),
		FirstName:               r.FirstName,
		LastName:                r.LastName,
		City:                    r.City,
		Phone:                   r.Phone,
		ConsultationPhoneNumber: r.ConsultationPhoneNumber,
		Email:                   r.Email,
		CreatedAt:               r.CreatedAt,
		UpdatedAt:               r.UpdatedAt,
	}
}
