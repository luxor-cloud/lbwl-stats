package flash

import (
	"context"

	"github.com/pkg/errors"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)


type sqlDataAccess struct {
	db sqlbuilder.Database
	session sqlbuilder.SQLBuilder
}

func newSQLDataAccess(session sqlbuilder.Database) DataAccess {
	return &sqlDataAccess{
		session:    session,
	}
}

func (sda *sqlDataAccess)  WithTX(ctx context.Context) (DataAccess, error) {
	tx, err := sda.db.NewTx(ctx)
	if err != nil {
		return nil, err
	}
	return &txSQLDataAccess{
		DataAccess: &sqlDataAccess{
			db: sda.db,
			session: sda.session,
		},
		tx: tx,
	}, nil
}

func (sda *sqlDataAccess) Close(_ error) error {
	return sda.db.Close()
}

type txSQLDataAccess struct {
	DataAccess
	tx sqlbuilder.Tx
}

// Close commits the transaction or rolls it back if there was an error
func (sda *txSQLDataAccess) Close(err error) error {
	if err != nil {
		return errors.Wrap(err, sda.tx.Rollback().Error())
	}
	return sda.tx.Commit()
}


func (sda *sqlDataAccess) GetPlayerStats(uuid string) (PlayerStats, error) {
	var stats PlayerStats
	if err := sda.session.
		SelectFrom("flash.player_stats").
		Where("uuid = ?", uuid).
		One(&stats); err != nil {
		return PlayerStats{}, err
	}

	return stats, nil
}

func (sda *sqlDataAccess) UpdatePlayerStats(diff PlayerStats) error {
	_, err := sda.session.Update("flash.player_stats").
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

func (sda *sqlDataAccess) Exists(uuid string) (bool, error) {
	stats, err := sda.GetPlayerStats(uuid)
	if err != nil {
		return false, err
	}

	if stats == (PlayerStats{}) {
		return false, nil
	}
	return true, nil
}

func (sda *sqlDataAccess) GetHighscorePerCheckpointForMapAndUUID(uuid, mapName string) ([]PlayerCheckpointScore, error) {
	var highscores []PlayerCheckpointScore
	if err := sda.session.Select("checkpoint", "uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_checkpoint_score").
		Where("uuid = ? AND map = ?", uuid, mapName).
		GroupBy("checkpoint").
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (sda *sqlDataAccess) GetBestHighscorePerCheckpointForMap(mapName string) ([]PlayerCheckpointScore, error) {
	var highscores []PlayerCheckpointScore
	if err := sda.session.Select("checkpoint", "uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_checkpoint_score").
		Where("map = ?", mapName).
		GroupBy("uuid").
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (sda *sqlDataAccess) AddCheckpointScore(score PlayerCheckpointScore) error {
	_, err := sda.session.InsertInto("flash.player_checkpoint_score").
		Columns("uuid", "map", "time_needed", "accomplished_at").
		Values(score.UUID, score.Map, score.TimeNeeded, score.AccomplishedAt).
		Exec()
	return err
}

func (sda *sqlDataAccess) GetHighscorePerMapByUUID(uuid string) ([]PlayerMapScore, error) {
	var highscores []PlayerMapScore
	if err := sda.session.Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_map_scores").
		Where("uuid = ?", uuid).
		GroupBy("map").
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (sda *sqlDataAccess) GetHighscoreForMapByUUID(uuid, mapName string) (PlayerMapScore, error) {
	var highscore PlayerMapScore
	if err := sda.session.Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_map_scores").
		Where("uuid = ?", uuid).
		And("map = ?", mapName).
		GroupBy("map").
		One(&highscore); err != nil {
			return PlayerMapScore{}, err
	}
	return highscore, nil
}

func (sda *sqlDataAccess) GetBestHighscore(mapName string) (PlayerMapScore, error) {
	var highscore PlayerMapScore
	if err := sda.session.Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_map_score").
		Where("map = ?", mapName).
		GroupBy("uuid").
		Limit(1).
		One(&highscore); err != nil {
			return PlayerMapScore{}, err
	}
	return highscore, nil
}

func (sda *sqlDataAccess) GetTopHighscores(mapName string, limit int) ([]PlayerMapScore, error) {
	var highscores []PlayerMapScore
	if err := sda.session.SelectFrom("flash.player_map_scores").
		Where("map = ?", mapName).
		OrderBy("time_needed ASC").
		Limit(limit).
		All(&highscores); err != nil {
			return nil, err
	}
	return highscores, nil
}

func (sda *sqlDataAccess) AddMapScore(score PlayerMapScore) error {
	_, err := sda.session.InsertInto("flash.player_map_scores").
		Columns("uuid", "map", "time_needed", "accomplished_at").
		Values(score.UUID, score.Map, score.TimeNeeded, score.AccomplishedAt).
		Exec()
	return err
}




