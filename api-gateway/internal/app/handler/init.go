package handler

import (
	"github.com/child6yo/rago/api-gateway/internal/app/client"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	grpclient *client.GRPClient
}

func NewHandler(grpclient *client.GRPClient) *Handler {
	return &Handler{grpclient: grpclient}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			auth := user.Group("/auth")
			{
				auth.POST("/sign-up", h.signUp)
				auth.POST("/sign-in", h.signIn)
			}

			apiKeys := user.Group("/api-keys", h.userIdentity)
			{
				apiKeys.POST("/", h.createAPIKey)
				apiKeys.GET("/", h.getAPIKeys)
				apiKeys.DELETE("/", h.deleteAPIKey)
			}
		}

		storage := api.Group("/storage", h.checkAPIKey)
		{
			storage.POST("/", h.loadDocuments)
		}

		api.GET("/generation", h.ssEventMiddleware(), h.generateAnswer)
	}

	return router
}
