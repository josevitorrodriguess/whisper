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
		if app == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "firebase app is not initialized"})
			return
		}

		authClient, err := app.Auth(context.Background())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to initialize firebase auth", "details": err.Error()})
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			return
		}

		idToken := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if idToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization token missing"})
			return
		}

		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "details": err.Error()})
			return
		}

		userRecord, err := authClient.GetUser(context.Background(), token.UID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch firebase user", "details": err.Error()})
			return
		}

		u := &models.User{
			ID:       userRecord.UID,
			Email:    userRecord.Email,
			Username: userRecord.DisplayName,
			PhotoURL: userRecord.PhotoURL,
		}

		c.Set("firebase_user_record", userRecord)
		c.Set("firebase_user", u)

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
