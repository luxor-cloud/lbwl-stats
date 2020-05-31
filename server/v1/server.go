package v1

import (
	"freggy.dev/stats/pkg/flash"
	servicev1 "freggy.dev/stats/service/v1"
)

type Server struct {
	servicev1.UnimplementedStatsServiceServer
	FlashDAO flash.DataAccess
}

func NewServer(flashDAO flash.DataAccess) *Server {
	return &Server{FlashDAO: flashDAO}
}
