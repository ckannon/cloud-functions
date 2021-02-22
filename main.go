package cloudfunctions

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, "mud-host")
	if err != nil {
		log.Fatal(err)
	}

	sub := client.Subscription("stop-server-sub")
	exists, err := sub.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		sub, err = client.CreateSubscription(context.Background(), "stop-server-sub",
			pubsub.SubscriptionConfig{Topic: client.Topic("vm.instance.stop")})
	}

	err = sub.Receive(context.Background(), func(ctx context.Context, message *pubsub.Message) {
		m := PubSubMessage{Data: string(message.Data)}
		ProcessStopMessage(ctx, m)
	})
}
