package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	fmt.Println("----------------- INIT DB CONNECTION ----------------")
	dsn := "root:@tcp(localhost:3306)/project_management?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connection databases")
	}

	DB = db
	return db
}

// func ConnectDB() (*gorm.DB, error) {
// 	// fmt.Println("----------------- INIT DB CONNECTION ----------------")
// 	dsn := "root:@tcp(localhost:3306)/project_management?charset=utf8&parseTime=True"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println("Error connection databases")
// 	}

// 	DB = db
// }

func CloseConnection(db *gorm.DB) {
	dbConn, _ := DB.DB()
	dbConn.Close()
	fmt.Println("close connection")
}
