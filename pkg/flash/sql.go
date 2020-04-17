package flash

import (
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
