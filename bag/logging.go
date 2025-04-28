package main

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

var (
	loglevel slog.LevelVar
)

type logHandler struct {
}

func (lh *logHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= loglevel.Level()
}

func (lh *logHandler) Handle(_ context.Context, r slog.Record) error {
	values := []string{}
	r.Attrs(
		func(a slog.Attr) bool {
			var val string
			switch t := a.Value.Any().(type) {
			case slog.LogValuer:
				val = t.LogValue().String()
				values = append(values, fmt.Sprintf("%s=%s", a.Key, val))
			default:
				t = a.Value
				val = a.Value.String()
				values = append(values, fmt.Sprintf("%s=%s", a.Key, a.Value))
			}
			return true
		})
	lparen := ""
	rparen := ""
	if len(values) > 0 {
		lparen = "("
		rparen = ")"
	}
	fmt.Println(r.Message, lparen+strings.Join(values, ",")+rparen)
	return nil
}

func (lh *logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return lh
}

func (lh *logHandler) WithGroup(name string) slog.Handler {
	return lh
}

func setupLog() {
	sl := slog.New(&logHandler{})
	slog.SetDefault(sl)
	return
}
