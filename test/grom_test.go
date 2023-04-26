package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(100);" json:"identity"`
	UserName string `gorm:"column:username;type:varchar(100);" json:"user_name"`
	Phone    string `gorm:"column:phone;type:varchar(20);" json:"phone" `
	Password string `gorm:"column:password;type:varchar(100);" json:"_"`
	Email    string `gorm:"column:email;type:varchar(100);"  json:"email"`
}

func (User) TableName() string {
	return "users"
}

//go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql

func TestGorm(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/online_exercise?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}
	data := make([]*User, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, datum := range data {
		fmt.Println(datum)
	}
}
