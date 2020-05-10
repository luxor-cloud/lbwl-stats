package statsserver

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"freggy.dev/stats/mock"
	"freggy.dev/stats/pkg/flash"
	"freggy.dev/stats/rpc/go/model"
	"freggy.dev/stats/rpc/go/service"
)

//
// GetFlashMapHighscoreForPlayer
//

func Test_GetFlashMapHighscoreForPlayer_WithoutCheckpoints_Successfuly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	damock := mock.NewMockDataAccess(ctrl)
	damock.EXPECT().GetHighscoreForMapByUUID(gomock.Any(), "1234", "some-map")
	damock.EXPECT().GetHighscorePerCheckpointForMapAndUUID(gomock.Any(), "1234", "some-map").Times(0)

	s := &Server{FlashDAO: damock}

	req := &service.GetFlashMapHighscoreForPlayerRequest{
		PlayerId:        "1234",
		WithCheckpoints: false,
		MapName:         "some-map",
	}

	s.GetFlashMapHighscoreForPlayer(context.Background(), req)
}

func TestServer_GetFlashMapHighscoreForPlayer_WithCheckpoints_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	damock := mock.NewMockDataAccess(ctrl)
	damock.EXPECT().GetHighscoreForMapByUUID(gomock.Any(), "1234", "some-map")
	damock.EXPECT().GetHighscorePerCheckpointForMapAndUUID(gomock.Any(), "1234", "some-map")

	s := &Server{FlashDAO: damock}
	req := &service.GetFlashMapHighscoreForPlayerRequest{
		PlayerId:        "1234",
		WithCheckpoints: true,
		MapName:         "some-map",
	}

	s.GetFlashMapHighscoreForPlayer(context.Background(), req)
}

//
// GetGlobalFlashMapHighscore
//

func TestServer_GetGlobalFlashMapHighscore_WithoutCheckpoints_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	damock := mock.NewMockDataAccess(ctrl)
	damock.EXPECT().GetBestHighscore(gomock.Any(), "some-map").Return(flash.PlayerMapScore{
		UUID:           "1234",
		Map:            "some-map",
		TimeNeeded:     1,
		AccomplishedAt: time.Time{},
	}, nil)
	damock.EXPECT().GetHighscorePerCheckpointForMapAndUUID(gomock.Any(), "1234", "some-map").Times(0)

	s := &Server{FlashDAO: damock}
	req := &service.GetGlobalFlashMapHighscoreRequest{
		WithCheckpoints: false,
		MapName:         "some-map",
	}
	s.GetGlobalFlashMapHighscore(context.Background(), req)
}

func TestServer_GetGlobalFlashMapHighscore_WithCheckpoints_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	damock := mock.NewMockDataAccess(ctrl)
	damock.EXPECT().GetBestHighscore(gomock.Any(), "some-map").Return(flash.PlayerMapScore{
		UUID:           "1234",
		Map:            "some-map",
		TimeNeeded:     1,
		AccomplishedAt: time.Time{},
	}, nil)
	damock.EXPECT().GetHighscorePerCheckpointForMapAndUUID(gomock.Any(), "1234", "some-map").Times(1)

	s := &Server{FlashDAO: damock}
	req := &service.GetGlobalFlashMapHighscoreRequest{
		WithCheckpoints: true,
		MapName:         "some-map",
	}
	s.GetGlobalFlashMapHighscore(context.Background(), req)
}

//
// GetTopFlashMapHighscores
//

func TestServer_GetTopFlashMapHighscores_WithoutCheckpoints_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mScores := []flash.PlayerMapScore{
		{
			"1",
			"some-map",
			1,
			time.Time{},
		},
		{
			"2",
			"some-map",
			2,
			time.Time{},
		},
		{
			"3",
			"some-map",
			3,
			time.Time{},
		},
		{
			"4",
			"some-map",
			4,
			time.Time{},
		},
		{
			"5",
			"some-map",
			5,
			time.Time{},
		},
	}

	limit := 5
	damock := mock.NewMockDataAccess(ctrl)
	damock.EXPECT().GetTopHighscores(gomock.Any(), "some-map", limit).Return(mScores, nil)

	s := &Server{FlashDAO: damock}
	req := &service.GetTopFlashMapHighscoresRequest{
		WithCheckpoints: false,
		MapName:         "some-map",
		Limit:           uint32(limit),
	}

	stats, err := s.GetTopFlashMapHighscores(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	expected := make([]*model.FlashMapStatisticCollection, 0)
	for _, score := range mScores {
		mstat := make([]*model.FlashMapStatistic, 0)
		mstat = append(mstat, wrapPlayerMapScore(score, nil))
		expected = append(expected, &model.FlashMapStatisticCollection{
			PlayerId: score.UUID,
			Maps:     mstat,
		})
	}
	assert.ElementsMatch(t, stats.Highscores, expected)
}

func TestServer_GetTopFlashMapHighscores_WithCheckpoints_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mScores := []flash.PlayerMapScore{
		{
			"1",
			"some-map",
			1,
			time.Time{},
		},
		{
			"2",
			"some-map",
			2,
			time.Time{},
		},
		{
			"3",
			"some-map",
			3,
			time.Time{},
		},
		{
			"4",
			"some-map",
			4,
			time.Time{},
		},
		{
			"5",
			"some-map",
			5,
			time.Time{},
		},
	}

	cScores := map[string]flash.PlayerCheckpointScore{
		"1": {
			UUID:           "1",
			Map:            "some-map",
			Checkpoint:     1,
			TimeNeeded:     1,
			AccomplishedAt: time.Time{},
		},
		"2": {
			UUID:           "2",
			Map:            "some-map",
			Checkpoint:     1,
			TimeNeeded:     1,
			AccomplishedAt: time.Time{},
		},
		"3": {
			UUID:           "3",
			Map:            "some-map",
			Checkpoint:     1,
			TimeNeeded:     1,
			AccomplishedAt: time.Time{},
		},
		"4": {
			UUID:           "4",
			Map:            "some-map",
			Checkpoint:     1,
			TimeNeeded:     1,
			AccomplishedAt: time.Time{},
		},
		"5": {
			UUID:           "5",
			Map:            "some-map",
			Checkpoint:     1,
			TimeNeeded:     1,
			AccomplishedAt: time.Time{},
		},
	}

	limit := 5
	damock := mock.NewMockDataAccess(ctrl)
	damock.EXPECT().GetTopHighscores(gomock.Any(), "some-map", limit).Return(mScores, nil)

	for _, score := range mScores {
		damock.EXPECT().
			GetHighscorePerCheckpointForMapAndUUID(gomock.Any(), score.UUID, score.Map).
			Return([]flash.PlayerCheckpointScore{cScores[score.UUID]}, nil)
	}

	s := &Server{FlashDAO: damock}
	req := &service.GetTopFlashMapHighscoresRequest{
		WithCheckpoints: true,
		MapName:         "some-map",
		Limit:           uint32(limit),
	}

	stats, err := s.GetTopFlashMapHighscores(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	expected := make([]*model.FlashMapStatisticCollection, 0)
	for _, score := range mScores {
		mstat := make([]*model.FlashMapStatistic, 0)
		mstat = append(mstat, wrapPlayerMapScore(score, []flash.PlayerCheckpointScore{cScores[score.UUID]}))
		expected = append(expected, &model.FlashMapStatisticCollection{
			PlayerId: score.UUID,
			Maps:     mstat,
		})
	}

	assert.ElementsMatch(t, stats.Highscores, expected)
}

//
// GetTopFlashPlayersByPoints
//


func topPlayersByPoints() []flash.Player {
	return []flash.Player{
		{
			Stats: flash.PlayerStats{
				UUID:          "1",
				Wins:          1,
				Deaths:        1,
				Checkpoints:   1,
				GamesPlayed:   1,
				InstantDeaths: 1,
				Points:        5,
			},
			MapScores: []flash.PlayerMapScore{
				{
					UUID:           "1",
					Map:            "some-map",
					TimeNeeded:     1,
					AccomplishedAt: time.Time{},
				},
			},
			CheckpointScores: []flash.PlayerCheckpointScore{
				{
					UUID:           "1",
					Map:            "some-map",
					Checkpoint:     1,
					TimeNeeded:     1,
					AccomplishedAt: time.Time{},
				},
			},
		},

		{
			Stats: flash.PlayerStats{
				UUID:          "2",
				Wins:          2,
				Deaths:        2,
				Checkpoints:   2,
				GamesPlayed:   2,
				InstantDeaths: 2,
				Points:        4,
			},
			MapScores: []flash.PlayerMapScore{
				{
					UUID:           "2",
					Map:            "some-map",
					TimeNeeded:     2,
					AccomplishedAt: time.Time{},
				},
			},
			CheckpointScores: []flash.PlayerCheckpointScore{
				{
					UUID:           "2",
					Map:            "some-map",
					Checkpoint:     2,
					TimeNeeded:     2,
					AccomplishedAt: time.Time{},
				},
			},
		},

		{
			Stats: flash.PlayerStats{
				UUID:          "3",
				Wins:          3,
				Deaths:        3,
				Checkpoints:   3,
				GamesPlayed:   3,
				InstantDeaths: 3,
				Points:        3,
			},
			MapScores: []flash.PlayerMapScore{
				{
					UUID:           "3",
					Map:            "some-map",
					TimeNeeded:     3,
					AccomplishedAt: time.Time{},
				},
			},
			CheckpointScores: []flash.PlayerCheckpointScore{
				{
					UUID:           "3",
					Map:            "some-map",
					Checkpoint:     3,
					TimeNeeded:     3,
					AccomplishedAt: time.Time{},
				},
			},
		},

		{
			Stats: flash.PlayerStats{
				UUID:          "4",
				Wins:          4,
				Deaths:        4,
				Checkpoints:   4,
				GamesPlayed:   4,
				InstantDeaths: 4,
				Points:        2,
			},
			MapScores: []flash.PlayerMapScore{
				{
					UUID:           "4",
					Map:            "some-map",
					TimeNeeded:     4,
					AccomplishedAt: time.Time{},
				},
			},
			CheckpointScores: []flash.PlayerCheckpointScore{
				{
					UUID:           "4",
					Map:            "some-map",
					Checkpoint:     4,
					TimeNeeded:     4,
					AccomplishedAt: time.Time{},
				},
			},
		},

		{
			Stats: flash.PlayerStats{
				UUID:          "5",
				Wins:          5,
				Deaths:        5,
				Checkpoints:   5,
				GamesPlayed:   5,
				InstantDeaths: 5,
				Points:        1,
			},
			MapScores: []flash.PlayerMapScore{
				{
					UUID:           "5",
					Map:            "some-map",
					TimeNeeded:     5,
					AccomplishedAt: time.Time{},
				},
			},
			CheckpointScores: []flash.PlayerCheckpointScore{
				{
					UUID:           "5",
					Map:            "some-map",
					Checkpoint:     5,
					TimeNeeded:     5,
					AccomplishedAt: time.Time{},
				},
			},
		},
	}
}

func allWithoutMapStats() []*model.FlashAllStatistics {
	players := topPlayersByPoints()
	return []*model.FlashAllStatistics{
		{
			PlayerId:    "1",
			PlayerStats: &model.FlashPlayerStatistic{
				Wins:          players[0].Stats.Wins,
				Deaths:        players[0].Stats.Deaths,
				GamesPlayed:   players[0].Stats.GamesPlayed,
				InstantDeaths: players[0].Stats.InstantDeaths,
				Checkpoints:   players[0].Stats.Checkpoints,
				Points:        players[0].Stats.Points,
			},
			MapStats:    nil,
		},
		{
			PlayerId:    "2",
			PlayerStats: &model.FlashPlayerStatistic{
				Wins:          players[1].Stats.Wins,
				Deaths:        players[1].Stats.Deaths,
				GamesPlayed:   players[1].Stats.GamesPlayed,
				InstantDeaths: players[1].Stats.InstantDeaths,
				Checkpoints:   players[1].Stats.Checkpoints,
				Points:        players[1].Stats.Points,
			},
			MapStats:    nil,
		},
		{
			PlayerId:    "3",
			PlayerStats: &model.FlashPlayerStatistic{
				Wins:          players[2].Stats.Wins,
				Deaths:        players[2].Stats.Deaths,
				GamesPlayed:   players[2].Stats.GamesPlayed,
				InstantDeaths: players[2].Stats.InstantDeaths,
				Checkpoints:   players[2].Stats.Checkpoints,
				Points:        players[2].Stats.Points,
			},
			MapStats:    nil,
		},
		{
			PlayerId:    "4",
			PlayerStats: &model.FlashPlayerStatistic{
				Wins:          players[3].Stats.Wins,
				Deaths:        players[3].Stats.Deaths,
				GamesPlayed:   players[3].Stats.GamesPlayed,
				InstantDeaths: players[3].Stats.InstantDeaths,
				Checkpoints:   players[3].Stats.Checkpoints,
				Points:        players[3].Stats.Points,
			},
			MapStats:    nil,
		},
		{
			PlayerId:    "5",
			PlayerStats: &model.FlashPlayerStatistic{
				Wins:          players[4].Stats.Wins,
				Deaths:        players[4].Stats.Deaths,
				GamesPlayed:   players[4].Stats.GamesPlayed,
				InstantDeaths: players[4].Stats.InstantDeaths,
				Checkpoints:   players[4].Stats.Checkpoints,
				Points:        players[4].Stats.Points,
			},
			MapStats:    nil,
		},
	}
}

func allWithMapStatsWithouCheckpointStats() []*model.FlashAllStatistics {
	players := topPlayersByPoints()
	mapStats := allWithoutMapStats()

	mapStats[0].MapStats = make([]*model.FlashMapStatistic, 0)
	mapStats[1].MapStats = make([]*model.FlashMapStatistic, 0)
	mapStats[2].MapStats = make([]*model.FlashMapStatistic, 0)
	mapStats[3].MapStats = make([]*model.FlashMapStatistic, 0)
	mapStats[4].MapStats = make([]*model.FlashMapStatistic, 0)

	mapStats[0].MapStats = append(mapStats[0].MapStats, wrapPlayerMapScore(players[0].MapScores[0], nil))
	mapStats[1].MapStats = append(mapStats[1].MapStats, wrapPlayerMapScore(players[1].MapScores[0], nil))
	mapStats[2].MapStats = append(mapStats[2].MapStats, wrapPlayerMapScore(players[2].MapScores[0], nil))
	mapStats[3].MapStats = append(mapStats[3].MapStats, wrapPlayerMapScore(players[3].MapScores[0], nil))
	mapStats[4].MapStats = append(mapStats[4].MapStats, wrapPlayerMapScore(players[4].MapScores[0], nil))

	return mapStats
}

func allStats() []*model.FlashAllStatistics {
	players := topPlayersByPoints()
	mapStats := allWithMapStatsWithouCheckpointStats()

	mapStats[0].MapStats[0] = wrapPlayerMapScore(players[0].MapScores[0], players[0].CheckpointScores)
	mapStats[1].MapStats[0] = wrapPlayerMapScore(players[1].MapScores[0], players[1].CheckpointScores)
	mapStats[2].MapStats[0] = wrapPlayerMapScore(players[2].MapScores[0], players[2].CheckpointScores)
	mapStats[3].MapStats[0] = wrapPlayerMapScore(players[3].MapScores[0], players[3].CheckpointScores)
	mapStats[4].MapStats[0] = wrapPlayerMapScore(players[4].MapScores[0], players[4].CheckpointScores)

	return mapStats
}

func TestServer_GetTopFlashPlayersByPoints_WithoutMapStats_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	players := topPlayersByPoints()

	damock := mock.NewMockDataAccess(ctrl)
	damock.EXPECT().GetTopPlayerByPoints(gomock.Any(), len(players)).Return(players, nil)
	damock.EXPECT().GetHighscorePerMapByUUID(gomock.Any(), gomock.Any()).Times(0)
	damock.EXPECT().GetHighscorePerCheckpointForMapAndUUID(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	s := &Server{FlashDAO: damock}
	req := &service.GetTopPlayersByPointsRequest{
		WithMapStats:    false,
		WithCheckpoints: false,
		Limit:           uint32(len(players)),
	}

	top, err := s.GetTopFlashPlayersByPoints(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, top.TopPlayers, allWithoutMapStats())
}

func TestServer_GetTopFlashPlayersByPoints_WithMapStats_WithoutCheckpointStats_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	players := topPlayersByPoints()

	damock := mock.NewMockDataAccess(ctrl)
	damock.EXPECT().GetTopPlayerByPoints(gomock.Any(), len(players)).Return(players, nil)

	for _, player := range players {
		damock.EXPECT().
			GetHighscorePerMapByUUID(gomock.Any(), player.Stats.UUID).
			Times(1).
			Return(player.MapScores, nil)
	}

	damock.EXPECT().GetHighscorePerCheckpointForMapAndUUID(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	s := &Server{FlashDAO: damock}
	req := &service.GetTopPlayersByPointsRequest{
		WithMapStats:    true,
		WithCheckpoints: false,
		Limit:           uint32(len(players)),
	}


	top, err := s.GetTopFlashPlayersByPoints(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, top.TopPlayers, allWithMapStatsWithouCheckpointStats())
}

func TestServer_GetTopFlashPlayersByPoints_WithAllStats_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	players := topPlayersByPoints()

	damock := mock.NewMockDataAccess(ctrl)
	damock.EXPECT().GetTopPlayerByPoints(gomock.Any(), len(players)).Return(players, nil)

	for _, player := range players {
		damock.EXPECT().
			GetHighscorePerMapByUUID(gomock.Any(), player.Stats.UUID).
			Times(1).
			Return(player.MapScores, nil)
	}

	for _, player := range players {
		damock.EXPECT().
			GetHighscorePerCheckpointForMapAndUUID(gomock.Any(), player.Stats.UUID, player.MapScores[0].Map).
			Times(1).
			Return(player.CheckpointScores, nil)

	}

	s := &Server{FlashDAO: damock}
	req := &service.GetTopPlayersByPointsRequest{
		WithMapStats:    true,
		WithCheckpoints: true,
		Limit:           uint32(len(players)),
	}

	top, err := s.GetTopFlashPlayersByPoints(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	r := allStats()

	assert.ElementsMatch(t, top.TopPlayers, r)
}
