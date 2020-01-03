package flash

import (
	"database/sql"
	"freggy.dev/stats/rpc/go/model"
	"time"
)

type DataAccess interface {
	GetMapStatistic(id string, mapName string) (*model.FlashMapStatistic, error)
	GetGameStatistic(id string) (*model.FlashGameStatistic, error)

	UpdateGameStatistic(id string, stats *model.FlashGameStatistic) error
	UpdateMapStatistic(id string, stats *model.FlashMapStatistic) error
}

const (
	GetCheckpointStats = `SELECT checkpoint, record_time, accomplished_at FROM checkpoint_records WHERE uuid = ? AND map = ?`
	GetMapStats        = `SELECT record_time, accomplished_at FROM map_records WHERE uuid = ? AND map = ?`
)

type SQLDataAccess struct {
	db *sql.DB
}

func (sda *SQLDataAccess) GetMapStatistic(id string, mapName string) (*model.FlashMapStatistic, error) {
	rows, err := sda.db.Query(GetCheckpointStats, id, mapName)
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
			Checkpoint: checkpoint,
			AccomplishedAt: accomplishedAt.String(),
			TimeNeeded: 0,
			RecordTime: recordTime,
		})
	}

	rows.Close() // Close before new assignment

	rows, err = sda.db.Query(GetMapStats, id, mapName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var recordTime uint64
	var accomplishedAt time.Time

	if rows.Next() {
		if err := rows.Scan(&recordTime, &accomplishedAt); err != nil {
			return nil, err
		}
	}

	return &model.FlashMapStatistic{
		Name:        mapName,
		RecordTime:  recordTime,
		AccomplishedAt: accomplishedAt.String(),
		Checkpoints: checkpoints,
	}, nil
}

func (sda *SQLDataAccess) GetGameStatistic(id string) (*model.FlashGameStatistic, error) {
	return &model.FlashGameStatistic{}, nil
}

func (sda *SQLDataAccess) UpdateGameStatistic(id string, stats *model.FlashGameStatistic) error {
	return nil
}

func (cda *SQLDataAccess) UpdateMapStatistic(id string, stats *model.FlashMapStatistic) error {
	return nil
}
