package main

import (
	"encoding/json"
	"log"

	"github.com/karashiiro/gacha/message"
	"github.com/streadway/amqp"
)

func main() {
	m := &message.Message{
		Command: "roll",
	}

	mBytes, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer ch.Close()

	mq, err := ch.QueueDeclare("gacha", false, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = ch.Publish("", mq.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        mBytes,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
