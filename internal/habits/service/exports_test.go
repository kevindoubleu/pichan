package service

func NewScorecardsServerWithStore(store Store) ScorecardsServer {
	scorecardsServer := NewScorecardsServer()
	scorecardsServer.store = store
	return scorecardsServer
}

func (s ScorecardsServer) GetStore() Store {
	return s.store
}
