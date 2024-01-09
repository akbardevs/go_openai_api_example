package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var logger *log.Logger

func main() {
	router := gin.Default()

	// Testing Health API
	router.GET("health", HealthCheckHandler)

	router.POST("message/send", SendMessageHandler)
	router.GET("message/", GetMessagesHandler)

	router.Run(DefaultPort)
}

func init() {
	// Check file env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	logger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
