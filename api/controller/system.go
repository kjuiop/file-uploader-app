package controller

import (
	"file-uploader-app/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

type SystemController struct {
}

func NewSystemController() *SystemController {
	return &SystemController{}
}

func (s *SystemController) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, models.HealthRes{Message: "pong"})
}

func (s *SystemController) OccurPanic(c *gin.Context) {
	panic("panic encounter")
}

func (s *SystemController) Print(c *gin.Context) {

	for i := 0; i < 20; i++ {
		slog.Debug("print log index", "index", i+1)
		time.Sleep(time.Second)
	}
}
