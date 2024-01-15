package transport

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

func decodeLeadBuilderV1(w *apiv1.LeadBuilderWrite) *domain.LeadBuilder {
	if w == nil {
		return nil
	}

	return &domain.LeadBuilder{
		BuilderID: domain.BuilderID(w.BuilderId),
		StatusID:  w.StatusId,
		IssuedAt:  w.DateStart,
		ExpiresAt: w.DateEnd,
	}
}

func encodeLeadBuilderV1(r *domain.LeadBuilder) *apiv1.LeadBuilderRead {
	if r == nil {
		return nil
	}

	return &apiv1.LeadBuilderRead{
		Id:          int64(r.ID),
		BuilderId:   int64(r.BuilderID),
		BuilderName: r.Builder.FirstName + r.Builder.LastName,
		StatusId:    r.StatusID,
		Status:      encodeStatusV1(r.Status),
		DateStart:   r.IssuedAt,
		DateEnd:     r.ExpiresAt,
	}
}

func encodeStatusV1(status *domain.LeadStatus) *apiv1.LeadStatus {
	if status == nil {
		return nil
	}

	return &apiv1.LeadStatus{
		Id:   status.ID,
		Name: status.Name,
	}
}

func encodeLeadBuildersV1(r []*domain.LeadBuilder) []*apiv1.LeadBuilderRead {
	if r == nil {
		return nil
	}

	leads := make([]*apiv1.LeadBuilderRead, len(r))
	for i, opt := range r {
		leads[i] = encodeLeadBuilderV1(opt)
	}

	return leads
}

func decodeLeadBuildersSearchCriteria(apiCriteria *apiv1.LeadBuilderSearchCriteria) domain.LeadBuilderSearchCriteria {

	if apiCriteria == nil {
		// default criteria.
		return domain.LeadBuilderSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.LeadBuilderSearchCriteria{
		Page:     helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		StatusID: apiCriteria.StatusId,
		Name:     apiCriteria.Name,
	}
}
