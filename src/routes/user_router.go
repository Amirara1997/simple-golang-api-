package routes

import (
	"github.com/labstack/echo"
	"goone/src/api/control"
)

func MasterUser(g *echo.Group) {

	DEFINE_URL := "/users"
	g.GET(DEFINE_URL+"/users",control.AllUsers)
	g.GET(DEFINE_URL+"/get_data/", control.GetUserController)
	g.POST(DEFINE_URL+"/add_data/", control.AddUserController)
	g.POST(DEFINE_URL+"/edite_data/:id/",control.EditeDataUsersController)
	g.POST(DEFINE_URL+"/delete_data/:id/",control.DeletDataUsersController)
	g.POST("/login",control.Login)
	g.POST("/logout",control.Logout)

}

