package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func setRoutes(basePath string, writer *kafka.Writer, brokers []string) {
	http.HandleFunc(basePath+"/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HealthResponse{
			Status: true,
		})
	})

	http.HandleFunc(basePath+"/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		topic := getTopic("user")

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Value: body,
			Topic: topic,
		})
		if err != nil {
			log.Println("Error writing message:", err)
			http.Error(w, "Error writing message", http.StatusInternalServerError)
			return
		}

		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
		})
		defer reader.Close()

		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			http.Error(w, "Error reading message", http.StatusBadRequest)
			return
		}

		var payload User
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(EventResponse[User]{
			Status:    "success",
			Partition: msg.Partition,
			Offset:    msg.Offset,
			Event: Event[User]{
				Id:        uuid.New().String(),
				Type:      "movie",
				Timestamp: time.Now(),
				Payload:   payload,
			},
		})

	})

	http.HandleFunc(basePath+"/movie", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		topic := getTopic("movie")

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Value: body,
			Topic: topic,
		})
		if err != nil {
			log.Println("Error writing message:", err)
			http.Error(w, "Error writing message", http.StatusInternalServerError)
			return
		}

		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
		})
		defer reader.Close()

		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			http.Error(w, "Error reading message", http.StatusBadRequest)
			return
		}

		var payload Movie
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(EventResponse[Movie]{
			Status:    "success",
			Partition: msg.Partition,
			Offset:    msg.Offset,
			Event: Event[Movie]{
				Id:        uuid.New().String(),
				Type:      "movie",
				Timestamp: time.Now(),
				Payload:   payload,
			},
		})
	})

	http.HandleFunc(basePath+"/payment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		topic := getTopic("payment")

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Value: body,
			Topic: topic,
		})
		if err != nil {
			log.Println("Error writing message:", err)
			http.Error(w, "Error writing message", http.StatusInternalServerError)
			return
		}

		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
		})
		defer reader.Close()

		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			http.Error(w, "Error reading message", http.StatusBadRequest)
			return
		}

		var payload Payment
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(EventResponse[Payment]{
			Status:    "success",
			Partition: msg.Partition,
			Offset:    msg.Offset,
			Event: Event[Payment]{
				Id:        uuid.New().String(),
				Type:      "payment",
				Timestamp: time.Now(),
				Payload:   payload,
			},
		})
	})
}
