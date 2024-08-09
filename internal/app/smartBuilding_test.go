package app

import (
	"SmartBuilding/internal/errorMessage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SmartBuildingTestSuite struct {
	suite.Suite
}

type MockMeteorologicalService struct {
	mock.Mock
}

func (ms *MockMeteorologicalService) RainingDayAmongTheNext(aNumberOfDays int) int {
	args := ms.Called(aNumberOfDays)
	return args.Int(0)
}

func TestSmartBuildingTestSuite(t *testing.T) {
	suite.Run(t, new(SmartBuildingTestSuite))
}

func (suite *SmartBuildingTestSuite) Test01FixedCapacityDaysToBuildAreCapacityOverArea() {
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	assert.Equal(suite.T(), 1, equipoRojo.DaysToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test02FixedCapacityDaysToBuildAreFullDays() {
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	assert.Equal(suite.T(), 2, equipoRojo.DaysToBuild(150))
}

func (suite *SmartBuildingTestSuite) Test03FixedPricePriceToBuildIsPriceTimesDaysToBuild() {
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	assert.Equal(suite.T(), 1000, equipoRojo.PriceToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test04FixedPriceChargesFullDays() {
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	assert.Equal(suite.T(), 2000, equipoRojo.PriceToBuild(150))
}

func (suite *SmartBuildingTestSuite) Test05VariableCapacityDaysToBuildUsesInitialCapacityTheInitialDays() {
	equipoRojo := NewConcreteTeam(NewVariableCapacity(2, 50, 25), NewFixedPrice(1000))
	assert.Equal(suite.T(), 2, equipoRojo.DaysToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test06VariableCapacityDaysToBuildGreaterThanInitialDaysUsesReminderCapacity() {
	equipoRojo := NewConcreteTeam(NewVariableCapacity(2, 50, 25), NewFixedPrice(1000))
	assert.Equal(suite.T(), 4, equipoRojo.DaysToBuild(150))
}

func (suite *SmartBuildingTestSuite) Test07VariableCapacityInitialDaysToBuildAreFullDays() {
	equipoRojo := NewConcreteTeam(NewVariableCapacity(2, 50, 25), NewFixedPrice(1000))
	assert.Equal(suite.T(), 2, equipoRojo.DaysToBuild(80))
}

func (suite *SmartBuildingTestSuite) Test08VariableCapacityRemainingDaysToBuildAreFullDays() {
	equipoRojo := NewConcreteTeam(NewVariableCapacity(2, 50, 25), NewFixedPrice(1000))
	assert.Equal(suite.T(), 4, equipoRojo.DaysToBuild(130))
}

func (suite *SmartBuildingTestSuite) Test09VariableCapacityInitialCapacityMustBeSmallerThanRemainingCapacity() {
	assert.PanicsWithError(suite.T(), errorMessage.InvalidCapacitiesErrorMessage, func() {
		NewConcreteTeam(NewVariableCapacity(2, 25, 50), NewFixedPrice(1000))
	})
}

func (suite *SmartBuildingTestSuite) Test10VariablePriceUsesRegularPriceWhenThereAreNotRainingDays() {
	mockService := new(MockMeteorologicalService)
	mockService.On("RainingDayAmongTheNext", mock.Anything).Return(0)
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewVariablePrice(1000, 2000, mockService))
	assert.Equal(suite.T(), 1000, equipoRojo.PriceToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test11VariablePriceUsesRainingPriceWhenThereAreRainingDays() {
	mockService := new(MockMeteorologicalService)
	mockService.On("RainingDayAmongTheNext", mock.Anything).Return(1)
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewVariablePrice(1000, 2000, mockService))
	assert.Equal(suite.T(), 2000, equipoRojo.PriceToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test12VariablePriceUsesRainingPriceOnRainingDaysAndRegularPriceOnRegularDays() {
	mockService := new(MockMeteorologicalService)
	mockService.On("RainingDayAmongTheNext", mock.Anything).Return(2)
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewVariablePrice(1000, 2000, mockService))
	assert.Equal(suite.T(), 6000, equipoRojo.PriceToBuild(400))
}

func (suite *SmartBuildingTestSuite) Test13RainingPriceIsGreaterThanRegularPrice() {
	assert.PanicsWithError(suite.T(), errorMessage.InvalidRainingPrice, func() {
		NewConcreteTeam(NewFixedCapacity(100), NewVariablePrice(1000, 500, new(MockMeteorologicalService)))
	})
}

func (suite *SmartBuildingTestSuite) Test14CombinedTeamCantBeEmpty() {
	assert.PanicsWithError(suite.T(), errorMessage.InvalidCombinedTeam, func() {
		NewCombinedTeam([]Team{})
	})
}

func (suite *SmartBuildingTestSuite) Test15CombinedTeamCantHaveDirectRepeatedTeams() {
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	assert.PanicsWithError(suite.T(), errorMessage.InvalidCombinedTeam, func() {
		NewCombinedTeam([]Team{equipoRojo, equipoRojo})
	})
}

func (suite *SmartBuildingTestSuite) Test16CombinedTeamDividesBuildingAreaEquallyBetweenConcreteTeams() {
	teamBlue := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamBeta := NewCombinedTeam([]Team{teamBlue, teamGreen})
	assert.Equal(suite.T(), 1, teamBlue.DaysToBuild(50))
	assert.Equal(suite.T(), 1, teamGreen.DaysToBuild(50))
	assert.Equal(suite.T(), 1, teamBeta.DaysToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test17ACombinedTeamDaysToBuildIsTheMaxBetweenSubteamsDaysToBuildAnEqualArea() {
	teamBlue := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	teamBeta := NewCombinedTeam([]Team{teamBlue, teamGreen})
	assert.Equal(suite.T(), 2, teamBlue.DaysToBuild(100))
	assert.Equal(suite.T(), 1, teamGreen.DaysToBuild(100))
	assert.Equal(suite.T(), 2, teamBeta.DaysToBuild(200))
}

func (suite *SmartBuildingTestSuite) Test17BCombinedTeamDaysToBuildIsTheMaxBetweenSubteamsDaysToBuildAnEqualArea() {
	teamRed := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamBlue := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamBeta := NewCombinedTeam([]Team{teamBlue, teamGreen})
	teamAlpha := NewCombinedTeam([]Team{teamRed, teamBeta})
	assert.Equal(suite.T(), 4, teamRed.DaysToBuild(200))
	assert.Equal(suite.T(), 1, teamBlue.DaysToBuild(100))
	assert.Equal(suite.T(), 2, teamGreen.DaysToBuild(100))
	assert.Equal(suite.T(), 2, teamBeta.DaysToBuild(200))
	assert.Equal(suite.T(), 4, teamAlpha.DaysToBuild(400))
}

func (suite *SmartBuildingTestSuite) Test18CombinedTeamPriceToBuildWithOneSubTeamIsSubteamPriceToBuild() {
	teamRed := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	teamBeta := NewCombinedTeam([]Team{teamRed})
	assert.Equal(suite.T(), 1000, teamRed.PriceToBuild(100))
	assert.Equal(suite.T(), 1000, teamBeta.PriceToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test19CombinedTeamPriceToBuildIsTheSumOfSubteamsPriceToBuildAnEqualArea() {
	teamRed := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	teamBlue := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(2000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(3000))
	teamBeta := NewCombinedTeam([]Team{teamRed, teamBlue, teamGreen})
	assert.Equal(suite.T(), 1000, teamRed.PriceToBuild(100))
	assert.Equal(suite.T(), 2000, teamBlue.PriceToBuild(100))
	assert.Equal(suite.T(), 3000, teamGreen.PriceToBuild(100))
	assert.Equal(suite.T(), 6000, teamBeta.PriceToBuild(300))
}

func (suite *SmartBuildingTestSuite) Test20CombinedTeamDividesBuildingAreaEquallyBetweenDirectTeams() {
	teamRed := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	teamBlue := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamBeta := NewCombinedTeam([]Team{teamBlue, teamGreen})
	teamAlpha := NewCombinedTeam([]Team{teamRed, teamBeta})
	assert.Equal(suite.T(), 2, teamRed.DaysToBuild(200))
	assert.Equal(suite.T(), 2, teamBlue.DaysToBuild(100))
	assert.Equal(suite.T(), 2, teamGreen.DaysToBuild(100))
	assert.Equal(suite.T(), 2, teamBeta.DaysToBuild(200))
	assert.Equal(suite.T(), 2, teamAlpha.DaysToBuild(400))
}

func (suite *SmartBuildingTestSuite) Test21CombinedTeamCantHaveRepeatedTeams() {
	teamBlue := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamBeta := NewCombinedTeam([]Team{teamBlue, teamGreen})
	assert.PanicsWithError(suite.T(), errorMessage.InvalidCombinedTeam, func() {
		NewCombinedTeam([]Team{teamBlue, teamBeta})
	})
}

func (suite *SmartBuildingTestSuite) Test22CombinedTeamDisplaysConcreteTeamsTimeToBuildCorrectly() {
	teamRed := NewConcreteTeam(NewFixedCapacity(200), NewFixedPrice(1500))
	teamBlue := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(500))
	teamBeta := NewCombinedTeam([]Team{teamBlue, teamGreen})
	teamAlpha := NewCombinedTeam([]Team{teamRed, teamBeta})
	expectedTimesToBuild := make(map[*ConcreteTeam]int)
	expectedTimesToBuild[teamRed] = 1
	expectedTimesToBuild[teamBlue] = 2
	expectedTimesToBuild[teamGreen] = 1
	actualTimesToBuild := make(map[*ConcreteTeam]int)
	teamAlpha.DisplayTimesToBuildOn(actualTimesToBuild, 400)
	assert.Equal(suite.T(), 1, teamRed.DaysToBuild(200))
	assert.Equal(suite.T(), 2, teamBlue.DaysToBuild(100))
	assert.Equal(suite.T(), 1, teamGreen.DaysToBuild(100))
	assert.Equal(suite.T(), expectedTimesToBuild, actualTimesToBuild)
}

func (suite *SmartBuildingTestSuite) Test23CombinedTeamDisplaysConcreteTeamsPriceToBuildCorrectly() {
	teamRed := NewConcreteTeam(NewFixedCapacity(200), NewFixedPrice(2000))
	teamBlue := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(500))
	teamBeta := NewCombinedTeam([]Team{teamBlue, teamGreen})
	teamAlpha := NewCombinedTeam([]Team{teamRed, teamBeta})
	expectedPricesToBuild := make(map[*ConcreteTeam]int)
	expectedPricesToBuild[teamRed] = 2000
	expectedPricesToBuild[teamBlue] = 1000
	expectedPricesToBuild[teamGreen] = 500
	actualPricesToBuild := make(map[*ConcreteTeam]int)
	teamAlpha.DisplayPricesToBuildOn(actualPricesToBuild, 400)
	assert.Equal(suite.T(), 2000, teamRed.PriceToBuild(200))
	assert.Equal(suite.T(), 1000, teamBlue.PriceToBuild(100))
	assert.Equal(suite.T(), 500, teamGreen.PriceToBuild(100))
	assert.Equal(suite.T(), expectedPricesToBuild, actualPricesToBuild)
}

func (suite *SmartBuildingTestSuite) Test24CheapestDirectTeamIsTeamWithLowestPriceToBuildAnEqualArea() {
	teamRed := NewConcreteTeam(NewFixedCapacity(200), NewFixedPrice(2000))
	teamBlue := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(500))
	teamBeta := NewCombinedTeam([]Team{teamBlue, teamGreen})
	teamAlpha := NewCombinedTeam([]Team{teamRed, teamBeta})
	assert.Equal(suite.T(), 1500, teamBeta.PriceToBuild(200))
	assert.Equal(suite.T(), 2000, teamRed.PriceToBuild(200))
	assert.Equal(suite.T(), teamBeta, teamAlpha.CheapestTeamToBuild(400))
}

func (suite *SmartBuildingTestSuite) Test25FastestDirectTeamIsTeamWithLowestDaysToBuildAnEqualArea() {
	teamRed := NewConcreteTeam(NewFixedCapacity(200), NewFixedPrice(2000))
	teamBlue := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(1000))
	teamGreen := NewConcreteTeam(NewFixedCapacity(50), NewFixedPrice(500))
	teamBeta := NewCombinedTeam([]Team{teamBlue, teamGreen})
	teamAlpha := NewCombinedTeam([]Team{teamRed, teamBeta})
	assert.Equal(suite.T(), 2, teamBeta.DaysToBuild(200))
	assert.Equal(suite.T(), 1, teamRed.DaysToBuild(200))
	assert.Equal(suite.T(), teamRed, teamAlpha.FastestTeamToBuild(400))
}
