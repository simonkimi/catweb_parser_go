package models

import "fmt"

const (
	InternalError = iota
	ParserError
	ElementNotFoundError
	ConverterError
)

type ParseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ParseError) Error() string {
	var errType string
	switch e.Code {
	case InternalError:
		errType = "InternalError"
	case ParserError:
		errType = "ParserError"
	case ElementNotFoundError:
		errType = "ElementNotFoundError"
	}
	return fmt.Sprintf("[%s] %s", errType, e.Message)
}

func NewParseError(code int, message string) *ParseError {
	return &ParseError{
		Code:    code,
		Message: message,
	}
}
