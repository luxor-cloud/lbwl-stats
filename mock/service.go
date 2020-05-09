// Code generated by MockGen. DO NOT EDIT.
// Source: freggy.dev/stats/rpc/go/service (interfaces: StatsService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	service "freggy.dev/stats/rpc/go/service"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStatsService is a mock of StatsService interface
type MockStatsService struct {
	ctrl     *gomock.Controller
	recorder *MockStatsServiceMockRecorder
}

// MockStatsServiceMockRecorder is the mock recorder for MockStatsService
type MockStatsServiceMockRecorder struct {
	mock *MockStatsService
}

// NewMockStatsService creates a new mock instance
func NewMockStatsService(ctrl *gomock.Controller) *MockStatsService {
	mock := &MockStatsService{ctrl: ctrl}
	mock.recorder = &MockStatsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatsService) EXPECT() *MockStatsServiceMockRecorder {
	return m.recorder
}

// GetFlashMapHighscoreForPlayer mocks base method
func (m *MockStatsService) GetFlashMapHighscoreForPlayer(arg0 context.Context, arg1 *service.GetFlashMapHighscoreForPlayerRequest) (*service.GetFlashMapHighscoreForPlayerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlashMapHighscoreForPlayer", arg0, arg1)
	ret0, _ := ret[0].(*service.GetFlashMapHighscoreForPlayerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlashMapHighscoreForPlayer indicates an expected call of GetFlashMapHighscoreForPlayer
func (mr *MockStatsServiceMockRecorder) GetFlashMapHighscoreForPlayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlashMapHighscoreForPlayer", reflect.TypeOf((*MockStatsService)(nil).GetFlashMapHighscoreForPlayer), arg0, arg1)
}

// GetGlobalFlashMapHighscore mocks base method
func (m *MockStatsService) GetGlobalFlashMapHighscore(arg0 context.Context, arg1 *service.GetGlobalFlashMapHighscoreRequest) (*service.GetGlobalFlashMapHighscoreResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGlobalFlashMapHighscore", arg0, arg1)
	ret0, _ := ret[0].(*service.GetGlobalFlashMapHighscoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGlobalFlashMapHighscore indicates an expected call of GetGlobalFlashMapHighscore
func (mr *MockStatsServiceMockRecorder) GetGlobalFlashMapHighscore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGlobalFlashMapHighscore", reflect.TypeOf((*MockStatsService)(nil).GetGlobalFlashMapHighscore), arg0, arg1)
}

// GetTopFlashMapHighscores mocks base method
func (m *MockStatsService) GetTopFlashMapHighscores(arg0 context.Context, arg1 *service.GetTopFlashMapHighscoresRequest) (*service.GetTopFlashMapHighscoresResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopFlashMapHighscores", arg0, arg1)
	ret0, _ := ret[0].(*service.GetTopFlashMapHighscoresResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopFlashMapHighscores indicates an expected call of GetTopFlashMapHighscores
func (mr *MockStatsServiceMockRecorder) GetTopFlashMapHighscores(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopFlashMapHighscores", reflect.TypeOf((*MockStatsService)(nil).GetTopFlashMapHighscores), arg0, arg1)
}

// GetTopFlashPlayersByPoints mocks base method
func (m *MockStatsService) GetTopFlashPlayersByPoints(arg0 context.Context, arg1 *service.GetTopPlayersByPointsRequest) (*service.GetTopPlayersByPointsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopFlashPlayersByPoints", arg0, arg1)
	ret0, _ := ret[0].(*service.GetTopPlayersByPointsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopFlashPlayersByPoints indicates an expected call of GetTopFlashPlayersByPoints
func (mr *MockStatsServiceMockRecorder) GetTopFlashPlayersByPoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopFlashPlayersByPoints", reflect.TypeOf((*MockStatsService)(nil).GetTopFlashPlayersByPoints), arg0, arg1)
}

// UpdateFlashStatistics mocks base method
func (m *MockStatsService) UpdateFlashStatistics(arg0 context.Context, arg1 *service.UpdateFlashStatisticsRequests) (*service.UpdateFlashStatisticsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFlashStatistics", arg0, arg1)
	ret0, _ := ret[0].(*service.UpdateFlashStatisticsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFlashStatistics indicates an expected call of UpdateFlashStatistics
func (mr *MockStatsServiceMockRecorder) UpdateFlashStatistics(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlashStatistics", reflect.TypeOf((*MockStatsService)(nil).UpdateFlashStatistics), arg0, arg1)
}
