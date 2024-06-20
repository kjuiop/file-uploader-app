package controller

import (
	"file-uploader-app/api/form"
	"file-uploader-app/config"
	"file-uploader-app/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type UploaderController struct {
	cfg config.Upload
}

func NewUploaderController(cfg config.Upload) *UploaderController {
	return &UploaderController{
		cfg: cfg,
	}
}

func (u *UploaderController) FileUpload(c *gin.Context) {

	req := models.FileUploadReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, form.FailRes{
			ErrorCode: form.ErrJsonParsing,
			Message:   form.GetErrMessage(form.ErrJsonParsing, err.Error()),
		})
		return
	}

	logger := slog.With(slog.String("request_param", fmt.Sprintf("%+v", req)))

	if err := req.CheckValid(); err != nil {
		logger.Error("check struct validation", slog.Int("status_code", http.StatusBadRequest), slog.Int("error_code", form.ErrStructValid), slog.String("err_message", err.Error()))
		c.JSON(http.StatusBadRequest, form.FailRes{
			ErrorCode: form.ErrStructValid,
			Message:   form.GetErrMessage(form.ErrStructValid, err.Error()),
		})
		return
	}

	targetDir := fmt.Sprintf("%s/%s", u.cfg.FileSourceDir, req.CustomerId)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, form.FailRes{
			ErrorCode: form.ErrFileRequest,
			Message:   form.GetErrMessage(form.ErrFileRequest, err.Error()),
		})
		return
	}

	if err := c.SaveUploadedFile(file, targetDir); err != nil {
		c.JSON(http.StatusBadRequest, form.FailRes{
			ErrorCode: form.ErrFileRequest,
			Message:   form.GetErrMessage(form.ErrFileRequest, err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, form.SuccessRes{
		ErrorCode: form.NoError,
		Message:   "successfully uploaded",
	})
}
