package pkg

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// Создание логгера: файл + консоль
func NewZapLogger() *zap.Logger {
	logsDirName := fmt.Sprintf("logs/%s", os.Getenv("SERVICE_NAME"))
	os.MkdirAll(logsDirName, os.ModePerm)
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout", logsDirName + "/service.log"}

	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build logger: %v", err))
	}

	return logger
}

// Unary interceptor с функционалом как HTTP-логгер
func ZapUnaryInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		start := time.Now()
		requestID := uuid.New().String()

		logger.Info("gRPC request started",
			zap.String("method", info.FullMethod),
			zap.String("request_id", requestID),
		)

		// panic-safe
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error("panic recovered",
					zap.Any("panic", rec),
					zap.String("method", info.FullMethod),
					zap.String("request_id", requestID),
				)
				err = status.Error(13, "internal server error") // INTERNAL
			}

			duration := time.Since(start)

			if err != nil {
				st, _ := status.FromError(err)

				logger.Error("gRPC request finished with error",
					zap.String("method", info.FullMethod),
					zap.String("request_id", requestID),
					zap.String("status", st.Code().String()),
					zap.String("error", st.Message()),
					zap.Duration("duration", duration),
				)
				return
			}

			logger.Info("gRPC request finished",
				zap.String("method", info.FullMethod),
				zap.String("request_id", requestID),
				zap.String("status", "OK"),
				zap.Duration("duration", duration),
			)
		}()

		return handler(ctx, req)
	}
}
