package main

import (
	"os"

	"github.com/Art0r/poc_motor/src/consumer"
	"github.com/Art0r/poc_motor/src/producer"
)

func main() {

	app := os.Getenv("APP")

	if app == "consumer" {
		consumer.Start()
		return
	}

	if app == "producer" {
		producer.Start()
		return
	}
}
