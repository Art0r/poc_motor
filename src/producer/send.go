package producer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Art0r/poc_motor/src/common"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessage(queue_name string, msg Data) {
	common.WrapMessageAction(
		queue_name,
		func(queue *amqp.Queue, channel *amqp.Channel) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			body, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error marshaling JSON: %s", err)
				return
			}

			err = channel.PublishWithContext(ctx,
				"",         // exchange
				queue.Name, // routing key
				false,      // mandatory
				false,      // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        []byte(body),
				})

			if err != nil {
				log.Printf("Error publishing message: %s", err)
				return
			}

			log.Printf(" [x] Sent %s\n", msg)
		},
	)
}
