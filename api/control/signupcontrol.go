package control

import (
	"github.com/labstack/echo"
	"goone/src/api/modelha"
	"goone/src/databaseConnetion"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Handler struct {
	db *mgo.Session
}

func (h *Handler) Signup(c echo.Context) (err error) {
	u := &modelha.MasterUser{Id:bson.NewObjectId()}
	if err = c.Bind(u);err != nil {
		return
	}
	if u.Username =="" || u.Passwd ==""{
		return c.JSON(http.StatusBadRequest,echo.Map{"message:":"invalid usename or password"})
	}
	db := databaseConnetion.Database().Begin()
	db.Create(u)
	if err !=nil {
		return c.JSON(http.StatusBadRequest,echo.Map{"error":err})
	}

	return c.JSON(http.StatusOK,u)
}

