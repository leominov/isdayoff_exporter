package httpclient

import (
	"net/http"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

func New() *http.Client {
	cli := cleanhttp.DefaultPooledClient()
	cli.Transport = &userAgentRoundTripper{
		userAgent: "isdayoff_exporter (+https://github.com/leominov/isdayoff_exporter)",
		inner:     cli.Transport,
	}
	return cli
}
