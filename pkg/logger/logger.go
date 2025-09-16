package logger

import "log/slog"

func Info(msg string, args ...any) {
	slog.Info("PDFApiGenerator middleware> "+msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn("PDFApiGenerator middleware> "+msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error("PDFApiGenerator middleware> "+msg, args...)
}
