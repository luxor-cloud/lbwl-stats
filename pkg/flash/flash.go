package flash

import (
	"database/sql"
	"fmt"
	"freggy.dev/stats/rpc/go/model"

	_ "github.com/go-sql-driver/mysql"
)

type DataAccess interface {
	GetMapStatistic(id string, mapName string) (*model.FlashMapStatistic, error)
	GetGameStatistic(id string) (*model.FlashGameStatistic, error)

	UpdateGameStatistic(id string, stats *model.FlashGameStatistic) error
	UpdateMapStatistic(id string, stats *model.FlashMapStatistic) error
}

// Add functions which contain database connection logic here

func ConnectSQL(addr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s?parseTime=true", addr))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectMongoDB(addr string) {

}
