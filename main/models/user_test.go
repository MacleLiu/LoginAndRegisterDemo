package models

import (
	"LoginAndRegisterDemo/main/utils"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

// DB test
var mysql utils.MysqlServer
var db *gorm.DB

func init() {
	mysql = utils.MysqlServer{UserName: "root", Password: "123456", Protocol: "tcp", Address: "localhost:3306", DBName: "gin"}
	var err error
	db, err = mysql.Open()
	if err != nil {
		fmt.Println("Open Database error")
	}
}

func TestGetUser(t *testing.T) {
	user, exists, err := GetUser(db, "asdaf")
	if err != nil {
		t.Errorf("GetUser error: %s", err.Error())
	}
	if exists {
		fmt.Printf("Get it: %v", user.UserName)
	} else {
		fmt.Println("Not Exists")
	}
}

func TestAddUser(t *testing.T) {
	user := User{UserName: "赵六", Password: "qqqq111d", Salt: "43s1"}
	err := AddUser(db, &user)
	if err != nil {
		t.Error("AddUser error: " + err.Error())
	}
}

func TestUpdatePasswd(t *testing.T) {
	user := User{UserName: "王五", Password: "newpassword", Salt: "newsalt"}
	err := UpdatePasswd(db, &user)
	if err != nil {
		t.Error("AddUser error: " + err.Error())
	}
}
