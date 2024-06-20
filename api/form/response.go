package form

type FailRes struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type SuccessRes struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}
