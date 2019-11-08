package main

import (
	"net/http"

	"github.com/leominov/isdayoff_exporter/httpclient"
	"github.com/leominov/isdayoff_exporter/isdayoff"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type Exporter struct {
	httpCli     *http.Client
	isDayOff    prometheus.Gauge
	totalErrors prometheus.Counter
}

func NewExporter() *Exporter {
	return &Exporter{
		httpCli: httpclient.New(),
		isDayOff: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "isdayoff",
				Help: "Is day off for current date.",
			},
		),
		totalErrors: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "isdayoff_errors_total",
				Help: "Number of errors while requesting data.",
			},
		),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.isDayOff.Desc()
	ch <- e.totalErrors.Desc()
}

func (e *Exporter) collect() {
	isDayOff, err := isdayoff.IsDayOffToday(e.httpCli)
	if err != nil {
		logrus.Error(err)
		e.totalErrors.Inc()
	}

	if isDayOff {
		e.isDayOff.Set(1.0)
		return
	}

	e.isDayOff.Set(0.0)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.collect()

	ch <- e.isDayOff
	ch <- e.totalErrors
}
