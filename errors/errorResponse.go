package errors

// Error Response Codes to track what happened internally

const (
	ErrorCodeBadRequest = 400

	ErrorCodeTooManyOptionsContracts         = 600
	ErrorCodeInvalidOptionsContractType      = 601
	ErrorCodeInvalidOptionsContractLongShort = 602
)

type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}
