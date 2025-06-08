package handler

import (
	"log"
	"strconv"

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

func (h *Handler) createAPIKey(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		// TODO
		c.JSON(500, err)
		return
	}

	key, err := h.grpclient.User.CreateAPIKey(id)
	if err != nil {
		// TODO
		c.JSON(500, err)
		return
	}

	Data := struct {
		APIKey string `json:"key"`
	}{
		APIKey: key,
	}

	c.JSON(200, Response{Status: "OK", Data: Data})
}

func (h *Handler) deleteAPIKey(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		// TODO
		log.Print(err)
		c.JSON(500, err)
		return
	}

	keyID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		// TODO
		log.Print(err)
		c.JSON(500, err)
		return
	}

	err = h.grpclient.User.DeleteAPIKey(keyID, userID)
	if err != nil {
		// TODO
		log.Print(err)
		c.JSON(500, err)
		return
	}
	
	c.JSON(200, nil)
}

func (h *Handler) getAPIKeys(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		// TODO
		log.Print(err)
		c.JSON(500, err)
		return
	}

	keys, err := h.grpclient.User.GetAPIKeys(userID)
		if err != nil {
		// TODO
		log.Print(err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, keys)
}
