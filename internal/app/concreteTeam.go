package app

type ConcreteTeam struct {
	capacity Capacity
	price    Price
}

func NewConcreteTeam(aCapacity Capacity, aPrice Price) *ConcreteTeam {
	ct := new(ConcreteTeam)
	ct.capacity = aCapacity
	ct.price = aPrice
	return ct
}

func (ct *ConcreteTeam) DaysToBuild(anArea float64) int {
	return ct.capacity.DaysToComplete(anArea)
}

func (ct *ConcreteTeam) PriceToBuild(anArea float64) int {
	return ct.price.PriceForWorking(ct.DaysToBuild(anArea))
}
