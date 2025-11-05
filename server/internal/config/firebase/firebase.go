package firebase

import (
	"context"
	"errors"
	"os"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var (
	firebaseAuth *auth.Client
	once         sync.Once
	initErr      error
)

func InitFirebase() error {
	once.Do(func() {
		credsPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
		if credsPath == "" {
			initErr = errors.New("FIREBASE_CREDENTIALS_PATH not set")
			return
		}

		opt := option.WithCredentialsFile(credsPath)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			initErr = err
			return
		}

		firebaseAuth, initErr = app.Auth(context.Background())
	})

	return initErr
}


func VerifyToken(idToken string) (*auth.Token, error) {
	if firebaseAuth == nil {
		return nil, errors.New("firebase auth client not initialized")
	}
	return firebaseAuth.VerifyIDToken(context.Background(), idToken)
}


func GetAuthClient() *auth.Client {
	return firebaseAuth
}
