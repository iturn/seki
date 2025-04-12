package seki

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func NewLogger() *slog.Logger {
	var output slog.Handler

	loggerOps := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			// format time on time key
			if attr.Key == slog.TimeKey {
				attr.Value = slog.StringValue(attr.Value.Time().Format("15:04:05"))
			}

			// replace log json key not in use
			// if attr.Key == slog.MessageKey {
			// 	attr.Key = "message"
			// }

			return attr
		},
	}

	if os.Getenv("ENVIROMENT") == "development" {
		// override log output json in development to pretty printed text
		output = slog.NewTextHandler(os.Stdout, loggerOps)
	} else {
		output = slog.NewJSONHandler(os.Stdout, loggerOps)
	}

	logger := slog.New(output)

	slog.SetDefault(logger)

	Logger = logger

	return logger
}
