package handler

import (
	"fmt"

	"github.com/child6yo/rago/api-gateway/internal"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input internal.User

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, "internal server error", 500, err)
		return
	}

	collection, err := h.grpclient.User.Register(c.Request.Context(), input)
	if err != nil {
		errorResponse(c, "failed to register", 500, err)
		return
	}

	err = h.grpclient.CreateCollection(c.Request.Context(), collection)
	if err != nil {
		errorResponse(c, fmt.Sprintf("failed to create collection: %s", collection), 500, err)
		return
	}

	data := struct {
		Collection string `json:"collection"`
	}{
		Collection: collection,
	}

	successResponse(c, fmt.Sprintf("new user successfully created with collection %s", data), nil)
}

func (h *Handler) signIn(c *gin.Context) {
	var input internal.User

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, "internal server error", 500, err)
		return
	}

	token, err := h.grpclient.Login(c.Request.Context(), input)
	if err != nil {
		errorResponse(c, "failed to login", 500, err)
		return
	}

	data := struct {
		Token string `json:"token"`
	}{
		Token: fmt.Sprintf("Bearer %s", token),
	}

	successResponse(c, "successfull user sign in", data)
}

func (h *Handler) createAPIKey(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		errorResponse(c, "failed to get auth token", 401, err)
		return
	}

	key, err := h.grpclient.User.CreateAPIKey(c.Request.Context(), id)
	if err != nil {
		errorResponse(c, "failed to create API key", 500, err)
		return
	}

	data := struct {
		APIKey string `json:"key"`
	}{
		APIKey: key,
	}

	successResponse(c, "new api key created successfully", data)
}

func (h *Handler) deleteAPIKey(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		errorResponse(c, "failed to get auth token", 401, err)
		return
	}

	keyID := c.Query("id")
	if keyID == "" {
		errorResponse(c, "failed to get api key id query", 400, nil)
		return
	}

	err = h.grpclient.User.DeleteAPIKey(c.Request.Context(), keyID, userID)
	if err != nil {
		errorResponse(c, "failed to delete api key", 500, err)
		return
	}

	successResponse(c, "api key successfully deleted", nil)
}

func (h *Handler) getAPIKeys(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		errorResponse(c, "failed to get auth token", 401, err)
		return
	}

	keys, err := h.grpclient.User.GetAPIKeys(c.Request.Context(), userID)
	if err != nil {
		errorResponse(c, "failed to get api keys", 500, err)
		return
	}

	successResponse(c, "api keys successfully got", keys)
}

func (h *Handler) getCollection(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		errorResponse(c, "failed to get auth token", 401, err)
		return
	}

	collection, err := h.grpclient.GetCollection(c.Request.Context(), userID)
	if err != nil {
		errorResponse(c, "failed to get user collection", 500, err)
		return
	}

	data := struct {
		Collection string `json:"collection"`
	}{
		Collection: collection,
	}

	successResponse(c, "collection successfully got", data)
}
