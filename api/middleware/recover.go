package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func RecoveryErrorReport(webhookUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errMsg := fmt.Sprintf("recovered from panic : %v", err)
				slog.Error(errMsg)
				sendSlackPanicReport(webhookUrl, errMsg)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// sendSlackNotification sends a Slack notification with the given message
func sendSlackPanicReport(webhookUrl, message string) {
	if len(webhookUrl) == 0 || len(message) == 0 {
		return
	}
	jsonStr := fmt.Sprintf(`{"text":"%s"}`, message)
	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	slog.Debug("report error message", "msg", message)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("fail report recover alarm", "error", err)
	}
	defer resp.Body.Close()
}
