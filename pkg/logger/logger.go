package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

const (
	traceKey = "TraceID"
)

var (
	logger   *zap.Logger
	initOnce sync.Once
)

func Init(profile string, config *Config) {
	initOnce.Do(func() {
		logger = NewZapLogger(profile, config)
	})
}

func Sync() {
	if logger == nil {
		return
	}
	_ = logger.Sync()
}

func Debug(c context.Context, msg string, fields ...zap.Field) {
	allFields := addTraceFields(c, fields...)
	logger.Debug(msg, allFields...)
}

func Debugf(c context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	var fields []zap.Field
	allFields := addTraceFields(c, fields...)
	logger.Debug(msg, allFields...)
}

func Info(c context.Context, msg string, fields ...zap.Field) {
	allFields := addTraceFields(c, fields...)
	logger.Info(msg, allFields...)
}

func Infof(c context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	var fields []zap.Field
	allFields := addTraceFields(c, fields...)
	logger.Info(msg, allFields...)
}

func Warn(c context.Context, msg string, fields ...zap.Field) {
	allFields := addTraceFields(c, fields...)
	logger.Warn(msg, allFields...)
}

func Warnf(c context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	var fields []zap.Field
	allFields := addTraceFields(c, fields...)
	logger.Warn(msg, allFields...)
}

func Error(c context.Context, msg string, fields ...zap.Field) {
	allFields := addTraceFields(c, fields...)
	logger.Error(msg, allFields...)
}

func Errorf(c context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	var fields []zap.Field
	allFields := addTraceFields(c, fields...)
	logger.Error(msg, allFields...)
}

func Panic(c context.Context, msg string, fields ...zap.Field) {
	allFields := addTraceFields(c, fields...)
	logger.Panic(msg, allFields...)
}

func Panicf(c context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	var fields []zap.Field
	allFields := addTraceFields(c, fields...)
	logger.Panic(msg, allFields...)
}

func Fatal(c context.Context, msg string, fields ...zap.Field) {
	allFields := addTraceFields(c, fields...)
	logger.Fatal(msg, allFields...)
}

func Fatalf(c context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	var fields []zap.Field
	allFields := addTraceFields(c, fields...)
	logger.Fatal(msg, allFields...)
}

func WithTrace(c context.Context, traceID string) context.Context {
	return context.WithValue(c, traceKey, traceID)
}

func addTraceFields(c context.Context, fields ...zap.Field) []zap.Field {
	if v := c.Value(traceKey); v != nil {
		if t, ok := v.(string); ok {
			var allFields []zap.Field
			if fields != nil {
				allFields = append(fields, zap.String(traceKey, t))
			} else {
				allFields = append(allFields, zap.String(traceKey, t))
			}

			return allFields
		}
	}

	return nil
}
