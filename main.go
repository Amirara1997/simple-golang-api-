package main

import (
	"goone/src/routes"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)

}

func main() {
	e := routes.Index()
	e.Validator = &CustomValidator{validator: validator.New()}

	//logger := customlogger.GetInstance("SYSTEM")
	//logger.Println("Starting Appliction")
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":4400"))
}
