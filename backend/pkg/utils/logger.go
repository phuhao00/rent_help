package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger interface
type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
}

// Field represents a key-value pair for structured logging
type Field struct {
	Key   string
	Value interface{}
}

// SimpleLogger is a basic logger implementation
type SimpleLogger struct {
	logger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() Logger {
	return &SimpleLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Info logs an info message
func (l *SimpleLogger) Info(msg string, fields ...Field) {
	l.logWithLevel("INFO", msg, fields...)
}

// Error logs an error message
func (l *SimpleLogger) Error(msg string, fields ...Field) {
	l.logWithLevel("ERROR", msg, fields...)
}

// Warn logs a warning message
func (l *SimpleLogger) Warn(msg string, fields ...Field) {
	l.logWithLevel("WARN", msg, fields...)
}

// Debug logs a debug message
func (l *SimpleLogger) Debug(msg string, fields ...Field) {
	l.logWithLevel("DEBUG", msg, fields...)
}

// logWithLevel logs a message with the specified level
func (l *SimpleLogger) logWithLevel(level, msg string, fields ...Field) {
	// Get caller information
	_, file, line, _ := runtime.Caller(3)

	// Format fields
	fieldStr := ""
	if len(fields) > 0 {
		fieldStr = " | "
		for i, field := range fields {
			if i > 0 {
				fieldStr += ", "
			}
			fieldStr += fmt.Sprintf("%s=%v", field.Key, field.Value)
		}
	}

	// Log the message
	l.logger.Printf("[%s] %s:%d | %s%s", level, file, line, msg, fieldStr)
}

// GinLogger returns a gin middleware for logging
func GinLogger(logger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate request duration
		latency := time.Since(start)

		// Get status code
		statusCode := c.Writer.Status()

		// Build log message
		if raw != "" {
			path = path + "?" + raw
		}

		// Log based on status code
		fields := []Field{
			{"method", c.Request.Method},
			{"path", path},
			{"status", statusCode},
			{"latency", latency},
			{"ip", c.ClientIP()},
			{"user_agent", c.Request.UserAgent()},
		}

		msg := fmt.Sprintf("%s %s", c.Request.Method, path)

		if statusCode >= 500 {
			logger.Error(msg, fields...)
		} else if statusCode >= 400 {
			logger.Warn(msg, fields...)
		} else {
			logger.Info(msg, fields...)
		}
	}
}

// GinRecovery returns a gin middleware for panic recovery
func GinRecovery(logger Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.Error("Panic recovered",
			Field{"error", recovered},
			Field{"path", c.Request.URL.Path},
			Field{"method", c.Request.Method},
		)
		c.AbortWithStatus(500)
	})
}
