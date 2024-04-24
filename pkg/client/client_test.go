package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	c := NewClient()
	cClient, ok := c.(*client)
	require.True(t, ok)
	assert.Equal(t, slog.Default(), cClient.logger)
}

func TestGet(t *testing.T) {
	t.Parallel()

	type Foo struct {
		Foo string `json:"foo"`
	}

	tests := []struct {
		name         string
		req          *http.Request
		body         any
		responseBody any
		statusCode   int
		expectedPath string
		expectedBody string
		expectedErr  error
		expectedResp any
	}{
		{
			name:         "default",
			req:          httptest.NewRequest(http.MethodGet, "http://example.com/foo/bar", nil),
			expectedPath: "/foo/bar",
		},
		{
			name:         "with body",
			req:          httptest.NewRequest(http.MethodGet, "http://example.com/foo/bar", nil),
			expectedPath: "/foo/bar",
			body:         Foo{Foo: "bar"},
			expectedBody: `{"foo":"bar"}`,
		},
		{
			name:         "with API error",
			req:          httptest.NewRequest(http.MethodGet, "http://example.com/foo/bar", nil),
			statusCode:   http.StatusTeapot,
			expectedPath: "/foo/bar",
			responseBody: Error{Code: 418, Message: "I'm a teapot", Error: true},
			expectedErr:  ErrAPIDefault,
		},
		{
			name:         "with API response",
			req:          httptest.NewRequest(http.MethodGet, "http://example.com/foo/bar", nil),
			statusCode:   http.StatusOK,
			expectedPath: "/foo/bar",
			responseBody: Foo{Foo: "bar"},
			expectedResp: &Foo{Foo: "bar"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, test.req.Method, r.Method)
				assert.Equal(t, test.expectedPath, r.URL.Path)

				if test.expectedBody != "" {
					defer r.Body.Close()
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.Equal(t, test.expectedBody, string(body))
				}

				statusCode := http.StatusOK
				if test.statusCode != 0 {
					statusCode = test.statusCode
				}

				w.WriteHeader(statusCode)

				if test.responseBody != nil {
					err := json.NewEncoder(w).Encode(test.responseBody)
					require.NoError(t, err)
				}
			}))
			defer s.Close()

			httpClient := &http.Client{
				Timeout: time.Second, // to make the test fail faster, might be flaky
			}

			c := NewClient(
				WithURL(s.URL),
				WithHTTPClient(httpClient),
			)

			var body any
			if test.expectedResp != nil {
				body = new(Foo)
			}

			b, err := c.Get(test.req.Context(), test.req.URL.Path, test.body, body)
			if test.expectedErr != nil {
				assert.ErrorIs(t, err, test.expectedErr)
			} else {
				assert.NoError(t, err)
			}

			if test.expectedResp != nil {
				bb, ok := b.(*Foo)
				require.True(t, ok)
				assert.Equal(t, test.expectedResp, bb)
			}
		})
	}
}

func TestPOST(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/foo/bar", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	}))
	defer s.Close()

	httpClient := &http.Client{
		Timeout: time.Second, // to make the test fail faster, might be flaky
	}

	c := NewClient(
		WithURL(s.URL),
		WithHTTPClient(httpClient),
	)

	_, err := c.Post(context.Background(), "/foo/bar", nil, nil)
	require.NoError(t, err)
}

func TestPUT(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/foo/bar", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	}))
	defer s.Close()

	httpClient := &http.Client{
		Timeout: time.Second, // to make the test fail faster, might be flaky
	}

	c := NewClient(
		WithURL(s.URL),
		WithHTTPClient(httpClient),
	)

	_, err := c.Put(context.Background(), "/foo/bar", nil, nil)
	require.NoError(t, err)
}

func TestDelete(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/foo/bar", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	}))
	defer s.Close()

	httpClient := &http.Client{
		Timeout: time.Second, // to make the test fail faster, might be flaky
	}

	c := NewClient(
		WithURL(s.URL),
		WithHTTPClient(httpClient),
	)

	_, err := c.Delete(context.Background(), "/foo/bar", nil, nil)
	require.NoError(t, err)
}

func TestGetURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		req         *http.Request
		expectedURL string
		expectedLog string
	}{
		{
			name: "needs to encode url",
			req: &http.Request{
				Method: http.MethodGet,
				URL:    &url.URL{Path: "/ !@#$%^&*()_+{}|:<>?"},
			},
			expectedURL: "https://example.com/%2520%2521@%2523$%2525%255E&%252A%2528%2529_+%257B%257D%257C:%253C%253E%253F",
			expectedLog: "failed to parse path",
		},
		{
			name: "with query params",
			req: &http.Request{
				Method: http.MethodGet,
				URL:    &url.URL{Path: "/foo?bar=baz"},
			},
			expectedURL: "https://example.com/foo?bar=baz",
		},
		{
			name: "with query params to encode",
			req: &http.Request{
				Method: http.MethodGet,
				URL:    &url.URL{Path: "/foo?bar=baz&b!z=b>z"},
			},
			expectedURL: "https://example.com/foo?b%21z=b%3Ez&bar=baz",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var b bytes.Buffer
			logger := slog.New(slog.NewTextHandler(&b, &slog.HandlerOptions{Level: slog.LevelDebug}))

			c := &client{}
			c.SetDefault()
			c = WithHost("example.com")(c)
			c = WithPath("")(c)
			c = WithLogger(logger)(c)

			u := c.getURL(test.req.URL.Path)
			assert.Equal(t, test.expectedURL, u.String())
			assert.Contains(t, b.String(), test.expectedLog)
		})
	}
}
