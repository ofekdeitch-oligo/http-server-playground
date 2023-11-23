package metrics

import "github.com/prometheus/client_golang/prometheus"

type MetricsModule struct {
	Registry *prometheus.Registry
	Metrics  MetricDefinitions
}

type MetricDefinitions struct {
	Count prometheus.Counter
}

func New() MetricsModule {
	reg := prometheus.NewRegistry()
	metrics := registerPrometheus(reg)

	return MetricsModule{
		Registry: reg,
		Metrics:  metrics,
	}
}

func registerPrometheus(reg *prometheus.Registry) MetricDefinitions {
	countMetric := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "playground_count_total",
	})

	reg.MustRegister(countMetric)
	return MetricDefinitions{Count: countMetric}
}
