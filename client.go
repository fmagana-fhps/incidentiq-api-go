package iiq

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type IncidentIQ struct {
	Domain string
	SiteID string
	Token  string
}

const url = "https://%s.incidentiq.com/api/v1.0%s"

var client = &http.Client{Timeout: 10 * time.Second}

// The request in the function uses the POST method but we do
// not provide a body to show that we are not expecting changes
// to the endpoint. It is how the API works *eye roll*
func (iiq IncidentIQ) doGet(endpoint string, model interface{}) error {
	url := fmt.Sprintf(url, iiq.Domain, endpoint)

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	request.Header.Add("SiteId", iiq.SiteID)
	request.Header.Add("Authorization", "Bearer "+iiq.Token)
	request.Header.Add("Client", "ApiClient")

	response, err := client.Do(request)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &model)
}
