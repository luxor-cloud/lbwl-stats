package statsserver

import (
	"context"
	"freggy.dev/stats/rpc/go/model"
	"freggy.dev/stats/rpc/go/service"
)

func (s *Server) GetFlashMapStats(ctx context.Context, mapReq *service.GetFlashMapStatsRequest) (*service.GetFlashStatsResponse, error) {
	// We could compute all of them in their own goroutine which would make this part faster,
	// but we just leave it like this for now.
	maps := make([]*model.FlashMapStatistic, 0)
	for _, m := range mapReq.GetMaps() {
		stats, err := s.flashDAO.GetMapStatistic(mapReq.GetPlayerId(), m)
		if err != nil {
			return nil, err
		}

		if stats != nil {
			maps = append(maps, stats)
		}
	}

	return &service.GetFlashStatsResponse{
		Stats: &model.FlashStatisticCompound{
			MapSummary: maps,
		},
	}, nil
}

func (s *Server) GetFlashGameStats(ctx context.Context, gameReq *service.GetFlashGameStatsRequest) (*service.GetFlashStatsResponse, error) {
	return nil, nil
}

func (s *Server) GetFlashStats(ctx context.Context, compReq *service.GetFlashStatsCompoundRequest) (*service.GetFlashStatsResponse, error) {
	return nil, nil
}
