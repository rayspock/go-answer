package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

// DB represents gorm.DB pointer
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig is for setting up db connection info
func BuildDBConfig() *DBConfig {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = 5432
		log.Printf("%s, use default port:%v instead \n", err, port)
	}
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

// DbURL retrieve db connection URL
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s port=%v user=%s dbname=%s sslmode=disable password=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Password,
	)
}
