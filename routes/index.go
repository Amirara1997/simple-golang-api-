package routes

import (
	"github.com/labstack/echo"
)

type M map[string]interface{}

func Index() *echo.Echo {
	e := echo.New()
	AppIndex(e.Group("/"))
	ApiGroup := e.Group("/api")
	MasterUser(ApiGroup)

	return e
}
