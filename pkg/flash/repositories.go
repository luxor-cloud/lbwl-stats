package flash

import (
	"time"
)

type PlayerStats struct {
	UUID          string `db:"uuid"`
	Wins          uint32 `db:"wins"`
	Deaths        uint32 `db:"deaths"`
	Checkpoints   uint32 `db:"checkpoints"`
	GamesPlayed   uint32 `db:"games_played"`
	InstantDeaths uint32 `db:"instant_deaths"`
	Points        uint32 `db:"points"`
}

type PlayerMapScore struct {
	UUID           string    `db:"uuid"`
	Map            string    `db:"map"`
	TimeNeeded     uint64    `db:"time_needed"`
	AccomplishedAt time.Time `db:"accomplished_at"`
}

type PlayerCheckpointScore struct {
	UUID           string    `db:"uuid"`
	Map            string    `db:"map"`
	Checkpoint     byte      `db:"checkpoint"`
	TimeNeeded     uint64    `db:"time_needed"`
	AccomplishedAt time.Time `db:"accomplished_at"`
}

type Player struct {
	Stats            PlayerStats
	CheckpointScores []PlayerCheckpointScore
	MapScores        []PlayerMapScore
}

type PlayerStatsRepository interface {

	// GetPlayerStats gets the player stats for a given player.
	GetPlayerStats(uuid string) (PlayerStats, error)

	// UpdatePlayerStats updates the stats for a given player.
	UpdatePlayerStats(stats PlayerStats) error
}

type PlayerCheckpointScoresRepository interface {

	// GetHighscorePerCheckpointForMapAndUUID gets all the best scores
	// for all distinct checkpoints specific to the player and map
	GetHighscorePerCheckpointForMapAndUUID(uuid, m string) ([]PlayerCheckpointScore, error)

	// GetBestHighscoresPerCheckpointForMap gets the best highscores of every checkpoint on the map.
	// Sorted from best to worst.
	GetBestHighscorePerCheckpointForMap(mapName string) ([]PlayerCheckpointScore, error)

	// AddCheckpointScore adds a score to the repository.
	AddCheckpointScore(score PlayerCheckpointScore) error
}

type PlayerMapScoresRepository interface {

	// GetHighscoresPerMapByUUID gets the highscore for every distinct map a given player has played.
	GetHighscorePerMapByUUID(uuid string) ([]PlayerMapScore, error)

	// GetHighscoreForMapByUUID gets the highscore a player achieved on the given map.
	GetHighscoreForMapByUUID(uuid, mapName string) (PlayerMapScore, error)

	// GetBestHighscore gets the best highscore achieved on the given map.
	GetBestHighscore(mapName string) (PlayerMapScore, error)

	// GetTopHighscores gets the top x highscores achieved on the given map. Sorted from best to worst.
	GetTopHighscores(mapName string, limit int) ([]PlayerMapScore, error)

	// TODO: add retrieval for all scores

	// AddCheckpointScore adds a score to the repository.
	AddMapScore(score PlayerMapScore) error
}
