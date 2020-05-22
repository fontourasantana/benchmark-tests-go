package main

import (
	"ameicosmeticos/app"
	"github.com/joho/godotenv"
)

func init() {
	if godotenv.Load() != nil {
		panic("Error loading .env file")
	}
}

func main() {
	app := app.Create()
	app.Run()

	// var logger log.Logger
	// logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	// logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// srv := server.New(log.With(logger, "component", "http"))
	// logger.Log("transport", "http", "address", ":8081", "msg", "listening")
	// http.ListenAndServe(":8081", srv)
}
