package controller

import (
	"file-uploader-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemController struct {
}

func NewSystemController() *SystemController {
	return &SystemController{}
}

func (s *SystemController) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, models.HealthRes{Message: "pong"})
}
