package errors

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewError(statusCode int, message string) *Error {
	return &Error{
		StatusCode: statusCode,
		Message:    message,
	}
}
