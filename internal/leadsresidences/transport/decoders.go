package transport

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

func decodeLeadsSearchCriteria(apiCriteria *apiv1.LeadResidenceSearchCriteria) domain.LeadResidenceSearchCriteria {

	if apiCriteria == nil {
		// default criteria.
		return domain.LeadResidenceSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.LeadResidenceSearchCriteria{
		Page:     helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		StatusID: apiCriteria.StatusId,
		Name:     apiCriteria.Name,
	}
}

func decodeLeadV1(w *apiv1.LeadWrite) *domain.LeadResidence {
	if w == nil {
		return nil
	}

	return &domain.LeadResidence{
		ResidenceID: domain.ResidenceID(w.ResidenceId),
		StatusID:    w.StatusId,
		IssuedAt:    w.DateStart,
		ExpiresAt:   w.DateEnd,
	}
}

func encodeLeadV1(r *domain.LeadResidence) *apiv1.LeadRead {
	if r == nil {
		return nil
	}

	residenceName := ""
	if r.Residence != nil {
		residenceName = r.Residence.Title
	}

	return &apiv1.LeadRead{
		Id:            int64(r.ID),
		ResidenceId:   int64(r.ResidenceID),
		ResidenceName: residenceName,
		StatusId:      r.StatusID,
		Status:        encodeStatusV1(r.Status),
		DateStart:     r.IssuedAt,
		DateEnd:       r.ExpiresAt,
	}
}

func encodeStatusV1(r *domain.LeadStatus) *apiv1.LeadStatus {
	if r == nil {
		return nil
	}

	return &apiv1.LeadStatus{
		Id:   r.ID,
		Name: r.Name,
	}
}

func encodeLeadsV1(r []*domain.LeadResidence) []*apiv1.LeadRead {
	if r == nil {
		return nil
	}

	leads := make([]*apiv1.LeadRead, len(r))
	for i, opt := range r {
		leads[i] = encodeLeadV1(opt)
	}

	return leads
}
