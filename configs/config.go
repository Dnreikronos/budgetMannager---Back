package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func init() {
	viper.SetDefault("api_port", "9000")
	viper.SetDefault("database.host", "${DB_HOST}")
	viper.SetDefault("database.port", "${DB_PORT}")
	viper.SetDefault("database.user", "${DB_USER}")
	viper.SetDefault("database.pass", "${DB_PASSWORD}")
}

func Load() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environments variables")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	err := viper.ReadConfig()
	if err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("darabase.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	}

	log.Printf("Database config: %+vZ\n", cfg.DB)
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func getAPI() APIConfig {
	return cfg.API
}
