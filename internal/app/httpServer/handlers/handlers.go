package handlers

import (
	"todo-user-service/internal/services"
	"todo-user-service/pkg/tokenManager"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *services.Services
	tokenManager *tokenManager.TokenManager
	router       *gin.Engine
}

func NewHandler(services *services.Services, tokenManager *tokenManager.TokenManager, router *gin.Engine) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		router:       router,
	}
}

func (h *Handler) RunHandler() {
	h.router.POST("/auth/register", h.Registration)
	h.router.POST("/auth/login", h.Login)
}
