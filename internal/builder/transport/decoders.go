package transport

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

func decodeBuilderV1(w *apiv1.BuilderWrite) *domain.Builder {
	if w == nil {
		return nil
	}

	return &domain.Builder{
		FirstName:               w.FirstName,
		LastName:                w.LastName,
		City:                    w.City,
		Phone:                   w.Phone,
		Email:                   w.Email,
		ConsultationPhoneNumber: w.ConsultationPhoneNumber,
		Image:                   w.Image,
		Password:                w.Password,
		IsBanned:                &w.IsBanned,
	}
}

func encodeBuilderV1(r *domain.Builder) *apiv1.BuilderRead {
	if r == nil {
		return nil
	}

	return &apiv1.BuilderRead{
		Id:                      int64(r.ID),
		FirstName:               r.FirstName,
		LastName:                r.LastName,
		Email:                   r.Email,
		City:                    r.City,
		Phone:                   r.Phone,
		ConsultationPhoneNumber: r.ConsultationPhoneNumber,
		Image:                   r.Image,
		IsFavourite:             r.IsFavourite,
		IsBanned:                r.IsBanned != nil && *r.IsBanned,
	}
}

func encodeBuildersV1(builders []*domain.Builder) []*apiv1.BuilderRead {
	if builders == nil {
		return nil
	}

	list := make([]*apiv1.BuilderRead, len(builders))
	for i, opt := range builders {
		list[i] = encodeBuilderV1(opt)
	}

	return list
}

func decodeBuilderSearchCriteriaV1(apiCriteria *apiv1.BuilderSearchCriteria) domain.BuilderSearchCriteria {

	if apiCriteria == nil {
		// default criteria.
		return domain.BuilderSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.BuilderSearchCriteria{
		Page:   helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		UserID: apiCriteria.UserId,
		ID:     apiCriteria.Id,
		Name:   apiCriteria.Name,
		Email:  apiCriteria.Email,
		Phone:  apiCriteria.Phone,
	}
}
