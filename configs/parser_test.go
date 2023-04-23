package configs_test

import (
	"os"
	"testing"

	"github.com/kevindoubleu/pichan/configs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParserSuite struct {
	suite.Suite
}

func TestParserSuite(t *testing.T) {
	suite.Run(t, new(ParserSuite))
}

func (s *ParserSuite) TestNewConfig_FileNotFound_ShouldReturnError() {
	config, err := configs.NewConfig("invalid path to file")

	assert.Nil(s.T(), config)
	assert.EqualError(s.T(), err, "open invalid path to file: no such file or directory")
}

func (s *ParserSuite) TestNewConfig_InvalidYaml_ShouldReturnError() {
	fileWithInvalidYaml := "empty_config_file.yaml"
	filePtr, err := os.Create(fileWithInvalidYaml)
	_, _ = filePtr.WriteString(":")
	assert.NoError(s.T(), err)
	defer os.Remove(fileWithInvalidYaml)

	config, err := configs.NewConfig(fileWithInvalidYaml)

	assert.Nil(s.T(), config)
	assert.EqualError(s.T(), err, "yaml: did not find expected key")
}

func (s *ParserSuite) TestNewConfig_ValidYaml_ShouldParse() {
	config, err := configs.NewConfig(configs.TestConfigFile)

	assert.NoError(s.T(), err)
	shouldParseServerConfigs(s.T(), config)
	shouldParseLogConfigs(s.T(), config)
	shouldParseHabitsConfigs(s.T(), config)
}

func shouldParseServerConfigs(t *testing.T, config *configs.Config) {
	assert.NotNil(t, config.Server)
	assert.NotNil(t, config.Server.Host)
	assert.NotNil(t, config.Server.Port)
}

func shouldParseLogConfigs(t *testing.T, config *configs.Config) {
	assert.NotNil(t, config.Log)
	assert.True(t, config.Log.SkipFatal)
}

func shouldParseHabitsConfigs(t *testing.T, config *configs.Config) {
	assert.NotNil(t, config.Habits)
	assert.NotNil(t, config.Habits.StoreUrl)
	assert.NotNil(t, config.Habits.StoreName)
}
