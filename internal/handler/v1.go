package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"access-system-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// V1Handler defines the interface for version 1 API handlers.
type V1Handler interface {
	AddEmbeddingHandler(c *gin.Context)
	ValidateEmbeddingHandler(c *gin.Context)
	DeleteEmbeddingHandler(c *gin.Context)
}

// v1Handler implements the V1Handler interface.
type v1Handler struct {
	embeddingService service.EmbeddingService
	log              *logrus.Logger
}

// NewV1Handler creates a new instance of v1Handler.
func NewV1Handler(embeddingService service.EmbeddingService, log *logrus.Logger) V1Handler {
	return &v1Handler{
		embeddingService: embeddingService,
		log:              log,
	}
}

// AddEmbeddingHandler handles the addition of a new embedding.
func (h *v1Handler) AddEmbeddingHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var data struct {
		Name   string    `json:"name" binding:"required"`
		Vector []float32 `json:"vector" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		h.log.Errorln("Error binding JSON:", err)
		c.String(http.StatusBadRequest, "Bad Request: %v", err)
		return
	}

	err := h.embeddingService.AddEmbedding(ctx, data.Name, data.Vector)
	if err != nil {
		h.log.Errorln("Error adding embedding:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error: %v", err)
		return
	}

	c.Status(http.StatusCreated)
}

// ValidateEmbeddingHandler handles the validation of an embedding.
func (h *v1Handler) ValidateEmbeddingHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var data struct {
		Vector []float32 `json:"vector" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		h.log.Errorln("Error binding JSON:", err)
		c.String(http.StatusBadRequest, "Bad Request: %v", err)
		return
	}

	err := h.embeddingService.ValidateEmbedding(ctx, data.Vector)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.log.Infoln("No relevant matches found:", err)
			c.Status(http.StatusNotFound)
			return
		}
		h.log.Errorln("Error validating embedding:", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	h.log.Infoln("Relevant match found")
	c.Status(http.StatusOK)
	return
}

// DeleteEmbeddingHandler handles the deletion of an embedding.
func (h *v1Handler) DeleteEmbeddingHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var data struct {
		ID int64 `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		h.log.Errorln("Error binding JSON:", err)
		c.String(http.StatusBadRequest, "Bad Request: %v", err)
		return
	}

	err := h.embeddingService.DeleteEmbedding(ctx, data.ID)
	if err != nil {
		h.log.Errorln("Error deleting embedding:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error: %v", err)
		return
	}

	c.Status(http.StatusOK)
}
