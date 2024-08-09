package app

type Price interface {
	PriceForWorking(aNumberOfDays int) int
}
