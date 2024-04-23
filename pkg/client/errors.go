package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrAPIDefault = errors.New("API error")

	ErrAPIBadRequest = errors.New("bad request")

	ErrAPIForbidden = errors.New("forbidden")

	ErrAPITooManyRequests = errors.New("too many requests")

	ErrAPIInternalServerError = errors.New("internal server error")

	ErrAPIServiceUnavailable = errors.New("service unavailable")
)

// https://sendbird.com/docs/chat/platform-api/v3/error-codes#2-error-codes
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	IsError bool   `json:"error"` //nolint: tagliatelle
}

func (e *Error) handleBadRequestError() error {
	return ErrAPIBadRequest
}

func (e *Error) handleForbidenError() error {
	panic("not implemented")
}

func (e *Error) handleTooManyRequestsError() error {
	panic("not implemented")
}

func (e *Error) handleInternalServerError() error {
	panic("not implemented")
}

func (e *Error) handleServiceUnavailableError() error {
	panic("not implemented")
}

func (c *client) handleError(status int, body io.Reader) error {
	var handledError Error
	if err := json.NewDecoder(body).Decode(&handledError); err != nil {
		b, _ := io.ReadAll(body)
		c.logger.Warn("failed to decode error", "body", string(b))

		return fmt.Errorf("failed to decode error: %w", err)
	}

	if !handledError.IsError {
		c.logger.Warn("error is not an error", "handledError", handledError)

		return nil
	}

	switch status {
	case http.StatusBadRequest:
		return handledError.handleBadRequestError()
	case http.StatusForbidden:
		return handledError.handleForbidenError()
	case http.StatusTooManyRequests:
		return handledError.handleTooManyRequestsError()
	case http.StatusInternalServerError:
		return handledError.handleInternalServerError()
	case http.StatusServiceUnavailable:
		return handledError.handleServiceUnavailableError()
	}

	return fmt.Errorf("unknown error %w: status: %d, code: %d, message: %q",
		ErrAPIDefault, status, handledError.Code, handledError.Message)
}
