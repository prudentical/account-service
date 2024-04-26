package configuration

import (
	"path"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

const (
	Prod string = "prod"
	Dev  string = "dev"
	Test string = "test"
)

const (
	Debug string = "debug"
	Info  string = "info"
	Warn  string = "warn"
	Error string = "error"
)

type Config struct {
	App struct {
		Name  string
		Env   string
		Debug bool
	}
	Server struct {
		Host        string
		Healthcheck string
		Port        int
	}
	Database struct {
		Name       string
		Host       string
		Port       int
		SSL        string
		User       string
		Password   string
		Timezone   string
		Connection struct {
			Idle int
			Open int
		}
	}
	Discovery struct {
		Server struct {
			Host string
			Port int
		}
	}
	Logging struct {
		Level string
	}
}

var config *Config

func setup() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(dir + "/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func NewConfig() Config {
	if config == nil {
		setup()
	}
	return *config
}
