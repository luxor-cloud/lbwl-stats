package flash

import (
	"database/sql"
	"freggy.dev/stats/rpc/go/model"
)

type DataAccess interface {
	GetMapStatistic(id string, mapName string) (*model.FlashMapStatistic, error)
	GetGameStatistic(id string) (*model.FlashGameStatistic, error)

	UpdateGameStatistic(id string, stats *model.FlashGameStatistic) error
	UpdateMapStatistic(id string, stats *model.FlashMapStatistic) error
}

const (
	GetCheckpointStats = `SELECT checkpoint, record_time FROM checkpoint_records WHERE id = ? AND map = ?`
	GetMapStats        = `SELECT record_time, accomplished_at FROM map_records WHERE id = ? AND map = ?`
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
		var recordTime uint64
		var checkpoint int32

		if err := rows.Scan(&recordTime, &checkpoint); err != nil {
			return nil, err
		}

		checkpoints = append(checkpoints, &model.FlashCheckpointStatistic{
			Checkpoint: checkpoint,
			TimeNeeded: 0,
			RecordTime: recordTime,
		})

		if !rows.Next() {
			break
		}
	}

	rows.Close() // Close before new assignment

	rows, err = sda.db.Query(GetMapStats, id, mapName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var recordTime uint64
	var accomplishedAt uint64

	if err := rows.Scan(&recordTime, &accomplishedAt); err != nil {
		return nil, err
	}

	return &model.FlashMapStatistic{
		Name:        mapName,
		RecordTime:  recordTime,
		Checkpoints: checkpoints,
	}, nil
}

func (sda *SQLDataAccess) GetGameStatistic(id string) (*model.FlashGameStatistic, error) {

}

func (sda *SQLDataAccess) UpdateGameStatistic(id string, stats *model.FlashGameStatistic) {

}

func (cda *SQLDataAccess) UpdateMapStatistic(id string, stats *model.FlashMapStatistic) {

}
