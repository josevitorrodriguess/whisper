package middlewares

import (
	"context"
	"net/http"
	"strings"

	fb "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/whisper/server/internal/config/firebase"
	"github.com/josevitorrodriguess/whisper/server/internal/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		token := strings.Replace(authHeader, "Bearer ", "", 1)

		decodedToken, err := firebase.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("uid", decodedToken.UID)

		c.Next()
	}
}

func ExtractUserInfos(app *fb.App, tokenID string) (models.User, error) {
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
