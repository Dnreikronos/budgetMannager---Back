package configs

import (
	"log"
	"os"
	"strings"

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
	viper.SetDefault("api.port", "9090")
	viper.SetDefault("database.host", "${DB_HOST}")
	viper.SetDefault("database.port", "${DB_PORT}")
	viper.SetDefault("database.user", "${DB_USER}")
	viper.SetDefault("database.pass", "${POSTGRES_PASSWORD}")
	viper.SetDefault("database.database", "${DB_NAME}") 
}

func Load() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs") 	
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = &config{
		API: APIConfig{
			Port: viper.GetString("api.port"),
		},
		DB: DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.pass"), 
			Database: viper.GetString("database.database"),
		},
	}

	log.Printf("Database config: %+v", cfg.DB)
	return nil
}

func replaceEnvVariables() {
	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			envKey := value[2 : len(value)-1]
			envValue := os.Getenv(envKey)
			if envValue != "" {
				viper.Set(key, envValue)
			}
		}
	}
}

func GetDB() DBConfig {
	if cfg == nil {
		log.Fatal("Configuration is not initialized")
	}
	return cfg.DB
}

func GetServerPort() string {
	if cfg == nil {
		log.Fatal("Configuration is not initialized")
	}
	return cfg.API.Port
}
