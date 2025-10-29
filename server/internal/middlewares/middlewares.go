package middlewares

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/whisper/server/internal/models"
)

func FirebaseAuthMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authClient, err := app.Auth(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Firebase Auth"})
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		c.Set("firebaseUID", token.UID)
		c.Next()
	}
}

func ExtractUserInfos(app *firebase.App, tokenID string) (models.User, error) {
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return models.User{}, err
	}

	token, err := authClient.VerifyIDToken(context.Background(), tokenID)
	if err != nil {
		return models.User{}, err
	}

	userRecord, err := authClient.GetUser(context.Background(), token.UID)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		ID:       userRecord.UID,
		Email:    userRecord.Email,
		Username: userRecord.DisplayName,
	}

	return user, nil
}
