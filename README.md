# sendbird-go

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/tomMoulard/sendbird-go)](https://github.com/tomMoulard/sendbird-go/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/tomMoulard/sendbird-go.svg)](https://pkg.go.dev/github.com/tomMoulard/sendbird-go)
[![go.mod](https://img.shields.io/github/go-mod/go-version/tomMoulard/sendbird-go)](go.mod)
[![LICENSE](https://img.shields.io/github/license/tomMoulard/sendbird-go)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/tomMoulard/sendbird-go/build.yml?branch=main)](https://github.com/tomMoulard/sendbird-go/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomMoulard/sendbird-go)](https://goreportcard.com/report/github.com/tomMoulard/sendbird-go)

Yet another go client for the [Sendbird](https://sendbird.com) chat API.

⭐ `Star` this repository if you find it valuable and worth maintaining.

👁 `Watch` this repository to get notified about new releases, issues, etc.

## User Guide

### Installation

```bash
go get github.com/tomMoulard/sendbird-go
```

### Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/tomMoulard/sendbird-go/pkg/client"
    "github.com/tomMoulard/sendbird-go/pkg/user"
)

func main() {
    opts := []client.Option{
        client.WithAPPID(os.Getenv("SENDBIRD_APP_ID")),
        client.WithAPIToken(os.Getenv("SENDBIRD_API_KEY")),
    }
    client := client.NewClient(opts...)

    // To create a user
    userClient := user.NewUser(client)

    u := user.CreateUserRequest{
        UserID: "user-id",
        Nickname: "nickname",
    }
    createdUser, err := userClient.Create(context.Background(), u)
    if err != nil {
        if errors.Is(err, client.ErrTooManyRequests) {
            log.Fatalf("rate limit exceeded: %v", err)
            return
        }
        log.Fatalf("failed to create user: %v", err)
    }
    fmt.Printf("created user: %v\n", createdUser)
}
```

See available methods in the corresponding package documentation.

### Usage in tests

See [the source](./pkg/message/message_test.go) for the full example.

```go
// iencli wraps the message.Message interface.
type iencli struct {
	message message.Message
}

// SendMessage wraps the sends message method.
func (c *iencli) SendMessage(ctx context.Context, channelType message.ChannelType, channelURL string, sendMessageRequest message.SendMessageRequest) (*message.SendMessageResponse, error) {
	got, err := c.message.SendMessage(ctx, channelType, channelURL, sendMessageRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	return got, nil
}

func TestSendMessage(t *testing.T) {
	t.Parallel()

	req := message.SendMessageRequest{
		Message: "hello",
	}
	messageMock := message.NewMessageMock(t).
		OnSendMessage(message.ChannelTypeGroup, "channelURL", req).TypedReturns(&message.SendMessageResponse{}, nil).Once().
		Parent

	c := &iencli{
		message: messageMock,
	}

	_, err := c.SendMessage(context.Background(), message.ChannelTypeGroup, "channelURL", req)
	if err != nil {
		t.Fatalf("failed to send message: %v", err)
	}
}
```

## Build

### Terminal

- `make` - execute the build pipeline.
- `make help` - print help for the [Make targets](Makefile).


## Contributing

Feel free to create an issue or propose a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
