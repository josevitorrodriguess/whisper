package handler

import (
	"net/http"
	"time"

	auth "firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/whisper/server/internal/models"
	"github.com/josevitorrodriguess/whisper/server/internal/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

type RegisterRequest struct {
	Name string `json:"name"`
}

func (h *UserHandler) RegisterUserHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	v, ok := c.Get("firebase_user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "firebase user not found in context"})
		return
	}
	fbUser := v.(*auth.UserRecord)

	user := &models.User{
		ID:        fbUser.UID,
		Username:  req.Name,
		Email:     fbUser.Email,
		PhotoURL:  fbUser.PhotoURL,
		CreatedAt: time.Now(),
	}

	created, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully", "user": created})
}


func (h *UserHandler) DeleteAccountHandler(c *gin.Context) {
	v, ok := c.Get("firebase_user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "firebase user not found in context"})
		return
	}
	fbUser := v.(*models.User)
	if err := h.service.DeleteUser(fbUser.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
