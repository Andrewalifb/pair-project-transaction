package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/Andrewalifb/pair-project-transaction/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var once sync.Once
	c := cron.New()
	once.Do(func() {
		c.AddFunc("* * * * *", func() {
			fmt.Println("EXECUTED")
		})
		c.Start()
	})
	c.Stop()
	e := echo.New()

	e.POST("/transactions", handler.CreateTransaction)
	e.GET("/transactions", handler.GetAllTransaction)
	e.GET("/transactions/:id", handler.GetTransactionByID)
	e.PUT("/transactions/:id", handler.UpdateDataTransaction)
	e.DELETE("/transactions/:id", handler.DeleteDataTransaction)
	e.Logger.Fatal(e.Start(":8081"))
}
