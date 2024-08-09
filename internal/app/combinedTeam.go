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
	return generics.MaxMapped(ct.teams, func(team Team) int { return team.DaysToBuild(ct.subteamEqualArea(anArea)) })
}

func (ct *CombinedTeam) PriceToBuild(anArea float64) int {
	totalPrice := 0
	for _, team := range ct.teams {
		totalPrice = totalPrice + team.PriceToBuild(ct.subteamEqualArea(anArea))
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
		team.DisplayTimesToBuildOn(timesToBuild, ct.subteamEqualArea(anArea))
	}
}

func (ct *CombinedTeam) DisplayPricesToBuildOn(pricesToBuild map[*ConcreteTeam]int, anArea float64) {
	for _, team := range ct.teams {
		team.DisplayPricesToBuildOn(pricesToBuild, ct.subteamEqualArea(anArea))
	}
}

func (ct *CombinedTeam) CheapestTeamToBuild(anArea float64) Team {
	return generics.MinimizeElementByComparer(ct.teams, func(a, b Team) bool {
		return a.PriceToBuild(ct.subteamEqualArea(anArea)) < b.PriceToBuild(ct.subteamEqualArea(anArea))
	})
}

func (ct *CombinedTeam) FastestTeamToBuild(anArea float64) Team {
	return generics.MinimizeElementByComparer(ct.teams, func(a, b Team) bool {
		return a.DaysToBuild(ct.subteamEqualArea(anArea)) < b.DaysToBuild(ct.subteamEqualArea(anArea))
	})
}

func (ct *CombinedTeam) subteamEqualArea(anArea float64) float64 {
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
