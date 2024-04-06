package main

import (
	"context"
	"flag"
	"os"

	"github.com/QuentinN42/xztester/pkg/logger"
	"github.com/QuentinN42/xztester/pkg/tester"
)

var (
	addr = flag.String("addr", "127.0.0.1:22", "ssh server address")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	logger.Info(ctx, "Testing the %s server", *addr)

	handshake, readkey, err := tester.Test(ctx, *addr)
	if err != nil {
		logger.Error(ctx, "Failed to test the server: %v", err)
		os.Exit(1)
	}

	logger.Info(ctx, "Server ping : %s", handshake)
	logger.Info(ctx, "Server time to read key : %s", readkey)
}
