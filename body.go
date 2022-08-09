package stdres

import "net/http"

type body struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	ServerTime int64       `json:"serverTime"`
}

func NewRawResponseBody(code, message string, data interface{}, timestamp ...int64) interface{} {
	return &body{
		Code:       code,
		Message:    message,
		Data:       data,
		ServerTime: getCurrentServerTime(timestamp...),
	}
}

// NewResponseBody return a response body with httpCode, message, and optional data.
func NewResponseBody(httpCode int, message string, data interface{}, timestamp ...int64) interface{} {
	return NewRawResponseBody(http.StatusText(httpCode), message, data, timestamp...)
}

// NewBadRequestResponse return a bad request response body.
func NewBadRequestResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusBadRequest, message, data, timestamp...)
}

// NewNotFoundResponse return a not found response body.
func NewNotFoundResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusNotFound, message, data, timestamp...)
}

// NewConflictResponse return a conflict response body.
func NewConflictResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusConflict, message, data, timestamp...)
}

// NewUnauthorizedResponse return an unauthorized response body.
func NewUnauthorizedResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusUnauthorized, message, data, timestamp...)
}

// NewForbiddenResponse return a forbidden response body.
func NewForbiddenResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusForbidden, message, data, timestamp...)
}

// NewInternalServerErrorResponse return an internal server error response body.
func NewInternalServerErrorResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusInternalServerError, message, data, timestamp...)
}

// NewSuccessResponse return a success response body.
func NewSuccessResponse(data interface{}, timestamp ...int64) interface{} {
	return NewRawResponseBody("SUCCESS", "SUCCESS", data, timestamp...)
}

// NewOKResponse return an OK response body.
func NewOKResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusOK, message, data, timestamp...)
}

// NewCreatedResponse return a created response body.
func NewCreatedResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusCreated, message, data, timestamp...)
}

// NewAcceptedResponse return an accepted response body.
func NewAcceptedResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusAccepted, message, data, timestamp...)
}

// NewNoContentResponse return a no content response body.
func NewNoContentResponse(message string, data interface{}, timestamp ...int64) interface{} {
	return NewResponseBody(http.StatusNoContent, message, data, timestamp...)
}
