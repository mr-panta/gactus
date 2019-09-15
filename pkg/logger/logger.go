package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

var gLog func(format string, args ...interface{})

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
	now := time.Now().Unix()
	logID = fmt.Sprintf("%s_%d", logID, now)
	return context.WithValue(ctx, logIDKey, ContextValue(logID))
}

// GetLogID is used to get log ID
// from context.
func GetLogID(ctx context.Context) string {
	logID, ok := ctx.Value(logIDKey).(ContextValue)
	if !ok {
		return ""
	}
	return string(logID)
}

// SetupLogger is used to setup logging function.
func SetupLogger(logFn func(format string, args ...interface{})) {
	gLog = logFn
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
	if gLog == nil {
		gLog = getDefaultLogFn()
	}
	logFormat := fmt.Sprintf("|%s|%s", mode, format)
	if ctx != nil {
		logID, ok := ctx.Value(logIDKey).(ContextValue)
		if ok {
			logFormat = fmt.Sprintf("|%s|log_id=%s|%s", mode, logID, format)
		}
	}
	gLog(logFormat, args...)
}
