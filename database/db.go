package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	// source/file import is required for migration files to read
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	// load pq as database driver
	_ "github.com/lib/pq"
)

var (
	LearnDocker *sqlx.DB
)

type SSLMode string

const (
	SSLModeEnable  SSLMode = "enable"
	SSLModeDisable SSLMode = "disable"
)

// ConnectAndMigrate function connects with a given database and returns error if there is any error
func ConnectAndMigrate(host, port, databaseName, user, password string, sslMode SSLMode) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, databaseName, sslMode)
	DB, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}
	LearnDocker = DB
	return nil
}

func ShutdownDatabase() error {
	return LearnDocker.Close()
}

func LoadEnv() {
	// loading env variables

	if err := godotenv.Load(); err != nil {
		err = godotenv.Load("./config/env.txt")
		if err != nil {
			logrus.Printf("Error getting env, %v", err)
		}
	} else {
		logrus.Println("We are getting the env values")
	}
}
