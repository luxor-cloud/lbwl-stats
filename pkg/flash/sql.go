package flash

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"freggy.dev/stats/rpc/go/model"
)

// Maybe read queries from files provided in the container?
// This way we would not have to rebuild the entire application if we make database changes
const (
	GetCheckpointStats = `SELECT checkpoint, record_time, accomplished_at FROM checkpoint_records WHERE uuid = ? AND map = ?`

	GetMapStatistics = `SELECT record_time, accomplished_at FROM map_records WHERE uuid = ? AND map = ?`

	GetGameStatistics = `SELECT wins, deaths, checkpoints, games_played, instant_deaths FROM global_game_data WHERE uuid = ?`

	UpdateGameStatistics = `
		INSERT INTO global_game_data (uuid, wins, deaths, instant_deaths, checkpoints, games_played)
		VALUES (?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE wins            = wins + 1,
                        	    games_played    = games_played + 1,
								deaths          = deaths + VALUES(deaths),
                        		instant_deaths  = instant_deaths + VALUES(instant_deaths),
                        		checkpoints     = checkpoints + VALUES(checkpoints)
	`

	UpdateMapStatistics = `
		INSERT INTO map_records (uuid, map, record_time, accomplished_at)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE record_time      = VALUES(record_time),
                        		accomplished_at  = VALUES(accomplished_at)
	`
)

// TODO: pass context to functions
￿
const SqlDateTimeLayout = "2006-01-02 15:04:05"￿

type SQLDataAccess struct {
	Conn *sql.DB
}

func (sda *SQLDataAccess) GetMapStatistic(id string, mapName string) (*model.FlashMapStatistic, error) {
	rows, err := sda.Conn.Query(GetCheckpointStats, id, mapName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	checkpoints := make([]*model.FlashCheckpointStatistic, 0)

	for {
		if !rows.Next() {
			break
		}

		var recordTime uint64
		var checkpoint int32
		var accomplishedAt time.Time

		if err := rows.Scan(&checkpoint, &recordTime, &accomplishedAt); err != nil {
			return nil, err
		}

		checkpoints = append(checkpoints, &model.FlashCheckpointStatistic{
			Checkpoint:     checkpoint,
			AccomplishedAt: accomplishedAt.Format(SqlDateTimeLayout),
			TimeNeeded:     0,
			RecordTime:     recordTime,
		})
	}

	var recordTime uint64
	var accomplishedAt time.Time

	if err := sda.Conn.QueryRow(GetMapStatistics, id, mapName).Scan(&recordTime, &accomplishedAt);
		err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &model.FlashMapStatistic{
		Name:           mapName,
		RecordTime:     recordTime,
		AccomplishedAt: accomplishedAt.Format(SqlDateTimeLayout),
		Checkpoints:    checkpoints,
	}, nil
}

func (sda *SQLDataAccess) GetGameStatistic(id string) (*model.FlashGameStatistic, error) {
	var (
		wins          uint32
		deaths        uint32
		checkpoints   uint32
		gamesPlayed   uint32
		instantDeaths uint32
	)

	if err := sda.Conn.QueryRow(GetGameStatistics, id).Scan(&wins, &deaths, &checkpoints, &gamesPlayed, &instantDeaths);
		err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &model.FlashGameStatistic{
		Wins:          wins,
		Deaths:        deaths,
		GamesPlayed:   gamesPlayed,
		InstantDeaths: instantDeaths,
		Checkpoints:   checkpoints,
	}, nil
}

func (sda *SQLDataAccess) UpdateGameStatistic(id string, stats *model.FlashGameStatistic) error {
	if _, err := sda.Conn.Exec(
		UpdateGameStatistics,
		id,
		stats.GetWins(),
		stats.GetDeaths(),
		stats.GetInstantDeaths(),
		stats.GetCheckpoints(),
		stats.GetGamesPlayed(),
	); err != nil {
		return err
	}
	return nil
}

func (sda *SQLDataAccess) UpdateMapStatistic(id string, stats *model.FlashMapStatistic) (err error) {
	tx, err := sda.Conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("Could not rollback transaction: %v", err)
				return
			}
		}
		if err := tx.Commit(); err != nil {
			log.Printf("Could not commit transaction: %v", err)
		}
	}()

	if _, err = tx.Exec(
		UpdateMapStatistics,
		id,
		stats.GetName(),
		stats.GetRecordTime(),
		stats.GetAccomplishedAt(),
	); err != nil {
		return err
	}

	if len(stats.GetCheckpoints()) <= 0 {
		return nil
	}

	// The string that will be built must have the following format.
	// This way we can insert multiple rows at once.
	// INSERT INTO checkpoint_records (uuid, map, checkpoint, record_time, accomplished_at)
	// VALUES
	//	      (?, ?, ?, ?, ?),
	//	      (?, ?, ?, ?, ?),
	//	      (?, ?, ?, ?, ?)
	// ON DUPLICATE KEY UPDATE record_time      = VALUES(record_time),
	//	                       accomplished_at  = VALUES(accomplished_at)
	query := new(strings.Builder)
	query.WriteString("INSERT INTO checkpoint_records (uuid, map, checkpoint, record_time, accomplished_at) VALUES")

	for i, checkpoint := range stats.GetCheckpoints() {
		s := "('%s', '%s', %d, %d, '%s'),"
		if i == len(stats.GetCheckpoints())-1 { // last entry must not have a comma at the end
			s = "('%s', '%s', %d, %d, '%s')"
		}
		query.WriteString(
			fmt.Sprintf(s, id, stats.GetName(), checkpoint.GetCheckpoint(), checkpoint.GetRecordTime(), checkpoint.GetAccomplishedAt()),
		)
	}

	query.WriteString(" ON DUPLICATE KEY UPDATE record_time = VALUES(record_time), accomplished_at = VALUES(accomplished_at)")

	log.Println(query.String())

	if _, err := tx.Exec(query.String()); err != nil {
		return err
	}

	return nil
}
