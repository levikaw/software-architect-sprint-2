package main

import "time"

type User struct {
	UserId    int       `json:"user_id"`
	Username  *string   `json:"username,omitempty"`
	Email     *string   `json:"email,omitempty"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
}

type Movie struct {
	MovieId     int       `json:"movie_id"`
	Title       string    `json:"title"`
	Action      string    `json:"action"`
	UserId      *int      `json:"user_id,omitempty"`
	Rating      *float64  `json:"rating,omitempty"`
	Genres      *[]string `json:"genres,omitempty"`
	Description *string   `json:"description,omitempty"`
}

type Payment struct {
	PaymentId  int       `json:"payment_id"`
	UserId     int       `json:"user_id"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	Timestamp  time.Time `json:"timestamp"`
	MethodType *string   `json:"method_type,omitempty"`
}

type EventResponse[T any] struct {
	Status    string   `json:"status"`
	Partition int      `json:"partition"`
	Offset    int64    `json:"offset"`
	Event     Event[T] `json:"event"`
}

type Event[T any] struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Payload   T         `json:"payload"`
}

type HealthResponse struct {
	Status bool `json:"status"`
}
