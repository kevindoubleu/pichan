package service

import "github.com/kevindoubleu/pichan/configs"

var config, _ = configs.NewConfig(configs.TestConfigFile)

func NewScorecardsServerWithStore(store Store) ScorecardsServer {
	scorecardsServer := NewScorecardsServer(config.Habits)
	scorecardsServer.store = store
	return scorecardsServer
}

func (s ScorecardsServer) GetStore() Store {
	return s.store
}
