package controller

import (
	"file-uploader-app/api/form"
	"file-uploader-app/models"
	"fmt"
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

	requestId, exists := c.Get("request_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, form.FailRes{Message: "request not exist"})
		return
	}

	panic(fmt.Errorf("panic encounter, request_id : %s", requestId))
}

func (s *SystemController) Print(c *gin.Context) {

	requestId, exists := c.Get("request_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, form.FailRes{Message: "request not exist"})
		return
	}

	for i := 0; i < 20; i++ {
		slog.Debug("print log index", "request_id", requestId, "index", i+1)
		time.Sleep(time.Second)
	}

	c.JSON(http.StatusOK, form.SuccessRes{Message: "ok"})
}
