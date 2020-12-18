package exception

import (
	"errors"
)

//RequestError ... Errors produced by an HTTP request within Status Code
type RequestError struct {
	Code int
	Err error
}

//HTTPError ... HTTP Error format
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

//NewWithError ... returns an request error that formats as the given text.
//Each call to NewWithError returns a distinct error value even if the text is identical.
func NewWithError(code int, err error) *RequestError {
	return &RequestError{code, err}
}

//New ... returns an request error that formats as default text "Exception".
//Each call to New returns a distinct error value even if the text is identical.
func New(code int) *RequestError {
	return &RequestError{code, errors.New("Exception")}
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}