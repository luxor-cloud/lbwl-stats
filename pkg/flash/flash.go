package flash

import (
	"freggy.dev/stats/rpc/go/model"
	"github.com/jmoiron/sqlx"
)

type DataAccess interface {
	GetMapStatistic(id string) (*model.FlashGameStatistic, error)
	GetGameStatistic(id string) (*model.FlashMapStatistic, error)

	UpdateGameStatistic(id string, stats *model.FlashGameStatistic) error
	UpdateMapStatistic(id string, stats *model.FlashMapStatistic) error
}

type SQLDataAccess struct {
	db *sqlx.DB
}

func (sda *SQLDataAccess) GetMapStatistic(id string) (*model.FlashMapStatistic, error) {

}

func (sda *SQLDataAccess) GetGameStatistic(id string) *model.FlashGameStatistic {

}

func (sda *SQLDataAccess) UpdateGameStatistic(id string, stats *model.FlashGameStatistic) {

}

func (cda *SQLDataAccess) UpdateMapStatistic(id string, stats *model.FlashMapStatistic) {

}
