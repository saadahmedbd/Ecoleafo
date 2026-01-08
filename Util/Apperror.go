package util

type AppError struct {
	Message string
	Code    int
}

func NewAppError(message string, code int) *AppError {
	return &AppError{Message: message, Code: code}
}

func (e *AppError) Error() string {
	return e.Message
}
