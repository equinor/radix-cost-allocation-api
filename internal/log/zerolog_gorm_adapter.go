package log

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

// Copied from https://github.com/mpalmer/gorm-zerolog/blob/master/logger.go @ 27.02.2024

type ZerologGormAdapterLogger struct {
}

func (l ZerologGormAdapterLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l ZerologGormAdapterLogger) Error(ctx context.Context, msg string, opts ...any) {
	zerolog.Ctx(ctx).Error().Msgf(msg, opts...)
}

func (l ZerologGormAdapterLogger) Warn(ctx context.Context, msg string, opts ...any) {
	zerolog.Ctx(ctx).Warn().Msgf(msg, opts...)
}

func (l ZerologGormAdapterLogger) Info(ctx context.Context, msg string, opts ...any) {
	zerolog.Ctx(ctx).Info().Msgf(msg, opts...)
}

func (l ZerologGormAdapterLogger) Trace(ctx context.Context, begin time.Time, f func() (string, int64), err error) {
	var level = zerolog.TraceLevel

	if err != nil {
		level = zerolog.DebugLevel
	}

	sql, rows := f()
	zerolog.Ctx(ctx).WithLevel(level).
		Dur("elapsed_ms", time.Since(begin)).
		Int64("rows", rows).
		Str("sql", sql).
		Send()
}
