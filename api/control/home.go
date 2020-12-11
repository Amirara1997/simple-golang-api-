package control

import (
	"github.com/labstack/echo"
)

//var logs = logger.GetInstance("SYSTEMS -")

type response_json map[string]interface{}
type DB struct {
	Value       interface{}
	Error       error
	RowAffected int64
}

func AppIndex(c echo.Context) error {
	return c.JSON(200, "High performance , minimalist Go web Framework running")
}
