package tester

import (
	"context"
	"time"
)

func Test(ctx context.Context, addr string) (time.Duration, time.Duration, error) {
	return time.Second, time.Second, nil
}
