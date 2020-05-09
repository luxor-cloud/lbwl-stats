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
	assert.ElementsMatch(t, stats.Highscores  , expected)
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

	assert.ElementsMatch(t, stats.Highscores  , expected)
}

//
// GetTopFlashPlayersByPoints
//



