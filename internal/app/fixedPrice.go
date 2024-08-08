package app

type FixedPrice struct {
	dailyPrice int
}

func NewFixedPrice(aPrice int) *FixedPrice {
	fp := new(FixedPrice)
	fp.dailyPrice = aPrice
	return fp
}

func (fp FixedPrice) PriceForWorking(aNumberOfDays int) int {
	return fp.dailyPrice * aNumberOfDays
}
