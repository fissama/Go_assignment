package main

import (
	"Go/services"
	"github.com/labstack/echo/v4"
)

type (
	HttpResponse struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func insertCsv(c echo.Context) error {
	// Get csv file
	fileInsert, err := c.FormFile("FileData")
	if err != nil {
		return c.JSON(200, HttpResponse{Code: "415", Message: "Insert failed. Err:" + err.Error()})
	}

	src, err := fileInsert.Open()
	if err != nil {
		return c.JSON(200, HttpResponse{Code: "415", Message: "Insert failed. Err:" + err.Error()})
	}
	defer src.Close()

	err = services.InsertListAssignments(src)
	if err != nil {
		return c.JSON(200, HttpResponse{Code: "400", Message: "Insert failed. Err:" + err.Error()})
	}

	return c.JSON(200, HttpResponse{Code: "201", Message: "Insert data into mongodb successfully"})
}

func main() {
	e := echo.New()
	e.POST("/assignment/insert", insertCsv)
	e.Logger.Fatal(e.Start(":1323"))
}
