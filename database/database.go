package database

import (
	"fmt"
	"log"

	// "gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// host := "localhost"
	// port := "5432"
	// database := "tolet"
	// user_name := "subhojyoti"
	// password := "Singha@12"

	// db, err := gorm.Open(postgres.Open("postgres://" + user_name + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"))
	dsn := "root:Singha@12@tcp(127.0.0.1:3306)/tolet?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err, " Database Connection Failed")
		log.Fatal("connection error: ", err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
	DB = db
}
