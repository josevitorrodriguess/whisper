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

	db := database.ConnectDatabase()

	if err := firebase.InitFirebase(); err != nil {
		log.Fatalf("failed to initialize Firebase: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, firebase.GetAuthClient())

	r := router.SetupRouter(userHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
