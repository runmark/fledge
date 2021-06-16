package main

import (
	"context"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/robfig/cron"
)

const (
	COMPONENT_PUBSUB_NAME = "pubsub"
)

func main() {

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	c:= cron.New()

	c.AddFunc("0/30 * * * * *", func() {
		NewFlParty().pub(ctx, client)
	})

	c.Start()

	select {}
	//for {}
}
