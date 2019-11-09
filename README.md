# isdayoff_exporter

Based on https://isdayoff.ru.

## Metrics

* `isdayoff`
* `isdayoff_errors_total`

## Environment variables

* `HTTP_PROXY`
* `HTTPS_PROXY`
* `NO_PROXY`

## Usage

```
Usage of ./isdayoff_exporter:
  -web.listen-address string
    	Address to listen on for web interface and telemetry. (default ":9393")
  -web.telemetry-path string
    	Path under which to expose metrics. (default "/metrics")
```
