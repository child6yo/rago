package handler

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// TODO
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		log.Print("empty authorization header")
		c.JSON(500, nil)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		log.Print("invalid authorizaton header")
		c.JSON(500, nil)
		return
	}

	if len(headerParts[1]) == 0 {
		log.Print("token is empty")
		c.JSON(500, nil)
		return
	}

	userId, err := h.grpclient.User.Auth(headerParts[1])
	if err != nil {
		log.Print("invalid token")
		c.JSON(500, nil)
		return
	}

	c.Set("userId", userId)
}

func (h *Handler) checkAPIKey(c *gin.Context) {
	key := c.Query("api-key")
	if key == "" {
		log.Print("empty api key")
		c.JSON(500, nil)
		return
	}

	err := h.grpclient.User.CheckAPIKey(key)
	if err != nil {
		log.Print("invalid api key: ", err)
		c.JSON(500, nil)
		return
	}
}
