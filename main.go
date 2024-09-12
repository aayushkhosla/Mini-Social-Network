package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aayushkhosla/Mini-Social-Network/controllers"
	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/middlewares"
	_ "github.com/aayushkhosla/Mini-Social-Network/migrations"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"

)

func init() {
	err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connetion to database .. ")
	}
	godotenv.Load()
	// if err := goose.SetDialect("postgres"); err != nil {
    //     panic(err)
    // }
    // if err := goose.Up(database.SQL_DB, "migrations"); err != nil {
    //     panic(err)
    // }
}

func main() {

	r := gin.Default()
	config := cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: false,
		ValidateHeaders: false,
	}
	r.Use(cors.Middleware(config))	
	
	//health check endpoint 
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.POST("/auth/signup", controllers.SignUp)
	r.POST("/auth/login", controllers.Login)  



	r.GET("/user/personaldetails" , middlewares.CheckAuth , controllers.Getuser )
	r.GET("/user/getalluser" , middlewares.CheckAuth , controllers.Userlist )
	r.GET("/user/follows/:id" , middlewares.CheckAuth , controllers.Follow )
	r.GET("/user/unfollows/:id" , middlewares.CheckAuth , controllers.Unfollow)
	r.POST("/user/UpdatePassword" , middlewares.CheckAuth , controllers.UpdatePassword)
	r.DELETE("user/DeleteUser" , middlewares.CheckAuth , controllers.Deleteuser)
	


	r.Run(":8089")

}