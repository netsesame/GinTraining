package initializers

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// func ReceiveMessage(queueName string, messageHandler func([]byte)) {
// 	conn, err := amqp.Dial("amqp://nawosix:wendy2009+@154.211.21.219:5672/")
// 	if err != nil {
// 		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatalf("Failed to open a channel: %v", err)
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		queueName, // name
// 		false,     // durable
// 		false,     // delete when unused
// 		false,     // exclusive
// 		false,     // no-wait
// 		nil,       // arguments
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to declare a queue: %v", err)
// 	}

// 	msgs, err := ch.Consume(
// 		q.Name, // queue
// 		"",     // consumer
// 		true,   // auto-ack
// 		false,  // exclusive
// 		false,  // no-local
// 		false,  // no-wait
// 		nil,    // args
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to register a consumer: %v", err)
// 	}

// 	fmt.Printf("Listening for messages on queue '%s'...\n", queueName)

//		for d := range msgs {
//			messageHandler(d.Body)
//		}
//	}

func ReceiveMessage(exchangeName, exchangeType, routingKey string) {
	conn, err := amqp.Dial("amqp://nawosix:wendy2009+@154.211.21.219:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare(exchangeName, exchangeType, false, false, false, false, nil); err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	if err := ch.QueueBind(q.Name, routingKey, exchangeName, false, nil); err != nil {
		log.Fatalf("Failed to bind the queue to the exchange: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	fmt.Println("Waiting for messages...")

	for d := range msgs {
		message := string(d.Body)
		data := loads([]byte(message))
		//decodedMessage, err := decodeUTF8(message)
		if err != nil {
			log.Fatalf("Failed to decode message: %v", err)
		}
		//"e"：表示消息的交换机名称。在这个例子中，该值为<nil>，表示该消息没有指定交换机。
		//"k"：表示消息的路由键（routing key）。在这个例子中，该值为"bnPositions"，表示该消息将被路由到名为"bnPositions"的队列。
		//"n"：表示消息的名称。在这个例子中，该值为<nil>，表示该消息没有指定名称。
		fmt.Printf("Received a message: %s\n", data)
	}
}
func loads(b []byte) map[string]interface{} {
	r, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	defer r.Close()
	var d map[string]interface{}
	err = json.NewDecoder(r).Decode(&d)
	if err != nil {
		panic(err)
	}
	return d
}

// func decodeUTF8(message string) (string, error) {
// 	decoded := ""
// 	for len(message) > 0 {
// 		r, size := utf8.DecodeRuneInString(message)
// 		if r == utf8.RuneError {
// 			return "", fmt.Errorf("invalid UTF-8 encoding")
// 		}
// 		decoded += string(r)
// 		message = message[size:]
// 	}
// 	return decoded, nil
// }
