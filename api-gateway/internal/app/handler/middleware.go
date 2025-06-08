package handler

import (
	"strings"
	"log"

	"github.com/gin-gonic/gin"
)

// TODO
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		log.Printf("empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		log.Printf("invalid authorizaton header")
		return
	}

	if len(headerParts[1]) == 0 {
		log.Printf("token is empty")
		return
	}

	userId, err := h.grpclient.User.Auth(headerParts[1])
	if err != nil {
		log.Printf("invalid token")
		return
	}

	c.Set("userId", userId)
}
