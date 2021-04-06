package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/karashiiro/gacha/message"
	"github.com/streadway/amqp"
)

func main() {
	dropsToAdd := []message.DropInsert{
		{
			ObjectID: 1,
			Rate:     0.2,
		},
		{
			ObjectID: 2,
			Rate:     0.2,
		},
		{
			ObjectID: 3,
			Rate:     0.2,
		},
		{
			ObjectID: 4,
			Rate:     0.4,
		},
	}

	dropsBytes, err := json.Marshal(dropsToAdd)
	if err != nil {
		log.Fatalln(err)
	}

	m := &message.Message{
		Command:    "set_drop_table",
		Parameters: []string{"test", string(dropsBytes)},
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

	mq, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	msgs, err := ch.Consume(mq.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	corrID := uuid.NewString()

	err = ch.Publish("", "gacha_v0", false, false, amqp.Publishing{
		ContentType:   "application/json",
		CorrelationId: corrID,
		ReplyTo:       mq.Name,
		Body:          mBytes,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for d := range msgs {
		if corrID == d.CorrelationId {
			if string(d.Body) != "Success" {
				log.Fatalln("Did not succeed")
			}
			break
		}
	}

	log.Println("Succeeded")
}
