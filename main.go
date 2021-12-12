package main

import (
	// "ameicosmeticos/app"

	"ameicosmeticos/infrastructure/queue"
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	if godotenv.Load() != nil {
		panic("Error loading .env file")
	}
}

func main() {
	// app := app.Create()
	// app.Run()

	config := queue.QueueConnectionConfig{
		User:     "guest",
		Password: "guest",
		Host:     "localhost",
		Port:     "5672",
	}

	queueHandler := queue.NewQueueHandler(config)
	message := queue.NewMessageQueue()
	message.SetRouterKey("order.email")
	message.SetBody("Testando mensagem")

	if err := queueHandler.Publish(message); err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	queue.NewQueueConsumer()
	<-forever

	// msgs, err := ch.Consume(
	// 	q.Name,          // quee
	// 	"email-service", // consumer
	// 	false,           // auto-act
	// 	false,           // exclusif
	// 	false,           // no-local
	// 	false,           // no-wait
	// 	nil,
	// )

	// if err != nil {
	// 	log.Fatalf("Failed to register a consumer : %s", err)
	// }

	// foever := make(chan bool)
	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("Received a massage: %s", d.Body)
	// 		d.Ack(false)
	// 		time.Sleep(time.Second * 2)
	// 	}
	// }()

	// log.Println(" [] Waiting For Massages. To Exit press CTRL+C")

	// <-foever

	// var logger log.Logger
	// logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	// logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// srv := server.New(log.With(logger, "component", "http"))
	// logger.Log("transport", "http", "address", ":8081", "msg", "listening")
	// http.ListenAndServe(":8081", srv)
}
