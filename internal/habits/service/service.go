package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kevindoubleu/pichan/configs"
	"github.com/kevindoubleu/pichan/internal/habits"
	"github.com/kevindoubleu/pichan/internal/habits/model"
	"github.com/kevindoubleu/pichan/pkg/logger"
	pb "github.com/kevindoubleu/pichan/proto/habits"
	"github.com/kevindoubleu/pichan/proto/pichan"
)

var log = logger.NewLogger("habits.service")

type Store interface {
	IsLive() bool
	Insert(scorecard habits.Scorecard) (*habits.Scorecard, error)
	List() ([]habits.Scorecard, error)
	Get(id int) (*habits.Scorecard, error)
}

type ScorecardsServer struct {
	store Store
	pb.UnimplementedScorecardsServer
}

func NewScorecardsServer(config configs.Habits) ScorecardsServer {
	log.SetSubLabel("NewScorecardsServer")

	store, err := model.GetScorecardStore(config.StoreUrl, config.StoreName)
	if err != nil || store == nil || !store.IsLive() {
		log.Fatalw("failed to get ScorecardStore",
			"err", err,
			"store", store,
			"live", store.IsLive(),
		)
	}

	return ScorecardsServer{
		store: *store,
	}
}

func (s ScorecardsServer) Describe(context.Context, *empty.Empty) (*pichan.Description, error) {
	return &pichan.Description{
		Title:    "Habits Scorecard",
		Subtitle: "Scorecard system for habits, referenced from the book Atomic Habits",
		Image:    "no image yet",
	}, nil
}

func (s ScorecardsServer) Insert(ctx context.Context, scorecard *pb.Scorecard) (*pb.Scorecard, error) {
	log.SetSubLabel("Insert")

	domainScorecard := protoScorecardToDomainScorecard(scorecard)
	insertedDomainScorecard, err := s.store.Insert(domainScorecard)
	if err != nil {
		log.Error("failed inserting to store", err)
		return nil, err
	}

	retrievedDomainScorecard, err := s.store.Get(insertedDomainScorecard.Id)
	if err != nil {
		log.Error("failed retrieving inserted scorecard", err)
		return nil, err
	}

	return domainScorecardToProtoScorecard(*retrievedDomainScorecard), nil
}

func (s ScorecardsServer) List(context.Context, *empty.Empty) (*pb.ScorecardList, error) {
	log.SetSubLabel("List")
	result := &pb.ScorecardList{
		Scorecards: make([]*pb.Scorecard, 0),
	}

	storeResult, err := s.store.List()
	if err != nil {
		log.Error("failed to list scorecards", err)
		return nil, err
	}

	for _, storeResult := range storeResult {
		pbScorecard := domainScorecardToProtoScorecard(storeResult)
		result.Scorecards = append(result.Scorecards, pbScorecard)
	}

	log.Infow("success")
	return result, nil
}

func domainScorecardToProtoScorecard(scorecard habits.Scorecard) *pb.Scorecard {
	return &pb.Scorecard{
		Id:          int32(scorecard.Id),
		Name:        scorecard.Name,
		Connotation: pb.Connotation(pb.Connotation_value[scorecard.Connotation]),
		Time:        scorecard.Time,
		Order:       int32(scorecard.Order),
	}
}

func protoScorecardToDomainScorecard(scorecard *pb.Scorecard) habits.Scorecard {
	return habits.Scorecard{
		Id:          int(scorecard.Id),
		Name:        scorecard.Name,
		Connotation: scorecard.Connotation.String(),
		Time:        scorecard.Time,
		Order:       int(scorecard.Order),
	}
}
