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

	// Get gets the player stats for a given player.
	Get(uuid string) (PlayerStats, error)

	// Update updates the stats for a given player.
	Update(stats PlayerStats) error
}

type PlayerCheckpointScoresRepository interface {

	// GetHighscorePerCheckpointForMapAndUUID gets all the best scores
	// for all distinct checkpoints specific to the player and map
	GetHighscorePerCheckpointForMapAndUUID(uuid, m string) ([]PlayerCheckpointScore, error)

	// GetBestHighscoresPerCheckpointForMap gets the best highscores of every checkpoint on the map.
	// Sorted from best to worst.
	GetBestHighscorePerCheckpointForMap(mapName string) ([]PlayerCheckpointScore, error)

	// Add adds a score to the repository.
	Add(score PlayerCheckpointScore) error
}

type PlayerMapScoreRepository interface {

	// GetHighscoresPerMapByUUID gets the highscore for every distinct map a given player has played.
	GetHighscorePerMapByUUID(uuid string) ([]PlayerMapScore, error)

	// GetHighscoreForMapByUUID gets the highscore a player achieved on the given map.
	GetHighscoreForMapByUUID(uuid, mapName string) (PlayerMapScore, error)

	// GetBestHighscore gets the best highscore achieved on the given map.
	GetBestHighscore(mapName string) (PlayerMapScore, error)

	// GetTopHighscores gets the top x highscores achieved on the given map. Sorted from best to worst.
	GetTopHighscores(mapName string, limit int) ([]PlayerMapScore, error)

	// TODO: add retrieval for all scores

	// Add adds a score to the repository.
	Add(score PlayerMapScore)
}

func display() {
	/*
		settings := mysql.ConnectionURL{
			User:     "root",
			Password: "secret",
			Database: "test",
			Host:     "localhost",
		}

		session, err := mysql.Open(settings)
		if err != nil {
			log.Fatal(err)
		}

		defer session.Close()

		var l GlobalGameData
		if err := session.SelectFrom("global_game_data").
			Where("uuid = ?", "92de217b-8b2b-403b-86a5-fe26fa3a9b5f").
			One(&l); err != nil {
			log.Fatal(err)
		}

		//for _, p := range l {

		log.Printf("UUID: %v\n", l.UUID)
		log.Printf("GamesPlayed: %v\n", l.GamesPlayed)
		log.Printf("Wins: %v\n", l.Wins)
		log.Printf("Deaths: %v\n", l.Deaths)
		log.Printf("Instant: %v\n", l.InstantDeaths)
		log.Printf("Check: %v\n", l.Checkpoints)



			for _, mr := range l.MapRecords {
				log.Printf("  Map: %v\n", mr.Map)
				log.Printf("  Accomp: %v\n", mr.AccomplishedAt)
				log.Printf("  Rec: %v\n", mr.RecordTime)

				for _, cr := range mr.Checkpoints {
					log.Printf("    Accomp: %v\n", cr.AccomplishedAt)
					log.Printf("    Rec: %v\n", cr.RecordTime)
				}
			}

		//}*/
}
