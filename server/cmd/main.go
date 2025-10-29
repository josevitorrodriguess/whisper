package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/whisper/server/internal/config/firebase"
	"github.com/josevitorrodriguess/whisper/server/internal/database"
)

func main() {
	godotenv.Load()
	database.ConnectDatabase()

	app, err := firebase.GetFireBaseApp()
	if err != nil {
		panic(err)
	}

	fmt.Println("Firebase App initialized:", app)
}
