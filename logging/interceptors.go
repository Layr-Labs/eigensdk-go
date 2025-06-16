package logging

import (
	"context"

	grpclogging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

func (s SLogger) InterceptorLogger() grpclogging.Logger {
	return grpclogging.LoggerFunc(func(ctx context.Context, lvl grpclogging.Level, msg string, fields ...any) {
		switch lvl {
		case grpclogging.LevelDebug:
			s.Debug(msg, fields...)
		case grpclogging.LevelInfo:
			s.Info(msg, fields...)
		case grpclogging.LevelWarn:
			s.Warn(msg, fields...)
		case grpclogging.LevelError:
			s.Error(msg, fields...)
		default:
			s.Info(msg, fields...)
		}
	})
}

func (z ZapLogger) InterceptorLogger() grpclogging.Logger {
	return grpclogging.LoggerFunc(func(ctx context.Context, lvl grpclogging.Level, msg string, fields ...any) {
		switch lvl {
		case grpclogging.LevelDebug:
			z.Debug(msg, fields...)
		case grpclogging.LevelInfo:
			z.Info(msg, fields...)
		case grpclogging.LevelWarn:
			z.Warn(msg, fields...)
		case grpclogging.LevelError:
			z.Error(msg, fields...)
		default:
			z.Info(msg, fields...)
		}
	})
}
