package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	Application struct {
		BasePath string
		BaseUrl  string
	}

	Logging struct {
		Level   string
		LogPath string
	}

	Server struct {
		Port            int
		BasePath        string
		CatchAllEnabled bool
		LogRequest      bool
	}

	Database struct {
		Host       string
		Port       int
		Schema     string
		User       string
		Password   string
		LogQueries bool
	}

	Security struct {
		Jwt struct {
			Secret string
		}
	}

	Mail struct {
		Smtp struct {
			Host     string
			Port     string
			User     string
			Password string
			Sender   string
		}
	}
}

var C config

func main() {
	Config := &C
	currentDir, _ := os.Getwd()

	//some default values
	viper.SetDefault("Application.BasePath", currentDir)
	viper.SetDefault("Logging.logPath", "log")

	viper.SetDefault("Server.Port", 8080)
	viper.SetDefault("Server.BasePath", "/api")
	viper.SetDefault("Server.CatchAllEnabled", false)
	viper.SetDefault("Server.LogRequest", false)

	viper.SetDefault("Database.Host", "localhost")
	viper.SetDefault("Database.Port", 3306)
	viper.SetDefault("Database.LogQueries", false)

	viper.SetDefault("Mail.Smtp.Port", 587)

	viper.SetConfigName("application-config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(C, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
