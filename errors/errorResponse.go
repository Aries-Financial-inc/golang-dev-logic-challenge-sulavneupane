package errors

// Error Response Codes to track what happened internally

const (
	ErrorCodeBadRequest = 400
)

type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}
