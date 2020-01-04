package stats

import (
	"context"
	"freggy.dev/stats/rpc/go/service"
	"github.com/twitchtv/twirp"
)

type Server struct {
	handler map[string]handler
}

type handler interface {
	GetGame() string
	GetStats(req *service.GetStatsRequest) (*service.GetStatsResponse, error)
	UpdateStats(req *service.UpdateStatsRequest) (*service.UpdateStatsResponse, error)
}

func (s *Server) GetStats(ctx context.Context, req *service.GetStatsRequest) (*service.GetStatsResponse, error) {
	handler := s.handler[req.GetGame()]
	if handler == nil {
		return nil, twirp.InvalidArgumentError("game", "Not a valid gamemode.")
	}

	resp, err := handler.GetStats(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) UpdateStats(ctx context.Context, req *service.UpdateStatsRequest) (*service.UpdateStatsResponse, error) {
	handler := s.handler[req.GetGame()]
	if handler == nil {
		return nil, twirp.InvalidArgumentError("game", "Not a valid gamemode.")
	}

	resp, err := handler.UpdateStats(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}



