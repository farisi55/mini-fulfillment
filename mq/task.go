package mq

import (
	"log"
	"os"
	"strings"
	"theapp/Models"

	"github.com/streadway/amqp"
)

func failOnErrorTask(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErrorTask(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErrorTask(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnErrorTask(err, "Failed to declare a queue")

	//body := bodyFrom(os.Args)
	user := Models.User{}
	body := "{\"name\":\""+user.Name+"\",\"email\":\""+user.Email+"\",\"phone\":\""+user.Phone+"\",\"address\":\""+user.Address+"\",\"password\":\""+user.Password+"\",\"operation\":\""+user.Operation+"\"}"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/json",
			Body:         []byte(body),
		})
	failOnErrorTask(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}