package errors

import "github.com/gin-gonic/gin"

// GinHTTPErrorHandler - Describe GinHTTPErrorHandler structure
type GinHTTPErrorHandler struct{
	context *gin.Context
}

// NewGinHTTPErrorHandler - Create new GinHTTPErrorHandler instance
func NewGinHTTPErrorHandler(context *gin.Context) (ginHTTPErrorHandler GinHTTPErrorHandler) {
	return GinHTTPErrorHandler{context: context}
}

// Response - Implements HTTPError interface
func (handler *GinHTTPErrorHandler) Response(httpStatus int, message string) {
	handler.context.JSON(
		httpStatus, 
		gin.H{
			"message": message,
		},
	)
}