package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/emirpasic/gods/sets"
)

var OGameBotGlobal OGameBot

type OGameBot struct {
	Ogamebot              *ogame.OGame
	MainPlanetCoord       ogame.Coordinate
	MainCelestitial ogame.Celestial
	MainPlanetCelestitial			  ogame.Celestial
	MainPlanetMoonCelestitial ogame.Celestial


	IsMainPlanetMoon      bool

	Class ogame.CharacterClass

	BuildRessSkipList sets.Set

	Universe string
	Username string
	Password string
	Language string
}
