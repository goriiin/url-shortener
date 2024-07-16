package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

// TODO :сделать настройки через env

type Server struct {
	Env        string `yaml:"env" env-required:"true"`
	HTTPServer `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost"`
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Server {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configPath := os.Getenv("SERVER_CONFIG_PATH")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Server
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("config file %s read error: %s", configPath, err)
	}
	return &cfg
}
