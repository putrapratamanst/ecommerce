package messaging

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
    channel *amqp.Channel
    conn    *amqp.Connection
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

    return &RabbitMQ{conn: conn, channel: ch}, nil

}

func (r *RabbitMQ) Publish(queue string, message []byte) error {
    _, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
    if err != nil {
        return err
    }

    err = r.channel.Publish(
        "",     // exchange
        queue,  // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        message,
        })
    return err
}

func (r *RabbitMQ) Consume(queue string, consumerFunc func([]byte)) error {
    msgs, err := r.channel.Consume(queue, "", true, false, false, false, nil)
    if err != nil {
        return err
    }

    go func() {
        for msg := range msgs {
            consumerFunc(msg.Body)
        }
    }()

    return nil
}

func (r *RabbitMQ) Close() {
    r.channel.Close()
    r.conn.Close()
}
