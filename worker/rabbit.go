package main

import (
	"eventExperiments/events"
	"github.com/wagslane/go-rabbitmq"
	"log"
	"time"
)

func amqpFn() {
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *rabbitmq.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	consumer, err := rabbitmq.NewConsumer(
		conn,
		func(d rabbitmq.Delivery) rabbitmq.Action {
			log.Printf("consumed: %v", string(d.Body))
			time.Sleep(time.Millisecond * 200)
			return rabbitmq.Ack
		},
		events.Queue,
		rabbitmq.WithConsumerOptionsRoutingKey(events.RoutingKey),
		rabbitmq.WithConsumerOptionsQOSPrefetch(1),
		rabbitmq.WithConsumerOptionsExchangeDurable,
		rabbitmq.WithConsumerOptionsExchangeName(events.Exchange),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	input := make(chan int)
	<-input
}
