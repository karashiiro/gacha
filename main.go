package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/karashiiro/gacha/ent"
	"github.com/karashiiro/gacha/message"
	"github.com/streadway/amqp"
)

var rng *rand.Rand

func checkRoll(drops []ent.Drop, val float32) (*ent.Drop, error) {
	agg := float32(0)
	for _, drop := range drops {
		agg += drop.Rate
		if agg > 1.0 {
			break
		}
		if agg > val {
			return &drop, nil
		}
	}
	return nil, errors.New("drop rates do not sum to 1.0")
}

func main() {
	// Set up logging
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	// Connect to database
	db, err := NewDatabase(sugar)
	if err != nil {
		sugar.Errorf("couldn't connect to database, aborting",
			"error", err,
		)
		panic(err)
	}
	defer db.edb.Close()

	// Open message queue
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		sugar.Errorf("couldn't open RabbitMQ connection, aborting",
			"error", err,
		)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		sugar.Errorf("couldn't open message channel, aborting",
			"error", err,
		)
		panic(err)
	}
	defer ch.Close()

	mq, err := ch.QueueDeclare("gacha_v0", false, false, false, false, nil)
	if err != nil {
		sugar.Errorf("couldn't open message queue, aborting",
			"error", err,
		)
		panic(err)
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		sugar.Errorf("couldn't set QoS, aborting",
			"error", err,
		)
		panic(err)
	}

	msgs, err := ch.Consume(mq.Name, "", false, false, false, false, nil)
	if err != nil {
		sugar.Errorw("couldn't begin consuming messages from queue, aborting",
			"error", err,
		)
		panic(err)
	}

	// Initialize randomizer with current time
	rngSource := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(rngSource)

	// Start message loop
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			m := &message.Message{}
			err := json.Unmarshal(d.Body, m)
			if err != nil {
				sugar.Errorw("failed to unmarshal message",
					"error", err,
					"correlation_id", d.CorrelationId,
				)
				err = d.Reject(false)
				if err != nil {
					sugar.Warnw("message ack could not be delivered to channel",
						"error", err,
						"correlation_id", d.CorrelationId,
					)
				}
				continue
			}

			switch m.Command {
			case "roll":
				rows, err := db.GetDropTable(m.Parameters[0])
				if err != nil {
					sugar.Errorw("failed to get rows",
						"error", err,
						"correlation_id", d.CorrelationId,
					)
					err = d.Reject(false)
					if err != nil {
						sugar.Warnw("message ack could not be delivered to channel",
							"error", err,
							"correlation_id", d.CorrelationId,
						)
					}
					continue
				}

				testValue := rng.Float32()
				roll, err := checkRoll(rows, testValue)
				if err != nil {
					sugar.Errorw("gacha roll failed",
						"error", err,
						"correlation_id", d.CorrelationId,
					)
					err = d.Reject(false)
					if err != nil {
						sugar.Warnw("message ack could not be delivered to channel",
							"error", err,
							"correlation_id", d.CorrelationId,
						)
					}
					continue
				}

				sugar.Infow(fmt.Sprintf("rolled %v", roll),
					"correlation_id", d.CorrelationId,
				)

				err = ch.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(fmt.Sprint(roll.ObjectID)),
				})
				if err != nil {
					sugar.Errorw("reply failed",
						"error", err,
						"correlation_id", d.CorrelationId,
					)
					err = d.Reject(false)
					if err != nil {
						sugar.Warnw("message ack could not be delivered to channel",
							"error", err,
							"correlation_id", d.CorrelationId,
						)
					}
					continue
				}
			case "set_table":
				dropInsertsRaw := m.Parameters[1:]
				dropInserts := make([]DropInsert, len(dropInsertsRaw))
				for i, param := range dropInsertsRaw {
					err := json.Unmarshal([]byte(param), &dropInserts[i])
					if err != nil {
						sugar.Errorw("failed to unmarshal message",
							"error", err,
							"correlation_id", d.CorrelationId,
						)
						err = d.Reject(false)
						if err != nil {
							sugar.Warnw("message ack could not be delivered to channel",
								"error", err,
								"correlation_id", d.CorrelationId,
							)
						}
						continue
					}
				}

				err := db.SetDropTable(m.Parameters[0], dropInserts)
				if err != nil {
					sugar.Errorw("failed to set drop table",
						"error", err,
						"correlation_id", d.CorrelationId,
					)
					err = d.Reject(false)
					if err != nil {
						sugar.Warnw("message ack could not be delivered to channel",
							"error", err,
							"correlation_id", d.CorrelationId,
						)
					}
					continue
				}

				sugar.Infow(fmt.Sprintf("set table %s", m.Parameters[0]),
					"correlation_id", d.CorrelationId,
				)

				err = ch.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte("Success"),
				})
				if err != nil {
					sugar.Errorw("reply failed",
						"error", err,
						"correlation_id", d.CorrelationId,
					)
					err = d.Reject(false)
					if err != nil {
						sugar.Warnw("message ack could not be delivered to channel",
							"error", err,
							"correlation_id", d.CorrelationId,
						)
					}
					continue
				}
			case "delete_drop_table":
				seriesName := m.Parameters[0]
				err := db.DeleteDropTable(seriesName)
				if err != nil {
					sugar.Errorw("series deletion failed",
						"error", err,
						"correlation_id", d.CorrelationId,
					)
					err = d.Reject(false)
					if err != nil {
						sugar.Warnw("message ack could not be delivered to channel",
							"error", err,
							"correlation_id", d.CorrelationId,
						)
					}
					continue
				}

				err = ch.Publish("", d.ReplyTo, false, false, amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte("Success"),
				})
				if err != nil {
					sugar.Errorw("reply failed",
						"error", err,
						"correlation_id", d.CorrelationId,
					)
					err = d.Reject(false)
					if err != nil {
						sugar.Warnw("message ack could not be delivered to channel",
							"error", err,
							"correlation_id", d.CorrelationId,
						)
					}
					continue
				}
			default:
				sugar.Warnw("received unknown message",
					"unk_msg", string(d.Body),
					"correlation_id", d.CorrelationId,
				)
				err = d.Reject(false)
				if err != nil {
					sugar.Warnw("message ack could not be delivered to channel",
						"error", err,
						"correlation_id", d.CorrelationId,
					)
				}
				continue
			}

			err = d.Ack(false)
			if err != nil {
				sugar.Warnw("message ack could not be delivered to channel",
					"error", err,
					"correlation_id", d.CorrelationId,
				)
			}
		}
	}()

	sugar.Infow("application started")

	<-forever
}
