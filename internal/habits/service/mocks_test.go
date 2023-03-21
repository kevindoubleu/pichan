package service_test

import (
	"github.com/kevindoubleu/pichan/internal/habits"
	"github.com/stretchr/testify/mock"
)

type mockScorecardStore struct {
	mock.Mock
}

func newMockScorecardStore() *mockScorecardStore {
	return &mockScorecardStore{}
}

func (m *mockScorecardStore) IsLive() bool {
	expectedReturns := m.Called()
	return expectedReturns.Bool(0)
}

func (m *mockScorecardStore) Insert(scorecard habits.Scorecard) (*habits.Scorecard, error) {
	expectedReturns := m.Called(scorecard)
	if expectedReturns.Get(0) != nil {
		return expectedReturns.Get(0).(*habits.Scorecard), expectedReturns.Error(1)
	}
	return nil, expectedReturns.Error(1)
}

func (m *mockScorecardStore) List() ([]habits.Scorecard, error) {
	expectedReturns := m.Called()
	if expectedReturns.Get(0) != nil {
		return expectedReturns.Get(0).([]habits.Scorecard), expectedReturns.Error(1)
	}
	return nil, expectedReturns.Error(1)
}

func (m *mockScorecardStore) Get(id int) (*habits.Scorecard, error) {
	expectedReturns := m.Called(id)
	if expectedReturns.Get(0) != nil {
		return expectedReturns.Get(0).(*habits.Scorecard), expectedReturns.Error(1)
	}
	return nil, expectedReturns.Error(1)
}
