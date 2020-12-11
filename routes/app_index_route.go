package routes

import (

	"goone/src/api/control"
	"github.com/labstack/echo"
)

func AppIndex(g *echo.Group) {
	g.GET("", control.AppIndex)
}
