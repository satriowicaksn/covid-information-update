package controllers

import (
	"covid-information-update/database"
	"covid-information-update/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetData(c echo.Context) error {
	result, err := database.FetchDataFromApi()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	database.PostData()
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Success fetch covid information data from API",
		Data:    result,
	})
}

func GetDataLogController(c echo.Context) error {
	data, err := database.GetDataLog()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Success fetch covid information data from database",
		Data:    data,
	})
}
