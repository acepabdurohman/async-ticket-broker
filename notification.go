package main

import(
	"log"
	"github.com/streadway/amqp"
	"fmt"
	"encoding/json"
)


func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

func main(){

	ReceiveMessage()
	
}

func ReceiveMessage(){

	conn, err := amqp.Dial("amqp://guest:guest@localhost:8083")
	failOnError(err, "Failed to connect AMQP")

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"booking", // queue name
		true, // durable
		false, // delete when used
		false, // exclusive
		false, // no-wait
		nil, // arguments
		)

	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			msg := string(d.Body)

			booking := Booking{}

			json.Unmarshal([]byte(msg), &booking)

			fmt.Println("Sending notification to user : " + booking.Username + " with booking code : " + booking.Code)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

type Booking struct{
	Code string `json:"code"`
	Username string `json:"username"`
	Destination string `json:"destination"`
}