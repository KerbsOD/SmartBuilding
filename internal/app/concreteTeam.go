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

func (ct *ConcreteTeam) AddTeamTo(aCollector *[]Team) {
	*aCollector = append(*aCollector, ct)
}

func (ct *ConcreteTeam) DisplayTimesToBuildOn(timesToBuild map[*ConcreteTeam]int, anArea float64) {
	timesToBuild[ct] = ct.DaysToBuild(anArea)
}

func (ct *ConcreteTeam) DisplayPricesToBuildOn(pricesToBuild map[*ConcreteTeam]int, anArea float64) {
	pricesToBuild[ct] = ct.price.PriceForWorking(ct.DaysToBuild(anArea))
}
