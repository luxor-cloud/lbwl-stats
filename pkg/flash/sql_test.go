// +build integration

package flash

import (
	"database/sql"
	"freggy.dev/stats/rpc/go/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// UUID that should be used to test retrieval functions
const RetrievalUuid = "92de217b-8b2b-403b-86a5-fe26fa3a9b5f"

// UUID that should be used to test update functions
const UpdateUuid = "aabb023c-66ef-4168-9bf9-b898bee9f73f"

//
// Connect methods for different databases
//

func connectMariaDB(t *testing.T) (DataAccess, *sql.DB) {
	db, err := sql.Open("mysql", "root:secret@/test?parseTime=true")
	if err != nil {
		t.Errorf("%v", err)
	}
	return &SQLDataAccess{db}, db
}

//
// SQL specific checks
//

func TestSQLDataAccess_Get_MapStatistic_Successfully(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	checkMapStatisticPlayloadSuccessfully(t, access)
}

func TestSQLDataAccess_GetMapStatistic_ID_Non_Existent(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	checkMapStatisticNonExistent(t, access, "abc", "Map1")
}

func TestSQLDataAccess_GetMapStatistic_Map_Non_Existent(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	checkMapStatisticNonExistent(t, access, RetrievalUuid, "Map42")
}

func TestSQLDataAccess_GetMapStatistic_Map_And_ID_Non_Existent(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	checkMapStatisticNonExistent(t, access, "abc", "Map42")
}

func TestSQLDataAccess_GetGameStatistic_Successfully(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	checkGameStatisticPayloadSuccessfully(t, access)
}

func TestSQLDataAccess_GetGameStatistic_ID_Non_Existent(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	checkEmptyGameStatisticPayload(t, access, "abc")
}

func TestSQLDataAccess_UpdateGameStatistic_Where_DataSet_Already_Exists(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	checkGameStatisticUpdateChanges(t, access, UpdateUuid, gameStatistic())
}

func TestSQLDataAccess_UpdateGameStatistic_Where_DataSet_Does_Not_Exist(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	// Expected behavior: New row should be inserted
	checkGameStatisticUpdateChanges(t, access, "99f5efbb-046f-4086-9a57-647959953d1f", gameStatistic())
	db.Exec("DELETE FROM test.map_records WHERE uuid = '99f5efbb-046f-4086-9a57-647959953d1f'")
}

func TestSQLDataAccess_UpdateMapStatistic_Where_DataSet_Already_Exists(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	m := mapStatisticWithCheckpointsAlreadyExisting()
	checkMapStatisticUpdateChanges(t, access, UpdateUuid, m.GetName(), m)
}

func TestSQLDataAccess_UpdateMapStatistic_Where_DataSet_Does_Not_Exist(t *testing.T) {
	access, db := connectMariaDB(t)
	defer db.Close()
	m := mapStatisticWithCheckpointsNotExisting()
	checkMapStatisticUpdateChanges(t, access, UpdateUuid, m.GetName(), m)

	// Cleanup
	db.Exec("DELETE FROM test.map_records WHERE uuid = ? AND map = ?", UpdateUuid, m.GetName())
	db.Exec("DELETE FROM test.checkpoint_records WHERE uuid = ? AND map = ?", UpdateUuid, m.GetName())
}

//
// MongoDB specific checks
//

// ...

//
// Map statistic checks
// These functions can be used regardless of the database behind
//

func checkMapStatisticNonExistent(t *testing.T, access DataAccess, id string, mapName string) {
	m := getMapStatistic(t, access, id, mapName)
	assert.Equal(t, uint64(0), m.GetRecordTime())
	assert.Empty(t, m.GetRecordTime())
	assert.Len(t, m.GetCheckpoints(), 0)
}

func checkMapStatisticPlayloadSuccessfully(t *testing.T, access DataAccess) {
	m := getMapStatistic(t, access, RetrievalUuid, "Map1")

	assert.Equal(t, uint64(22324234), m.GetRecordTime())
	assert.Equal(t, "2019-01-01 02:59:23", m.GetAccomplishedAt())
	assert.Equal(t, "Map1", m.GetName())

	assert.Len(t, m.GetCheckpoints(), 3)

	c1 := m.GetCheckpoints()[0]
	assert.Equal(t, int32(1), c1.GetCheckpoint())
	assert.Equal(t, uint64(13424), c1.GetRecordTime())
	assert.Equal(t, "2018-03-05 14:00:23", c1.GetAccomplishedAt())

	c2 := m.GetCheckpoints()[1]
	assert.Equal(t, int32(2), c2.GetCheckpoint())
	assert.Equal(t, uint64(33424), c2.GetRecordTime())
	assert.Equal(t, "2018-03-05 15:00:23", c2.GetAccomplishedAt())

	c3 := m.GetCheckpoints()[2]
	assert.Equal(t, int32(3), c3.GetCheckpoint())
	assert.Equal(t, uint64(53424), c3.GetRecordTime())
	assert.Equal(t, "2018-03-05 16:00:23", c3.GetAccomplishedAt())
}

func checkMapStatisticUpdateChanges(t *testing.T, access DataAccess, id string, mapName string, toUpdate *model.FlashMapStatistic) {
	if err := access.UpdateMapStatistic(id, toUpdate); err != nil {
		t.Errorf("%v", err)
	}

	newStats := getMapStatistic(t, access, id, mapName)

	assert.Equal(t, toUpdate.GetName(), newStats.GetName())
	assert.Equal(t, toUpdate.GetAccomplishedAt(), newStats.GetAccomplishedAt())
	assert.Equal(t, toUpdate.GetRecordTime(), newStats.GetRecordTime())

	for i, c := range newStats.GetCheckpoints() {
		assert.Equal(t, toUpdate.GetCheckpoints()[i].GetCheckpoint(), c.GetCheckpoint())
		assert.Equal(t, toUpdate.GetCheckpoints()[i].GetAccomplishedAt(), c.GetAccomplishedAt())
		assert.Equal(t, toUpdate.GetCheckpoints()[i].GetRecordTime(), c.GetRecordTime())
	}
}

func mapStatisticWithCheckpointsAlreadyExisting() *model.FlashMapStatistic {
	c := make([]*model.FlashCheckpointStatistic, 0)
	c = append(c, &model.FlashCheckpointStatistic{
		Checkpoint:     1,
		AccomplishedAt: time.Now().Format(SqlDateTimeLayout),
		RecordTime:     25565,
	})

	return &model.FlashMapStatistic{
		Name:           "Map1",
		RecordTime:     128,
		AccomplishedAt: time.Now().Format(SqlDateTimeLayout),
		Checkpoints:    c,
	}
}

func mapStatisticWithCheckpointsNotExisting() *model.FlashMapStatistic {
	c := make([]*model.FlashCheckpointStatistic, 0)
	c = append(c, &model.FlashCheckpointStatistic{
		Checkpoint:     1,
		AccomplishedAt: time.Now().Format(SqlDateTimeLayout),
		RecordTime:     255234,
	})
	c = append(c, &model.FlashCheckpointStatistic{
		Checkpoint:     2,
		AccomplishedAt: time.Now().Format(SqlDateTimeLayout),
		RecordTime:     25565,
	})
	return &model.FlashMapStatistic{
		Name:           "Map3",
		RecordTime:     128,
		AccomplishedAt: time.Now().Format("SqlDateTimeLayout"),
		Checkpoints:    c,
	}
}

func getMapStatistic(t *testing.T, access DataAccess, id string, mapName string) *model.FlashMapStatistic {
	m, err := access.GetMapStatistic(id, mapName)
	if err != nil {
		t.Errorf("%v", err)
	}
	return m
}

//
// Game statistic payload checks
//

func checkGameStatisticPayloadSuccessfully(t *testing.T, access DataAccess) {
	m := getGameStatistic(t, access, RetrievalUuid)
	assert.Equal(t, uint32(1337), m.GetWins())
	assert.Equal(t, uint32(12315252), m.GetDeaths())
	assert.Equal(t, uint32(245245), m.GetCheckpoints())
	assert.Equal(t, uint32(5245245), m.GetGamesPlayed())
	assert.Equal(t, uint32(3534535), m.GetInstantDeaths())
}

func checkEmptyGameStatisticPayload(t *testing.T, access DataAccess, id string) {
	m := getGameStatistic(t, access, id)
	assert.Equal(t, uint32(0), m.GetWins())
	assert.Equal(t, uint32(0), m.GetDeaths())
	assert.Equal(t, uint32(0), m.GetCheckpoints())
	assert.Equal(t, uint32(0), m.GetGamesPlayed())
	assert.Equal(t, uint32(0), m.GetInstantDeaths())
}

func checkGameStatisticUpdateChanges(t *testing.T, access DataAccess, id string, toUpdate *model.FlashGameStatistic) {
	oldStats := getGameStatistic(t, access, id)

	if err := access.UpdateGameStatistic(id, toUpdate); err != nil {
		t.Errorf("%v", err)
	}

	newStats := getGameStatistic(t, access, id)

	assert.Equal(t, oldStats.GetWins()+toUpdate.GetWins(), newStats.GetWins())
	assert.Equal(t, oldStats.GetDeaths()+toUpdate.GetDeaths(), newStats.GetDeaths())
	assert.Equal(t, oldStats.GetCheckpoints()+toUpdate.GetCheckpoints(), newStats.GetCheckpoints())
	assert.Equal(t, oldStats.GetGamesPlayed()+toUpdate.GetGamesPlayed(), newStats.GetGamesPlayed())
	assert.Equal(t, oldStats.GetInstantDeaths()+toUpdate.GetInstantDeaths(), newStats.GetInstantDeaths())
}

func getGameStatistic(t *testing.T, access DataAccess, id string) *model.FlashGameStatistic {
	m, err := access.GetGameStatistic(id)
	if err != nil {
		t.Errorf("%v", err)
	}
	return m
}

func gameStatistic() *model.FlashGameStatistic {
	return &model.FlashGameStatistic{
		Wins:          1,
		Deaths:        40,
		GamesPlayed:   1,
		InstantDeaths: 100,
		Checkpoints:   10,
	}
}
