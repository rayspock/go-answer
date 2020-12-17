package main

import (
	"log"

	"github.com/rayspock/go-answer/config"
	"github.com/rayspock/go-answer/models"
	"github.com/rayspock/go-answer/routes"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

func main() {
	log.Println("Load .env")
	config.LoadENV()

	log.Println("Connect to database...")
	config.DB, err = gorm.Open("postgres", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		log.Println("Status:", err)
	}
	defer config.DB.Close()

	log.Println("Initialize database...")
	models.Init(config.DB)

	log.Println("Spin up Server...")
	r := routes.SetupRouter()
	r.Run()
}
