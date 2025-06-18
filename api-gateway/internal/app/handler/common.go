package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: полностью переработать логирование

// Response определяет структуру HTTP ответа.
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get("userID")
	if !ok {
		return 0, errors.New("unknown jwt")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("unknown jwt")
	}

	return idInt, nil
}

func errorResponse(c *gin.Context, description string, httpError int, err error) {
	log.Printf("http request handling error: %s: %v", description, err)
	c.AbortWithStatusJSON(httpError, Response{Status: httpError, Data: description})
}

func successResponse(c *gin.Context, description string, data any) {
	log.Printf("http request successfully handled: %s", description)
	c.JSON(http.StatusOK, Response{Status: http.StatusOK, Data: data})
}
