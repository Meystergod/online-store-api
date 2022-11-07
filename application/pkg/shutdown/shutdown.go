package shutdown

import (
	"context"
	"io"
	"os"
	"os/signal"

	"github.com/Meystergod/online-store-api/pkg/logging"
)

func Graceful(ctx context.Context, signals []os.Signal, closeItems ...io.Closer) {
	logger := logging.GetLogger(ctx)

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, signals...)
	sig := <-sigChannel
	logger.Infof("caught signal %s. shutting down...", sig)

	for _, item := range closeItems {
		if err := item.Close(); err != nil {
			logger.Errorf("failed to close %v: %v", item, err)
		}
	}
}
