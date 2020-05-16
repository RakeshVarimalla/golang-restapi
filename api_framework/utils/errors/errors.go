package errors

import "net/http"

type ApiErrors interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct{
	AStatus int `json:"status""`
	AMessage string `json:"message""`
	AError string `json:"error,omitempty"`
}


func (s *apiError) Status() int{
	return s.AStatus
}
func (s *apiError) Message() string{
	return s.AMessage
}
func (s *apiError) Error() string{
	return s.AError
}

func NewApiError(status int, message string) ApiErrors{
return &apiError{
AStatus:status,
AMessage: message,
}
}
func NewNotFoundError(message string) ApiErrors{
	return &apiError{
		AStatus:http.StatusNotFound,
		AMessage: message,
	}
}

func NewInternalServerError(message string) ApiErrors{
return &apiError{
AStatus:http.StatusInternalServerError,
AMessage: message,
}
}

func NewBadRequestError(message string) ApiErrors{
return &apiError{
AStatus:http.StatusBadRequest,
AMessage: message,
}
}