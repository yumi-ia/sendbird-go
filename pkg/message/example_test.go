package message_test

import (
	"github.com/yumi-ia/sendbird-go/pkg/client"
	"github.com/yumi-ia/sendbird-go/pkg/message"
)

func ExampleNewMessage() {
	// Initialize a client.
	opts := []client.Option{}
	c := client.NewClient(opts...)

	// Initialize a message service.
	m := message.NewMessage(c)

	// the message client is ready to be used.
	_ = m
	// m.DoWork()
}
