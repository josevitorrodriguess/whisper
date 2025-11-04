package firebase

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var firebaseAuth *auth.Client

func InitFirebase() error {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_PATH"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	firebaseAuth, err = app.Auth(context.Background())
	return err
}

func VerifyToken(idToken string) (*auth.Token, error) {
	return firebaseAuth.VerifyIDToken(context.Background(), idToken)
}
