package messaging

import (
	"fmt"

    amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
    Channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    // Declare exchange
    err = ch.ExchangeDeclare(
        "order_exchange", // name
        "direct",         // type
        true,             // durable
        false,            // auto-deleted
        false,            // internal
        false,            // no-wait
        nil,              // arguments
    )
    if err != nil {
        return nil, err
    }

    return &RabbitMQ{Channel: ch}, nil
}

func (r *RabbitMQ) PublishOrderPlaced(orderID uint) error {
    body := fmt.Sprintf("Order %d placed", orderID)

    err := r.Channel.Publish(
        "order_exchange", // exchange
        "order_placed",   // routing key
        false,            // mandatory
        false,            // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )

    return err
}
