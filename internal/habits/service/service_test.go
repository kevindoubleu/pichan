package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kevindoubleu/pichan/configs"
	"github.com/kevindoubleu/pichan/internal/habits"
	"github.com/kevindoubleu/pichan/internal/habits/service"
	pb "github.com/kevindoubleu/pichan/proto/habits"
	testdata "github.com/kevindoubleu/pichan/test/data/habits"
	mock "github.com/kevindoubleu/pichan/test/mocks/habits/service/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
)

func TestNewScorecardsServer(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		server := service.NewScorecardsServer()

		assert.NotNil(t, server)
		assert.True(t, server.GetStore().IsLive())
	})

	t.Run("fail to get scorecard store", func(t *testing.T) {
		configs.LOG_SKIP_FATAL = true
		validUrl := configs.HABITS_STORE
		configs.HABITS_STORE = "invalid url"

		server := service.NewScorecardsServer()

		assert.False(t, server.GetStore().IsLive())

		configs.LOG_SKIP_FATAL = false
		configs.HABITS_STORE = validUrl
	})
}

func TestDescribe(t *testing.T) {
	svc := service.NewScorecardsServer()
	description, err := svc.Describe(ctx, nil)

	assert.NoError(t, err)
	assert.NotEmpty(t, description.Title)
	assert.NotEmpty(t, description.Subtitle)
	assert.NotEmpty(t, description.Image)
}

func TestInsert(t *testing.T) {
	domainScorecard := testdata.DomainScorecard1
	protoScorecard := testdata.ProtoScorecard1

	t.Run("success: happy case", func(t *testing.T) {
		m := mock.NewStore(t)
		svc := service.NewScorecardsServerWithStore(m)

		m.On("Insert", domainScorecard).Return(&domainScorecard, nil)
		m.On("Get", 0).Return(&domainScorecard, nil)

		expected := protoScorecard
		actual, actualErr := svc.Insert(ctx, expected)

		assert.Equal(t, expected, actual)
		assert.NoError(t, actualErr)
		m.AssertExpectations(t)
	})

	t.Run("fail: Store.Insert failure", func(t *testing.T) {
		m := mock.NewStore(t)
		svc := service.NewScorecardsServerWithStore(m)

		expectedErr := errors.New("error on insert")
		m.On("Insert", domainScorecard).Return(nil, expectedErr)

		actual, actualErr := svc.Insert(ctx, protoScorecard)

		assert.EqualError(t, expectedErr, actualErr.Error())
		assert.Nil(t, actual)
		m.AssertExpectations(t)
	})

	t.Run("fail: Store.Get failure", func(t *testing.T) {
		m := mock.NewStore(t)
		svc := service.NewScorecardsServerWithStore(m)

		expectedErr := errors.New("example error on Store.Get")
		m.On("Insert", domainScorecard).Return(&domainScorecard, nil)
		m.On("Get", 0).Return(nil, expectedErr)

		actual, actualErr := svc.Insert(ctx, protoScorecard)

		assert.EqualError(t, expectedErr, actualErr.Error())
		assert.Nil(t, actual)
		m.AssertExpectations(t)
	})
}

func TestList(t *testing.T) {
	domainScorecards := []habits.Scorecard{
		testdata.DomainScorecard1,
		testdata.DomainScorecard1,
	}
	protoScorecardList := &pb.ScorecardList{
		Scorecards: []*pb.Scorecard{
			testdata.ProtoScorecard1,
			testdata.ProtoScorecard1,
		},
	}

	t.Run("success: happy case", func(t *testing.T) {
		m := mock.NewStore(t)
		svc := service.NewScorecardsServerWithStore(m)

		m.On("List").Return(domainScorecards, nil)

		expected := protoScorecardList
		actual, actualErr := svc.List(ctx, nil)

		assert.Equal(t, expected, actual)
		assert.NoError(t, actualErr)
		m.AssertExpectations(t)
	})

	t.Run("success: empty list", func(t *testing.T) {
		m := mock.NewStore(t)
		svc := service.NewScorecardsServerWithStore(m)

		emptyScorecardSlice := []habits.Scorecard{}
		m.On("List").Return(emptyScorecardSlice, nil)

		expected := &pb.ScorecardList{
			Scorecards: []*pb.Scorecard{},
		}
		actual, actualErr := svc.List(ctx, nil)

		assert.Equal(t, expected, actual)
		assert.NoError(t, actualErr)
		m.AssertExpectations(t)
	})

	t.Run("fail: Store.List failure", func(t *testing.T) {
		m := mock.NewStore(t)
		svc := service.NewScorecardsServerWithStore(m)

		expectedErr := errors.New("example error on Store.List")
		m.On("List").Return(nil, expectedErr)

		actual, actualErr := svc.List(ctx, nil)

		assert.Nil(t, actual)
		assert.EqualError(t, expectedErr, actualErr.Error())
		m.AssertExpectations(t)
	})
}
