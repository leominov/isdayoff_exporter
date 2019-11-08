package main

import (
	"flag"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	listenAddress = flag.String("web.listen-address", ":9393", "Address to listen on for web interface and telemetry.")
	metricPath    = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
)

func main() {
	flag.Parse()

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Starting IsDayOff Exporter...")

	exp := NewExporter()
	err := prometheus.Register(exp)
	if err != nil {
		logrus.Fatal(err)
	}

	http.Handle(*metricPath, promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
			<head><title>IsDayOff Exporter</title></head>
			<body>
			<h1>IsDayOff Exporter</h1>
			<p><a href='` + *metricPath + `'>Metrics</a></p>
			</body>
			</html>
		`))
	})

	logrus.Printf("Providing metrics at %s%s", *listenAddress, *metricPath)
	logrus.Fatal(http.ListenAndServe(*listenAddress, nil))
}