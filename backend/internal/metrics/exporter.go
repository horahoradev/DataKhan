package metrics

import (
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/kit/metrics/prometheus"
)

// Abstraction layer over go-kit metrics, used for testing
type MetricsExporter interface {
	Gauge(namespace, subsystem, name, help string) prometheus.Gauge
}

type ConcExporter struct {
}

func (c *ConcExporter) Gauge(namespace, subsystem, name, help string, constLabels map[string]string, runtimeLabels []string) *prometheus.Gauge {
	return prometheus.NewGauge(stdprometheus.NewGaugeVec(stdprometheus.GaugeOpts{
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}, runtimeLabels))
}

type MockExporter struct {
}
