package flash

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DataAccess interface {

	GetPlayerRepository() PlayerStatsRepository

	GetCheckpointScoresRepository() PlayerCheckpointScoresRepository

	GetPlayerMapScoresRepository() PlayerMapScoresRepository
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
