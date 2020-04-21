package statsserver

import (
	"context"

	"github.com/twitchtv/twirp"

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
	return nil, nil
}

func (server *Server) GetGlobalFlashMapHighscore(
	ctx context.Context, r *service.GetGlobalFlashMapHighscoreRequest,
) (*service.GetGlobalFlashMapHighscoreResponse, error) {
	return nil, nil
}

func (server *Server) GetTopFlashMapHighscores(
	ctx context.Context, r *service.GetTopFlashMapHighscoresRequest,
) (*service.GetTopFlashMapHighscoresResponse, error) {

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
