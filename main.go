package main

import (
	"log"

	"github.com/rayspock/go-answer/config"
	"github.com/rayspock/go-answer/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	log.Println("Load .env")
	config.LoadENV()

	log.Println("Connect to Database...")
	config.DB, err = gorm.Open("postgres", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		log.Println("Status:", err)
	}
	defer config.DB.Close()

	log.Println("Spin up Server...")
	r := routes.SetupRouter()
	r.Run()
}
