package log

import (
	"context"
	"log/slog"
	"testing"
)

func Test(t *testing.T) {
	ctx := context.Background()
	t.Run("Logger", func(t *testing.T) {
		Global()

		slog.Log(ctx, LogLevelWarning, "Test")
	})
}
