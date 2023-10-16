package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	_ "github.com/spf13/viper"
)

const (
	defaultHTTPServerPort           = "8000"
	defaultServerRWTimeout          = 10 * time.Second
	defaultServerMaxHeaderMegabytes = 1
)

type Config struct {
	Server     HTTPConfig
	SMTPConfig SMTPConfig
}

type SMTPConfig struct {
	Username string
	Password string
}

type HTTPConfig struct {
	Addr           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func Init() (*Config, error) {
	err := godotenv.Load("./app.env")
	if err != nil {
		logrus.Fatal(err)
	}
	cfg := Config{}
	cfg.Server = HTTPConfig{
		Addr:           defaultHTTPServerPort,
		ReadTimeout:    defaultServerRWTimeout,
		WriteTimeout:   defaultServerRWTimeout,
		MaxHeaderBytes: defaultServerMaxHeaderMegabytes,
	}
	cfg.SMTPConfig = SMTPConfig{
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}
	return &cfg, nil

	// viper.AddConfigPath(path)

	// viper.SetConfigName("app")
	// viper.SetConfigType("env")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	return nil, err
	// }
	// viper.AutomaticEnv()
}
