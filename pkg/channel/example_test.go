package channel_test

import (
	"github.com/yumi-ia/sendbird-go/pkg/channel"
	"github.com/yumi-ia/sendbird-go/pkg/client"
)

func ExampleNewChannel() {
	// Initialize a client.
	opts := []client.Option{}
	c := client.NewClient(opts...)

	// Initialize a channel service.
	m := channel.NewChannel(c)

	// the channel client is ready to be used.
	_ = m
	// m.DoWork()
}
