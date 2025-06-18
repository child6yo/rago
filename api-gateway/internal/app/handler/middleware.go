package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		errorResponse(c, "empty authorization header", 401, nil)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		errorResponse(c, "invalid authorizaton header", 401, nil)
		return
	}

	if len(headerParts[1]) == 0 {
		errorResponse(c, "empty authorization token", 401, nil)
		return
	}

	userID, err := h.grpclient.User.Auth(c.Request.Context(), headerParts[1])
	if err != nil {
		errorResponse(c, "invalid authorization token", 401, nil)
		return
	}

	c.Set("userID", userID)
}

func (h *Handler) checkAPIKey(c *gin.Context) {
	key := c.Query("api-key")
	if key == "" {
		errorResponse(c, "empty api key", 401, nil)
		return
	}

	err := h.grpclient.User.CheckAPIKey(c.Request.Context(), key)
	if err != nil {
		errorResponse(c, "invalid api key", 401, err)
		return
	}
}

func (h *Handler) ssEventMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Next()
	}
}
