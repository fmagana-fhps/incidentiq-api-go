package iiq

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Site struct {
	Domain string
	Id     string
	Token  string
}

type Options[T any] struct {
	Method   string
	Endpoint string
	Response T
}

const url = "https://%s.incidentiq.com/api/v1.0%s"

var client = &http.Client{Timeout: 10 * time.Second}

// The request in the function uses the POST method but we do
// not provide a body to show that we are not expecting changes
// to the endpoint. It is how the API works *eye roll* '|' for bulk
func do[T any](site Site, options Options[T]) (T, error) {
	url := fmt.Sprintf(url, site.Domain, options.Endpoint)

	request, err := http.NewRequest(options.Method, url, nil)
	if err != nil {
		return options.Response, err
	}

	request.Header.Add("SiteId", site.Id)
	request.Header.Add("Authorization", "Bearer "+site.Token)
	request.Header.Add("Client", "ApiClient")

	response, err := client.Do(request)
	if err != nil {
		return options.Response, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return options.Response, err
	}

	return options.Response, json.Unmarshal(body, &options.Response)
}

func get[T any](site Site, endpoint string) (T, error) {
	return do(site,
		Options[T]{
			Method:   http.MethodGet,
			Endpoint: endpoint,
		})
}
