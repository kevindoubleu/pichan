package habits_test

import (
	"context"
	"net"
	"testing"

	"github.com/kevindoubleu/pichan/configs"
	"github.com/kevindoubleu/pichan/internal/habits"
	"github.com/kevindoubleu/pichan/internal/habits/model"
	"github.com/kevindoubleu/pichan/internal/habits/service"
	"github.com/kevindoubleu/pichan/pkg/db"
	pb "github.com/kevindoubleu/pichan/proto/habits"
	testdata "github.com/kevindoubleu/pichan/test/data/habits"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HabitsIntegrationSuite struct {
	suite.Suite
	ctx context.Context

	config *configs.Config
	store  *model.ScorecardStore

	listener   *bufconn.Listener
	grpcClient pb.ScorecardsClient
}

func TestHabitsIntegrationSuite(t *testing.T) {
	suite.Run(t, new(HabitsIntegrationSuite))
}

func (s *HabitsIntegrationSuite) SetupSuite() {
	s.config, _ = configs.NewConfig("../../configs/" + configs.TestConfigFile)
	s.store, _ = model.GetScorecardStore(s.config.StoreUrl, s.config.StoreName)

	s.ctx = context.Background()
	s.listener = createBufConnListener()
	dialer := createBufConnDialer(s.listener)
	s.grpcClient = getGrpcClient(dialer)

	startTestScorecardGrpcServer(s.config, s.listener)
}

func createBufConnListener() *bufconn.Listener {
	bufSize := 1024 * 1024
	return bufconn.Listen(bufSize)
}

func createBufConnDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, s string) (net.Conn, error) {
		return listener.Dial()
	}
}

func getGrpcClient(dialer func(context.Context, string) (net.Conn, error)) pb.ScorecardsClient {
	grpcConn, err := grpc.Dial("",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return pb.NewScorecardsClient(grpcConn)
}

func startTestScorecardGrpcServer(config *configs.Config, listener *bufconn.Listener) {
	grpcServer := grpc.NewServer()
	pb.RegisterScorecardsServer(grpcServer, service.NewScorecardsServer(*config))
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()
}

func (s *HabitsIntegrationSuite) SetupTest() {
	s.store.GetDb().Exec(db.DropTable(s.config.StoreName))
	s.store.GetDb().Exec(db.CreateTable(s.config.StoreName, habits.ScorecardSchema))
}

func (s *HabitsIntegrationSuite) TearDownSuite() {
	s.store.GetDb().Exec(db.DropTable(s.config.StoreName))
	s.listener.Close()
}

func (s *HabitsIntegrationSuite) TestDescribe_ShouldReturnDescription() {
	desc, err := s.grpcClient.Describe(s.ctx, &emptypb.Empty{})

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), desc)
	assert.NotNil(s.T(), desc.Title)
	assert.NotNil(s.T(), desc.Subtitle)
	assert.NotNil(s.T(), desc.Image)
}

func (s *HabitsIntegrationSuite) TestInsert_ShouldInsertIntoScorecardStore() {
	insertedScorecard, err := s.grpcClient.Insert(s.ctx, testdata.ProtoScorecard1)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), insertedScorecard)
	assertEqualProtoAndDomainScorecard(s.T(), testdata.ProtoScorecard1, insertedScorecard)
}

func assertEqualProtoAndDomainScorecard(t *testing.T, sc1, sc2 *pb.Scorecard) {
	assert.Equal(t, sc1.Name, sc2.Name)
	assert.Equal(t, sc1.Time, sc2.Time)
	assert.Equal(t, sc1.Order, sc2.Order)
	assert.Equal(t, sc1.Connotation, sc2.Connotation)
}

func (s *HabitsIntegrationSuite) TestList_ShouldReturnListOfScorecards() {
	scorecards, err := s.grpcClient.List(s.ctx, &emptypb.Empty{})
	shouldReturnEmptyList(s.T(), scorecards, err)

	s.store.Insert(testdata.DomainScorecard1)
	scorecards, err = s.grpcClient.List(s.ctx, &emptypb.Empty{})
	shouldReturnListWithScorecard(s.T(), scorecards, err)

	s.store.Insert(testdata.DomainScorecard1)
	scorecards, err = s.grpcClient.List(s.ctx, &emptypb.Empty{})
	shouldReturnListWithMultipleScorecards(s.T(), scorecards, err)
}

func shouldReturnEmptyList(t *testing.T, scorecards *pb.ScorecardList, err error) {
	assert.NoError(t, err)
	assert.NotNil(t, scorecards)
	assert.Empty(t, scorecards.Scorecards)
}

func shouldReturnListWithScorecard(t *testing.T, scorecards *pb.ScorecardList, err error) {
	assert.NoError(t, err)
	assert.NotNil(t, scorecards)
	assert.Len(t, scorecards.Scorecards, 1)
	assertEqualProtoAndDomainScorecard(t, testdata.ProtoScorecard1, scorecards.Scorecards[0])
}

func shouldReturnListWithMultipleScorecards(t *testing.T, scorecards *pb.ScorecardList, err error) {
	assert.NoError(t, err)
	assert.NotNil(t, scorecards)
	assert.Len(t, scorecards.Scorecards, 2)
	assertEqualProtoAndDomainScorecard(t, testdata.ProtoScorecard1, scorecards.Scorecards[0])
	assertEqualProtoAndDomainScorecard(t, testdata.ProtoScorecard1, scorecards.Scorecards[1])
}
