package main

import (
	"log"

	"github.com/josevitorrodriguess/whisper/server/internal/config/firebase"
	"github.com/josevitorrodriguess/whisper/server/internal/database"
	"github.com/josevitorrodriguess/whisper/server/internal/handler"
	"github.com/josevitorrodriguess/whisper/server/internal/repository"
	"github.com/josevitorrodriguess/whisper/server/internal/router"
	"github.com/josevitorrodriguess/whisper/server/internal/services"
)

func main() {
	app, err := firebase.GetFireBaseApp()
	if err != nil {
		log.Fatalf("error to initialize firebase: %v", err)
	}

	db := database.ConnectDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(*userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.SetupRouter(userHandler, app)

	r.Run(":8080")
}
