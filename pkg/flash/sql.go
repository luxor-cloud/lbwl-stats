package flash

import (
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type SQLPlayerStatsRepository struct {
	session sqlbuilder.Database
}

func (repo *SQLPlayerStatsRepository) Get(uuid string) (*PlayerStats, error) {
	var stats PlayerStats
	if err := repo.session.
		SelectFrom("player_stats").
		Where("uuid = ?", uuid).
		One(&stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

func (repo *SQLPlayerStatsRepository) Update(diff *PlayerStats) error {
	_, err := repo.session.Update("player_stats").
		Set("wins = ? + wins", diff.Wins).
		Set("deaths = ? + deaths", diff.Deaths).
		Set("checkpoints = ? + checkpoints", diff.Checkpoints).
		Set("games_played = ? + games_played", diff.GamesPlayed).
		Set("instant_deaths = ? + instant_deaths", diff.InstantDeaths).
		Set("points = ? + points", diff.Points).
		Where("uuid = ?", diff.UUID).
		Exec()
	return err
}

func (repo *SQLPlayerStatsRepository) Exists(uuid string) (bool, error) {
	stats, err := repo.Get(uuid)
	if err != nil {
		return false, err
	}

	if stats == nil {
		return false, nil
	}
	return true, nil
}

type SQLPlayerCheckpointScoresRepository struct {
	Session sqlbuilder.Database
}

func (repo *SQLPlayerCheckpointScoresRepository) GetHighscorePerCheckpointForMapAndUUID(uuid, mapName string) ([]PlayerCheckpointScore, error) {
	var highscores []PlayerCheckpointScore
	if err := repo.Session.Select("checkpoint", "uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("player_checkpoint_score").
		Where("uuid = ? AND map = ?", uuid, mapName).
		GroupBy("checkpoint").
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (repo *SQLPlayerCheckpointScoresRepository) GetBestHighscorePerCheckpointForMap(mapName string) ([]PlayerCheckpointScore, error) {
	var highscores []PlayerCheckpointScore
	if err := repo.Session.Select("checkpoint", "uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("player_checkpoint_score").
		Where("map = ?", mapName).
		GroupBy("uuid").
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (repo *SQLPlayerCheckpointScoresRepository) Add(score PlayerCheckpointScore) error {
	_, err := repo.Session.InsertInto("player_checkpoint_score").
		Columns("uuid", "map", "time_needed", "accomplished_at").
		Values(score.UUID, score.Map, score.TimeNeeded, score.AccomplishedAt).
		Exec()
	return err
}

type SQLPlayerMapScoreRepository struct {
	Session sqlbuilder.Database
}

func (repo *SQLPlayerMapScoreRepository) GetHighscorePerMapByUUID(uuid string) ([]PlayerMapScore, error) {
	var highscores []PlayerMapScore
	if err := repo.Session.Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("player_map_scores").
		Where("uuid = ?", uuid).
		GroupBy("map").
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (repo *SQLPlayerMapScoreRepository) GetHighscoreForMapByUUID(uuid, mapName string) (PlayerMapScore, error) {
	var highscore PlayerMapScore
	if err := repo.Session.Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("player_map_scores").
		Where("uuid = ?", uuid).
		And("map = ?", mapName).
		GroupBy("map").
		One(&highscore); err != nil {
			return PlayerMapScore{}, err
	}
	return highscore, nil
}

func (repo *SQLPlayerMapScoreRepository) GetBestHighscore(mapName string) (PlayerMapScore, error) {
	var highscore PlayerMapScore
	if err := repo.Session.Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("player_map_score").
		Where("map = ?", mapName).
		GroupBy("uuid").
		Limit(1).
		One(&highscore); err != nil {
			return PlayerMapScore{}, err
	}
	return highscore, nil
}

func (repo *SQLPlayerMapScoreRepository) GetTopHighscores(mapName string, limit int) ([]PlayerMapScore, error) {
	var highscores []PlayerMapScore
	if err := repo.Session.SelectFrom("player_map_scores").
		Where("map = ?", mapName).
		OrderBy("time_needed ASC").
		Limit(limit).
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}




