package modelha

import (
	"api/src/api/security"
	"goone/src/databaseConnetion"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Users struct {
	Id       int    `gorm:"AUTO_INCREMENT"`
	Username string `gorm:"username"`
	Passwd   string `gorm:"passwd"`
	CreateAt time.Time `gorm:"default:current_timestamp()" json:"create_at"`
	UpdateAt time.Time  `gorm:"default:current_timestamp()" json:"update_at"`
}

func (Users) TableName() string {
	return "users"
}

type MasterUser struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty" gorm:"type:varchar(40);unique"`
	Username string `json:"username" bson"username" gorm:"type:varchar(40);unique"`
	Passwd   string `json:"passwd" bson:"passwd"`
}

type UserModel struct {
	UserModel []MasterUser `json:"users"`
}

func (u *Users) BeforeSave() error  {
	hashedPasswd , err := security.Hash(u.Passwd)
	if err != nil{
		return err
	}
	u.Passwd = string(hashedPasswd)
	return nil
}


func GetUserByUsername (username string) (MasterUser,error) {
	var (
		user MasterUser
		err error
	)
	db := databaseConnetion.Database().Begin()
	if err = db.Find(&user,"username = ?",username ).Error;err != nil {
		return user,err
	}
	db.Commit()
	return user,err
}

func GetUser() ([]MasterUser, error) {
	var (
		user []MasterUser
		err	error

	)
	tx := databaseConnetion.Database().Begin()
	if err = tx.Find(&user).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, err
}
