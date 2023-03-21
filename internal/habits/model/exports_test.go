package model

import "database/sql"

// exported only when running tests

func (s ScorecardStore) GetDb() *sql.DB {
	return s.db
}

var GetLastOrder = ScorecardStore.getLastOrder
