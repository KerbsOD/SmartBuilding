package app

import (
	"SmartBuilding/internal/errorMessage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SmartBuildingTestSuite struct {
	suite.Suite
}

func TestSmartBuildingTestSuite(t *testing.T) {
	suite.Run(t, new(SmartBuildingTestSuite))
}

func (suite *SmartBuildingTestSuite) SetupTest() {

}

func (suite *SmartBuildingTestSuite) Test01FixedCapacityDaysToBuildAreCapacityOverArea() {
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	assert.Equal(suite.T(), int(1), equipoRojo.DaysToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test02FixedCapacityDaysToBuildAreFullDays() {
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	assert.Equal(suite.T(), int(2), equipoRojo.DaysToBuild(150))
}

func (suite *SmartBuildingTestSuite) Test03FixedPricePriceToBuildIsPriceTimesDaysToBuild() {
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	assert.Equal(suite.T(), int(1000), equipoRojo.PriceToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test04FixedPriceChargesFullDays() {
	equipoRojo := NewConcreteTeam(NewFixedCapacity(100), NewFixedPrice(1000))
	assert.Equal(suite.T(), int(2000), equipoRojo.PriceToBuild(150))
}

func (suite *SmartBuildingTestSuite) Test05VariableCapacityDaysToBuildUsesInitialCapacityTheInitialDays() {
	equipoRojo := NewConcreteTeam(NewVariableCapacity(2, 50, 25), NewFixedPrice(1000))
	assert.Equal(suite.T(), int(2), equipoRojo.DaysToBuild(100))
}

func (suite *SmartBuildingTestSuite) Test06VariableCapacityDaysToBuildGreaterThanInitialDaysUsesReminderCapacity() {
	equipoRojo := NewConcreteTeam(NewVariableCapacity(2, 50, 25), NewFixedPrice(1000))
	assert.Equal(suite.T(), int(4), equipoRojo.DaysToBuild(150))
}

func (suite *SmartBuildingTestSuite) Test07VariableCapacityInitialDaysToBuildAreFullDays() {
	equipoRojo := NewConcreteTeam(NewVariableCapacity(2, 50, 25), NewFixedPrice(1000))
	assert.Equal(suite.T(), int(2), equipoRojo.DaysToBuild(80))
}

func (suite *SmartBuildingTestSuite) Test08VariableCapacityRemainingDaysToBuildAreFullDays() {
	equipoRojo := NewConcreteTeam(NewVariableCapacity(2, 50, 25), NewFixedPrice(1000))
	assert.Equal(suite.T(), int(4), equipoRojo.DaysToBuild(130))
}

func (suite *SmartBuildingTestSuite) Test09VariableCapacityInitialCapacityMustBeSmallerThanRemainingCapacity() {
	assert.PanicsWithError(suite.T(), errorMessage.InvalidCapacitiesErrorMessage, func() {
		NewConcreteTeam(NewVariableCapacity(2, 25, 50), NewFixedPrice(1000))
	})
}
