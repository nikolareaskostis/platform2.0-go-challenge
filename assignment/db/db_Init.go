package db

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	ErrSQLConnection = errors.New("sql connection timeout")
)

var dbConnection *gorm.DB
var connectionString string

func Init(connectionStr string) {
	connectionString = connectionStr
}

func GetGormConnection(connectionString string) (*gorm.DB, error) {

	if dbConnection == nil {
		connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  true,
				},
			),
		})

		if err != nil {
			return nil, ErrSQLConnection
		}

		dbConnection = connection

		Initialize(dbConnection)
	}

	return dbConnection, nil
}
