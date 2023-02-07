package main

import (
	"zuzustack/learn/api/controllers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main(){
	dsn := "host=localhost user=postgres password=password dbname=testgolang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err == nil {
		db.AutoMigrate(&controllers.User{})
	}
}