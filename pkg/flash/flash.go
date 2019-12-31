package flash

import flash "freggy.dev/stats/pkg/flash/model"

type DataAccess interface {
	GetMapStatistic(id string) *flash.GameStatistic
	GetGameStatistic(id string) *flash.MapStatistic

	UpdateGameStatistic(id string, game *flash.GameStatistic)
	UpdateMapStatistic(id string, game *flash.MapStatistic)
}

type CassandraDataAccess struct {

}

func (cda *CassandraDataAccess) GetMapStatistic(id string) *flash.MapStatistic {

}

func (cda *CassandraDataAccess) GetGameStatistic(id string) *flash.GameStatistic {

}

func (cda *CassandraDataAccess) UpdateGameStatistic(id string, game *flash.GameStatistic) {

}

func (cda *CassandraDataAccess) UpdateMapStatistic(id string, game *flash.MapStatistic) {

}
