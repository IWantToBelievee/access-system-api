package router

import (
	"access-system-api/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Router struct to hold the Gin engine and handlers
type Router struct {
	engine *gin.Engine
	v1     handler.V1Handler
	log    *logrus.Logger
}

// NewRouter initializes a new Router instance
func NewRouter(v1 handler.V1Handler, log *logrus.Logger) *Router {
	return &Router{
		engine: gin.Default(),
		v1:     v1,
		log:    log,
	}
}

// Run starts the Gin server and sets up the routes
func (r *Router) Run() {
	v1 := r.engine.Group("/api/v1")
	{
		v1.POST("/embedding", r.v1.AddEmbeddingHandler)
		v1.POST("/embedding/validate", r.v1.ValidateEmbeddingHandler)
		v1.DELETE("/embedding", r.v1.DeleteEmbeddingHandler)
	}

	gin.SetMode(gin.ReleaseMode)
	if err := r.engine.Run(":8081"); err != nil {
		r.log.Fatalf("Failed to start server: %v", err)
	}
}
