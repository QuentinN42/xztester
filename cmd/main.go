package main

import (
	"context"

	"github.com/QuentinN42/xztester/pkg/logger"
)

func main() {
	ctx := context.Background()
	logger.Info(ctx, "Running")
}
