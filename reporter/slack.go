package reporter

import (
	"bytes"
	"encoding/json"
	"file-uploader-app/config"
	"file-uploader-app/models"
	"log/slog"
	"net/http"
)

var Client *Slack

type Slack struct {
	cfg config.Slack
}

func NewSlackReporter(cfg config.Slack) {
	slack := &Slack{
		cfg: cfg,
	}

	Client = slack
}

func (s *Slack) SendSlackPanicReport(message string) {
	if message == "" {
		slog.Error("fail send panic report, message is empty")
		return
	}

	webhookRes := models.NewWebhookRes(message)
	body, err := json.Marshal(webhookRes)
	if err != nil {
		slog.Error("fail send panic report, message is empty", "error", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, s.cfg.WebhookReportUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("fail report recover alarm", "error", err)
	}
	defer resp.Body.Close()
}
