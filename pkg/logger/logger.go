package logger

import (
	"context"
	"fmt"
	"log"
	"os"
)

var g_log func(format string, args ...interface{})

// ContextKey is type for context key.
type ContextKey string

// ContextValue is type for context value.
type ContextValue string

const logIDKey ContextKey = "log_id"

// GetContextWithLogID is used to setup context
// and set log ID into it.
func GetContextWithLogID(ctx context.Context, logID string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, logIDKey, ContextValue(logID))
}

// SetupLogger is used to setup logging function.
func SetupLogger(logFn func(format string, args ...interface{})) {
	g_log = logFn
}

// Infof is used to log information message.
func Infof(ctx context.Context, format string, args ...interface{}) {
	printf(ctx, "INFO", format, args...)
}

// Warnf is used to log warning message.
func Warnf(ctx context.Context, format string, args ...interface{}) {
	printf(ctx, "WARN", format, args...)
}

// Errorf is used to log error message.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	printf(ctx, "ERROR", format, args...)
}

// Fatalf is used to log error message
// then call os.Exit(1).
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	printf(ctx, "FATAL", format, args...)
	os.Exit(1)
}

func getDefaultLogFn() func(format string, args ...interface{}) {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	return func(format string, args ...interface{}) {
		log.Printf(fmt.Sprintf("\b\b%s", format), args...)
	}
}

func printf(ctx context.Context, mode, format string, args ...interface{}) {
	if g_log == nil {
		g_log = getDefaultLogFn()
	}
	logFormat := fmt.Sprintf("|%s|%s", mode, format)
	if ctx != nil {
		logId, ok := ctx.Value(logIDKey).(ContextValue)
		if ok {
			logFormat = fmt.Sprintf("|%s|log_id=%s|%s", mode, logId, format)
		}
	}
	g_log(logFormat, args...)
}
