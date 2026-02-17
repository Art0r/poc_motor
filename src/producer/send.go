package producer

import (
	"context"
	"log"
	"time"

	"github.com/Art0r/poc_motor/src/common"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessage(queue_name string, msg_text string) {
	common.WrapMessageAction(
		queue_name,
		func(queue *amqp.Queue, channel *amqp.Channel) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			channel.PublishWithContext(ctx,
				"",         // exchange
				queue.Name, // routing key

				false, // mandatory
				false, // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(msg_text),
				})
			log.Printf(" [x] Sent %s\n", msg_text)
		},
	)
}
