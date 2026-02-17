package consumer

import (
	"log"

	"github.com/Art0r/poc_motor/src/common"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Start() {

	common.WrapMessageAction(
		"hello",
		func(queue *amqp.Queue, channel *amqp.Channel) {
			msgs, err := channel.Consume(
				queue.Name, // queue
				"",         // consumer
				true,       // auto-ack
				false,      // exclusive
				false,      // no-local
				false,      // no-wait
				nil,        // args
			)
			failOnError(err, "Failed to register a consumer")

			var forever chan struct{}

			go func() {
				for d := range msgs {
					log.Printf("Received a message: %s", d.Body)
				}
			}()

			log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
			<-forever
		},
	)
}
