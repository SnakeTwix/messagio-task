package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"messagio/core/enums/env"
	"messagio/utils"
	"os"
	"time"
)

type DSNConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	timezone string
}

func (d *DSNConfig) String() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s sslmode=disable", d.host, d.user, d.password, d.dbname, d.port, d.timezone)
}

func initDSN() *DSNConfig {
	host := utils.GetEnv(env.PostgresHost)
	user := utils.GetEnv(env.PostgresUser)
	password := utils.GetEnv(env.PostgresPassword)
	dbname := utils.GetEnv(env.PostgresDb)
	port := utils.GetEnv(env.PostgresPort)
	timezone := utils.GetEnv(env.PostgresTimezone)

	return &DSNConfig{
		host:     host,
		user:     user,
		password: password,
		dbname:   dbname,
		port:     port,
		timezone: timezone,
	}
}

func InitDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Enable Color
		},
	)

	dsn := initDSN().String()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	return db
}
