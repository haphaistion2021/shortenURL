package config

import (
	"flag"
	"log"
	"testing"
	"time"

	"github.com/spf13/viper"
)

type ConfigSetting struct {
	configName string
	configType string
	configPath string
}

type Configure struct {
	Server struct {
		Host    *string `yaml:"host"` // the local machine IP Address to bind the HTTP Server to
		Port    *string `yaml:"port"` // the local machine TCP Port to bind the HTTP Server to
		Timeout struct {
			Server *time.Duration `yaml:"server"` // the general server timeout to use for graceful shutdowns
			Read   *time.Duration `yaml:"read"`   // the amount of time to wait until an HTTP server read operation is cancelled
			Write  *time.Duration `yaml:"write"`  // the amount of time to wait until an HTTP server write opperation is cancelled
			Idle   *time.Duration `yaml:"idle"`   // the amount of time to wait until an IDLE HTTP session is closed
		} `yaml:"timeout"`
	}
	Database struct {
		Host     *string `yaml:"host"`
		User     *string `yaml:"user"`
		Name     *string `yaml:"name"`
		Password *string `yaml:"password"`
		Debug    *bool   `yaml:"debug"`
	} `yaml:"database"`
	Redis struct {
		Address *string
	} `yaml:"redis"`
}

var (
	Config  Configure
	setting ConfigSetting
)

func init() {
	// workaround for go test
	var _ = func() bool {
		testing.Init()
		return true
	}()

	flag.StringVar(&setting.configName, "configName", "config", "configuration file name")
	flag.StringVar(&setting.configType, "configType", "yaml", "configuration file type")
	flag.StringVar(&setting.configPath, "configPath", "./config/", "configuration file path")
	flag.Parse()
	log.Printf("setting: %+v\n", setting)

	viper.AutomaticEnv()
	viper.SetConfigName(setting.configName)
	viper.SetConfigType(setting.configType)
	viper.AddConfigPath(setting.configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("err: %s\n", err)
	}

	if err := viper.GetViper().Unmarshal(&Config); err != nil {
		log.Printf("err: %s\n", err)
	}
}
