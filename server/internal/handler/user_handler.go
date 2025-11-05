package handler

import (
	"context"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/whisper/server/internal/services"
)

type UserHandler struct {
	service *services.UserService
	fbAuth  *auth.Client
}

func NewUserHandler(service *services.UserService, fbAuth *auth.Client) *UserHandler {
	return &UserHandler{service: service, fbAuth: fbAuth}
}

func (h *UserHandler) SignIn(c *gin.Context) {
	uidValue, ok := c.Get("uid")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "uid not found in context"})
		return
	}

	uid, _ := uidValue.(string)

	userRecord, err := h.fbAuth.GetUser(context.Background(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to get user from Firebase",
			"details": err.Error(),
		})
		return
	}

	err = h.service.EnsureUserExists(
		c.Request.Context(),
		userRecord.UID,
		userRecord.Email,
		userRecord.DisplayName,
		userRecord.PhotoURL,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to ensure user",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user signed in successfully",
	})
}
