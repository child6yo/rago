package handler

import (
	"github.com/child6yo/rago/api-gateway/internal/app/client"
	"github.com/child6yo/rago/api-gateway/internal/app/kafka/producer"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	grpclient     *client.GRPClient
	kafkaProducer producer.Producer
}

func NewHandler(grpclient *client.GRPClient, kafkaProducer producer.Producer) *Handler {
	return &Handler{grpclient: grpclient, kafkaProducer: kafkaProducer}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(prometheusMiddleware(initPrometheus()))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

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
			storage.POST("/:collection", h.loadDocuments)
		}

		api.GET("/generation", h.ssEventMiddleware(), h.generateAnswer)
	}

	return router
}
