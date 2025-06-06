package handler

import (
	"github.com/child6yo/rago/api-gateway/internal"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input internal.User

	if err := c.BindJSON(&input); err != nil {
		// TODO
		c.JSON(500, err)
		return
	}

	if err := h.grpclient.User.Register(input); err != nil {
		// TODO
		c.JSON(500, err)
		return
	}

	c.JSON(200, Response{Status: "OK"})
}

func (h *Handler) signIn(c *gin.Context) {
	var input internal.User

	if err := c.BindJSON(&input); err != nil {
		// TODO
		c.JSON(500, err)
		return
	}

	token, err := h.grpclient.Login(input)
	if err != nil {
		// TODO
		c.JSON(500, err)
		return
	}

	Data := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	c.JSON(200, Response{Status: "OK", Data: Data})
}
