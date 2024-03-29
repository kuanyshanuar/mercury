package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewConnection creates a new database connection.
func NewConnection(dsn string) (*gorm.DB, error) {

	if len(dsn) == 0 {
		return nil, fmt.Errorf("dsn required")
	}

	gormDB, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to setup sql database: %s", err.Error())
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
