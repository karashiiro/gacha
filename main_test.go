package main_test

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/karashiiro/gacha/message"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestRoll(t *testing.T) {
	m := &message.Message{
		Command:    "roll",
		Parameters: []string{"test"},
	}

	mBytes, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := amqp.Dial(os.Getenv("GACHA_RMQ_CONNECTION_STRING"))
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

	var roll int
	for d := range msgs {
		if corrID == d.CorrelationId {
			roll, err = strconv.Atoi(string(d.Body))
			if err != nil {
				log.Fatalln(err)
			}
			break
		}
	}

	log.Printf("Rolled object with ID: %d", roll)
}
