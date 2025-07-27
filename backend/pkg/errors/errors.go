package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string      `json:"error"`
	Message string      `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
	Code    int         `json:"code"`
}

// AppError represents an application error
type AppError struct {
	Message    string
	StatusCode int
	Details    interface{}
}

// Error implements the error interface
func (e AppError) Error() string {
	return e.Message
}

// Common errors
var (
	ErrBadRequest         = AppError{Message: "Bad request", StatusCode: http.StatusBadRequest}
	ErrUnauthorized       = AppError{Message: "Unauthorized", StatusCode: http.StatusUnauthorized}
	ErrForbidden          = AppError{Message: "Forbidden", StatusCode: http.StatusForbidden}
	ErrNotFound           = AppError{Message: "Resource not found", StatusCode: http.StatusNotFound}
	ErrConflict           = AppError{Message: "Resource already exists", StatusCode: http.StatusConflict}
	ErrInternalServer     = AppError{Message: "Internal server error", StatusCode: http.StatusInternalServerError}
	ErrValidation         = AppError{Message: "Validation failed", StatusCode: http.StatusBadRequest}
	ErrInvalidCredentials = AppError{Message: "Invalid credentials", StatusCode: http.StatusUnauthorized}
	ErrTokenExpired       = AppError{Message: "Token expired", StatusCode: http.StatusUnauthorized}
	ErrInvalidToken       = AppError{Message: "Invalid token", StatusCode: http.StatusUnauthorized}
)

// NewAppError creates a new application error
func NewAppError(message string, statusCode int, details interface{}) AppError {
	return AppError{
		Message:    message,
		StatusCode: statusCode,
		Details:    details,
	}
}

// NewValidationError creates a validation error
func NewValidationError(details interface{}) AppError {
	return AppError{
		Message:    "Validation failed",
		StatusCode: http.StatusBadRequest,
		Details:    details,
	}
}

// NewNotFoundError creates a not found error
func NewNotFoundError(resource string) AppError {
	return AppError{
		Message:    resource + " not found",
		StatusCode: http.StatusNotFound,
	}
}

// NewUnauthorizedError creates an unauthorized error
func NewUnauthorizedError(message string) AppError {
	return AppError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

// NewConflictError creates a conflict error
func NewConflictError(message string) AppError {
	return AppError{
		Message:    message,
		StatusCode: http.StatusConflict,
	}
}

// HandleError handles application errors and sends appropriate response
func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case AppError:
		c.JSON(e.StatusCode, ErrorResponse{
			Error:   e.Message,
			Details: e.Details,
			Code:    e.StatusCode,
		})
	default:
		// For unexpected errors, don't expose internal details
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Internal server error",
			Code:  http.StatusInternalServerError,
		})
	}
}

// ErrorHandler returns a gin middleware for centralized error handling
func ErrorHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(error); ok {
			HandleError(c, err)
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error: "Internal server error",
				Code:  http.StatusInternalServerError,
			})
		}
	})
}

// BadRequest sends a bad request error response
func BadRequest(c *gin.Context, message string, details interface{}) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Error:   message,
		Details: details,
		Code:    http.StatusBadRequest,
	})
}

// Unauthorized sends an unauthorized error response
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Error: message,
		Code:  http.StatusUnauthorized,
	})
}

// Forbidden sends a forbidden error response
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, ErrorResponse{
		Error: message,
		Code:  http.StatusForbidden,
	})
}

// NotFound sends a not found error response
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Error: message,
		Code:  http.StatusNotFound,
	})
}

// Conflict sends a conflict error response
func Conflict(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, ErrorResponse{
		Error: message,
		Code:  http.StatusConflict,
	})
}

// InternalServerError sends an internal server error response
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error: message,
		Code:  http.StatusInternalServerError,
	})
}

// Success sends a success response
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// Created sends a created response
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}
