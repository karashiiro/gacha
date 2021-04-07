package main_test

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/karashiiro/gacha/message"
	"github.com/streadway/amqp"
)

func TestSetTable(t *testing.T) {
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
		t.Fatal(err)
	}

	m := &message.Message{
		Command:    "set_drop_table",
		Parameters: []string{"test", string(dropsBytes)},
	}

	mBytes, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}

	var conn *amqp.Connection
	for conn == nil {
		conn, err = amqp.Dial(os.Getenv("GACHA_RMQ_CONNECTION_STRING"))
		if err != nil {
			log.Println("couldn't open RabbitMQ connection, retrying in 5 seconds")
			time.Sleep(5 * time.Second)
		}
	}
	log.Println("opened RabbitMQ connection")
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("opened RabbitMQ channel")
	defer ch.Close()

	mq, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	msgs, err := ch.Consume(mq.Name, "", false, false, false, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	corrID := uuid.NewString()

	err = ch.Publish("", os.Getenv("GACHA_RMQ_CHANNEL"), false, false, amqp.Publishing{
		ContentType:   "application/json",
		CorrelationId: corrID,
		ReplyTo:       mq.Name,
		Body:          mBytes,
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Println("published application request")

	for d := range msgs {
		if corrID == d.CorrelationId {
			if string(d.Body) != "Success" {
				t.Fatalf("Did not succeed")
			}
			break
		}
	}

	log.Println("Succeeded")
}

func TestRoll(t *testing.T) {
	m := &message.Message{
		Command:    "roll",
		Parameters: []string{"test"},
	}

	mBytes, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}

	var conn *amqp.Connection
	for conn == nil {
		conn, err = amqp.Dial(os.Getenv("GACHA_RMQ_CONNECTION_STRING"))
		if err != nil {
			log.Println("couldn't open RabbitMQ connection, retrying in 5 seconds")
			time.Sleep(5 * time.Second)
		}
	}
	log.Println("opened RabbitMQ connection")
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("opened RabbitMQ channel")
	defer ch.Close()

	mq, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	msgs, err := ch.Consume(mq.Name, "", false, false, false, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	corrID := uuid.NewString()

	err = ch.Publish("", os.Getenv("GACHA_RMQ_CHANNEL"), false, false, amqp.Publishing{
		ContentType:   "application/json",
		CorrelationId: corrID,
		ReplyTo:       mq.Name,
		Body:          mBytes,
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Println("published application request")

	var roll int
	for d := range msgs {
		if corrID == d.CorrelationId {
			roll, err = strconv.Atoi(string(d.Body))
			if err != nil {
				t.Fatal(err)
			}
			break
		}
	}

	log.Printf("Rolled object with ID: %d", roll)
}

func TestDeleteTable(t *testing.T) {
	m := &message.Message{
		Command:    "delete_drop_table",
		Parameters: []string{"test"},
	}

	mBytes, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}

	var conn *amqp.Connection
	for conn == nil {
		conn, err = amqp.Dial(os.Getenv("GACHA_RMQ_CONNECTION_STRING"))
		if err != nil {
			log.Println("couldn't open RabbitMQ connection, retrying in 5 seconds")
			time.Sleep(5 * time.Second)
		}
	}
	log.Println("opened RabbitMQ connection")
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("opened RabbitMQ channel")
	defer ch.Close()

	mq, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	msgs, err := ch.Consume(mq.Name, "", false, false, false, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	corrID := uuid.NewString()

	err = ch.Publish("", os.Getenv("GACHA_RMQ_CHANNEL"), false, false, amqp.Publishing{
		ContentType:   "application/json",
		CorrelationId: corrID,
		ReplyTo:       mq.Name,
		Body:          mBytes,
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Println("published application request")

	for d := range msgs {
		if corrID == d.CorrelationId {
			if string(d.Body) != "Success" {
				t.Fatalf("Did not succeed")
			}
			break
		}
	}

	log.Println("Succeeded")
}
