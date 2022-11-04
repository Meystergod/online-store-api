package apperror

import "encoding/json"

var (
	ErrorNotFound   = NewAppError(nil, "not found", "SA-00101")
	ErrorBadRequest = NewAppError(nil, "bad request", "SA-00102")
)

type AppError struct {
	Err     error  `json:"-"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func NewAppError(err error, message string, code string) *AppError {
	return &AppError{
		Err:     err,
		Message: message,
		Code:    code,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func systemError(err error) *AppError {
	return NewAppError(err, "internal system error", "SA-00001")
}
