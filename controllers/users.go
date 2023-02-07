package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type User struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	Name string
}



func GetUsers(c *gin.Context) {
	dsn := "host=localhost user=postgres password=password dbname=testgolang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var user []User

	if err != nil {
		fmt.Print("Error")
	} else {
		db.Raw("Select * From users").Scan(&user)

		c.JSON(200, gin.H{
			"status": 200,
			"data": user,
		})
	}

}



type FormCreateUser struct{
	Name string
}
func CreateUser(c *gin.Context){
	dsn := "host=localhost user=postgres password=password dbname=testgolang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	var request FormCreateUser

	c.BindJSON(&request)

	var user  = User{
		Name: request.Name,
	}

	if err != nil {
		fmt.Print("Error")
	} else {
		db.Create(&user)
		c.JSON(200, gin.H{
			"status": 200,
			"message": "success create user",
			"data": user,

		})
	}
}



type requestById struct{
	Id string `uri:"id" binding:"required"`
}

func DeleteUser(c *gin.Context){
	dsn := "host=localhost user=postgres password=password dbname=testgolang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	var request requestById

	c.BindUri(&request)

	if err != nil {
		fmt.Print("Error")
	} else {
		db.Where("id = ?", request.Id).Delete(&User{})
		c.JSON(200, gin.H{
			"status": 200,
			"message": "success delete user",
			"request": request,
		})
	}

}


func UpdateUser(c *gin.Context){
	dsn := "host=localhost user=postgres password=password dbname=testgolang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	var request requestById
	var form FormCreateUser

	c.BindUri(&request)
	c.BindJSON(&form)

	if err != nil {
		fmt.Print("Error")
	} else {
		db.Model(User{}).Where("id = ?", request.Id).Update("name", form.Name)
		c.JSON(200, gin.H{
			"status": 200,
			"message": "success update user",
			"request": request,
			"Test" : form.Name,
		})
	}
}