package otel

import "log/slog"

var opts *slog.HandlerOptions = &slog.HandlerOptions{
	Level: slog.LevelDebug,
}

var log *slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
slog.SetDefault(log)
