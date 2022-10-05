package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewDBConnection() DBConnection {
	return DBConnection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}

func (c DBConnection) ToConnectionString() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		c.Host,
		c.Username,
		c.Password,
		c.Database,
		c.Port,
	)
}

func OpenConnection() *gorm.DB {
	dsn := NewDBConnection().ToConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func CloseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}
