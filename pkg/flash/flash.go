package flash

import (
	"freggy.dev/stats/rpc/go/model"
)

type DataAccess interface {
	GetMapStatistic(id string) *model.FlashGameStatistic
	GetGameStatistic(id string) *model.FlashMapStatistic

	UpdateGameStatistic(id string, game *model.FlashGameStatistic)
	UpdateMapStatistic(id string, game *model.FlashMapStatistic)
}

type CassandraDataAccess struct {

}

func (cda *CassandraDataAccess) GetMapStatistic(id string) *model.FlashMapStatistic {

}

func (cda *CassandraDataAccess) GetGameStatistic(id string) *model.FlashGameStatistic {

}

func (cda *CassandraDataAccess) UpdateGameStatistic(id string, game *model.FlashGameStatistic) {

}

func (cda *CassandraDataAccess) UpdateMapStatistic(id string, game *model.FlashMapStatistic) {

}
