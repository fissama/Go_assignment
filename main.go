package main

import (
	"Go/configs"
	"Go/packages/mongodb"
	"Go/services"
	"context"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)

type (
	HttpResponse struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func InsertCsv(c echo.Context) error {
	//Get config from config_dev.json
	config := configs.GetConfig()

	//Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Open mongodb connection
	mongoClient, err := mongodb.NewMongoDBConnect(ctx, config)
	if err != nil {
		return c.JSON(200, HttpResponse{Code: "503", Message: "Failed to connect CosmosDB. Err:" + err.Error()})
	}

	//Get numbers collections from mongoClient
	numbersCollection := mongoClient.Database(config.DB).Collection(config.Collection)

	// Get csv file
	fileInsert, err := c.FormFile("FileData")
	if err != nil {
		return c.JSON(200, HttpResponse{Code: "417", Message: "Insert failed. Err:" + err.Error()})
	}

	//Check if file is not .csv => return
	nameFile := fileInsert.Filename
	if !strings.HasSuffix(nameFile, ".csv") {
		return c.JSON(200, HttpResponse{Code: "415", Message: "Insert failed. File is not supported"})
	}

	src, err := fileInsert.Open()
	if err != nil {
		return c.JSON(200, HttpResponse{Code: "400", Message: "Insert failed. Err:" + err.Error()})
	}
	defer src.Close()

	err = services.InsertListAssignments(ctx, src, numbersCollection)
	if err != nil {
		return c.JSON(200, HttpResponse{Code: "400", Message: "Insert failed. Err:" + err.Error()})
	}

	return c.JSON(200, HttpResponse{Code: "201", Message: "Insert data into mongodb successfully"})
}

func main() {
	e := echo.New()
	e.POST("/assignment/insert", InsertCsv)
	e.Logger.Fatal(e.Start(":1323"))
}
