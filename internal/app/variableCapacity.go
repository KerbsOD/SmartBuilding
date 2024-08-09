package app

import (
	"SmartBuilding/internal/errorMessage"
	"errors"
)

type VariableCapacity struct {
	initialDays     int
	initialCapacity float64
	regularCapacity float64
}

func NewVariableCapacity(aNumberOfDays int, initialCapacity, regularCapacity float64) *VariableCapacity {
	assertValidCapacities(initialCapacity, regularCapacity)
	vc := new(VariableCapacity)
	vc.initialDays = aNumberOfDays
	vc.initialCapacity = initialCapacity
	vc.regularCapacity = regularCapacity
	return vc
}

func (fc VariableCapacity) DaysToComplete(anArea float64) int {
	daysToCompleteUsingInitialCapacity := FullAreaDivision(anArea, fc.initialCapacity)
	if daysToCompleteUsingInitialCapacity <= fc.initialDays {
		return daysToCompleteUsingInitialCapacity
	}

	remainingArea := anArea - (fc.initialCapacity * float64(fc.initialDays))
	daysToCompleteUsingRegularCapacity := FullAreaDivision(remainingArea, fc.regularCapacity)

	return fc.initialDays + daysToCompleteUsingRegularCapacity
}

func assertValidCapacities(initialCapacity float64, regularCapacity float64) {
	if initialCapacity < regularCapacity {
		panic(errors.New(errorMessage.InvalidCapacitiesErrorMessage))
	}
}
