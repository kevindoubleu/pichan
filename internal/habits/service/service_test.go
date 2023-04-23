package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kevindoubleu/pichan/internal/habits"
	pb "github.com/kevindoubleu/pichan/proto/habits"
	"github.com/stretchr/testify/suite"

	"github.com/kevindoubleu/pichan/configs"
	"github.com/kevindoubleu/pichan/internal/habits/service"
	testdata "github.com/kevindoubleu/pichan/test/data/habits"
	mock "github.com/kevindoubleu/pichan/test/mocks/habits/service/mocks"
	"github.com/stretchr/testify/assert"
)

type ServiceSuite struct {
	suite.Suite

	config configs.Config
	ctx    context.Context

	mockStore *mock.Store
	service   service.ScorecardsServer
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupSuite() {
	config, _ := configs.NewConfig("../../../configs/" + configs.TestConfigFile)
	s.config = *config
	s.ctx = context.Background()
}

func (s *ServiceSuite) SetupTest() {
	s.mockStore = &mock.Store{}
	s.service = service.NewScorecardsServerWithStore(s.config, s.mockStore)
}

func (s *ServiceSuite) TearDownTest() {
	s.mockStore.AssertExpectations(s.T())
}

var (
	// TestInsert
	domainScorecard = testdata.DomainScorecard1
	protoScorecard  = testdata.ProtoScorecard1

	// TestList
	domainScorecards = []habits.Scorecard{
		testdata.DomainScorecard1,
		testdata.DomainScorecard1,
	}
	protoScorecardList = &pb.ScorecardList{
		Scorecards: []*pb.Scorecard{
			testdata.ProtoScorecard1,
			testdata.ProtoScorecard1,
		},
	}
)

func (s *ServiceSuite) TestNewScorecardsServer_ValidConfig_ShouldGetLiveServer() {
	server := service.NewScorecardsServer(s.config)

	assert.NotNil(s.T(), server)
	assert.True(s.T(), server.GetStore().IsLive())
}

func (s *ServiceSuite) TestNewScorecardsServer_InvalidConfig_ShouldNotGetLiveServer() {
	configWithInvalidUrl, _ := configs.NewConfig("../../../configs/" + configs.TestConfigFile)
	configWithInvalidUrl.Habits.StoreUrl = "invalid url"

	server := service.NewScorecardsServer(*configWithInvalidUrl)

	assert.False(s.T(), server.GetStore().IsLive())
}

func (s *ServiceSuite) TestDescribe_ShouldGetAllDescriptionFields() {
	svc := service.NewScorecardsServer(s.config)
	description, err := svc.Describe(s.ctx, nil)

	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), description.Title)
	assert.NotEmpty(s.T(), description.Subtitle)
	assert.NotEmpty(s.T(), description.Image)
}

func (s *ServiceSuite) TestInsert_ValidDomainScorecard_ShouldInsertValidPbScorecard() {
	s.mockStore.On("Insert", domainScorecard).Return(&domainScorecard, nil)
	s.mockStore.On("Get", 0).Return(&domainScorecard, nil)

	expected := protoScorecard
	actual, actualErr := s.service.Insert(s.ctx, expected)

	assert.Equal(s.T(), expected, actual)
	assert.NoError(s.T(), actualErr)
}

func (s *ServiceSuite) TestInsert_StoreInsertFails_ShouldReturnError() {
	expectedErr := errors.New("error on insert")
	s.mockStore.On("Insert", domainScorecard).Return(nil, expectedErr)

	actual, actualErr := s.service.Insert(s.ctx, protoScorecard)

	assert.EqualError(s.T(), expectedErr, actualErr.Error())
	assert.Nil(s.T(), actual)
}

func (s *ServiceSuite) TestInsert_StoreGetFails_ShouldReturnError() {
	expectedErr := errors.New("example error on Store.Get")
	s.mockStore.On("Insert", domainScorecard).Return(&domainScorecard, nil)
	s.mockStore.On("Get", 0).Return(nil, expectedErr)

	actual, actualErr := s.service.Insert(s.ctx, protoScorecard)

	assert.EqualError(s.T(), expectedErr, actualErr.Error())
	assert.Nil(s.T(), actual)
}

func (s *ServiceSuite) TestList_StoreHasScorecards_ShouldReturnListOfScorecards() {
	s.mockStore.On("List").Return(domainScorecards, nil)

	expected := protoScorecardList
	actual, actualErr := s.service.List(s.ctx, nil)

	assert.Equal(s.T(), expected, actual)
	assert.NoError(s.T(), actualErr)
}

func (s *ServiceSuite) TestList_StoreHasNoScorecards_ShouldReturnEmptyListOfScorecards() {
	emptyScorecardSlice := []habits.Scorecard{}
	s.mockStore.On("List").Return(emptyScorecardSlice, nil)

	expected := &pb.ScorecardList{
		Scorecards: []*pb.Scorecard{},
	}
	actual, actualErr := s.service.List(s.ctx, nil)

	assert.Equal(s.T(), expected, actual)
	assert.NoError(s.T(), actualErr)
}

func (s *ServiceSuite) TestList_StoreListFails_ShouldReturnError() {
	expectedErr := errors.New("example error on Store.List")
	s.mockStore.On("List").Return(nil, expectedErr)

	actual, actualErr := s.service.List(s.ctx, nil)

	assert.Nil(s.T(), actual)
	assert.EqualError(s.T(), expectedErr, actualErr.Error())
}
