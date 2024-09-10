package main

import (
	"fmt"

	"github.com/aayushkhosla/Mini-Social-Network/controllers"
	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/middlewares"
	_ "github.com/aayushkhosla/Mini-Social-Network/migrations"
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose"
)

func init() {
	err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connetion to database .. ")
	}


	if err := goose.SetDialect("postgres"); err != nil {
        panic(err)
    }
	// goose.AddMigration(migrations.UpAddressDetails, migrations.DownAddressDetails)

    if err := goose.Up(database.SQL_DB, "migrations"); err != nil {
        panic(err)
    }
}

func main() {
	r := gin.Default()

	
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello to go lang",
	// 	})
	// })
	
	 

	r.POST("/auth/signup", controllers.SignUp)
	r.POST("/auth/login", controllers.Login)  
	r.GET("/user/personaldetails" , middlewares.CheckAuth , controllers.Getuser )
	r.GET("/user/getalluser" , middlewares.CheckAuth , controllers.Userlist )
	r.GET("/user/follows/:id" , middlewares.CheckAuth , controllers.Follow )
	
	fmt.Println("Hello")
	r.Run()

}