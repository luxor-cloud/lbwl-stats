package statsserver

import (
	"freggy.dev/stats/pkg/flash"
	"freggy.dev/stats/rpc/go/service"
)

type FlashHandler struct {
	dao flash.DataAccess
}

func NewFlashHandler(dao flash.DataAccess) *FlashHandler {
	return &FlashHandler{dao}
}

func (fh FlashHandler) GetGame() string {
	return "flash"
}

func (fh *FlashHandler) GetStats(req *service.GetStatsRequest) (*service.GetStatsResponse, error) {

}

func (fh *FlashHandler) UpdateStats(req *service.UpdateStatsRequest) (*service.UpdateStatsResponse, error) {

}

