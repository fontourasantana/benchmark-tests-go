package queue

import (
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

type QueueConnectionConfig struct {
	User     string
	Password string
	Host     string
	Port     string
}

type QueueHandler struct {
	ch *amqp.Channel
}

type MessageQueue struct {
	exchange    string
	routeKey    string
	contentType string
	body        string
	persistent  bool
}

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		"amq.topic",
		"",
		"text/plain",
		"",
		true,
	}
}

func (this *MessageQueue) SetExchange(exchange string) {
	this.exchange = exchange
}

func (this *MessageQueue) SetRouterKey(routeKey string) {
	this.routeKey = routeKey
}

func (this *MessageQueue) SetContentType(contentType string) {
	this.contentType = contentType
}

func (this *MessageQueue) SetBody(body string) {
	this.body = body
}

func (this *MessageQueue) SetPersistent(persistent bool) {
	this.persistent = persistent
}

var (
	queue     *QueueHandler
	queueOnce sync.Once
)

func NewQueueHandler(config QueueConnectionConfig) *QueueHandler {
	if queue == nil {
		queueOnce.Do(func() {
			log.SetFlags(0)
			log.Println("> [ infrastructure ] Creating queue handler ...")
			conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.User, config.Password, config.Host, config.Port))
			if err != nil {
				log.Fatal("> [ infrastructure ] Não foi possivel estabelecer conexão com o RabbitMQ")
			}

			ch, err := conn.Channel()
			if err != nil {
				log.Fatalf("> [ infrastructure ] QueueHandler error: %s", err)
			}

			// // RECEIVED_PAYMENTS
			// q, err := ch.QueueDeclare(
			// 	"EMAILS", // name
			// 	true,     // durable
			// 	false,    // delete when unused
			// 	false,    // exclusif
			// 	false,    // no-wait
			// 	nil,      // arguments
			// )

			// if err != nil {
			// 	log.Fatalf("Failed to Declare a queue : %s", err)
			// }

			queue = &QueueHandler{ch}

			log.Println("> [ infrastructure ] Queue handler created")
		})
	}

	return queue
}

func NewQueueConsumer() {
	log.SetFlags(0)
	if queue == nil {
		log.Fatal("> [ infrastructure ] Queue handler not started")
	}

	msgs, err := queue.ch.Consume(
		"EMAILS",        // quee
		"email-service", // consumer
		false,           // auto-act
		false,           // exclusif
		false,           // no-local
		false,           // no-wait
		nil,
	)

	if err != nil {
		log.Fatalf("> [ infrastructure ] Erro ao criar consumer")
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a massage: %v", d)
			d.Ack(false)
		}
	}()
}

func (this *QueueHandler) Publish(message *MessageQueue) error {
	var deliveryMode uint8

	if message.persistent {
		deliveryMode = amqp.Persistent
	} else {
		deliveryMode = amqp.Transient
	}

	return this.ch.Publish(
		message.exchange, // exchange
		message.routeKey, //routing key
		false,            //mandatory
		false,            // immediati
		amqp.Publishing{
			DeliveryMode: deliveryMode,
			ContentType:  message.contentType,
			Body:         []byte(message.body),
		})
}
