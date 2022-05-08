package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	TotalHits prometheus.Counter
	Hits      *prometheus.CounterVec
	Errors    *prometheus.CounterVec
	Durations *prometheus.HistogramVec
}

func CreateNewMetric(name string) (*Metric, error) {
	var m Metric

	m.TotalHits = prometheus.NewCounter(prometheus.CounterOpts{
		Name: name + "_total",
	})

	if err := prometheus.Register(m.TotalHits); err != nil {
		return nil, err
	}

	m.Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: name + "_hits",
	}, []string{"status", "method", "path"},
	)
	if err := prometheus.Register(m.Hits); err != nil {
		return nil, err
	}

	m.Errors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: name + "_errors",
	}, []string{"status", "method", "path"},
	)
	if err := prometheus.Register(m.Errors); err != nil {
		return nil, err
	}

	m.Durations = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: name + "_durations",
	}, []string{"status", "method", "path"},
	)
	if err := prometheus.Register(m.Durations); err != nil {
		return nil, err
	}

	return &m, nil
}
