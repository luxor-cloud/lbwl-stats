package flash

import (
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)


type sqlDataAccess struct {
	session sqlbuilder.Database
	pStatsRepo *sqlPlayerStatsRepository
	mStatsRepo *sqlPlayerMapScoreRepository
	cStatsRepo *sqlPlayerCheckpointScoresRepository
}

func newSQLDataAccess(session sqlbuilder.Database) DataAccess {
	return &sqlDataAccess{
		session:    session,
		pStatsRepo: &sqlPlayerStatsRepository{session: session},
		mStatsRepo: &sqlPlayerMapScoreRepository{session: session},
		cStatsRepo: &sqlPlayerCheckpointScoresRepository{session: session},
	}
}

func (sda *sqlDataAccess)  GetPlayerRepository() PlayerStatsRepository {
	return sda.pStatsRepo
}

func (sda *sqlDataAccess)  GetCheckpointScoresRepository() PlayerCheckpointScoresRepository {
	return sda.cStatsRepo
}

func (sda *sqlDataAccess)  GetPlayerMapScoresRepository() PlayerMapScoresRepository {
	return sda.mStatsRepo
}

func (sda *sqlDataAccess) Close() error {
	return sda.session.Close()
}


type sqlPlayerStatsRepository struct {
	session sqlbuilder.Database
}

func (repo *sqlPlayerStatsRepository) Get(uuid string) (PlayerStats, error) {
	var stats PlayerStats
	if err := repo.session.
		SelectFrom("flash.player_stats").
		Where("uuid = ?", uuid).
		One(&stats); err != nil {
		return PlayerStats{}, err
	}

	return stats, nil
}

func (repo *sqlPlayerStatsRepository) Update(diff PlayerStats) error {
	_, err := repo.session.Update("flash.player_stats").
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

func (repo *sqlPlayerStatsRepository) Exists(uuid string) (bool, error) {
	stats, err := repo.Get(uuid)
	if err != nil {
		return false, err
	}

	if stats == (PlayerStats{}) {
		return false, nil
	}
	return true, nil
}

type sqlPlayerCheckpointScoresRepository struct {
	session sqlbuilder.Database
}

func (repo *sqlPlayerCheckpointScoresRepository) GetHighscorePerCheckpointForMapAndUUID(uuid, mapName string) ([]PlayerCheckpointScore, error) {
	var highscores []PlayerCheckpointScore
	if err := repo.session.Select("checkpoint", "uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_checkpoint_score").
		Where("uuid = ? AND map = ?", uuid, mapName).
		GroupBy("checkpoint").
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (repo *sqlPlayerCheckpointScoresRepository) GetBestHighscorePerCheckpointForMap(mapName string) ([]PlayerCheckpointScore, error) {
	var highscores []PlayerCheckpointScore
	if err := repo.session.Select("checkpoint", "uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_checkpoint_score").
		Where("map = ?", mapName).
		GroupBy("uuid").
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (repo *sqlPlayerCheckpointScoresRepository) Add(score PlayerCheckpointScore) error {
	_, err := repo.session.InsertInto("flash.player_checkpoint_score").
		Columns("uuid", "map", "time_needed", "accomplished_at").
		Values(score.UUID, score.Map, score.TimeNeeded, score.AccomplishedAt).
		Exec()
	return err
}

type sqlPlayerMapScoreRepository struct {
	session sqlbuilder.Database
}

func (repo *sqlPlayerMapScoreRepository) GetHighscorePerMapByUUID(uuid string) ([]PlayerMapScore, error) {
	var highscores []PlayerMapScore
	if err := repo.session.Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_map_scores").
		Where("uuid = ?", uuid).
		GroupBy("map").
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (repo *sqlPlayerMapScoreRepository) GetHighscoreForMapByUUID(uuid, mapName string) (PlayerMapScore, error) {
	var highscore PlayerMapScore
	if err := repo.session.Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_map_scores").
		Where("uuid = ?", uuid).
		And("map = ?", mapName).
		GroupBy("map").
		One(&highscore); err != nil {
			return PlayerMapScore{}, err
	}
	return highscore, nil
}

func (repo *sqlPlayerMapScoreRepository) GetBestHighscore(mapName string) (PlayerMapScore, error) {
	var highscore PlayerMapScore
	if err := repo.session.Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_map_score").
		Where("map = ?", mapName).
		GroupBy("uuid").
		Limit(1).
		One(&highscore); err != nil {
			return PlayerMapScore{}, err
	}
	return highscore, nil
}

func (repo *sqlPlayerMapScoreRepository) GetTopHighscores(mapName string, limit int) ([]PlayerMapScore, error) {
	var highscores []PlayerMapScore
	if err := repo.session.SelectFrom("flash.player_map_scores").
		Where("map = ?", mapName).
		OrderBy("time_needed ASC").
		Limit(limit).
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (repo *sqlPlayerMapScoreRepository) Add(score PlayerMapScore) error {
	_, err := repo.session.InsertInto("flash.player_map_scores").
		Columns("uuid", "map", "time_needed", "accomplished_at").
		Values(score.UUID, score.Map, score.TimeNeeded, score.AccomplishedAt).
		Exec()
	return err
}




