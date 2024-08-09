package app

type MeteorologicalService interface {
	RainingDayAmongTheNext(aNumberOfDays int) int
}
