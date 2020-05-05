package flash

import "time"

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
