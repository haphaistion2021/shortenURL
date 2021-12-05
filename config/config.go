package config

import (
	"flag"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func init() {
	// workaround for go test
	var _ = func() bool {
		testing.Init()
		return true
	}()

	viper.AutomaticEnv()
	setFlag()

	initDefault()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("err:%s\n", err)
	}

	// Unmarshal
	parseYaml(viper.GetViper())
}

type configSetting struct {
	fileName string
	filePath string
	fileType string
}

// Configure struct for webapp config
type Configure struct {
	Server struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host *string

		// Port is the local machine TCP Port to bind the HTTP Server to
		Port    *string
		Timeout struct {
			// Server is the general server timeout to use
			// for graceful shutdowns
			Server *time.Duration

			// Write is the amount of time to wait until an HTTP server
			// write opperation is cancelled
			Write *time.Duration

			// Read is the amount of time to wait until an HTTP server
			// read operation is cancelled
			Read *time.Duration

			// Read is the amount of time to wait
			// until an IDLE HTTP session is closed
			Idle *time.Duration
		}
	}
	Database struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host     *string
		User     *string
		Name     *string
		Password *string
		Debug    *bool
	}
	Redis struct {
		Address *string
	}
}

var (
	// Config for config setting
	Config  Configure
	setting configSetting
)

func parseYaml(v *viper.Viper) {
	if err := v.Unmarshal(&Config); err != nil {
		log.Printf("err:%s", err)
	}
	log.Println("config:\n ", Config)
}

func initDefault() {
	// set config file name
	viper.SetConfigName(setting.fileName)
	// add config file path
	viper.AddConfigPath(setting.filePath)
	viper.AddConfigPath("./")
	viper.AddConfigPath("../config")
	// set config file type
	viper.SetConfigType(setting.fileType)

	fmt.Printf("setting: %+v\n", setting)
}

func setFlag() {
	flag.StringVar(&setting.fileName, "configName", "config", "Configuration file name.")
	flag.StringVar(&setting.filePath, "configPath", "./config/", "Configuration file path.")
	flag.StringVar(&setting.fileType, "configType", "yaml", "Configuration file type.")

	flag.Parse()
}
