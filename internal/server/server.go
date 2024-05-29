package server

import (
	"errors"
	"file-uploader-app/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	srv *http.Server
	cfg config.Server
}

func SetupGinServer(cfg config.Server) *Server {
	router := getGinEngine(cfg.IsDryRun)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	return &Server{
		srv: srv,
		cfg: cfg,
	}
}

func (s *Server) Run() {

	err := s.srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server close")
	} else {
		log.Fatalf("run server error : %s\n", err.Error())
	}
}

func getGinEngine(isDryRun bool) *gin.Engine {
	if isDryRun {
		return gin.Default()
	}
	return gin.New()
}
