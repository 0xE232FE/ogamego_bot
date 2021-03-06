package OGameBotRoutine

import "bitbucket.org/jc01rho/ogame"

func (bot *OGameBot) GetNextFacilityInBuildRess(planet ogame.Planet) (bool, ogame.Planet, ogame.ID, int64) {

	//targetBuilding = ogame.MetalMine.Base
	//lowsestPrice = OgameUtil.ResourcePricesSum(ogame.MetalMine.GetPrice(resbuildings.MetalMine))
	//currentLevel = resbuildings.MetalMine
	facilities, _ := planet.GetFacilities()

	if facilities.RoboticsFactory < 2 {
		return true, planet, ogame.RoboticsFactoryID, facilities.RoboticsFactory
	} else if facilities.Shipyard < 2 {
		return true, planet, ogame.ShipyardID, facilities.Shipyard
	} else if facilities.Shipyard < 4 {
		return true, planet, ogame.ShipyardID, facilities.Shipyard
	} else if facilities.RoboticsFactory < 7 {
		return true, planet, ogame.RoboticsFactoryID, facilities.RoboticsFactory
	} else if facilities.Shipyard < 7 {
		return true, planet, ogame.ShipyardID, facilities.Shipyard
	} else if facilities.MissileSilo < 4 {
		return true, planet, ogame.MissileSiloID, facilities.MissileSilo
	} else {
		//fail fall
		return false, planet, ogame.ShipyardID, facilities.Shipyard
	}

}
