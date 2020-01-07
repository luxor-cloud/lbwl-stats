package statsserver

import (
	"freggy.dev/stats/rpc/go/model"
	"freggy.dev/stats/rpc/go/service"
	"testing"
)

type flashDataAccessMock struct {
	mapStats  map[string]map[string]*model.FlashMapStatistic
	gameStats map[string]*model.FlashGameStatistic
}

func NewFlashMockDAO() *flashDataAccessMock {
	mock := flashDataAccessMock{}
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
		AccomplishedAt: "",
		RecordTime:     1,
	})

	c = append(c, &model.FlashCheckpointStatistic{
		Checkpoint:     2,
		AccomplishedAt: "",
		RecordTime:     2,
	})

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
		return &model.FlashMapStatistic{}, nil
	}
	return stats, nil
}

func (fda flashDataAccessMock) GetGameStatistic(id string) (*model.FlashGameStatistic, error) {
	stats := fda.gameStats[id]
	if stats == nil {
		return &model.FlashGameStatistic{}, nil
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

func TestServer_GetFlashGameStats_Valid_Response(t *testing.T) {
	s := getMockServer()

	req := &service.GetFlashGameStatsRequest{
		PlayerId: "92de217b-8b2b-403b-86a5-fe26fa3a9b5f",
	}

	resp, err := s.GetFlashGameStats(nil, req)
	if err != nil {
		t.Errorf("%v", err)
	}

	resp.Stats
}
