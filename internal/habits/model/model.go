package model

import (
	"database/sql"
	"time"

	"github.com/kevindoubleu/pichan/internal/habits"
	"github.com/kevindoubleu/pichan/pkg/logger"
	"github.com/lib/pq"
)

var log = logger.NewLogger("habits.model.ScorecardStore")

type ScorecardStore struct {
	db        *sql.DB
	tableName string
}

func GetScorecardStore(conn, table string) (*ScorecardStore, error) {
	log.SetSubLabel("GetScorecardStore")

	db, err := sql.Open("postgres", conn)
	if err != nil || db == nil {
		log.FatalError("unable to open db conn", err)
		return nil, err
	}

	return &ScorecardStore{
		db:        db,
		tableName: pq.QuoteIdentifier(table),
	}, nil
}

func (s ScorecardStore) GetDb() *sql.DB {
	return s.db
}

func (s ScorecardStore) IsLive() bool {
	log.SetSubLabel("Live")

	err := s.db.Ping()

	return err == nil
}

func (s ScorecardStore) Insert(scorecard habits.Scorecard) (*habits.Scorecard, error) {
	log.SetSubLabel("Insert")

	if scorecard.Order == 0 {
		scorecard.Order = s.getLastOrder() + 500
	}

	if scorecard.Time == "" {
		scorecard.Time = "00:00"
	} else if err := validateTime(scorecard.Time); err != nil {
		return nil, err
	}

	var lastInsertedId int
	err := s.db.QueryRow(`
		INSERT INTO
    		`+s.tableName+` (name, connotation, time, sortOrder)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id`,
		scorecard.Name, scorecard.Connotation, scorecard.Time, scorecard.Order).
		Scan(&lastInsertedId)
	if err != nil || lastInsertedId == 0 {
		log.Errorw("error inserting",
			"err", err,
			"last inserted id", lastInsertedId,
		)
		return nil, err
	}

	rows, err := s.db.Query(`
		SELECT
			*
		FROM
			`+s.tableName+`
		WHERE
			id = $1`,
		lastInsertedId)
	if err != nil {
		log.Error("error querying inserted row", err)
		return nil, err
	}

	scorecardList, err := parse(rows)
	if err != nil {
		log.Error("error parsing", err)
		return nil, err
	}

	return &scorecardList[0], nil
}

func (s ScorecardStore) getLastOrder() int {
	var lastId int
	row := s.db.QueryRow(`
		SELECT
			id
		FROM
			` + s.tableName + `
		ORDER BY
			id DESC
		LIMIT 1`)
	err := row.Scan(&lastId)

	if lastId == 0 {
		return 0
	}

	if err != nil {
		log.Error("error getting last id", err)
		return 0
	}

	lastScorecard, err := s.Get(lastId)
	if err != nil {
		log.Errorw("error getting last scorecard",
			"last id", lastId,
			"last scorecard", lastScorecard,
			"err", err,
		)
	}

	return lastScorecard.Order
}

func validateTime(timeStr string) error {
	_, err := time.Parse("15:04", timeStr)
	if err != nil {
		log.Error("invalid time", err)
		return err
	}

	return nil
}

func (s ScorecardStore) List() ([]habits.Scorecard, error) {
	log.SetSubLabel("List")

	rows, err := s.db.Query("SELECT * FROM " + s.tableName + " ORDER BY sortorder")
	if err != nil {
		log.Error("error querying", err)
		return nil, err
	}

	scorecardList, err := parse(rows)
	if err != nil {
		log.Error("error parsing", err)
		return nil, err
	}

	return scorecardList, nil
}

func (s ScorecardStore) Get(id int) (*habits.Scorecard, error) {
	log.SetSubLabel("Get")

	rows, err := s.db.Query("SELECT * FROM "+s.tableName+" WHERE id = $1", id)
	if err != nil {
		log.Error("error querying", err)
		return nil, err
	}

	scorecardList, err := parse(rows)
	if err != nil {
		log.Error("error parsing", err)
		return nil, err
	}

	if len(scorecardList) == 0 {
		return nil, sql.ErrNoRows
	}
	return &scorecardList[0], nil
}

func parse(rows *sql.Rows) ([]habits.Scorecard, error) {
	log.SetSubLabel("parse")
	scorecardList := make([]habits.Scorecard, 0)

	for rows.Next() {
		scorecard := habits.Scorecard{}
		rowTime := sql.NullString{}

		err := rows.Scan(&scorecard.Id, &scorecard.Name, &scorecard.Connotation, &rowTime, &scorecard.Order)
		if err != nil {
			log.Error("rows.Scan error scanning", err)
			continue
		}

		if rowTime.Valid {
			scorecard.Time = formatSqlTimeToDomainTime(rowTime.String)
		}

		scorecardList = append(scorecardList, scorecard)
	}

	return scorecardList, nil
}

func formatSqlTimeToDomainTime(sqlTime string) string {
	log.SetSubLabel("formatSqlTimeToDomainTime")

	if sqlTime == "" {
		return ""
	}

	timestamp, err := time.Parse("2006-01-02T15:04:05Z", sqlTime)
	if err != nil {
		log.Error("error parsing time string from sql", err)
		return ""
	}

	return timestamp.Format("15:04")
}
