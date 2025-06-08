package handler

import (
	"log"

	"github.com/child6yo/rago/api-gateway/internal"
	"github.com/gin-gonic/gin"
)

func (h *Handler) loadDocuments(c *gin.Context) {
	var docs []internal.Document

	if err := c.BindJSON(&docs); err != nil {
		// TODO
		log.Print(err)
		c.JSON(500, nil)
		return
	}

	if err := h.grpclient.Storage.LoadDocuments(docs); err != nil {
		// TODO
		log.Print(err)
		c.JSON(500, nil)
		return
	}

	log.Print("success")
	c.JSON(201, nil)
}
