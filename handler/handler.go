package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/Andrewalifb/pair-project-transaction/config"
	"github.com/Andrewalifb/pair-project-transaction/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTransaction(c echo.Context) error {
	collectionName := os.Getenv("TRANSACTION_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	e := new(model.Transaction)
	err = c.Bind(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := collection.InsertOne(context.Background(), e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func GetAllTransaction(c echo.Context) error {
	collectionName := os.Getenv("TRANSACTION_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	var datas []model.Transaction


	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}


	for cursor.Next(context.Background()) {
		var data model.Transaction
		if err := cursor.Decode(&data); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}


		datas = append(datas, data)
	}

	if err := cursor.Err(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, datas)
}