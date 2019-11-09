package main

import (
	"net/http"

	"github.com/leominov/isdayoff_exporter/httpclient"
	"github.com/leominov/isdayoff_exporter/isdayoff"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

const exporter = "isdayoff"

type Exporter struct {
	httpCli      *http.Client
	IsDayOff     prometheus.Gauge
	ScrapeErrors prometheus.Counter
	Error        prometheus.Gauge
}

func NewExporter() *Exporter {
	namespace := exporter
	return &Exporter{
		httpCli: httpclient.New(),
		IsDayOff: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: exporter,
				Help: "Is day off for current date (1 for day off).",
			},
		),
		ScrapeErrors: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "scrape_errors_total",
				Help:      "Total number of times an error occurred.",
			},
		),
		Error: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "last_scrape_error",
				Help:      "Whether the last scrape of metrics from isdayoff.ru resulted in an error (1 for error, 0 for success).",
			},
		),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.IsDayOff.Desc()
	ch <- e.ScrapeErrors.Desc()
	ch <- e.Error.Desc()
}

func (e *Exporter) collect() {
	isDayOff, err := isdayoff.IsDayOffToday(e.httpCli)
	if err != nil {
		logrus.Error(err)
		e.ScrapeErrors.Inc()
		e.Error.Set(1.0)
	} else {
		e.Error.Set(0.0)
	}

	if isDayOff {
		e.IsDayOff.Set(1.0)
		return
	}

	e.IsDayOff.Set(0.0)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.collect()

	ch <- e.IsDayOff
	ch <- e.ScrapeErrors
	ch <- e.Error
}
