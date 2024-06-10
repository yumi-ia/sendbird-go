package client_test

import (
	"log/slog"
	"os"

	"github.com/yumi-ia/sendbird-go/pkg/client"
)

func ExampleNewClient() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	opts := []client.Option{
		client.WithHost("example.com"),
		client.WithPath("v3"),
		client.WithLogger(logger),
		client.WithAPPID(os.Getenv("SENDBIRD_APP_ID")),
		client.WithAPIToken(os.Getenv("SENDBIRD_API_KEY")),
	}
	c := client.NewClient(opts...)

	// the client is ready to be used.
	_ = c
	// c.DoWork()
}
