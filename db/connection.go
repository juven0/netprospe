package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type SQLServerConfig struct {
	userName string
	password string
	host     string
	dbName   string
}

func LoadConfig() *SQLServerConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return &SQLServerConfig{
		userName: os.Getenv("DBUSERNAME"),
		password: os.Getenv("PASSWORD"),
		host:     os.Getenv("HOST"),
		dbName:   os.Getenv("DBNAME"),
	}
}

func Connection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	connConfig := LoadConfig()

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", connConfig.userName, connConfig.password, connConfig.host, connConfig.dbName)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, err
	}
	fmt.Println("Connected to SQLServer !")
	return db, nil
}
