package controller

import (
	"encoding/json"
	"file-uploader-app/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHealthSuccess(t *testing.T) {

	tests := []struct {
		expectedCode int
		title        string
		response     models.HealthRes
	}{
		{
			http.StatusOK, "health-check api test", models.HealthRes{Message: "pong"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testAssert := assert.New(t)
			resp := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(resp)

			c.Request = httptest.NewRequest(http.MethodGet, "/health-check", nil)

			systemController.GetHealth(c)
			testAssert.Equal(tc.expectedCode, resp.Code)

			var responseBody models.HealthRes
			err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
			testAssert.NoError(err)
			testAssert.Equal(tc.response.Message, responseBody.Message)
		})
	}
}
