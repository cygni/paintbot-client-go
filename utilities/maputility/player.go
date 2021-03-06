package maputility

import "paintbot-client/models"

type Player struct {
	info    *models.CharacterInfo
	utility *MapUtility
}

func (p Player) GetPos() models.Coordinates {
	return p.utility.ConvertPositionToCoordinates(p.info.Position)
}

func (p Player) HasPowerUp() bool {
	return p.info.CarryingPowerUp
}

func (p Player) StunnedForTicks() int {
	return p.info.StunnedForGameTicks
}

func (p Player) GetPoints() int {
	return p.info.Points
}

func (p Player) GetName() string {
	return p.info.Name
}

func (p Player) GetID() string {
	return p.info.ID
}

func (p Player) GetColouredPositions() []models.Coordinates {
	return p.utility.ConvertPositionsToCoordinates(p.info.ColouredPosition)
}
