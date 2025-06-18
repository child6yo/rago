package handler

import (
	"fmt"

	"github.com/child6yo/rago/api-gateway/internal"
	"github.com/gin-gonic/gin"
)

func (h *Handler) loadDocuments(c *gin.Context) {
	var docs internal.DocumentArray

	collection := c.Param("collection")
	if collection == "" {
		errorResponse(c, "empty collection parameter", 400, nil)
		return
	}

	if err := c.BindJSON(&docs); err != nil {
		errorResponse(c, "internal server error", 500, err)
		return
	}

	docs.Collection = collection

	if err := h.kafkaProducer.SendMessage(docs); err != nil {
		errorResponse(c, "failed to send document array JSON", 500, err)
		return
	}

	c.JSON(201, nil)
}

func (h *Handler) deleteDocument(c *gin.Context) {
	collection := c.Param("collection")
	if collection == "" {
		errorResponse(c, "empty collection parameter", 400, nil)
		return
	}

	docID := c.Param("id")
	if docID == "" {
		errorResponse(c, "empty docID parameter", 400, nil)
		return
	}

	err := h.grpclient.DeleteDocument(c.Request.Context(), collection, docID)
	if err != nil {
		errorResponse(c, fmt.Sprintf("failed to delete document from collection %s", collection), 500, err)
	}

	successResponse(c, fmt.Sprintf("document from collection %s successfull deleted", collection), nil)
}

func (h *Handler) getDocument(c *gin.Context) {
	collection := c.Param("collection")
	if collection == "" {
		errorResponse(c, "empty collection parameter", 400, nil)
		return
	}

	docID := c.Param("id")
	if docID == "" {
		errorResponse(c, "empty docID parameter", 400, nil)
		return
	}

	doc, err := h.grpclient.GetDocument(c.Request.Context(), collection, docID)
	if err != nil {
		errorResponse(c, fmt.Sprintf("failed to get document from collection %s", collection), 500, err)
	}

	successResponse(c, "document successfully got", doc)
}

func (h *Handler) getAllDocuments(c *gin.Context) {
	collection := c.Param("collection")
	if collection == "" {
		errorResponse(c, "empty collection parameter", 400, nil)
		return
	}

	docs, err := h.grpclient.GetAllDocuments(c.Request.Context(), collection)
	if err != nil {
		errorResponse(c, fmt.Sprintf("failed to get documents from collection %s", collection), 500, err)
	}

	docs.Collection = collection

	successResponse(c, "documents successfully got", docs)
}
