package statsserver

import (
	"freggy.dev/stats/pkg/flash"
)

type Server struct {
	flashDAO flash.DataAccess
}


func NewServer(flashDAO flash.DataAccess) *Server {
	return &Server{flashDAO: flashDAO}
}


/*
type handler interface {
	GetGame() string
	GetStats(req *service.GetStatsRequest) (*service.GetStatsResponse, error)
	UpdateStats(req *service.UpdateStatsRequest) (*service.UpdateStatsResponse, error)
}

func NewServer(handlers ...handler) *Server {
	handlerMap := make(map[string]handler)
	for _, h := range handlers {
		handlerMap[h.GetGame()] = h
	}

	return &Server{handler: handlerMap}
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
}*/



