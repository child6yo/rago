package handler

import (
	"log"

	"github.com/child6yo/rago/api-gateway/internal"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input internal.User

	if err := c.BindJSON(&input); err != nil {
		// TODO
		log.Print(err)
		return
	}

	if err := h.grpclient.User.Register(input); err != nil {
		// TODO
		log.Print(err)
		return
	}
}
