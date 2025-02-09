package rest

import "fmt"

type OKXError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewOKXError(code, message string) OKXError {
	return OKXError{
		Code:    code,
		Message: message,
	}
}

func (e OKXError) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}
