package iiq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	ctx  context.Context
	opts *Options
}

type Options struct {
	Domain     string
	BaseURL    string
	SiteId     string
	Token      string
	HTTPClient *http.Client
}

type Parameters struct {
	PageSize   int
	PageNumber int
	Filter     string
	OrderBy    string
	Display    string
}

func (p Parameters) encode() string {
	v := url.Values{}

	if p.PageSize != 0 {
		v.Add("$s", fmt.Sprintf("%d", p.PageSize))
	}

	if p.PageNumber != 0 {
		v.Add("$p", fmt.Sprintf("%d", p.PageNumber))
	}

	if p.OrderBy != "" {
		v.Add("$o", p.OrderBy)
	}

	if p.Display != "" {
		v.Add("$d", p.Display)
	}

	if p.Filter != "" {
		v.Add("$filter", p.Filter)
	}

	return v.Encode()
}

const DefaultBaseURL = "https://%s.incidentiq.com/api/v1.0%s"

func NewClient(options *Options) (*Client, error) {
	if options.Domain == "" {
		return nil, errors.New("a domain has not been provided but is required")
	}

	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}

	if options.BaseURL == "" {
		options.BaseURL = DefaultBaseURL
	}

	client := &Client{
		ctx:  context.Background(),
		opts: options,
	}

	return client, nil
}

func (c *Client) request(method, url string, data io.Reader, respData interface{}) (interface{}, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	request.Header.Add("SiteId", c.opts.SiteId)
	request.Header.Add("Authorization", "Bearer "+c.opts.Token)
	request.Header.Add("Client", "ApiClient")

	if data != nil {
		request.Header.Set("Accept", "application/json, text/plain, */*")
		request.Header.Add("Pragma", "no-cache")
		request.Header.Add("Cache-Control", "no-cache")
		request.Header.Add("Accept-Encoding", "gzip deflate, br")
		request.Header.Add("Accept-Language", "en-US,en;q=0.9")
		request.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.opts.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respData, json.Unmarshal(body, &respData)
}

func (c *Client) do(method, endpoint string, reqData, respData interface{}) (interface{}, error) {
	url := fmt.Sprintf(c.opts.BaseURL, c.opts.Domain, endpoint)

	if reqData != nil {
		b, err := json.Marshal(reqData)
		if err != nil {
			return nil, err
		}

		return c.request(method, url, strings.NewReader(string(b)), respData)
	}

	return c.request(method, url, nil, respData)
}

func (c *Client) get(endpoint string, respData interface{}) (interface{}, error) {
	return c.do(http.MethodGet, endpoint, nil, respData)
}

func (c *Client) post(endpoint string, reqData, respData interface{}) (interface{}, error) {
	return c.do(http.MethodPost, endpoint, reqData, respData)
}
