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

// InitDB for database initialization
func InitDB(config *config.Configure) *gorm.DB {
	return initORM(
		*config.Database.Host,
		*config.Database.User,
		*config.Database.Name,
		*config.Database.Password,
		*config.Database.Debug,
	)
}

func initORM(host, user, name, password string, isDebugMode bool) *gorm.DB {
	logLevel := logger.Warn
	hasColor := false
	if isDebugMode {
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
		host,
		user,
		name,
		password)
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
