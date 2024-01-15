package transport

import (
	identityapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/identityserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

func decodeProfileV1(w *identityapiv1.ProfileWrite) *domain.Profile {
	if w == nil {
		return nil
	}

	return &domain.Profile{
		FirstName:               w.FirstName,
		LastName:                w.LastName,
		City:                    w.City,
		Phone:                   w.Phone,
		Email:                   w.Email,
		Password:                w.Password,
		ConsultationPhoneNumber: w.ConsultationPhoneNumber,
	}
}

func encodeProfileV1(r *domain.Profile) *identityapiv1.ProfileRead {
	if r == nil {
		return nil
	}

	return &identityapiv1.ProfileRead{
		Id:                      int64(r.ID),
		RoleId:                  int64(r.RoleID),
		FirstName:               r.FirstName,
		LastName:                r.LastName,
		City:                    r.City,
		Phone:                   r.Phone,
		Email:                   r.Email,
		ConsultationPhoneNumber: r.ConsultationPhoneNumber,
		CreatedAt:               r.CreatedAt,
	}
}
