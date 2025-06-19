package handler

import (
	"io"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) generateAnswer(c *gin.Context) {
	c.Writer.Flush()

	query := c.Query("query")
	collection := c.Param("collection")

	// получение потока токенов
	stream, err := h.grpclient.Generator.Generate(c, query, collection)
	if err != nil {
		log.Printf("Generation error: %v", err)
		c.SSEvent("error", "Failed to start generation")
		c.Writer.Flush()
		return
	}

	// потоковая передача
	c.Stream(func(w io.Writer) bool {
		select {
		case <-c.Writer.CloseNotify():
			log.Println("Client disconnected")
			return false

		case chunk, ok := <-stream:
			if !ok {
				c.SSEvent("end", "[DONE]")
				log.Println("Stream finished")
				return false
			}

			// отправление чанка данных

			chunkWithNbsp := strings.TrimLeftFunc(chunk, func(r rune) bool { return r == ' ' })
			leadingSpaces := len(chunk) - len(chunkWithNbsp)
			spaces := strings.Repeat("&nbsp;", leadingSpaces)
			finalChunk := spaces + chunkWithNbsp

			c.SSEvent("message", finalChunk)
			return true
		}
	})
}
