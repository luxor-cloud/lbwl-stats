package flash

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/postgresql"
)

type DataAccess interface {

	// GetPlayerRepository returns the PlayerStatsRepository
	GetPlayerRepository() PlayerStatsRepository

	// GetCheckpointScoresRepository returns the PlayerCheckpointScoresRepository
	GetCheckpointScoresRepository() PlayerCheckpointScoresRepository

	// GetPlayerMapScoresRepository returns the PlayerMapScoresRepository
	GetPlayerMapScoresRepository() PlayerMapScoresRepository

	// Creates a new data access using a transaction as underlying mechanism.
	WithTX(ctx context.Context) (DataAccess, error)

	// Close closes the connection to the database, commits or rolls back the underlying transaction
	// if there was an error.
	Close(err error) error
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
