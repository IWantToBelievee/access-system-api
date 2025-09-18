package cfg

import (
	"os"

	"github.com/joho/godotenv"
)

// DbCfg holds the database configuration parameters.
type DbCfg struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// LoadDbCfg loads database configuration from environment variables.
func LoadDbCfg() (*DbCfg, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &DbCfg{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}, nil
}

// LoadTestDbCfg loads test database configuration from environment variables.
func LoadTestDbCfg() (*DbCfg, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, err
	}

	return &DbCfg{
		Host:     os.Getenv("POSTGRES_TEST_HOST"),
		Port:     os.Getenv("POSTGRES_TEST_PORT"),
		User:     os.Getenv("POSTGRES_TEST_USER"),
		Password: os.Getenv("POSTGRES_TEST_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_TEST_DB"),
	}, nil
}
