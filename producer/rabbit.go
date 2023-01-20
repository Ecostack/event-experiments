package main

import (
	"eventExperiments/events"
	"github.com/wagslane/go-rabbitmq"
	"log"
	"strconv"
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

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(events.Exchange),
		rabbitmq.WithPublisherOptionsExchangeDurable,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer publisher.Close()

	for i := 0; i < 200; i++ {
		err = publisher.Publish(
			[]byte("hello, world "+strconv.FormatInt(int64(i), 10)),
			[]string{events.RoutingKey},
			rabbitmq.WithPublishOptionsContentType("application/json"),
			rabbitmq.WithPublishOptionsPersistentDelivery,
			rabbitmq.WithPublishOptionsExchange(events.Exchange),
		)
		if err != nil {
			log.Println(err)
		}
	}
}
