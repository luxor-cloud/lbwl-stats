package flash

import (
	_ "github.com/go-sql-driver/mysql"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/postgresql"
)

type DataAccess interface {
	GetPlayerRepository() PlayerStatsRepository

	GetCheckpointScoresRepository() PlayerCheckpointScoresRepository

	GetPlayerMapScoresRepository() PlayerMapScoresRepository

	Close() error
}

// Add functions which contain database connection logic here

func ConnectSQL(user, password, host, database string) (DataAccess, error) {
	settings := postgresql.ConnectionURL{
		User:     user,
		Password: password,
		Host:     host,
		Database: database,
	}

	session, err := mysql.Open(settings)
	if err != nil {
		return nil, err
	}
	return newSQLDataAccess(session), nil
}
