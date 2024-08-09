package app

import (
	"SmartBuilding/internal/errorMessage"
	"errors"
)

type VariablePrice struct {
	regularDailyPrice     int
	rainingDailyPrice     int
	meteorologicalService MeteorologicalService
}

func NewVariablePrice(aRegularPrice, aRainingPrice int, aService MeteorologicalService) *VariablePrice {

	if aRegularPrice > aRainingPrice {
		panic(errors.New(errorMessage.InvalidRainingPrice))
	}

	vp := new(VariablePrice)
	vp.regularDailyPrice = aRegularPrice
	vp.rainingDailyPrice = aRainingPrice
	vp.meteorologicalService = aService
	return vp
}

func (vp VariablePrice) PriceForWorking(aNumberOfDays int) int {
	rainingDays := vp.meteorologicalService.RainingDayAmongTheNext(aNumberOfDays)
	regularDays := aNumberOfDays - rainingDays
	return (rainingDays * vp.rainingDailyPrice) + (regularDays * vp.regularDailyPrice)
}
