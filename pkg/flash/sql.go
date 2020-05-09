package flash

import (
	"context"

	"github.com/pkg/errors"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type sqlDataAccess struct {
	db      sqlbuilder.Database
	session sqlbuilder.SQLBuilder
}

func newSQLDataAccess(db sqlbuilder.Database) DataAccess {
	return &sqlDataAccess{
		db:      db,
		session: db,
	}
}

func (sda *sqlDataAccess) WithTX(ctx context.Context) (DataAccess, error) {
	tx, err := sda.db.NewTx(ctx)
	if err != nil {
		return nil, err
	}
	return &txSQLDataAccess{
		DataAccess: &sqlDataAccess{
			db:      sda.db,
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

// context creates
func (sda *sqlDataAccess) context(ctx context.Context) sqlbuilder.SQLBuilder {
	if dbSession, ok := sda.session.(sqlbuilder.Database); ok {
		return dbSession.WithContext(ctx)
	}

	if txSession, ok := sda.session.(sqlbuilder.Tx); ok {
		return txSession.WithContext(ctx)
	}
	return sda.session
}

func (sda *sqlDataAccess) GetPlayerStats(ctx context.Context, uuid string) (PlayerStats, error) {
	var stats PlayerStats
	if err := sda.context(ctx).
		SelectFrom("flash.player_stats").
		Where("uuid = ?", uuid).
		One(&stats); err != nil {
		return PlayerStats{}, err
	}

	return stats, nil
}

func (sda *sqlDataAccess) GetTopPlayerByPoints(ctx context.Context, limit int) ([]Player, error) {
	var players []Player
	if err := sda.context(ctx).
		SelectFrom("flash.player_stats").
		OrderBy("points DESC").
		Limit(limit).
		All(&players); err != nil {
			return nil, err
	}
	return players, nil
}

func (sda *sqlDataAccess) UpdatePlayerStats(ctx context.Context, diff PlayerStats) error {
	_, err := sda.context(ctx).
		Update("flash.player_stats").
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

func (sda *sqlDataAccess) GetHighscorePerCheckpointForMapAndUUID(ctx context.Context, uuid, mapName string) ([]PlayerCheckpointScore, error) {
	var highscores []PlayerCheckpointScore
	if err := sda.context(ctx).
		Select("checkpoint", "uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_checkpoint_score").
		Where("uuid = ? AND map = ?", uuid, mapName).
		GroupBy("checkpoint").
		All(&highscores); err != nil {
		return nil, err
	}
	return highscores, nil
}

func (sda *sqlDataAccess) GetBestHighscorePerCheckpointForMap(ctx context.Context, mapName string) ([]PlayerCheckpointScore, error) {
	var highscores []PlayerCheckpointScore
	if err := sda.context(ctx).
		Select("checkpoint", "uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_checkpoint_score").
		Where("map = ?", mapName).
		GroupBy("uuid").
		All(&highscores); err != nil {
		return nil, err
	}
	return highscores, nil
}

func (sda *sqlDataAccess) AddCheckpointScore(ctx context.Context, score PlayerCheckpointScore) error {
	_, err := sda.context(ctx).
		InsertInto("flash.player_checkpoint_score").
		Columns("uuid", "map", "time_needed", "accomplished_at").
		Values(score.UUID, score.Map, score.TimeNeeded, score.AccomplishedAt).
		Exec()
	return err
}

func (sda *sqlDataAccess) GetHighscorePerMapByUUID(ctx context.Context, uuid string) ([]PlayerMapScore, error) {
	var highscores []PlayerMapScore
	if err := sda.context(ctx).
		Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_map_scores").
		Where("uuid = ?", uuid).
		GroupBy("map").
		All(&highscores); err != nil {
		return nil, err
	}
	return highscores, nil
}

func (sda *sqlDataAccess) GetHighscoreForMapByUUID(ctx context.Context, uuid, mapName string) (PlayerMapScore, error) {
	var highscore PlayerMapScore
	if err := sda.context(ctx).
		Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_map_scores").
		Where("player_map_scores.uuid = ?", uuid).
		And("player_map_scores.map = ?", mapName).
		GroupBy("player_map_scores.map").
		One(&highscore); err != nil {
		return PlayerMapScore{}, err
	}
	return highscore, nil
}

func (sda *sqlDataAccess) GetBestHighscore(ctx context.Context, mapName string) (PlayerMapScore, error) {
	var highscore PlayerMapScore
	if err := sda.context(ctx).
		Select("uuid", "map", "accomplished_at", db.Raw("MIN(time_needed) as time_needed")).
		From("flash.player_map_score").
		Where("map = ?", mapName).
		GroupBy("uuid").
		Limit(1).
		One(&highscore); err != nil {
		return PlayerMapScore{}, err
	}
	return highscore, nil
}

func (sda *sqlDataAccess) GetTopHighscores(ctx context.Context, mapName string, limit int) ([]PlayerMapScore, error) {
	var highscores []PlayerMapScore
	if err := sda.context(ctx).
		SelectFrom("flash.player_map_scores").
		Where("map = ?", mapName).
		OrderBy("time_needed ASC").
		Limit(limit).
		All(&highscores); err != nil {
		return nil, err
	}
	return highscores, nil
}

func (sda *sqlDataAccess) AddMapScore(ctx context.Context, score PlayerMapScore) error {
	_, err := sda.context(ctx).
		InsertInto("flash.player_map_scores").
		Columns("uuid", "map", "time_needed", "accomplished_at").
		Values(score.UUID, score.Map, score.TimeNeeded, score.AccomplishedAt).
		Exec()
	return err
}
