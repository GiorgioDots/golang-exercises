package intermediate

import (
	"fmt"
	"net/http"
	"time"
)

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	// Do "before sending requests" actions here.
	fmt.Printf("Sending request to %v\n", req.URL)

	start := time.Now()
	// Send the request, get the response (or the error)
	res, e = lrt.Proxied.RoundTrip(req)
	duration := time.Since(start)

	fmt.Printf("Request took %v\n", duration)

	// Handle the result.
	if e != nil {
		fmt.Printf("Error: %v", e)
	} else {
		fmt.Printf("Received %v response\n", res.Status)
	}

	return
}

func ExecuteHttpClient() {
	client := &http.Client{
		Transport: &LoggingRoundTripper{http.DefaultTransport},
	}
	client.Get("https://google.com")
}
