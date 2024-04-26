package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/Andrewalifb/pair-project-transaction/config"
	"github.com/Andrewalifb/pair-project-transaction/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ANDREW
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

// Yohanes
// GET transcation by ID
func GetTransactionByID(c echo.Context) error {
	collectionName := os.Getenv("TRANSACTION_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var data model.Transaction

	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

// UPDATE
func UpdateDataTransaction(c echo.Context) error {
	collectionName := os.Getenv("TRANSACTION_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p := new(model.Transaction)
	err = c.Bind(p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Proses Update Data
	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": p},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

// stephen
func DeleteDataTransaction(c echo.Context) error {
	collectionName := os.Getenv("TRANSACTION_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.DeletedCount == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Data Not Found")
	}

	return c.JSON(http.StatusCreated, result)
}
