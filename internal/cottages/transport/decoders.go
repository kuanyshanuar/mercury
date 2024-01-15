package transport

import (
	cottageserviceapi "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

func decodeHousePlanV1(w *cottageserviceapi.HousePlanWrite) *domain.HousePlan {
	if w == nil {
		return nil
	}

	return &domain.HousePlan{
		CottageID:      w.CottageId,
		Title:          w.Title,
		NumberOfRooms:  w.NumberOfRooms,
		Area:           w.Area,
		Longitude:      w.Longitude,
		Territory:      w.Territory,
		CeilingHeight:  w.CeilingHeight,
		Price:          w.Price,
		PricePerSquare: w.PricePerSquare,
		PlanImages:     w.PlanImages,
		HouseImages:    w.HouseImages,
		HousingClassID: w.HousingClassId,
	}
}

func decodeCriteriaV1(apiCriteria *cottageserviceapi.CottageSearchCriteria) domain.CottageSearchCriteria {
	if apiCriteria == nil {
		// default criteria.
		return domain.CottageSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   domain.DefaultPageSize,
			},
		}
	}

	return domain.CottageSearchCriteria{
		Page:               helpers.DecodePageRequestV1(apiCriteria.Page),
		Sorts:              helpers.DecodeSorts(apiCriteria.Sort),
		Title:              apiCriteria.Title,
		BuilderIDs:         apiCriteria.BuilderIds,
		CityID:             apiCriteria.CityId,
		DistrictID:         apiCriteria.DistrictId,
		StatusID:           apiCriteria.StatusId,
		RoomsMin:           apiCriteria.RoomsMin,
		RoomsMax:           apiCriteria.RoomsMax,
		CeilingHeightMin:   apiCriteria.CeilingHeightMin,
		CeilingHeightMax:   apiCriteria.CeilingHeightMax,
		AreaMin:            apiCriteria.AreaMin,
		AreaMax:            apiCriteria.AreaMax,
		PriceMin:           apiCriteria.PriceMin,
		PriceMax:           apiCriteria.PriceMax,
		InteriorDecoration: apiCriteria.InteriorDecoration,
		HeatingTypes:       apiCriteria.HeatingTypes,
		PurchaseMethods:    apiCriteria.PurchaseMethods,
		ElevatorTypes:      apiCriteria.ElevatorTypes,
		HouseType:          apiCriteria.HousingClass,
		FloorsMin:          apiCriteria.FloorsMin,
		FloorsMax:          apiCriteria.FloorsMax,
		UserID:             apiCriteria.UserId,
	}
}

func decodeCottageWriteV1(w *cottageserviceapi.CottageWrite) *domain.Cottage {
	if w == nil {
		return nil
	}

	// Decode windows
	windows := make([]*domain.WindowType, len(w.WindowTypes))
	for i, val := range w.WindowTypes {
		windows[i] = &domain.WindowType{
			ID: val,
		}
	}

	// Decode cottage plans
	cottagePlans := make([]*domain.HousePlan, len(w.HousePlans))
	for i, val := range w.HousePlans {
		plan := &domain.HousePlan{
			CottageID:      val.CottageId,
			Title:          val.Title,
			NumberOfRooms:  val.NumberOfRooms,
			Area:           val.Area,
			Longitude:      val.Longitude,
			Territory:      val.Territory,
			CeilingHeight:  val.CeilingHeight,
			Price:          val.Price,
			PricePerSquare: val.PricePerSquare,
			PlanImages:     val.PlanImages,
			HouseImages:    val.HouseImages,
			HousingClassID: val.HousingClassId,
		}
		cottagePlans[i] = plan
	}

	// Decode wall types
	wallType := make([]*domain.WallType, len(w.WallTypes))
	for i, val := range w.WallTypes {
		wallType[i] = &domain.WallType{
			ID: val,
		}
	}

	// Decode wall types
	elevatorTypes := make([]*domain.ElevatorType, len(w.ElevatorTypes))
	for i, val := range w.ElevatorTypes {
		elevatorTypes[i] = &domain.ElevatorType{
			ID: domain.ElevatorTypeID(val),
		}
	}

	// Decode warming types
	warmingTypes := make([]*domain.WarmingType, len(w.WarmingTypes))
	for i, val := range w.WarmingTypes {
		warmingTypes[i] = &domain.WarmingType{
			ID: val,
		}
	}

	// Decode interior decoration types
	interiorDecorationTypes := make([]*domain.InteriorDecoration, len(w.InteriorDecorations))
	for i, val := range w.InteriorDecorations {
		interiorDecorationTypes[i] = &domain.InteriorDecoration{
			ID: domain.InteriorDecorationID(val),
		}
	}

	// Decode purchase methods
	purchaseMethods := make([]*domain.PurchaseMethod, len(w.PurchaseMethods))
	for i, val := range w.PurchaseMethods {
		purchaseMethods[i] = &domain.PurchaseMethod{
			ID: domain.PurchaseMethodID(val),
		}
	}

	// Decode heating types
	heatingTypes := make([]*domain.HeatingType, len(w.HeatingTypes))
	for i, val := range w.HeatingTypes {
		heatingTypes[i] = &domain.HeatingType{
			ID: domain.HeatingTypeID(val),
		}
	}
	// Decode parking types
	parkingTypes := make([]*domain.ParkingType, len(w.ParkingTypes))
	for i, val := range w.ParkingTypes {
		parkingTypes[i] = &domain.ParkingType{
			ID: domain.ParkingTypeID(val),
		}
	}
	return &domain.Cottage{
		CityID:              w.CityId,
		UserID:              w.UserId,
		StatusID:            w.StatusId,
		DistrictID:          w.DistrictId,
		HousingClassID:      domain.HousingClassID(w.HousingClassId),
		Title:               w.Title,
		Description:         w.Description,
		Address:             w.Address,
		Latitude:            w.Latitude,
		Longitude:           w.Longitude,
		Territory:           w.Territory,
		CeilingHeightMin:    w.CeilingHeightMin,
		CeilingHeightMax:    w.CeilingHeightMax,
		BuildingArea:        w.BuildingArea,
		HouseAmount:         w.HouseAmount,
		FloorsCount:         w.FloorsCount,
		Facade:              w.Facade,
		CanRePlan:           w.CanReplan,
		AreaMin:             w.AreaMin,
		AreaMax:             w.AreaMax,
		PricePerSquareMin:   w.PricePerSquareMin,
		PricePerSquareMax:   w.PricePerSquareMax,
		WindowTypes:         windows,
		HousePlans:          cottagePlans,
		WallTypes:           wallType,
		ElevatorTypes:       elevatorTypes,
		WarmingTypes:        warmingTypes,
		ParkingTypes:        parkingTypes,
		InteriorDecorations: interiorDecorationTypes,
		PurchaseMethods:     purchaseMethods,
		HeatingTypes:        heatingTypes,
		Images:              w.Images,
		RoomsMin:            w.RoomsMin,
		RoomsMax:            w.RoomsMax,
	}
}

func encodeCottageReadV1(r *domain.Cottage) *cottageserviceapi.CottageRead {
	if r == nil {
		return nil
	}

	return &cottageserviceapi.CottageRead{
		Id:                  r.ID,
		CityId:              r.CityID,
		UserId:              r.UserID,
		StatusId:            r.StatusID,
		Status:              encodeStatusV1(r.Status),
		SaleStatus:          encodeSaleStatusV1(r.SaleStatus),
		DistrictId:          r.DistrictID,
		District:            encodeDistrictV1(r.District),
		Title:               r.Title,
		Description:         r.Description,
		Address:             r.Address,
		Latitude:            r.Latitude,
		Longitude:           r.Longitude,
		Territory:           r.Territory,
		HousingClassId:      int64(r.HousingClassID),
		CeilingHeightMin:    r.CeilingHeightMin,
		CeilingHeightMax:    r.CeilingHeightMax,
		BuildingArea:        r.BuildingArea,
		HouseAmount:         r.HouseAmount,
		AreaMin:             r.AreaMin,
		AreaMax:             r.AreaMax,
		FloorsCount:         r.FloorsCount,
		Facade:              r.Facade,
		WindowTypes:         encodeWindowV1(r.WindowTypes),
		CanReplan:           r.CanRePlan,
		PricePerSquareMin:   r.PricePerSquareMin,
		PricePerSquareMax:   r.PricePerSquareMax,
		RoomsMin:            r.RoomsMin,
		RoomsMax:            r.RoomsMax,
		HousePlans:          encodeHousePlansV1(r.HousePlans),
		WallTypes:           encodeWallTypesV1(r.WallTypes),
		ElevatorTypes:       encodeElevatorTypesV1(r.ElevatorTypes),
		WarmingTypes:        encodeWarmingTypesV1(r.WarmingTypes),
		HousingClass:        encodeHouseTypeV1(r.HousingClass),
		InteriorDecorations: encodeInteriorDecorationsV1(r.InteriorDecorations),
		PurchaseMethods:     encodePurchaseMethodsV1(r.PurchaseMethods),
		HeatingTypes:        encodeHeatingTypesV1(r.HeatingTypes),
		ParkingTypes:        encodeParkingTypesV1(r.ParkingTypes),
		Images:              r.Images,
		CreatedAt:           r.CreatedAt,
		UpdatedAt:           r.UpdatedAt,
		DeletedAt:           int64(r.DeletedAt),
	}
}

func encodeHousePlansV1(plans []*domain.HousePlan) []*cottageserviceapi.HousePlanRead {
	if plans == nil {
		return nil
	}

	returnVal := make([]*cottageserviceapi.HousePlanRead, len(plans))
	for i, val := range plans {
		returnVal[i] = encodeHousePlanV1(val)
	}

	return returnVal
}

func encodeCottageListV1(request []*domain.Cottage) []*cottageserviceapi.CottageRead {
	if request == nil {
		return nil
	}

	returnVal := make([]*cottageserviceapi.CottageRead, len(request))
	for i, val := range request {
		returnVal[i] = encodeCottageReadV1(val)
	}
	return returnVal
}

func encodeHousePlanV1(r *domain.HousePlan) *cottageserviceapi.HousePlanRead {
	if r == nil {
		return nil
	}

	return &cottageserviceapi.HousePlanRead{
		Id:             r.ID,
		CottageId:      r.CottageID,
		Title:          r.Title,
		NumberOfRooms:  r.NumberOfRooms,
		Area:           r.Area,
		Longitude:      r.Longitude,
		Territory:      r.Territory,
		CeilingHeight:  r.CeilingHeight,
		Price:          r.Price,
		PricePerSquare: r.PricePerSquare,
		PlanImages:     r.PlanImages,
		HouseImages:    r.HouseImages,
		HousingClassId: r.HousingClassID,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
		DeletedAt:      int64(r.DeletedAt),
	}
}
func encodeParkingTypesV1(request []*domain.ParkingType) []*cottageserviceapi.ParkingTypeRead {
	if request == nil {
		return nil
	}

	response := make([]*cottageserviceapi.ParkingTypeRead, len(request))
	for i, val := range request {
		response[i] = &cottageserviceapi.ParkingTypeRead{
			Id:   int64(val.ID),
			Name: val.Name,
		}
	}
	return response
}
func encodeWindowV1(request []*domain.WindowType) []*cottageserviceapi.WindowTypeRead {
	if request == nil {
		return nil
	}

	windowReturn := make([]*cottageserviceapi.WindowTypeRead, len(request))
	for i, val := range request {
		windowReturn[i] = &cottageserviceapi.WindowTypeRead{
			Id:   val.ID,
			Name: val.Name,
		}
	}
	return windowReturn
}
func encodeStatusV1(request *domain.Status) *cottageserviceapi.StatusRead {
	if request == nil {
		return nil
	}

	return &cottageserviceapi.StatusRead{
		Id:   int64(request.ID),
		Name: request.Name,
	}
}
func encodeSaleStatusV1(request *domain.SaleStatus) *cottageserviceapi.SaleStatusRead {
	if request == nil {
		return nil
	}
	return &cottageserviceapi.SaleStatusRead{
		Id:   int64(request.ID),
		Name: request.Name,
	}
}

func encodeHouseTypeV1(request domain.HousingClass) *cottageserviceapi.HouseClassRead {

	return &cottageserviceapi.HouseClassRead{
		Id:   int64(request.ID),
		Name: request.Name,
	}
}

func encodeWallTypesV1(wallTypes []*domain.WallType) []*cottageserviceapi.WallTypeRead {
	if wallTypes == nil {
		return nil
	}

	returnVal := make([]*cottageserviceapi.WallTypeRead, len(wallTypes))
	for i, val := range wallTypes {
		returnVal[i] = &cottageserviceapi.WallTypeRead{
			Id:   val.ID,
			Name: val.Name,
		}
	}
	return returnVal
}

func encodeElevatorTypesV1(elevators []*domain.ElevatorType) []*cottageserviceapi.ElevatorTypeRead {
	if elevators == nil {
		return nil
	}

	returnVal := make([]*cottageserviceapi.ElevatorTypeRead, len(elevators))
	for i, val := range elevators {
		returnVal[i] = &cottageserviceapi.ElevatorTypeRead{
			Id:   int64(val.ID),
			Name: val.Name,
		}
	}
	return returnVal
}

func encodeWarmingTypesV1(request []*domain.WarmingType) []*cottageserviceapi.WarmingTypeRead {
	if request == nil {
		return nil
	}

	returnVal := make([]*cottageserviceapi.WarmingTypeRead, len(request))
	for i, val := range request {
		returnVal[i] = &cottageserviceapi.WarmingTypeRead{
			Id:   val.ID,
			Name: val.Name,
		}
	}
	return returnVal
}

func encodeInteriorDecorationsV1(request []*domain.InteriorDecoration) []*cottageserviceapi.InteriorDecorationRead {
	if request == nil {
		return nil
	}

	returnVal := make([]*cottageserviceapi.InteriorDecorationRead, len(request))
	for i, val := range request {
		returnVal[i] = &cottageserviceapi.InteriorDecorationRead{
			Id:   int64(val.ID),
			Name: val.Name,
		}
	}
	return returnVal
}

func encodePurchaseMethodsV1(request []*domain.PurchaseMethod) []*cottageserviceapi.PurchaseMethodRead {
	if request == nil {
		return nil
	}

	returnVal := make([]*cottageserviceapi.PurchaseMethodRead, len(request))
	for i, val := range request {
		returnVal[i] = &cottageserviceapi.PurchaseMethodRead{
			Id:   int64(val.ID),
			Name: val.Name,
		}
	}
	return returnVal
}

func encodeHeatingTypesV1(request []*domain.HeatingType) []*cottageserviceapi.HeatingTypeRead {
	if request == nil {
		return nil
	}

	returnVal := make([]*cottageserviceapi.HeatingTypeRead, len(request))
	for i, val := range request {
		returnVal[i] = &cottageserviceapi.HeatingTypeRead{
			Id:   int64(val.ID),
			Name: val.Name,
		}
	}

	return returnVal
}

func encodeDistrictV1(opt *domain.District) *cottageserviceapi.DistrictRead {
	if opt == nil {
		return nil
	}

	return &cottageserviceapi.DistrictRead{
		Id:   int64(opt.ID),
		Name: opt.Name,
	}
}
