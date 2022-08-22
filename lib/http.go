package lib

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RequestParameters struct {
	Url     string
	Method  string
	Body    string
	Headers types.Map
}

func (c *HttpClient) Request(request RequestParameters) *http.Response {

	requestBody := strings.NewReader(request.Body)
	httpRequest, err := http.NewRequestWithContext(context.TODO(), request.Method, request.Url, requestBody)

	if err != nil {
		log.Fatal(err)
	}

	for name, value := range request.Headers.Elems {
		var headerValue string
		tfsdk.ValueAs(context.TODO(), value, &headerValue)
		httpRequest.Header.Set(name, headerValue)
	}

	resp, err := c.client.Do(httpRequest)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}
