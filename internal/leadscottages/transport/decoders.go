package transport

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

func decodeLeadsSearchCriteria(apiCriteria *apiv1.LeadCottageSearchCriteria) domain.LeadCottageSearchCriteria {

	if apiCriteria == nil {
		// default criteria.
		return domain.LeadCottageSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.LeadCottageSearchCriteria{
		Page:     helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		StatusID: apiCriteria.StatusId,
		Name:     apiCriteria.Name,
	}
}

func decodeLeadV1(w *apiv1.LeadCottageWrite) *domain.LeadCottage {
	if w == nil {
		return nil
	}

	return &domain.LeadCottage{
		CottageID: w.CottageId,
		StatusID:  w.StatusId,
		IssuedAt:  w.IssueDate,
		ExpiresAt: w.ExpireDate,
	}
}

func encodeLeadV1(r *domain.LeadCottage) *apiv1.LeadCottageRead {
	if r == nil {
		return nil
	}

	cottageName := ""
	if r.Cottage != nil {
		cottageName = r.Cottage.Title
	}

	return &apiv1.LeadCottageRead{
		Id:          r.ID,
		CottageId:   r.CottageID,
		CottageName: cottageName,
		StatusId:    r.StatusID,
		Status:      encodeStatusV1(r.Status),
		DateStart:   r.IssuedAt,
		DateEnd:     r.ExpiresAt,
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

func encodeLeadsV1(r []*domain.LeadCottage) []*apiv1.LeadCottageRead {
	if r == nil {
		return nil
	}

	leads := make([]*apiv1.LeadCottageRead, len(r))
	for i, opt := range r {
		leads[i] = encodeLeadV1(opt)
	}

	return leads
}
