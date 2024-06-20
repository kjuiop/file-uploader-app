package form

import "fmt"

const (
	NoError        int = 0
	ErrJsonParsing int = 4001
	ErrStructValid int = 4002
	ErrFileRequest int = 4003
)

var codeToMessage = map[int]string{
	NoError:        "ok",
	ErrJsonParsing: "invalid request body",
	ErrStructValid: "invalid request body",
	ErrFileRequest: "upload file err",
}

func GetErrMessage(code int, error string) string {
	message, exists := codeToMessage[code]
	if !exists {
		return "Unknown error"
	}

	return fmt.Sprintf("%s, err: %s", message, error)
}
