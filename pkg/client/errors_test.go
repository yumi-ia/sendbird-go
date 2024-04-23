package client

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		statusCode int
		body       string
		err        error
	}{
		{
			name: "not an error",
			body: `{"code":418,"message":"I'm a teapot","error":false}`,
		},
		{
			name: "no body",
			err:  io.EOF,
		},
		{
			name:       "unknown status code",
			statusCode: http.StatusTeapot,
			body:       `{"code":418,"message":"I'm a teapot","error":true}`,
			err:        ErrAPIDefault,
		},
		{
			name:       "bad request",
			statusCode: http.StatusBadRequest,
			body:       `{"code":418,"message":"I'm a teapot","error":true}`,
			err:        ErrAPIBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := &client{}
			c.SetDefault()

			err := c.handleError(test.statusCode, strings.NewReader(test.body))
			if err != nil {
				assert.ErrorIs(t, err, test.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
