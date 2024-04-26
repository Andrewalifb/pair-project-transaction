package main

import (
	"log"

	"github.com/Andrewalifb/pair-project-transaction/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.POST("/transactions", handler.CreateTransaction)
	e.GET("/transactions", handler.GetAllTransaction)
	e.Logger.Fatal(e.Start(":8081"))
}