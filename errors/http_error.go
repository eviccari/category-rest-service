package errors

// HTTPError - Data structure that defines error payload
type HTTPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// HTTPErrorHandler - To generate http response with status and error message
type HTTPErrorHandler interface {
	Response(httpStatus int, message string)
}