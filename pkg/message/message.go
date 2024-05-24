package message

import (
	"context"
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type amqpConnT struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    *amqp.Queue
}

// Global variable but private
var amqpConn *amqpConnT = nil

// Create a new connection to a AMQP server.
func CreateConnection(url string) error {
	conn, err := amqp.Dial(url)

	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect to RabbitMQ: %s", err.Error()))
	}

	ch, err := conn.Channel()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to open a channel: %s", err.Error()))
	}

	q, err := ch.QueueDeclare(
		"acme_messages", // name
		false,           // durable
		true,            // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)

	amqpConn = &amqpConnT{
		conn: conn,
		ch:   ch,
		q:    &q,
	}

	return nil
}

// Return the instance or error if the RabbitMQ broker is not loaded yet
func GetConnection() (*amqpConnT, error) {
	if amqpConn == nil {
		return nil, errors.New("you must call `CreateConnection(<url>)` first.")
	}

	return amqpConn, nil
}

// Close global RabbitMQ connection if not nil
func CloseConnection() error {
	if amqpConn == nil {
		return errors.New("you're AMQP connection is empty")
	}

	amqpConn.conn.Close()
	amqpConn.ch.Close()

	return nil
}

// Send a message to the selected queue using a body
func SendMessage(body []byte) error {
	conn, err := GetConnection()

	if err != nil {
		return err
	}

	ctx := context.Background()
	err = conn.ch.PublishWithContext(ctx,
		"",
		conn.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
