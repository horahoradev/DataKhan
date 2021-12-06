package dkmetrics

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	"github.com/go-kit/kit/metrics/statsd"
	"os"
	"time"
)

var stats *statsd.Statsd

func StartExporter(ctx context.Context) {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = level.NewFilter(logger, level.AllowDebug())
	logger = log.With(logger,
		"svc", "order",
		"ts", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	stats = statsd.New("DataKhan", logger)
	report := time.NewTicker(5 * time.Second)
	//defer report.Stop()
	go stats.SendLoop(context.Background(), report.C, "udp", "graphite:8125")
}

func MockCounter(name string) metrics.Counter {
	return discard.NewCounter()
}

func ConcreteCounter(name string) metrics.Counter {
	return stats.NewCounter(name, 1.00)
}
