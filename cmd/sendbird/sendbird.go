package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)
	client := client.NewClient(
		client.WithLogger(logger),
	)

	fmt.Println(client)
}
