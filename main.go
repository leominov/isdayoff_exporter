package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	version string
	commit  string
	date    string

	listenAddress = flag.String("web.listen-address", ":9393", "Address to listen on for web interface and telemetry.")
	metricPath    = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	printsVersion = flag.Bool("version", false, "Prints version and exit.")
)

func main() {
	flag.Parse()

	if *printsVersion {
		fmt.Printf("version=%s\ncommit=%s\ndate=%s\n", version, commit, date)
		return
	}

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
			<h1>IsDayOff Exporter ` + version + `</h1>
			<p><a href='` + *metricPath + `'>Metrics</a></p>
			</body>
			</html>
		`))
	})

	logrus.Printf("Providing metrics at %s%s", *listenAddress, *metricPath)
	logrus.Fatal(http.ListenAndServe(*listenAddress, nil))
}
