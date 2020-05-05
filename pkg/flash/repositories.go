package flash

import "context"

type PlayerStatsRepository interface {

	// GetPlayerStats gets the player stats for a given player.
	GetPlayerStats(ctx context.Context, uuid string) (PlayerStats, error)

	// UpdatePlayerStats updates the stats for a given player.
	UpdatePlayerStats(ctx context.Context, stats PlayerStats) error

	GetTopPlayerByPoints(ctx context.Context, limit int) ([]Player, error)
}

type PlayerCheckpointScoresRepository interface {

	// GetHighscorePerCheckpointForMapAndUUID gets all the best scores
	// for all distinct checkpoints specific to the player and map
	GetHighscorePerCheckpointForMapAndUUID(ctx context.Context, uuid, mapName string) ([]PlayerCheckpointScore, error)

	// GetBestHighscoresPerCheckpointForMap gets the best highscores of every checkpoint on the map.
	// Sorted from best to worst.
	GetBestHighscorePerCheckpointForMap(ctx context.Context, mapName string) ([]PlayerCheckpointScore, error)

	// AddCheckpointScore adds a score to the repository.
	AddCheckpointScore(ctx context.Context, score PlayerCheckpointScore) error
}

type PlayerMapScoresRepository interface {

	// GetHighscoresPerMapByUUID gets the highscore for every distinct map a given player has played.
	GetHighscorePerMapByUUID(ctx context.Context, uuid string) ([]PlayerMapScore, error)

	// GetHighscoreForMapByUUID gets the highscore a player achieved on the given map.
	GetHighscoreForMapByUUID(ctx context.Context, uuid, mapName string) (PlayerMapScore, error)

	// GetBestHighscore gets the best highscore achieved on the given map.
	GetBestHighscore(ctx context.Context, mapName string) (PlayerMapScore, error)

	// GetTopHighscores gets the top x highscores achieved on the given map. Sorted from best to worst.
	GetTopHighscores(ctx context.Context, mapName string, limit int) ([]PlayerMapScore, error)

	// TODO: add retrieval for all scores

	// AddCheckpointScore adds a score to the repository.
	AddMapScore(ctx context.Context, score PlayerMapScore) error
}
