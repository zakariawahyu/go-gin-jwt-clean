package response

import "strings"

type ResponsesSuccess struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsesError struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func BuildResponse(status bool, message string, data interface{}) ResponsesSuccess {
	return ResponsesSuccess{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func BuildErrorResponse(message string, err string) ResponsesError {
	splittedError := strings.Split(err, "; ")

	return ResponsesError{
		Status:  false,
		Message: message,
		Errors:  splittedError,
	}
}
