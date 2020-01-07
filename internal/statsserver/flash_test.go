package statsserver

import (
	"freggy.dev/stats/rpc/go/model"
	"freggy.dev/stats/rpc/go/service"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type flashDataAccessMock struct {
	mapStats  map[string]map[string]*model.FlashMapStatistic
	gameStats map[string]*model.FlashGameStatistic
}

// TODO: use same data in mock as in integration tests

func NewFlashMockDAO() *flashDataAccessMock {
	mock := flashDataAccessMock{
		mapStats: make(map[string]map[string]*model.FlashMapStatistic),
		gameStats: make(map[string]*model.FlashGameStatistic),
	}

	log.Println(mock.mapStats)
	log.Println(mock.gameStats)

	mock.gameStats["92de217b-8b2b-403b-86a5-fe26fa3a9b5f"] = &model.FlashGameStatistic{
		Wins:          1200,
		Deaths:        1337,
		GamesPlayed:   200,
		InstantDeaths: 200,
		Checkpoints:   500,
	}

	c := make([]*model.FlashCheckpointStatistic, 0)
	c = append(c, &model.FlashCheckpointStatistic{
		Checkpoint:     1,
		AccomplishedAt: "2011-03-05 16:00:23",
		RecordTime:     1,
	})

	c = append(c, &model.FlashCheckpointStatistic{
		Checkpoint:     2,
		AccomplishedAt: "2012-03-05 16:00:23",
		RecordTime:     2,
	})

	mock.mapStats["92de217b-8b2b-403b-86a5-fe26fa3a9b5f"] = make(map[string]*model.FlashMapStatistic)
	mock.mapStats["92de217b-8b2b-403b-86a5-fe26fa3a9b5f"]["Map1"] = &model.FlashMapStatistic{
		Name:           "Map1",
		RecordTime:     200,
		AccomplishedAt: "2018-03-05 16:00:23",
		Checkpoints:    c,
	}

	return &mock
}

func (fda flashDataAccessMock) GetMapStatistic(id string, mapName string) (*model.FlashMapStatistic, error) {
	stats := fda.mapStats[id][mapName]
	if stats == nil {
		return nil, nil
	}
	return stats, nil
}

func (fda flashDataAccessMock) GetGameStatistic(id string) (*model.FlashGameStatistic, error) {
	stats := fda.gameStats[id]
	if stats == nil {
		return nil, nil
	}
	return stats, nil
}

func (flashDataAccessMock) UpdateGameStatistic(id string, stats *model.FlashGameStatistic) error {
	return nil
}

func (flashDataAccessMock) UpdateMapStatistic(id string, stats *model.FlashMapStatistic) error {
	return nil
}

func getMockServer() *Server {
	return &Server{
		flashDAO: NewFlashMockDAO(),
	}
}

func TestServer_GetFlashMapStats_Valid_Response(t *testing.T) {
	s := getMockServer()
	req := &service.GetFlashMapStatsRequest{
		PlayerId: "92de217b-8b2b-403b-86a5-fe26fa3a9b5f",
		Maps: []string{"Map1"},
	}

	resp, err := s.GetFlashMapStats(nil, req)
	if err != nil {
		t.Errorf("%v", err)
	}

	maps := resp.GetStats().GetMapSummary()
	assert.Len(t, maps, 1)

	assert.Equal(t, "Map1", maps[0].GetName())
	assert.Equal(t, "2018-03-05 16:00:23", maps[0].GetAccomplishedAt())
	assert.Equal(t, uint64(200), maps[0].GetRecordTime())

	checkpoints := maps[0].GetCheckpoints()
	assert.Len(t, checkpoints, 2)

	assert.Equal(t, int32(1), checkpoints[0].GetCheckpoint())
	assert.Equal(t, uint64(1), checkpoints[0].GetRecordTime())
	assert.Equal(t, "2011-03-05 16:00:23", checkpoints[0].GetAccomplishedAt())

	assert.Equal(t, int32(2), checkpoints[1].GetCheckpoint())
	assert.Equal(t, uint64(2), checkpoints[1].GetRecordTime())
	assert.Equal(t, "2012-03-05 16:00:23", checkpoints[1].GetAccomplishedAt())
}

func TestServer_GetFlashMapStats_Invalid_Map(t *testing.T) {
	s := getMockServer()
	req := &service.GetFlashMapStatsRequest{
		PlayerId: "92de217b-8b2b-403b-86a5-fe26fa3a9b5f",
		Maps: []string{"Map42"},
	}

	resp, err := s.GetFlashMapStats(nil, req)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Empty(t, resp.GetStats().GetMapSummary())
}

func TestServer_GetFlashGameStats_Valid_Response(t *testing.T) {
	s := getMockServer()

	req := &service.GetFlashGameStatsRequest{
		PlayerId: "92de217b-8b2b-403b-86a5-fe26fa3a9b5f",
	}

	resp, err := s.GetFlashGameStats(nil, req)
	if err != nil {
		t.Errorf("%v", err)
	}

	g := resp.GetStats().GetGameSummary()

	assert.NotNil(t, g)
	assert.Equal(t, 1200, g.GetWins())
	assert.Equal(t, 1337, g.GetDeaths())
	assert.Equal(t, 200, g.GetGamesPlayed())
	assert.Equal(t, 200, g.GetInstantDeaths())
	assert.Equal(t, 500, g.GetCheckpoints())
}
