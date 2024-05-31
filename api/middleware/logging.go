package middleware

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func LoggingMiddleware(c *gin.Context) {

	start := time.Now() // Start timer
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	// Process request
	c.Next()

	// Fill the params
	param := gin.LogFormatterParams{}

	param.TimeStamp = time.Now() // Stop timer
	param.Latency = param.TimeStamp.Sub(start)
	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	param.ClientIP = c.ClientIP()
	param.Method = c.Request.Method
	param.StatusCode = c.Writer.Status()
	param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
	param.BodySize = c.Writer.Size()
	if raw != "" {
		path = path + "?" + raw
	}
	param.Path = path

	logger := slog.With(
		"client_ip", param.ClientIP,
		"method", param.Method,
		"status_code", param.StatusCode,
		"body_size", param.BodySize,
		"path", param.Path,
		"latency", param.Latency.String(),
	)

	if c.Writer.Status() >= 200 && c.Writer.Status() < 300 {
		logger.Info("success")
	} else {
		logger.Error(param.ErrorMessage)
	}
}
