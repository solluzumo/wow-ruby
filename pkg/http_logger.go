package pkg

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	status int
}

func (w *ResponseWriterWrapper) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rw := &ResponseWriterWrapper{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			logger.Info("request started",
				zap.String("method", r.Method),
				zap.String("url", r.URL.Path),
			)

			defer func() {
				if rec := recover(); rec != nil {
					logger.Error("panic recovered",
						zap.Any("error", rec),
						zap.String("method", r.Method),
						zap.String("url", r.URL.Path),
					)
					http.Error(rw, "internal server error", http.StatusInternalServerError)
				}

				duration := time.Since(start)
				logger.Info("request finished",
					zap.String("method", r.Method),
					zap.String("url", r.URL.Path),
					zap.Int("status", rw.status),
					zap.Duration("duration", duration),
				)
			}()

			next.ServeHTTP(rw, r)
		})
	}
}

func NewLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	logsDirName := fmt.Sprintf("logs/%s", os.Getenv("SERVICE_NAME"))
	os.MkdirAll(logsDirName, os.ModePerm)
	cfg.OutputPaths = []string{"stdout", logsDirName + "/service.log"}
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to create logger: %v", err))
	}
	return logger
}
