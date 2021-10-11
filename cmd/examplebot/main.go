package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"paintbot-client/basebot"
	"paintbot-client/models"
	"paintbot-client/utilities/maputility"
)

func main() {
	basebot.Start("Simple Go Bot", models.Training, desiredGameSettings, calculateMove)
}

var (
	moves   = []models.Action{models.Right, models.Down, models.Left, models.Up} // models.Explode, models.Stay}
	lastDir = 0
)

// Implement your paintbot here
func calculateMove(settings models.GameSettings, updateEvent models.MapUpdateEvent) models.Action {
	utility := maputility.New(updateEvent.Map, *updateEvent.ReceivingPlayerID)
	me := utility.GetMe()

	if me.StunnedForTicks() > 0 {
		return models.Stay
	}

	if me.HasPowerUp() {
		return models.Explode
	}

	for !utility.CanIMoveInDirection(moves[lastDir]) {
		lastDir = (lastDir + 1) % len(moves)
	}

	return moves[lastDir]
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		ForceQuote:             true,
		FullTimestamp:          true,
		TimestampFormat:        "15:04:05.999",
		DisableLevelTruncation: true,
		PadLevelText:           true,
		QuoteEmptyFields:       true,
	})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
}

// desired game settings can be changed to nil to get default settings
var desiredGameSettings = &models.GameSettings{
	MaxNOOFPlayers:                 5,
	TimeInMSPerTick:                250,
	ObstaclesEnabled:               true,
	PowerUpsEnabled:                true,
	AddPowerUpLikelihood:           38,
	RemovePowerUpLikelihood:        5,
	TrainingGame:                   true,
	PointsPerTileOwned:             1,
	PointsPerCausedStun:            5,
	NOOFTicksInvulnerableAfterStun: 3,
	MinNOOFTicksStunned:            8,
	MaxNOOFTicksStunned:            10,
	StartObstacles:                 40,
	StartPowerUps:                  41,
	GameDurationInSeconds:          15,
	ExplosionRange:                 4,
	PointsPerTick:                  false,
}
