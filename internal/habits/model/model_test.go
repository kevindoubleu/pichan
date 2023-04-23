package model_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/kevindoubleu/pichan/configs"
	"github.com/kevindoubleu/pichan/internal/habits"
	"github.com/kevindoubleu/pichan/internal/habits/model"
	"github.com/kevindoubleu/pichan/pkg/db"
	"github.com/kevindoubleu/pichan/pkg/logger"
)

var (
	log    = logger.NewLogger("habits.model_test")
	config *configs.Config
)

func TestMain(m *testing.M) {
	config, _ = configs.NewConfig("../../../configs/" + configs.TestConfigFile)
	m.Run()
}

func TestGetScorecardStore(t *testing.T) {
	testCases := []struct {
		desc string
		conn string
		live bool
	}{
		{
			desc: "invalid conn string",
			conn: "invalid://invalid",
			live: false,
		}, {
			desc: "valid conn string",
			conn: config.StoreUrl,
			live: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			store, _ := model.GetScorecardStore(tC.conn, config.StoreName)
			if tC.live != store.IsLive() {
				t.Errorf("TestGetScorecardStore: liveness %t != %t", tC.live, store.IsLive())
			}
		})
	}
}

func resetTestDb() {
	log.SetSubLabel("setup")

	store, err := model.GetScorecardStore(config.StoreUrl, config.StoreName)
	if err != nil && store.IsLive() {
		log.FatalError("error initializing test db", err)
	}
	sqlDb := store.GetDb()
	defer sqlDb.Close()

	result, err := sqlDb.Exec(db.DropTable(config.StoreName))
	log.Infow("dropping table",
		"result", result,
		"err", err,
	)
	result, err = sqlDb.Exec(db.CreateTable(config.StoreName, habits.ScorecardSchema))
	log.Infow("creating table",
		"result", result,
		"err", err,
	)
}

func TestInsert(t *testing.T) {
	resetTestDb()
	store, err := model.GetScorecardStore(config.StoreUrl, config.StoreName)
	if err != nil || store == nil {
		t.Fatal("TestInsert: cannot get ScorecardStore", err)
	}

	testCases := []struct {
		desc string
		row  *habits.Scorecard
		want *habits.Scorecard
		err  error
	}{
		{
			desc: "all fields filled",
			row: &habits.Scorecard{
				Name:        "test-name",
				Connotation: "Neutral",
				Time:        "01:02",
				Order:       1,
			},
			want: &habits.Scorecard{
				Name:        "test-name",
				Connotation: "Neutral",
				Time:        "01:02",
				Order:       1,
			},
		}, {
			desc: "empty order field filled as last order + 500",
			row: &habits.Scorecard{
				Name:        "test-name",
				Connotation: "Neutral",
				Time:        "01:02",
			},
			want: &habits.Scorecard{
				Name:        "test-name",
				Connotation: "Neutral",
				Time:        "01:02",
				Order:       501,
			},
		}, {
			desc: "empty time is filled as 00:00",
			row: &habits.Scorecard{
				Name:        "test-name",
				Connotation: "Neutral",
				Order:       3,
			},
			want: &habits.Scorecard{
				Name:        "test-name",
				Connotation: "Neutral",
				Time:        "00:00",
				Order:       3,
			},
		}, {
			desc: "invalid time range",
			row: &habits.Scorecard{
				Name:        "test-name",
				Connotation: "Neutral",
				Time:        "99:00",
			},
			err: errors.New("parsing time \"99:00\": hour out of range"),
		}, {
			desc: "invalid connotation",
			row: &habits.Scorecard{
				Name:        "test-name",
				Connotation: "invalid-connotation",
			},
			err: errors.New("pq: invalid input value for enum habit_connotation: \"invalid-connotation\""),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := store.Insert(*tC.row)

			if tC.want != nil && !equalNotNilScorecards(tC.want, result) {
				t.Errorf("TestInsert: %+v != %+v", tC.want, result)
			}

			if tC.err != nil && tC.err.Error() != err.Error() {
				t.Errorf("TestInsert err: %s != %s", tC.err, err)
			}
		})
	}
}

func TestList(t *testing.T) {
	resetTestDb()
	store, err := model.GetScorecardStore(config.StoreUrl, config.StoreName)
	if err != nil || store == nil {
		t.Fatal("TestInsert: cannot get ScorecardStore", err)
	}
	store.Insert(habits.Scorecard{
		Name:        "TestList-test-name-1",
		Connotation: "Positive",
	})
	store.Insert(habits.Scorecard{
		Name:        "TestList-test-name-2",
		Connotation: "Negative",
		Time:        "23:59",
	})

	t.Run("gets all scorecards", func(t *testing.T) {
		result, err := store.List()
		lenWant := 2

		if lenWant != len(result) || err != nil {
			t.Errorf("TestList: %d != %d", lenWant, len(result))
		}
	})

	t.Run("parses and returns all fields", func(t *testing.T) {
		result, err := store.List()
		want := []habits.Scorecard{
			{
				Id:          1,
				Name:        "TestList-test-name-1",
				Connotation: "Positive",
				Time:        "00:00",
				Order:       500,
			}, {
				Id:          2,
				Name:        "TestList-test-name-2",
				Connotation: "Negative",
				Time:        "23:59",
				Order:       1000,
			},
		}

		if !equalNotNilScorecardList(want, result) || err != nil {
			t.Errorf("TestList: %+v != %+v", want, result)
		}
	})
}

func TestGet(t *testing.T) {
	resetTestDb()
	store, err := model.GetScorecardStore(config.StoreUrl, config.StoreName)
	if err != nil || store == nil {
		t.Fatal("TestInsert: cannot get ScorecardStore", err)
	}
	store.Insert(habits.Scorecard{
		Name:        "TestGet-test-name",
		Connotation: "Positive",
	})

	testCases := []struct {
		desc string
		id   int
		want *habits.Scorecard
		err  error
	}{
		{
			desc: "get found",
			id:   1,
			want: &habits.Scorecard{
				Name:        "TestGet-test-name",
				Connotation: "Positive",
				Time:        "00:00",
				Order:       500,
			},
		}, {
			desc: "get not found",
			id:   -1,
			err:  sql.ErrNoRows,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := store.Get(tC.id)

			if tC.want != nil && !equalNotNilScorecards(tC.want, result) {
				t.Errorf("TestInsert: %+v != %+v", tC.want, result)
			}

			if tC.err != nil && tC.err.Error() != err.Error() {
				t.Errorf("TestInsert err: %s != %s", tC.err, err)
			}
		})
	}
}

func equalNotNilScorecardList(expecteds, actuals []habits.Scorecard) bool {
	if len(expecteds) != len(actuals) {
		return false
	}

	for i := range expecteds {
		if !equalNotNilScorecards(&expecteds[i], &actuals[i]) {
			return false
		}
	}

	return true
}

func equalNotNilScorecards(expected, actual *habits.Scorecard) bool {
	return expected.Name == actual.Name &&
		expected.Connotation == actual.Connotation &&
		expected.Time == actual.Time &&
		expected.Order == actual.Order
}
