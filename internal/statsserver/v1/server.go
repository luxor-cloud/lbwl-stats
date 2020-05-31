package v1

import (
	servicev1 "freggy.dev/stats/internal/service/v1"
	"freggy.dev/stats/pkg/flash"
)

type Server struct {
	FlashDAO flash.DataAccess
	servicev1.UnimplementedStatsServiceServer

}

func NewServer(flashDAO flash.DataAccess) *Server {
	return &Server{FlashDAO: flashDAO}
}
