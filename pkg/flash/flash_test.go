package flash

import (
	"database/sql"
	"freggy.dev/stats/rpc/go/model"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const UUID = "92de217b-8b2b-403b-86a5-fe26fa3a9b5f"

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
	checkMapStatisticNonExistent(t, access, UUID, "Map42")
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

//
// MongoDB specific checks
//

// ...


//
// Map statistic payload checks
//

func checkMapStatisticNonExistent(t *testing.T,access DataAccess, id string, mapName string) {
	m := getMapStatistic(t, access, id, mapName)
	assert.Equal(t, uint64(0), m.GetRecordTime())
	assert.Empty(t, m.GetRecordTime())
	assert.Len(t, m.GetCheckpoints(), 0)
}

func checkMapStatisticPlayloadSuccessfully(t *testing.T, access DataAccess) {
	m := getMapStatistic(t, access, UUID, "Map1")

	assert.Equal(t, uint64(22324234), m.GetRecordTime())
	assert.Equal(t, "2019-01-01 02:59:23 +0000 UTC", m.GetAccomplishedAt())
	assert.Equal(t, "Map1", m.GetName())

	assert.Len(t, m.GetCheckpoints(), 3)

	c1 := m.GetCheckpoints()[0]
	assert.Equal(t, int32(1), c1.GetCheckpoint())
	assert.Equal(t, uint64(13424), c1.GetRecordTime())
	assert.Equal(t, "2018-03-05 14:00:23 +0000 UTC", c1.GetAccomplishedAt())

	c2 := m.GetCheckpoints()[1]
	assert.Equal(t, int32(2), c2.GetCheckpoint())
	assert.Equal(t, uint64(33424), c2.GetRecordTime())
	assert.Equal(t, "2018-03-05 15:00:23 +0000 UTC", c2.GetAccomplishedAt())

	c3 := m.GetCheckpoints()[2]
	assert.Equal(t, int32(3), c3.GetCheckpoint())
	assert.Equal(t, uint64(53424), c3.GetRecordTime())
	assert.Equal(t, "2018-03-05 16:00:23 +0000 UTC", c3.GetAccomplishedAt())
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
	m := getGameStatistic(t, access, UUID)
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

func getGameStatistic(t *testing.T, access DataAccess, id string) *model.FlashGameStatistic {
	m, err := access.GetGameStatistic(id)
	if err != nil {
		t.Errorf("%v", err)
	}
	return m
}
