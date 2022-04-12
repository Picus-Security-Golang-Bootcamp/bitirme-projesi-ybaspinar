package main

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/database"
	logger "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/logging"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("Starting the application...")
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatal("Error loading config file")
	}
	logger.NewLogger(cfg)
	defer logger.Close()

	DB := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

}
