package client

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
)

// client is the implementation of the Client interface.
type client struct {
	// logger is the logger of the client.
	logger *slog.Logger
	// httpClient is the http client of the client.
	httpClient *http.Client
	// baseURL is the base URL of the sendbird API.
	// See https://sendbird.com/docs/chat/platform-api/v3/prepare-to-use-api
	baseURL *url.URL
	// header is the header of the client.
	// See https://sendbird.com/docs/chat/platform-api/v3/prepare-to-use-api#2-headers
	header http.Header
}

func (c *client) SetDefault() {
	c.logger = slog.Default()
	c.httpClient = http.DefaultClient
	c.baseURL = &url.URL{
		Scheme: "https",
		Path:   "/v3",
	}
	c.header = http.Header{}
	c.header.Set("Content-Type", "application/json; charset=utf-8")
}

// Option is the interface for the options of the client.
type Option func(client *client) *client

// WithLogger is the option for the logger of the client.
func WithLogger(l *slog.Logger) Option {
	return func(client *client) *client {
		client.logger = l

		return client
	}
}

// WithHTTPClient is the option for the http client of the client.
func WithHTTPClient(c *http.Client) Option {
	return func(client *client) *client {
		client.httpClient = c

		return client
	}
}

// WithScheme is the option for the scheme of the client.
func WithScheme(scheme string) Option {
	return func(client *client) *client {
		client.baseURL.Scheme = scheme

		return client
	}
}

// WithHost is the option for the host of the client.
func WithHost(host string) Option {
	return func(client *client) *client {
		client.baseURL.Host = host

		return client
	}
}

// WithPath is the option for the path of the client.
func WithPath(path string) Option {
	return func(client *client) *client {
		client.baseURL.Path = path

		return client
	}
}

// WithURL is the option for the url of the client.
func WithURL(u string) Option {
	return func(client *client) *client {
		baseURL, err := url.Parse(u)
		if err != nil {
			panic(fmt.Sprintf("invalid url: %v", err))
		}

		client.baseURL = baseURL

		return client
	}
}

// WithAPIToken is the option for the api-token of the client.
func WithAPIToken(apiToken string) Option {
	return func(client *client) *client {
		client.header.Set("Api-Token", apiToken)

		return client
	}
}
