package configs

const (
	ConfigFile     = "config.yaml"
	TestConfigFile = "test_config.yaml"
)

type Config struct {
	Server `yaml:"server"`
	Log    `yaml:"log"`
	Habits `yaml:"habits"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Log struct {
	SkipFatal bool `yaml:"skip_fatal"`
}

type Habits struct {
	StoreUrl  string `yaml:"store_url"`
	StoreName string `yaml:"store_name"`
}
