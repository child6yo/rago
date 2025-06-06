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
		}
	}

	return router
}
