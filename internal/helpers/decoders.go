package helpers

import (
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

// DecodePageRequestV1 - decodes page request.
func DecodePageRequestV1(page *apiv1.PageRequest) domain.PageRequest {
	if page == nil {
		return domain.PageRequest{
			Offset: 0,
			Size:   domain.DefaultPageSize,
		}
	}
	return domain.PageRequest{
		Offset: int(page.Offset),
		Size:   int(page.Size),
	}
}

// DecodeSort - decodes the sorting structure.
func DecodeSort(sort *apiv1.Sort) domain.Sort {
	if sort == nil {
		return domain.Sort{Order: domain.Asc}
	}
	order := domain.Asc
	if !sort.Asc {
		order = domain.Desc
	}
	return domain.Sort{
		FieldName: sort.FieldName,
		Order:     order,
	}
}

// DecodeSorts - decodes the sorting structure.
func DecodeSorts(sorts []*apiv1.Sort) []domain.Sort {
	var result []domain.Sort
	for _, s := range sorts {
		result = append(result, DecodeSort(s))
	}
	return result
}

// PointerBool - returns the pointer of the boolean value.
func PointerBool(b bool) *bool {
	return &b
}

// DecodeHasHGF - decodes has hgf
func DecodeHasHGF(filter bool) *bool {
	switch filter {
	case true:
		return PointerBool(true)
	default:
		return nil
	}
}
