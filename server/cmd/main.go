package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/whisper/server/internal/config/firebase"
	"github.com/josevitorrodriguess/whisper/server/internal/database"
	"github.com/josevitorrodriguess/whisper/server/internal/handler"
	"github.com/josevitorrodriguess/whisper/server/internal/repository"
	"github.com/josevitorrodriguess/whisper/server/internal/router"
	"github.com/josevitorrodriguess/whisper/server/internal/services"
)

func main() {
	godotenv.Load()
	app, err := firebase.GetFireBaseApp()
	if err != nil {
		log.Fatalf("error to initialize firebase: %v", err)
	}

	db := database.ConnectDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.SetupRouter(userHandler, app)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
