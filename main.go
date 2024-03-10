package main

import (
	"ToLet/controller"
	"ToLet/database"
	"ToLet/model"
	"ToLet/util"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DBConnect() {
	database.ConnectDB()
	database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.Room{})
	database.DB.AutoMigrate(&model.Rental{})
	database.DB.AutoMigrate(&model.Payment{})
}

func main() {
	router := gin.Default()

	loadEnv()
	DBConnect()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Adjust with your React app's address
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	authRouter := router.Group("/auth")
	authRouter.POST("/register", controller.Register)
	authRouter.POST("/login", controller.Login)

	ownerRouter := router.Group("/owner")
	ownerRouter.Use(util.JWTAuth())
	ownerRouter.GET("/users", controller.GetUsers)

	adminRouter := router.Group("/admin")
	adminRouter.Use(util.JWTAuthAdmin())
	adminRouter.POST("/register-room", controller.CreateRoom)

	openRouter := router.Group("/api")
	openRouter.Use(util.JWTAuthMiddleware())
	openRouter.GET("/user/:id", controller.GetUserById)
	// openRouter.PUT("/update-user/:id", controller.UpdateUser)
	openRouter.PATCH("/update-user/:id", controller.UpdateTheUser)
	openRouter.GET("/rooms", controller.GetAllRooms)
	openRouter.GET("/room/:id", controller.GetRoom)

	router.Run("localhost:8080")
}
