package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	port := os.Getenv("PORT")
	brokerUrl := os.Getenv("KAFKA_BROKERS")

	for _, v := range []string{port, brokerUrl} {
		if v == "" {
			log.Fatal("One of env vars not set (PORT, KAFKA_BROKERS)")
		}
	}

	basePath := "/api/events"

	brokers := []string{brokerUrl}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:     brokers,
		Logger:      kafka.LoggerFunc(logf),
		ErrorLogger: kafka.LoggerFunc(logf),
		Async:       true,
	})
	defer writer.Close()

	setRoutes(basePath, writer, brokers)

	fmt.Printf("Starting server at port %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func logf(msg string, a ...any) {
	fmt.Printf(msg, a...)
	fmt.Println()
}
