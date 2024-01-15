package transport

import (
	residenceapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

func decodeContactDetails(w *residenceapiv1.ContactDetailsWrite) *domain.ContactDetails {
	if w == nil {
		return nil
	}

	return &domain.ContactDetails{
		FullName: w.FullName,
		Phone:    w.Phone,
		Message:  w.Message,
	}
}

func decodeResidenceContactDetails(w *residenceapiv1.ResidenceContactDetailsWrite) *domain.ResidenceContactDetails {
	if w == nil {
		return nil
	}

	return &domain.ResidenceContactDetails{
		ResidenceID: w.GetResidenceId(),
		FullName:    w.GetFullName(),
		Phone:       w.GetPhone(),
	}
}

func encodeListContactDetails(r []*domain.ContactDetails) []*residenceapiv1.ContactDetailsRead {
	if r == nil {
		return nil
	}

	list := make([]*residenceapiv1.ContactDetailsRead, len(r))
	for i, opt := range r {
		list[i] = encodeContactDetails(opt)
	}

	return list
}

func encodeContactDetails(r *domain.ContactDetails) *residenceapiv1.ContactDetailsRead {
	if r == nil {
		return nil
	}

	return &residenceapiv1.ContactDetailsRead{
		Id:       r.ID,
		FullName: r.FullName,
		Phone:    r.Phone,
		Message:  r.Message,
	}
}

func encodeListResidenceContactDetails(r []*domain.ResidenceContactDetails) []*residenceapiv1.ResidenceContactDetailsRead {
	if r == nil {
		return nil
	}

	list := make([]*residenceapiv1.ResidenceContactDetailsRead, len(r))
	for i, opt := range r {
		list[i] = encodeResidenceContactDetails(opt)
	}

	return list
}

func encodeResidenceContactDetails(r *domain.ResidenceContactDetails) *residenceapiv1.ResidenceContactDetailsRead {
	if r == nil {
		return nil
	}

	return &residenceapiv1.ResidenceContactDetailsRead{
		Id:          r.ID,
		FullName:    r.FullName,
		ResidenceId: r.ResidenceID,
		Phone:       r.Phone,
	}
}
