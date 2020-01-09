package statsserver

import (
	"context"
	"freggy.dev/stats/rpc/go/model"
	"freggy.dev/stats/rpc/go/service"
	"github.com/twitchtv/twirp"
)

var (
	InvalidIDError = twirp.InvalidArgumentError("playerId", "cannot be empty")
	InvalidMapError = twirp.InvalidArgumentError("maps", "cannot be empty")
)

func (s *Server) GetFlashMapStats(ctx context.Context, mapReq *service.GetFlashMapStatsRequest) (*service.GetFlashStatsResponse, error) {
	if mapReq.GetPlayerId() == "" {
		return nil, InvalidIDError
	}

	if len(mapReq.GetMaps()) == 0 {
		return nil, InvalidMapError
	}

	maps, err := s.getMapStats(mapReq.GetPlayerId(), mapReq.GetMaps())
	if err != nil {
		return nil, err
	}

	return &service.GetFlashStatsResponse{
		Stats: &model.FlashStatisticCompound{
			MapSummary: maps,
		},
	}, nil
}

func (s *Server) GetFlashGameStats(ctx context.Context, gameReq *service.GetFlashGameStatsRequest) (*service.GetFlashStatsResponse, error) {
	if gameReq.GetPlayerId() == "" {
		return nil, InvalidIDError
	}

	stats, err := s.flashDAO.GetGameStatistic(gameReq.GetPlayerId())
	if err != nil {
		return nil, err
	}

	return &service.GetFlashStatsResponse{
		Stats: &model.FlashStatisticCompound{
			GameSummary: stats,
		},
	}, nil
}

func (s *Server) GetFlashStats(ctx context.Context, compReq *service.GetFlashStatsCompoundRequest) (*service.GetFlashStatsResponse, error) {
	if compReq.GetPlayerId() == "" {
		return nil, InvalidIDError
	}

	if len(compReq.GetMaps()) == 0 {
		return nil, InvalidMapError
	}

	gameStats, err := s.flashDAO.GetGameStatistic(compReq.GetPlayerId())
	if err != nil {
		return nil, err
	}

	mapStats, err := s.getMapStats(compReq.GetPlayerId(), compReq.GetMaps())
	if err != nil {
		return nil, err
	}

	return &service.GetFlashStatsResponse{
		Stats: &model.FlashStatisticCompound{
			GameSummary: gameStats,
			MapSummary:  mapStats,
		},
	}, nil
}

func (s *Server) getMapStats(id string, mapNames []string) ([]*model.FlashMapStatistic, error) {
	// We could compute all of them in their own goroutine which would make this part faster,
	// but we just leave it like this for now.
	maps := make([]*model.FlashMapStatistic, 0)
	for _, m := range mapNames {
		stats, err := s.flashDAO.GetMapStatistic(id, m)
		if err != nil {
			return nil, err
		}

		if stats != nil {
			maps = append(maps, stats)
		}
	}
	return maps, nil
}
