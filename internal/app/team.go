package app

type Team interface {
	DaysToBuild(anArea float64) int
	PriceToBuild(anArea float64) int
	AddTeamTo(aCollector *[]Team)
	DisplayTimesToBuildOn(timesToBuild map[*ConcreteTeam]int, anArea float64)
	DisplayPricesToBuildOn(pricesToBuild map[*ConcreteTeam]int, anArea float64)
}
