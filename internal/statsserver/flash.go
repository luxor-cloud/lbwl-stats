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
	stats := &model.FlashMapStatisticCollection{}
	var cScores []flash.PlayerCheckpointScore

	if r.WithCheckpoints {
		scores, err := server.FlashDAO.GetCheckpointScoresRepository().GetHighscorePerCheckpointForMapAndUUID(r.PlayerId, r.MapName)
		if err != nil {
			return nil, err // TODO: maybe wrap err
		}
		cScores = scores
	}

	mScores, err := server.FlashDAO.GetPlayerMapScoresRepository().GetHighscoreForMapByUUID(r.PlayerId, r.MapName)
	if err != nil {
		return nil, err // TODO: maybe wrap err
	}

	stats.MapSummary = make([]*model.FlashMapStatistic, 0)
	stats.MapSummary = append(stats.MapSummary, createMapStats(mScores, cScores))

	return &service.GetFlashMapHighscoreForPlayerResponse{
		Highscores: stats,
	}, nil
}

func (server *Server) GetGlobalFlashMapHighscore(
	ctx context.Context, r *service.GetGlobalFlashMapHighscoreRequest,
) (*service.GetGlobalFlashMapHighscoreResponse, error) {
	return nil, nil
}

func (server *Server) GetTopFlashMapHighscores(
	ctx context.Context, r *service.GetTopFlashMapHighscoresRequest,
) (*service.GetTopFlashMapHighscoresResponse, error) {
	return nil, nil
}

func (server *Server) GetTopFlashPlayersByPoints(
	ctx context.Context, r *service.GetTopPlayersByPointsRequest,
) (*service.GetTopPlayersByPointsResponse, error) {
	return nil, nil
}

func (server *Server) UpdateFlashStatistics(
	ctx context.Context, r *service.UpdateFlashStatisticsRequests,
) (*service.UpdateFlashStatisticsResponse, error) {
	return nil, nil
}


func createMapStats(mScore flash.PlayerMapScore, cScore []flash.PlayerCheckpointScore) *model.FlashMapStatistic {
	checkpoints := make([]*model.FlashCheckpointStatistic, 0)
	for _, c := range cScore {
		checkpoints = append(checkpoints, &model.FlashCheckpointStatistic{
			Checkpoint:     int32(c.Checkpoint),
			TimeNeeded:     c.TimeNeeded,
			AccomplishedAt: c.AccomplishedAt.String(),
		})
	}
	return &model.FlashMapStatistic{
		Name:           mScore.Map,
		TimeNeeded:     mScore.TimeNeeded,
		AccomplishedAt: mScore.AccomplishedAt.String(),
		Checkpoints: checkpoints,
	}
}
