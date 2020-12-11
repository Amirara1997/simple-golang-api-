package cookieha

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"

	"time"
)

func setSession(c echo.Context)   {
	cookie := &http.Cookie{
		Name: "session",
		Value: "some_string",
		Path: "/",
		Expires: time.Now().Add(48 *time.Hour),
	}
	c.SetCookie(cookie)
}


func clearSession(c echo.Context){
	cookie := &http.Cookie{
		Name: "Authorization",
		Value: "",
		Path: "/",
		MaxAge: -1,
		Expires: time.Now().Add(24 *time.Hour),

	}
	c.SetCookie(cookie)

}



func readCookie(c echo.Context)error  {
	cookie , err := c.Cookie("username")
	if err != nil {
		return c.JSON(http.StatusBadRequest,echo.Map{"error :":err})
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)

	return c.String(http.StatusOK,"read a cookie")
}

func readAllCookie(c echo.Context) error  {
	for _,cookie := range c.Cookies(){
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}

	return c.String(http.StatusOK,"reade All cookie")
}
