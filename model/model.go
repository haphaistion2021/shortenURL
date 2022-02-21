package model

import (
	"fmt"
	"log"
	"os"
	"short-url/config"
	"short-url/domain"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB

// InitDB for database initialization
func InitDB(configure *config.Configure) {
	dbInstance = initORM(configure)
}

func initORM(configure *config.Configure) *gorm.DB {
	setting := configure.Database

	logLevel := logger.Warn
	hasColor := false
	if *setting.Debug {
		logLevel = logger.Info
		hasColor = true
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // IO writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshhold
			LogLevel:      logLevel,
			Colorful:      hasColor,
		},
	)
	// Build connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		*setting.Host,
		*setting.User,
		*setting.Name,
		*setting.Password)
	conn, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Panic(err)
	}
	err = migrateDB(conn)
	if err != nil {
		log.Panic(err)
	}

	return conn
}

// MigrateDB for Migrate the schema
func migrateDB(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(domain.Paste{})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return dbInstance
}
