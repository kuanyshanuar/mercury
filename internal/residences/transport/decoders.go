package transport

import (
	residencesapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

func decodeResidenceV1(residence *residencesapiv1.ResidenceWrite) *domain.Residence {
	if residence == nil {
		return nil
	}

	flatPlans := make([]*domain.FlatPlan, len(residence.FlatPlans))
	for i, flatPlan := range residence.FlatPlans {
		if flatPlan == nil {
			continue
		}
		flatPlans[i] = &domain.FlatPlan{
			ResidenceID:   domain.ResidenceID(flatPlan.ResidenceId),
			NumberOfRooms: int(flatPlan.NumberOfRooms),
			Area:          flatPlan.Area,
			Price:         int(flatPlan.Price),
			Images:        flatPlan.Images,
		}
	}

	parkingTypes := make([]*domain.ParkingType, len(residence.ParkingTypeIds))
	for i, parkingTypeID := range residence.ParkingTypeIds {
		parkingTypes[i] = &domain.ParkingType{
			ID: domain.ParkingTypeID(parkingTypeID),
		}
	}

	interiorDecorationTypes := make([]*domain.InteriorDecoration, len(residence.InteriorDecorationIds))
	for i, interiorDecorationTypeID := range residence.InteriorDecorationIds {
		interiorDecorationTypes[i] = &domain.InteriorDecoration{
			ID: domain.InteriorDecorationID(interiorDecorationTypeID),
		}
	}

	heatingTypes := make([]*domain.HeatingType, len(residence.HeatingTypeIds))
	for i, heatingTypeID := range residence.HeatingTypeIds {
		heatingTypes[i] = &domain.HeatingType{
			ID: domain.HeatingTypeID(heatingTypeID),
		}
	}

	elevatorTypes := make([]*domain.ElevatorType, len(residence.ElevatorTypesIds))
	for i, elevatorTypeID := range residence.ElevatorTypesIds {
		elevatorTypes[i] = &domain.ElevatorType{
			ID: domain.ElevatorTypeID(elevatorTypeID),
		}
	}

	purchaseMethods := make([]*domain.PurchaseMethod, len(residence.PurchaseMethodIds))
	for i, methodID := range residence.PurchaseMethodIds {
		purchaseMethods[i] = &domain.PurchaseMethod{
			ID: domain.PurchaseMethodID(methodID),
		}
	}

	return &domain.Residence{
		StatusID:            residence.StatusId,
		UserID:              residence.UserId,
		CityID:              residence.CityId,
		HousingClassID:      domain.HousingClassID(residence.HousingClassId),
		ConstructionTypeID:  residence.ConstructionTypeId,
		Title:               residence.Title,
		Description:         residence.Description,
		Address:             residence.Address,
		DistrictID:          domain.DistrictID(residence.DistrictId),
		Latitude:            residence.Latitude,
		Longitude:           residence.Longitude,
		DeadlineYear:        residence.DeadlineYear,
		DeadlineQuarter:     residence.DeadlineQuarter,
		FlatsCount:          residence.FlatsCount,
		FloorsMax:           residence.FloorsMax,
		RoomsMin:            residence.RoomsMin,
		RoomsMax:            residence.RoomsMax,
		CeilingHeight:       residence.CeilingHeight,
		PricePerSquareMin:   residence.PricePerSquareMin,
		PriceMin:            residence.PriceMin,
		PriceMax:            residence.PriceMax,
		AreaMin:             residence.AreaMin,
		AreaMax:             residence.AreaMax,
		ParkingTypes:        parkingTypes,
		HeatingTypes:        heatingTypes,
		InteriorDecorations: interiorDecorationTypes,
		ElevatorTypes:       elevatorTypes,
		PurchaseMethods:     purchaseMethods,
		TitleImage:          residence.TitleImage,
		Images:              residence.Images,
		FlatPlans:           flatPlans,
		Slug:                residence.Slug,
		SaleStatusID:        residence.SaleStatusId,
	}
}

func encodeResidenceV1(residence *domain.Residence) *residencesapiv1.ResidenceRead {
	if residence == nil {
		return nil
	}

	return &residencesapiv1.ResidenceRead{
		Id:                  int64(residence.ID),
		StatusId:            residence.StatusID,
		Status:              encodeStatusV1(&residence.Status),
		SaleStatus:          encodeSaleStatusV1(&residence.SaleStatus),
		Builder:             encodeBuilderV1(residence.User),
		UserId:              residence.UserID,
		CityId:              residence.CityID,
		City:                encodeCityV1(residence.City),
		HousingClassId:      int64(residence.HousingClassID),
		HousingClass:        encodeHousingClassV1(residence.HousingClass),
		ConstructionTypeId:  residence.ConstructionTypeID,
		ConstructionType:    encodeConstructionTypeV1(residence.ConstructionType),
		Title:               residence.Title,
		TitleImage:          residence.TitleImage,
		Description:         residence.Description,
		Address:             residence.Address,
		District:            encodeDistrictV1(residence.District),
		Latitude:            residence.Latitude,
		Longitude:           residence.Longitude,
		DeadlineYear:        residence.DeadlineYear,
		DeadlineQuarter:     residence.DeadlineQuarter,
		FlatsCount:          residence.FlatsCount,
		FloorsMax:           residence.FloorsMax,
		RoomsMin:            residence.RoomsMin,
		RoomsMax:            residence.RoomsMax,
		CeilingHeight:       residence.CeilingHeight,
		PricePerSquareMin:   residence.PricePerSquareMin,
		PriceMin:            residence.PriceMin,
		PriceMax:            residence.PriceMax,
		AreaMin:             residence.AreaMin,
		AreaMax:             residence.AreaMax,
		Views:               residence.Views,
		Likes:               residence.Likes,
		FlatPlans:           encodeFlatPlansV1(residence.FlatPlans),
		ParkingTypes:        encodeParkingTypesV1(residence.ParkingTypes),
		InteriorDecorations: encodeInteriorDecorationsV1(residence.InteriorDecorations),
		HeatingTypes:        encodeHeatingTypesV1(residence.HeatingTypes),
		ElevatorTypes:       encodeElevatorTypesV1(residence.ElevatorTypes),
		PurchaseMethods:     encodePurchaseMethodsV1(residence.PurchaseMethods),
		Images:              residence.Images,
		IsFavourite:         residence.IsFavourite,
		CreatedAt:           residence.CreatedAt,
		UpdatedAt:           residence.UpdatedAt,
		DeletedAt:           int64(residence.DeletedAt),
	}
}

func encodePurchaseMethodsV1(r []*domain.PurchaseMethod) []*residencesapiv1.PurchaseMethodRead {
	if r == nil {
		return nil
	}

	methods := make([]*residencesapiv1.PurchaseMethodRead, len(r))
	for i, opt := range r {
		methods[i] = encodePurchaseMethodV1(opt)
	}

	return methods
}

func encodePurchaseMethodV1(r *domain.PurchaseMethod) *residencesapiv1.PurchaseMethodRead {
	if r == nil {
		return nil
	}

	return &residencesapiv1.PurchaseMethodRead{
		Id:   int64(r.ID),
		Name: r.Name,
	}
}

func encodeBuilderV1(r *domain.User) *residencesapiv1.BuilderRead {
	if r == nil {
		return nil
	}

	return &residencesapiv1.BuilderRead{
		Id:                      int64(r.ID),
		FirstName:               r.FirstName,
		LastName:                r.LastName,
		Email:                   r.Email,
		ConsultationPhoneNumber: r.ConsultationPhoneNumber,
		Image:                   r.Image,
	}
}

func encodeCityV1(r *domain.City) *residencesapiv1.City {
	if r == nil {
		return nil
	}

	return &residencesapiv1.City{
		Id:   int64(r.ID),
		Name: r.Name,
	}
}

func encodeListResidencesV1(residences []*domain.Residence) []*residencesapiv1.ResidenceRead {

	result := make([]*residencesapiv1.ResidenceRead, len(residences))
	for i, residence := range residences {
		result[i] = encodeResidenceV1(residence)
	}

	return result
}

func encodeStatusV1(opt *domain.Status) *residencesapiv1.StatusRead {
	if opt == nil {
		return nil
	}

	return &residencesapiv1.StatusRead{
		Id:   int64(opt.ID),
		Name: opt.Name,
	}
}

func encodeSaleStatusV1(opt *domain.SaleStatus) *residencesapiv1.SaleStatusRead {
	if opt == nil {
		return nil
	}

	return &residencesapiv1.SaleStatusRead{
		Id:   int64(opt.ID),
		Name: opt.Name,
	}
}

func encodeHousingClassV1(opt domain.HousingClass) *residencesapiv1.HouseClassRead {

	return &residencesapiv1.HouseClassRead{
		Id:   int64(opt.ID),
		Name: opt.Name,
	}
}

func encodeDistrictV1(opt domain.District) *residencesapiv1.DistrictRead {

	return &residencesapiv1.DistrictRead{
		Id:   int64(opt.ID),
		Name: opt.Name,
	}
}

func encodeConstructionTypeV1(opt domain.ConstructionType) *residencesapiv1.ConstructionTypeRead {

	return &residencesapiv1.ConstructionTypeRead{
		Id:   int64(opt.ID),
		Name: opt.Name,
	}
}

func encodeFlatPlansV1(flatPlans []*domain.FlatPlan) []*residencesapiv1.ResidenceFlatPlanRead {
	if flatPlans == nil {
		return nil
	}

	result := make([]*residencesapiv1.ResidenceFlatPlanRead, len(flatPlans))
	for i, opt := range flatPlans {
		result[i] = &residencesapiv1.ResidenceFlatPlanRead{
			Id:            int64(opt.ID),
			ResidenceId:   int64(opt.ResidenceID),
			NumberOfRooms: int32(opt.NumberOfRooms),
			Area:          opt.Area,
			Price:         int32(opt.Price),
			Images:        opt.Images,
			CreatedAt:     opt.CreatedAt,
			UpdatedAt:     opt.UpdatedAt,
			DeletedAt:     int64(opt.DeletedAt),
		}
	}

	return result
}

func encodeParkingTypesV1(parkingTypes []*domain.ParkingType) []*residencesapiv1.ParkingTypeRead {
	if parkingTypes == nil {
		return nil
	}

	result := make([]*residencesapiv1.ParkingTypeRead, len(parkingTypes))
	for i, opt := range parkingTypes {
		result[i] = &residencesapiv1.ParkingTypeRead{
			Id:   int64(opt.ID),
			Name: opt.Name,
		}
	}

	return result
}

func encodeInteriorDecorationsV1(interiorDecorationTypes []*domain.InteriorDecoration) []*residencesapiv1.InteriorDecorationRead {
	if interiorDecorationTypes == nil {
		return nil
	}

	result := make([]*residencesapiv1.InteriorDecorationRead, len(interiorDecorationTypes))
	for i, opt := range interiorDecorationTypes {
		result[i] = &residencesapiv1.InteriorDecorationRead{
			Id:   int64(opt.ID),
			Name: opt.Name,
		}
	}

	return result
}

func encodeHeatingTypesV1(heatingTypes []*domain.HeatingType) []*residencesapiv1.HeatingTypeRead {
	if heatingTypes == nil {
		return nil
	}

	result := make([]*residencesapiv1.HeatingTypeRead, len(heatingTypes))
	for i, opt := range heatingTypes {
		result[i] = &residencesapiv1.HeatingTypeRead{
			Id:   int64(opt.ID),
			Name: opt.Name,
		}
	}

	return result
}

func encodeElevatorTypesV1(elevatorTypes []*domain.ElevatorType) []*residencesapiv1.ElevatorTypeRead {
	if elevatorTypes == nil {
		return nil
	}

	result := make([]*residencesapiv1.ElevatorTypeRead, len(elevatorTypes))
	for i, opt := range elevatorTypes {
		result[i] = &residencesapiv1.ElevatorTypeRead{
			Id:   int64(opt.ID),
			Name: opt.Name,
		}
	}

	return result
}

func decodeResidenceSearchCriteria(apiCriteria *residencesapiv1.ResidencesSearchCriteria) domain.ResidenceSearchCriteria {

	if apiCriteria == nil {
		// default criteria.
		return domain.ResidenceSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.ResidenceSearchCriteria{
		Page:                  helpers.DecodePageRequestV1(apiCriteria.PageRequest),
		Sorts:                 helpers.DecodeSorts(apiCriteria.Sorts),
		Title:                 apiCriteria.Title,
		BuilderIDs:            apiCriteria.BuilderIds,
		CityID:                apiCriteria.CityId,
		DistrictID:            apiCriteria.DistrictId,
		HousingClassID:        apiCriteria.HousingClassId,
		StatusID:              apiCriteria.StatusId,
		RoomsMin:              apiCriteria.RoomsMin,
		RoomsMax:              apiCriteria.RoomsMax,
		CeilingHeightMin:      apiCriteria.CeilingHeightMin,
		CeilingHeightMax:      apiCriteria.CeilingHeightMax,
		HasHGF:                helpers.DecodeHasHGF(apiCriteria.HasHgf),
		AreaMin:               float64(apiCriteria.AreaMin),
		AreaMax:               float64(apiCriteria.AreaMax),
		PriceMin:              apiCriteria.PriceMin,
		PriceMax:              apiCriteria.PriceMax,
		FloorsMin:             apiCriteria.FloorsMin,
		FloorsMax:             apiCriteria.FloorsMax,
		ConstructionTypesIDs:  apiCriteria.ConstructionTypeIds,
		ParkingTypesIds:       apiCriteria.ParkingTypeIds,
		InteriorDecorationIDs: apiCriteria.InteriorDecorationIds,
		HeatingTypesIDs:       apiCriteria.HeatingTypesIds,
		PurchaseMethodsIDs:    apiCriteria.PurchaseMethodsIds,
		ElevatorTypesIDs:      apiCriteria.ElevatorTypesIds,
		UserID:                apiCriteria.UserId,
	}
}

func decodeResidenceIDs(ids []int64) []domain.ResidenceID {

	residenceIDs := make([]domain.ResidenceID, len(ids))
	for i, opt := range ids {
		residenceIDs[i] = domain.ResidenceID(opt)
	}

	return residenceIDs
}

func decodeFlatPlanV1(w *residencesapiv1.ResidenceFlatPlanWrite) *domain.FlatPlan {
	if w == nil {
		return nil
	}

	return &domain.FlatPlan{
		ResidenceID:   domain.ResidenceID(w.ResidenceId),
		NumberOfRooms: int(w.NumberOfRooms),
		Area:          w.Area,
		Price:         int(w.Price),
		Images:        w.Images,
	}
}

func encodeFlatPlanV1(r *domain.FlatPlan) *residencesapiv1.ResidenceFlatPlanRead {
	if r == nil {
		return nil
	}

	return &residencesapiv1.ResidenceFlatPlanRead{
		Id:            int64(r.ID),
		ResidenceId:   int64(r.ResidenceID),
		NumberOfRooms: int32(r.NumberOfRooms),
		Area:          r.Area,
		Price:         int32(r.Price),
		Images:        r.Images,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
		DeletedAt:     int64(r.DeletedAt),
	}
}
