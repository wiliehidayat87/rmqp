package rmqp

// this package is the initiation of AMQP
// and handle of any create connection of Rabbit MQ Framework

import (
	"github.com/streadway/amqp"
)

// IntegratePublish : publish a message
func (rabbit *AMQP) IntegratePublish(exch string, queue string, contentType string, correlationId string, requestBody string) bool {

	err := rabbit.Channel.Publish(
		exch,  // exchange name
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:   contentType,
			DeliveryMode:  amqp.Persistent,
			CorrelationId: correlationId,
			Body:          []byte(requestBody),
		},
	)

	//input := Lib.ReduceWords(requestBody, 0, 30)

	if err != nil {
		//fmt.Printf("[x] Failed published: %s, Data: %s ...\n", correlationId, requestBody)

		return false

	} else {
		//fmt.Printf("[v] Published: %s, Data: %s ...\n", correlationId, requestBody)

		return true
	}

}

// Publish V2 : publish a message with several new added parameter
func (rabbit *AMQP) PublishMsg(pi PublishItems) bool {

	err := rabbit.Channel.Publish(
		pi.ExchangeName, // exchange name
		pi.QueueName,    // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType:   pi.ContentType,
			DeliveryMode:  amqp.Persistent,
			Priority:      pi.Priority,
			CorrelationId: pi.CorId,
			Body:          []byte(pi.Payload),
		},
	)

	//input := Lib.ReduceWords(requestBody, 0, 30)

	if err != nil {
		//fmt.Printf("[x] Failed published: %s, Data: %s ...\n", pi.CorId, pi.Payload)

		return false

	} else {
		//fmt.Printf("[v] Published: %s, Data: %s ...\n", pi.CorId, pi.Payload)

		return true
	}

}
