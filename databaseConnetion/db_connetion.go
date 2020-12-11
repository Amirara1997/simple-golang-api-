package databaseConnetion

import (

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var logs = logger.GetInstance("SYSTEMS -")

var (

	db  *gorm.DB
	err error
)


const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = ""
	DB_USER = "root"
	DB_PASS = ""
)

func CreateCon(){
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", DB_USER, DB_PASS, DB_HOST, DB_NAME))
	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
}

func Database() *gorm.DB  {
	return db
}
