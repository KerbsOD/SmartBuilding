package app

import (
	"SmartBuilding/internal/errorMessage"
	"SmartBuilding/internal/generics"
	"errors"
)

type CombinedTeam struct {
	teams []Team
}

func NewCombinedTeam(teams []Team) *CombinedTeam {
	assertValidTeam(teams)
	ct := new(CombinedTeam)
	ct.teams = teams
	return ct
}

func (ct *CombinedTeam) DaysToBuild(anArea float64) int {
	anEqualAreaPerTeam := ct.areaDividedQuantityOfTeamMembers(anArea)
	maxDaysToBuild := ct.maxDaysToBuildTheSameAreaBetweenSubteams(anEqualAreaPerTeam)
	return maxDaysToBuild
}

func (ct *CombinedTeam) PriceToBuild(anArea float64) int {
	equalArea := ct.areaDividedQuantityOfTeamMembers(anArea)
	totalPrice := 0
	for _, team := range ct.teams {
		totalPrice = totalPrice + team.PriceToBuild(equalArea)
	}
	return totalPrice
}

func (ct *CombinedTeam) AddTeamTo(aCollector *[]Team) {
	for _, team := range ct.teams {
		team.AddTeamTo(aCollector)
	}
}

func (ct *CombinedTeam) DisplayTimesToBuildOn(timesToBuild map[*ConcreteTeam]int, anArea float64) {
	for _, team := range ct.teams {
		team.DisplayTimesToBuildOn(timesToBuild, ct.areaDividedQuantityOfTeamMembers(anArea))
	}
}

func (ct *CombinedTeam) DisplayPricesToBuildOn(pricesToBuild map[*ConcreteTeam]int, anArea float64) {
	for _, team := range ct.teams {
		team.DisplayPricesToBuildOn(pricesToBuild, ct.areaDividedQuantityOfTeamMembers(anArea))
	}
}

func (ct *CombinedTeam) CheapestTeamToBuild(anArea float64) Team {
	return generics.MinimizeElementByComparer(ct.teams, func(a, b Team) bool {
		equalArea := ct.areaDividedQuantityOfTeamMembers(anArea)
		teamAPrice := a.PriceToBuild(equalArea)
		teamBPrice := b.PriceToBuild(equalArea)
		return teamAPrice < teamBPrice
	})
}

func (ct *CombinedTeam) FastestTeamToBuild(anArea float64) Team {
	return generics.MinimizeElementByComparer(ct.teams, func(a, b Team) bool {
		equalArea := ct.areaDividedQuantityOfTeamMembers(anArea)
		teamADays := a.DaysToBuild(equalArea)
		teamBDays := b.DaysToBuild(equalArea)
		return teamADays < teamBDays
	})
}

func (ct *CombinedTeam) maxDaysToBuildTheSameAreaBetweenSubteams(anArea float64) int {
	return generics.MaxMapped(ct.teams, func(team Team) int { return team.DaysToBuild(anArea) })
}

func (ct *CombinedTeam) areaDividedQuantityOfTeamMembers(anArea float64) float64 {
	return anArea / float64(len(ct.teams))
}

func assertValidTeam(teams []Team) {
	assertValidTeamSize(teams)
	assertNotRepeatedTeams(teams)
}

func assertNotRepeatedTeams(teams []Team) {
	aTeamCollector := []Team{}
	for _, team := range teams {
		team.AddTeamTo(&aTeamCollector)
	}

	if len(generics.RepeatedElements(aTeamCollector)) > 0 {
		panic(errors.New(errorMessage.InvalidCombinedTeam))
	}
}

func assertValidTeamSize(teams []Team) {
	if len(teams) == 0 {
		panic(errors.New(errorMessage.InvalidCombinedTeam))
	}
}
