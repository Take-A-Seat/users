package main

import (
	"github.com/Take-A-Seat/auth/validatorAuth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

var mongoHost = "takeaseat.knilq.mongodb.net"
var mongoUser = "admin"
var mongoPass = "p4r0l4"
var mongoDatabase = "TakeASeat"


func main() {
	port := os.Getenv("TAKEASEAT_USERS_PORT")
	if port == "" {
		port = "9210"
	}

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accepts", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           1 * time.Minute,
		AllowCredentials: true,
	}))

	users := router.Group("/users")
	{
		users.POST("/validateUser", handleValidateUser)
		users.POST("/", handleCreateUser)
	}

	//privateRoutesUsers need Authorization token in header
	protectedUsers := router.Group("/users")
	protectedUsers.Use(validatorAuth.AuthMiddleware("http://54.93.123.171/auth/isAuthenticated"))
	{

	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Port already in use!")
	}

}
