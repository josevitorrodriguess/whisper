package router

import (
	"net/http"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/whisper/server/internal/handler"
	"github.com/josevitorrodriguess/whisper/server/internal/middlewares"
)

func SetupRouter(userHandler *handler.UserHandler, firebaseApp *firebase.App) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "API is running 🚀"})
	})

	user := r.Group("/user")
	{
		user.Use(middlewares.FirebaseAuthMiddleware(firebaseApp))

		user.POST("/register", userHandler.RegisterUserHandler)
		user.DELETE("/delete", userHandler.DeleteAccountHandler)
	}

	return r
}
