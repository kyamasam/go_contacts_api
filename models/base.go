package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.DB
func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	//dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username,password, dbName ) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

func GetDB() *gorm.DB {
	return db
}
//	dbUri := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username,password, dbName ) //Build connection string