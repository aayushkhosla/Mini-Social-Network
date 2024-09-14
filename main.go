package main

import (
	"fmt"
	"log"
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
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}
	
	err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connetion to database .. ")
	}
	
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

		//Grouping routes
	auth := r.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/login", controllers.Login)
		
	}
	user := r.Group("/user")
	{
		user.GET("/getalluser" , middlewares.CheckAuth , controllers.Userlist )
		user.GET("/getuser" , middlewares.CheckAuth , controllers.Getuser )
		user.GET("/follows/:id" , middlewares.CheckAuth , controllers.Follow )
		user.GET("/unfollows/:id" , middlewares.CheckAuth , controllers.Unfollow)
		user.POST("/updatePassword" , middlewares.CheckAuth , controllers.UpdatePassword)
		user.DELETE("/deleteUser" , middlewares.CheckAuth , controllers.Deleteuser)
		user.GET("/getfollowinglist" , middlewares.CheckAuth ,controllers.FollowingList)
		user.GET("/getfollowslist" , middlewares.CheckAuth ,controllers.FollowersList)
		user.PATCH("/updateUser",middlewares.CheckAuth,controllers.UpdateUser)
	}
	r.Run(":8089")

}