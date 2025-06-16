package handler

import (
	"log"

	"github.com/child6yo/rago/api-gateway/internal"
	"github.com/gin-gonic/gin"
)

func (h *Handler) loadDocuments(c *gin.Context) {
	var docs internal.DocumentArray

	collection := c.Param("collection")
	if collection == "" {
		// TODO
		log.Print("empty coll param")
		c.JSON(500, nil)
		return
	}

	if err := c.BindJSON(&docs); err != nil {
		// TODO
		log.Print(err)
		c.JSON(500, nil)
		return
	}

	docs.Collection = collection

	if err := h.kafkaProducer.SendMessage(docs); err != nil {
		// TODO
		log.Print(err)
		c.JSON(500, nil)
		return
	}

	log.Print("success")
	c.JSON(201, nil)
}
