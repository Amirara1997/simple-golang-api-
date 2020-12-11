package control

import (
	"crypto/md5"
	"encoding/hex"
	"gopkg.in/mgo.v2/bson"

	"net/http"

	"goone/src/databaseConnetion"

	"github.com/labstack/echo"
	"goone/src/api/modelha"
)

func AllUsers(c echo.Context) error  {
	var (
		user []modelha.MasterUser
		err 	error
	)
	user, err = modelha.GetUser()
	if err != nil {
		return c.JSON(http.StatusForbidden,err)
	}

	return c.JSON(http.StatusOK,user)
}


func GetUserController(c echo.Context) error {
	db := databaseConnetion.Database().Begin()
	defer db.Close()

	rows, err := db.Raw("SELECT id, username , passwd FROM users").Rows()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	defer rows.Close()

	each := modelha.MasterUser{}
	results := []modelha.MasterUser{}

	for rows.Next() {
		var id, username, passwd []byte
		err := rows.Scan(&id, &username, &passwd)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		var str string = string(id)
		hasher := md5.New()
		hasher.Write([]byte(str))
		coverid := hex.EncodeToString(hasher.Sum(nil))

		each.Id = bson.ObjectId(string(id))
		each.Username = string(username)
		each.Passwd = coverid

		results = append(results, each)

	}

	response := response_json{
		"data":   results,
		"status": status_200,
	}
	return c.JSON(http.StatusOK, response)
}

func AddUserController(c echo.Context) error  {
	db := databaseConnetion.Database().Begin()
	defer db.Close()

	username := c.FormValue("username")
	passwd := c.FormValue("passwd")

	user := modelha.Users{
		Username: username,
		Passwd: passwd,
	}
	db.NewRecord(user)

	if error_insert := db.Create(&user);error_insert.Error != nil {
		return c.JSON(http.StatusBadRequest,echo.Map{"error":error_insert})
	}
	db.NewRecord(user)

	response := response_json{
		"status" : status_200,
	}
	return c.JSON(200,response)
}

func EditeDataUsersController(c echo.Context) error  {
	db := databaseConnetion.Database().Begin()
	defer db.Close()

	id := c.Param("id")
	username := c.FormValue("username")
	passwd := c.FormValue("passwd")

	var user modelha.Users

	data := db.Model(&user).Where("id = ?",id).Updates(map[string]interface{}{
		"usename": username,
		"passwd":passwd,
	})
	if data.Error != nil {
		return c.JSON(http.StatusBadRequest,echo.Map{"error":data.Error})
	} else if data.RowsAffected == 0 {
		not_modified(c)
		return nil
	}

	response := response_json{
		"status": status_200,
	}
	return c.JSON(200,response)

}

func DeletDataUsersController(c echo.Context) error  {
	db := databaseConnetion.Database().Begin()
	defer db.Close()

	id := c.Param("id")
	username := c.FormValue("username")
	passwd := c.FormValue("passwd")

	var user modelha.Users

	data := db.Model(&user).Where("id = ?",id).Delete(map[string]interface{}{
		"usename": username,
		"passwd":passwd,
	})
	if data.Error != nil {
		return c.JSON(http.StatusBadRequest,echo.Map{"error":data.Error})
	} else if data.RowsAffected == 0 {
		not_modified(c)
		return nil
	}

	response := response_json{
		"status": status_200,
	}
	return c.JSON(200,response)

}


