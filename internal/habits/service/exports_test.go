package service

import "github.com/kevindoubleu/pichan/configs"

func NewScorecardsServerWithStore(config configs.Config, store Store) ScorecardsServer {
	scorecardsServer := NewScorecardsServer(config)
	scorecardsServer.store = store
	return scorecardsServer
}

func (s ScorecardsServer) GetStore() Store {
	return s.store
}
