package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"
)

func (bot *OGameBot) BuildBaseFacility(planet ogame.Planet) (*ogame.Planet, *ogame.BaseBuilding, int64){

	var targetPlanet *ogame.Planet = &planet
	var targetBuilding *ogame.BaseBuilding = nil
	var currentLevel int64 = -1


	Facilities,_ := planet.GetFacilities()

	if Facilities.RoboticsFactory < 2 {
		targetBuilding = &ogame.RoboticsFactory.BaseBuilding
		currentLevel = Facilities.RoboticsFactory
	} else if Facilities.Shipyard < 4 {
		targetBuilding = &ogame.Shipyard.BaseBuilding
		currentLevel = Facilities.Shipyard
	}else if Facilities.RoboticsFactory < 7 {
		targetBuilding = &ogame.RoboticsFactory.BaseBuilding
		currentLevel = Facilities.RoboticsFactory
	} else if Facilities.Shipyard < 7 {
		targetBuilding = &ogame.Shipyard.BaseBuilding
		currentLevel = Facilities.Shipyard
	} else if Facilities.RoboticsFactory < 10 {
		targetBuilding = &ogame.RoboticsFactory.BaseBuilding
		currentLevel = Facilities.RoboticsFactory
	} else {
		 currentLevel = -1
	}





	return targetPlanet, targetBuilding, currentLevel




}
