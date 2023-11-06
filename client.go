package iiq

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
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
	PageSize  int
	PageIndex int
	Filter    string
	OrderBy   string
	Display   string
}

func (p Parameters) encode() string {
	v := url.Values{}

	if p.PageSize != 0 {
		v.Add("$s", fmt.Sprintf("%d", p.PageSize))
	}

	if p.PageIndex != 0 {
		v.Add("$p", fmt.Sprintf("%d", p.PageIndex))
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

func (c *Client) request(method, url string, data io.Reader, respData interface{}) (*m.Response[any], error) {
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

	response := &m.Response[any]{
		Header:     resp.Header.Clone(),
		StatusCode: resp.StatusCode,
		Body:       respData,
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, json.Unmarshal(body, &response.Body)
}

func (c *Client) do(method, endpoint string, reqData, respData interface{}) (*m.Response[any], error) {
	url := fmt.Sprintf(c.opts.BaseURL, c.opts.Domain, endpoint)

	if reqData != nil {
		b, err := json.Marshal(reqData)
		if err != nil {
			return nil, err
		}

		return c.request(method, url, bytes.NewReader(b), respData)
	}

	return c.request(method, url, nil, respData)
}

func (c *Client) get(endpoint string, respData interface{}) (*m.Response[any], error) {
	return c.do(http.MethodGet, endpoint, nil, respData)
}

func (c *Client) post(endpoint string, reqData, respData interface{}) (*m.Response[any], error) {
	return c.do(http.MethodPost, endpoint, reqData, respData)
}

func (c *Client) delete(endpoint string, reqData, respData interface{}) (*m.Response[any], error) {
	return c.do(http.MethodDelete, endpoint, reqData, respData)
}

func convertBody[T any](resp *m.Response[any], newResponse T) m.Response[T] {
	body, ok := resp.Body.(*T)
	if !ok {
		panic("unexpected type during conversion")
	}

	return m.Response[T]{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       *body,
	}
}
