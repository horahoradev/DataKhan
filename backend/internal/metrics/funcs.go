package dkmetrics

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	"github.com/go-kit/kit/metrics/graphite"
	"logrus"
	"time"
)

func StartExporter(ctx context.Context) error {
	g := graphite.New("", nil)
	report := time.NewTicker(5 * time.Second)
	//defer report.Stop()
	go g.SendLoop(context.Background(), report.C, "tcp", "graphite:8125")
}

func MockCounter(name string) metrics.Counter {
	return discard.NewCounter()
}

func ConcreteCounter(name string) metrics.Counter {
	return graphite.NewCounter(name)
}
