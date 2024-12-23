package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestLoggerCreation(t *testing.T) {
	t.Run("test logger creation with file output", func(t *testing.T) {
		logger := NewLogger("file", "info")
		if logger == nil {
			t.Fatal("expected logger to be created, but got nil")
		}
		if logger.logLevel != "info" {
			t.Fatalf("expected log level to be 'info', but got %v", logger.logLevel)
		}
	})

	t.Run("test logger creation with stdout", func(t *testing.T) {
		logger := NewLogger("stdout", "debug")
		if logger == nil {
			t.Fatal("expected logger to be created, but got nil")
		}
		if logger.logLevel != "debug" {
			t.Fatalf("expected log level to be 'debug', but got %v", logger.logLevel)
		}
	})
}

func TestLogLevels(t *testing.T) {
	tests := []struct {
		name     string
		logLevel string
		message  string
		expected string
	}{
		{"debug level", "debug", "debug message", "DEBUG"},
		{"info level", "info", "info message", "INFO"},
		{"warn level", "warn", "warn message", "WARN"},
		{"error level", "error", "error message", "ERROR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture output
			var buf bytes.Buffer
			logger := NewLogger("stdout", tt.logLevel)
			logger.Logger.SetOutput(&buf)

			// Log the message
			logger.logMessage(tt.logLevel, tt.message)

			// Check if the output contains the expected log level
			if !strings.Contains(buf.String(), tt.expected) {
				t.Fatalf("expected log message to contain %s, but got: %v", tt.expected, buf.String())
			}
		})
	}
}

func TestLogMessageFormatting(t *testing.T) {
	tests := []struct {
		logLevel string
		expected string
	}{
		{"info", "INFO"},
		{"debug", "DEBUG"},
		{"warn", "WARN"},
		{"error", "ERROR"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("test %s log message", tt.logLevel), func(t *testing.T) {
			// Capture output
			var buf bytes.Buffer
			logger := NewLogger("stdout", tt.logLevel)
			logger.Logger.SetOutput(&buf)

			// Log a message
			logger.logMessage(tt.logLevel, "Test message")

			// Check for log level in output
			if !strings.Contains(buf.String(), tt.expected) {
				t.Fatalf("expected log message to contain %s, but got: %v", tt.expected, buf.String())
			}

			// Check for timestamp in output
			if !strings.Contains(buf.String(), time.Now().Format("2006-01-02")) {
				t.Fatalf("expected log message to contain timestamp, but got: %v", buf.String())
			}
		})
	}
}

func TestSerializationErrorHandling(t *testing.T) {
	t.Run("test serialization error", func(t *testing.T) {
		// Capture output
		var buf bytes.Buffer
		logger := NewLogger("stdout", "debug")
		logger.Logger.SetOutput(&buf)

		// Log a message with a non-serializable argument (e.g., a channel)
		ch := make(chan int)
		logger.Debug("Test message with non-serializable arg", ch)

		// Check for the error in the output
		if !strings.Contains(buf.String(), "serialization error") {
			t.Fatalf("expected log message to contain 'serialization error', but got: %v", buf.String())
		}
	})
}
