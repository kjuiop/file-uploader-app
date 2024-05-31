package server

import (
	"context"
	"errors"
	"file-uploader-app/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sync"
)

type Gin struct {
	srv *http.Server
	cfg config.Server
}

func NewGinServer(cfg config.Server) Client {

	router := getGinEngine(cfg.Mode)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	return &Gin{
		srv: srv,
		cfg: cfg,
	}
}

func (g *Gin) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	err := g.srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		slog.Debug("server close")
	} else {
		slog.Error("run server error", "error", err)
	}
}

func (g *Gin) Shutdown(ctx context.Context) {
	if err := g.srv.Shutdown(ctx); err != nil {
		slog.Error("error during server shutdown", "error", err)
	}
}

func getGinEngine(mode string) *gin.Engine {
	switch mode {
	case "prod":
		return gin.New()
	case "test":
		gin.SetMode(gin.TestMode)
		return gin.Default()
	default:
		return gin.Default()
	}
}
