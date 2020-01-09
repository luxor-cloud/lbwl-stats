package statsserver

import (
	"context"
	"freggy.dev/stats/rpc/go/model"
	"freggy.dev/stats/rpc/go/service"
	"github.com/twitchtv/twirp"
)

var (
	EmptyIDError  = twirp.InvalidArgumentError("playerId", "cannot be empty")
	EmptyMapError = twirp.InvalidArgumentError("maps", "cannot be empty")
	EmptyCompoundError = twirp.InvalidArgumentError("stats", "cannot be empty")
)

func (s *Server) GetFlashMapStats(ctx context.Context, mapReq *service.GetFlashMapStatsRequest) (*service.GetFlashStatsResponse, error) {
	if mapReq.GetPlayerId() == "" {
		return nil, EmptyIDError
	}

	if len(mapReq.GetMaps()) == 0 {
		return nil, EmptyMapError
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
		return nil, EmptyIDError
	}

	stats, err := s.FlashDAO.GetGameStatistic(gameReq.GetPlayerId())
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
		return nil, EmptyIDError
	}

	if len(compReq.GetMaps()) == 0 {
		return nil, EmptyMapError
	}

	gameStats, err := s.FlashDAO.GetGameStatistic(compReq.GetPlayerId())
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

func (s *Server) UpdateFlashStats(ctx context.Context, req *service.UpdateFlashStatsRequest) (*service.UpdateFlashStatsResponse, error) {
	if req.GetPlayerId() == "" {
		return nil, EmptyIDError
	}

	cmp := req.GetStats()
	if cmp == nil {
		return nil, EmptyCompoundError
	}

	if cmp.GetGameSummary() != nil {
		game := cmp.GetGameSummary()
		if err := s.FlashDAO.UpdateGameStatistic(req.GetPlayerId(), game); err != nil {
			return nil, err
		}
	}

	if cmp.GetMapSummary() != nil {
		maps := cmp.GetMapSummary()
		for _, m := range maps {
			if err := s.FlashDAO.UpdateMapStatistic(req.GetPlayerId(), m); err != nil {
				return nil, err
			}
		}
	}

	return &service.UpdateFlashStatsResponse{}, nil
}


func (s *Server) getMapStats(id string, mapNames []string) ([]*model.FlashMapStatistic, error) {
	// We could compute all of them in their own goroutine which would make this part faster,
	// but we just leave it like this for now.
	maps := make([]*model.FlashMapStatistic, 0)
	for _, m := range mapNames {
		stats, err := s.FlashDAO.GetMapStatistic(id, m)
		if err != nil {
			return nil, err
		}

		if stats != nil {
			maps = append(maps, stats)
		}
	}
	return maps, nil
}
