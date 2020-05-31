package v1

import (
	"context"
	"time"

	"freggy.dev/stats/pkg/flash"
	servicev1 "freggy.dev/stats/service/v1"
)

func (server *Server) GetFlashMapHighscoreForPlayer(
	ctx context.Context, r *servicev1.GetFlashMapHighscoreForPlayerRequest,
) (*servicev1.GetFlashMapHighscoreForPlayerResponse, error) {
	var cScores []flash.PlayerCheckpointScore
	if r.WithCheckpoints {
		timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		scores, err := server.FlashDAO.GetHighscorePerCheckpointForMapAndUUID(timeout, r.PlayerId, r.MapName)
		if err != nil {
			return nil, err
		}
		cScores = scores
	}

	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	mScores, err := server.FlashDAO.GetHighscoreForMapByUUID(timeout, r.PlayerId, r.MapName)
	if err != nil {
		return nil, err
	}

	return &servicev1.GetFlashMapHighscoreForPlayerResponse{
		Highscore: wrapPlayerMapScore(mScores, cScores),
	}, nil
}

func (server *Server) GetGlobalFlashMapHighscore(
	ctx context.Context, r *servicev1.GetGlobalFlashMapHighscoreRequest,
) (*servicev1.GetGlobalFlashMapHighscoreResponse, error) {
	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var cScores []flash.PlayerCheckpointScore
	mScore, err := server.FlashDAO.GetBestHighscore(timeout, r.MapName)
	if err != nil {
		return nil, err
	}

	if r.WithCheckpoints {
		timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		scores, err := server.FlashDAO.GetHighscorePerCheckpointForMapAndUUID(timeout, mScore.UUID, mScore.Map)
		if err != nil {
			return nil, err
		}
		cScores = scores
	}

	return &servicev1.GetGlobalFlashMapHighscoreResponse{
		Highscore: wrapPlayerMapScore(mScore, cScores),
	}, nil
}

func (server *Server) GetTopFlashMapHighscores(
	ctx context.Context, r *servicev1.GetTopFlashMapHighscoresRequest,
) (*servicev1.GetTopFlashMapHighscoresResponse, error) {
	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Array index corresponds to the players placement.
	// 1. entry best highscore
	// 2. entry second best highscore
	// etc.
	mScores, err := server.FlashDAO.GetTopHighscores(timeout, r.MapName, int(r.Limit))
	if err != nil {
		return nil, err
	}

	stats := make([]*servicev1.FlashMapStatisticCollection, 0)

	for i, mScore := range mScores {
		stats = append(stats, &servicev1.FlashMapStatisticCollection{
			PlayerId: mScore.UUID, Maps: make([]*servicev1.FlashMapStatistic, 0),
		})
		var cScores []flash.PlayerCheckpointScore
		if r.WithCheckpoints {
			timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
			cs, err := server.FlashDAO.GetHighscorePerCheckpointForMapAndUUID(timeout, mScore.UUID, mScore.Map)
			cancel()

			if err != nil {
				return nil, err
			}
			cScores = cs
		}
		stats[i].Maps = append(stats[i].Maps, wrapPlayerMapScore(mScore, cScores))
	}

	return &servicev1.GetTopFlashMapHighscoresResponse{Highscores: stats}, nil
}

func (server *Server) GetTopFlashPlayersByPoints(
	ctx context.Context, r *servicev1.GetTopPlayersByPointsRequest,
) (*servicev1.GetTopPlayersByPointsResponse, error) {

	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Array index corresponds to the players placement.
	// 1. entry most points
	// 2. entry second most points
	// etc.
	players, err := server.FlashDAO.GetTopPlayerByPoints(timeout, int(r.Limit))
	if err != nil {
		return nil, err
	}

	stats := make([]*servicev1.FlashAllStatistics, 0)

	for _, player := range players {
		stats = append(stats, &servicev1.FlashAllStatistics{
			PlayerId:    player.Stats.UUID,
			PlayerStats: wrapPlayerStats(player.Stats),
			MapStats:    nil,
		})
	}

	if r.WithMapStats {
		for i := 0; i < len(players); i++ {
			timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
			mScores, err := server.FlashDAO.GetHighscorePerMapByUUID(timeout, players[i].Stats.UUID)
			cancel()
			if err != nil {
				return nil, err
			}
			stats[i].MapStats = make([]*servicev1.FlashMapStatistic, 0)
			for _, mScore := range mScores {
				var cScores []flash.PlayerCheckpointScore
				if r.WithCheckpoints {
					timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
					cScores, err = server.FlashDAO.GetHighscorePerCheckpointForMapAndUUID(timeout, players[i].Stats.UUID, mScore.Map)
					cancel()
					if err != nil {
						return nil, err
					}
				}
				stats[i].MapStats = append(stats[i].MapStats, wrapPlayerMapScore(mScore, cScores))
			}
		}
	}

	return &servicev1.GetTopPlayersByPointsResponse{TopPlayers: stats}, nil
}

func (server *Server) UpdateFlashStatistics(
	ctx context.Context, r *servicev1.UpdateFlashStatisticsRequests,
) (resp *servicev1.UpdateFlashStatisticsResponse, err error) {
	dao, err := server.FlashDAO.WithTX(ctx)
	if err != nil {
		return nil, err
	}

	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer func() {
		cancel()
		err = dao.Close(err)
	}()

	if err := dao.UpdatePlayerStats(timeout, flash.PlayerStats{
		UUID:          r.Stats.PlayerId,
		Wins:          r.Stats.PlayerStats.Wins,
		Deaths:        r.Stats.PlayerStats.Deaths,
		Checkpoints:   r.Stats.PlayerStats.Checkpoints,
		GamesPlayed:   r.Stats.PlayerStats.GamesPlayed,
		InstantDeaths: r.Stats.PlayerStats.InstantDeaths,
		Points:        r.Stats.PlayerStats.Points,
	}); err != nil {
		return nil, dao.Close(err)
	}

	for _, mScore := range r.Stats.MapStats {
		t, err := time.Parse("2006-01-02 15:04:05", mScore.AccomplishedAt)
		if err != nil {
			return nil, err
		}

		timeout, cancel := context.WithTimeout(ctx, 5*time.Second)

		if err := dao.AddMapScore(timeout, flash.PlayerMapScore{
			UUID:           r.Stats.PlayerId,
			Map:            mScore.Name,
			TimeNeeded:     mScore.TimeNeeded,
			AccomplishedAt: t,
		}); err != nil {
			cancel()
			return nil, err
		}
		cancel()

		for _, cScore := range mScore.Checkpoints {
			t, err := time.Parse("2006-01-02 15:04:05", cScore.AccomplishedAt)
			if err != nil {
				return nil, err
			}

			timeout, cancel := context.WithTimeout(ctx, 5*time.Second)

			if err := dao.AddCheckpointScore(timeout, flash.PlayerCheckpointScore{
				UUID:           r.Stats.PlayerId,
				Map:            mScore.Name,
				Checkpoint:     byte(cScore.Checkpoint),
				TimeNeeded:     cScore.TimeNeeded,
				AccomplishedAt: t,
			}); err != nil {
				cancel()
				return nil, err
			}
			cancel()
		}
	}

	return &servicev1.UpdateFlashStatisticsResponse{}, nil
}

func wrapPlayerStats(stats flash.PlayerStats) *servicev1.FlashPlayerStatistic {
	return &servicev1.FlashPlayerStatistic{
		Wins:          stats.Wins,
		Deaths:        stats.Deaths,
		GamesPlayed:   stats.GamesPlayed,
		InstantDeaths: stats.InstantDeaths,
		Checkpoints:   stats.Checkpoints,
		Points:        stats.Points,
	}
}

func wrapPlayerMapScore(mScore flash.PlayerMapScore, cScores []flash.PlayerCheckpointScore) *servicev1.FlashMapStatistic {
	var checkpoints []*servicev1.FlashCheckpointStatistic = nil
	if cScores != nil {
		checkpoints = wrapPlayerCheckpointScore(cScores)
	}
	return &servicev1.FlashMapStatistic{
		Name:           mScore.Map,
		TimeNeeded:     mScore.TimeNeeded,
		AccomplishedAt: mScore.AccomplishedAt.String(),
		Checkpoints:    checkpoints,
	}
}

func wrapPlayerCheckpointScore(scores []flash.PlayerCheckpointScore) []*servicev1.FlashCheckpointStatistic {
	checkpoints := make([]*servicev1.FlashCheckpointStatistic, 0)
	for _, c := range scores {
		checkpoints = append(checkpoints, &servicev1.FlashCheckpointStatistic{
			Checkpoint:     int32(c.Checkpoint),
			TimeNeeded:     c.TimeNeeded,
			AccomplishedAt: c.AccomplishedAt.String(),
		})
	}
	return checkpoints
}
