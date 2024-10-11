package lago

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	DefaultBaseURL       = "https://api.getlago.com"
	DefaultBaseIngestURL = "https://ingest.getlago.com"
	ApiV1Path            = "/api/v1/"
)

// HTTPClient is an interface that only takes the Do method from http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is a struct that holds the configuration for the client.
// Zero values are not valid. Use New to create a new client.
// The client is safe for concurrent use by multiple goroutines.
//
// TODO(nikola-jokic): Should defaults be applied and functional options used?
// For now, I think it might be better to force specifying all the fields.
type Client struct {
	baseURL          *url.URL
	baseIngestURL    *url.URL
	useIngestService bool
	debug            bool
	userAgent        string
	bearerAuth       string
	client           HTTPClient
}

func New(cfg Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	baseURL, _ := url.Parse(cfg.BaseURL)
	baseIngestURL, _ := url.Parse(cfg.BaseIngestURL)

	return &Client{
		baseURL:          baseURL,
		baseIngestURL:    baseIngestURL,
		useIngestService: cfg.UseIngestService,
		debug:            cfg.Debug,
		userAgent:        "lago-go github.com/nikola-jokic/lago-go",
		bearerAuth:       "Bearer " + cfg.APIKey,
		client:           cfg.Client,
	}, nil
}

func get[R any](ctx context.Context, c *Client, path string) (*R, *Error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		path,
		nil,
	)
	if err != nil {
		return nil, &Error{Err: err}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Authorization", c.bearerAuth)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, &Error{Err: err}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, readError(res.Body)
	}

	var result R
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, &Error{Err: err}
	}

	return &result, nil
}

func delete[R any](ctx context.Context, c *Client, path string) (*R, *Error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		path,
		nil,
	)
	if err != nil {
		return nil, &Error{Err: err}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Authorization", c.bearerAuth)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, &Error{Err: err}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, readError(res.Body)
	}

	var result R
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, &Error{Err: err}
	}

	return &result, nil
}

func post[B, R any](ctx context.Context, client *Client, path string, body *B) (*R, *Error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, &Error{Err: err}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		path,
		&buf,
	)
	if err != nil {
		return nil, &Error{Err: err}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", client.userAgent)
	req.Header.Set("Authorization", client.bearerAuth)

	res, err := client.client.Do(req)
	if err != nil {
		return nil, &Error{Err: err}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, readError(res.Body)
	}

	var result R
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, &Error{Err: err}
	}

	return &result, nil
}

func postWithoutBody[R any](ctx context.Context, client *Client, path string) (*R, *Error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		path,
		nil,
	)
	if err != nil {
		return nil, &Error{Err: err}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", client.userAgent)
	req.Header.Set("Authorization", client.bearerAuth)

	res, err := client.client.Do(req)
	if err != nil {
		return nil, &Error{Err: err}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, readError(res.Body)
	}

	var result R
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, &Error{Err: err}
	}

	return &result, nil
}

func put[B, R any](ctx context.Context, client *Client, path string, body *B) (*R, *Error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, &Error{Err: err}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		path,
		&buf,
	)
	if err != nil {
		return nil, &Error{Err: err}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", client.userAgent)
	req.Header.Set("Authorization", client.bearerAuth)

	res, err := client.client.Do(req)
	if err != nil {
		return nil, &Error{Err: err}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, readError(res.Body)
	}

	var result R
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, &Error{Err: err}
	}

	return &result, nil
}

func putWithoutBody[R any](ctx context.Context, client *Client, path string) (*R, *Error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		path,
		nil,
	)
	if err != nil {
		return nil, &Error{Err: err}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", client.userAgent)
	req.Header.Set("Authorization", client.bearerAuth)

	res, err := client.client.Do(req)
	if err != nil {
		return nil, &Error{Err: err}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, readError(res.Body)
	}

	var result R
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, &Error{Err: err}
	}

	return &result, nil
}

func (c *Client) url(path string, q url.Values) string {
	u := *c.baseURL
	u.Path = ApiV1Path + path
	u.RawQuery = q.Encode()
	return u.String()
}
