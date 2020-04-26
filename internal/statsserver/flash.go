package statsserver

import (
	"context"

	"github.com/twitchtv/twirp"

	"freggy.dev/stats/pkg/flash"
	"freggy.dev/stats/rpc/go/model"
	"freggy.dev/stats/rpc/go/service"
)

var (
	EmptyIDError       = twirp.InvalidArgumentError("playerId", "cannot be empty")
	EmptyMapError      = twirp.InvalidArgumentError("maps", "cannot be empty")
	EmptyCompoundError = twirp.InvalidArgumentError("stats", "cannot be empty")
)

func (server *Server) GetFlashMapHighscoreForPlayer(
	ctx context.Context, r *service.GetFlashMapHighscoreForPlayerRequest,
) (*service.GetFlashMapHighscoreForPlayerResponse, error) {
	var cScores []flash.PlayerCheckpointScore

	if r.WithCheckpoints {
		scores, err := server.FlashDAO.GetCheckpointScoresRepository().GetHighscorePerCheckpointForMapAndUUID(r.PlayerId, r.MapName)
		if err != nil {
			return nil, err
		}
		cScores = scores
	}

	mScores, err := server.FlashDAO.GetPlayerMapScoresRepository().GetHighscoreForMapByUUID(r.PlayerId, r.MapName)
	if err != nil {
		return nil, err
	}

	return &service.GetFlashMapHighscoreForPlayerResponse{
		Highscore: wrapPlayerMapScore(mScores, cScores),
	}, nil
}

func (server *Server) GetGlobalFlashMapHighscore(
	ctx context.Context, r *service.GetGlobalFlashMapHighscoreRequest,
) (*service.GetGlobalFlashMapHighscoreResponse, error) {
	var cScores []flash.PlayerCheckpointScore
	mScore, err := server.FlashDAO.GetPlayerMapScoresRepository().GetBestHighscore(r.MapName)
	if err != nil {
		return nil, err
	}

	if r.WithCheckpoints {
		scores, err := server.FlashDAO.GetCheckpointScoresRepository().GetHighscorePerCheckpointForMapAndUUID(mScore.UUID, mScore.Map)
		if err != nil {
			return nil, err
		}
		cScores = scores
	}

	return &service.GetGlobalFlashMapHighscoreResponse{
		Highscore: wrapPlayerMapScore(mScore, cScores),
	}, nil
}

func (server *Server) GetTopFlashMapHighscores(
	ctx context.Context, r *service.GetTopFlashMapHighscoresRequest,
) (*service.GetTopFlashMapHighscoresResponse, error) {
	// Array index corresponds to the players placement.
	// 1. entry best highscore
	// 2. entry second best highscore
	// etc.
	mScores, err := server.FlashDAO.GetPlayerMapScoresRepository().GetTopHighscores(r.MapName, int(r.Limit))
	if err != nil {
		return nil, err
	}

	stats := &model.FlashMapStatisticCollection{MapSummary: nil}

	for _, mScore := range mScores {
		stats.MapSummary = append(stats.MapSummary, wrapPlayerMapScore(mScore, nil))
	}

	if r.WithCheckpoints {
		for i := 0; i < len(mScores); i++ {
			cs, err := server.FlashDAO.GetCheckpointScoresRepository().GetHighscorePerCheckpointForMapAndUUID(mScores[i].UUID, mScores[i].Map)
			if err != nil {
				return nil, err
			}
			// -> see comment above
			// applies for cs as well
			stats.MapSummary[i].Checkpoints = wrapPlayerCheckpointScore(cs)
		}
	}

	return nil, nil
}

func (server *Server) GetTopFlashPlayersByPoints(
	ctx context.Context, r *service.GetTopPlayersByPointsRequest,
) (*service.GetTopPlayersByPointsResponse, error) {
	// TODO: implement
	return nil, nil
}

func (server *Server) UpdateFlashStatistics(
	ctx context.Context, r *service.UpdateFlashStatisticsRequests,
) (*service.UpdateFlashStatisticsResponse, error) {
	return nil, nil
}

func wrapPlayerMapScore(mScore flash.PlayerMapScore, cScores []flash.PlayerCheckpointScore) *model.FlashMapStatistic {
	var checkpoints []*model.FlashCheckpointStatistic = nil
	if cScores != nil {
		checkpoints = wrapPlayerCheckpointScore(cScores)
	}
	return &model.FlashMapStatistic{
		Name:           mScore.Map,
		TimeNeeded:     mScore.TimeNeeded,
		AccomplishedAt: mScore.AccomplishedAt.String(),
		Checkpoints:    checkpoints,
	}
}

func wrapPlayerCheckpointScore(scores []flash.PlayerCheckpointScore) []*model.FlashCheckpointStatistic {
	checkpoints := make([]*model.FlashCheckpointStatistic, 0)
	for _, c := range scores {
		checkpoints = append(checkpoints, &model.FlashCheckpointStatistic{
			Checkpoint:     int32(c.Checkpoint),
			TimeNeeded:     c.TimeNeeded,
			AccomplishedAt: c.AccomplishedAt.String(),
		})
	}
	return checkpoints
}
