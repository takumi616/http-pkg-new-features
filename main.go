package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	//Get path value
	id := r.PathValue("id")

	msg := &Message{
		Content: fmt.Sprintf("Test PathValue method and new way of routing definition. Got id is: %s", id),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		log.Fatalf("Failed to encode: %v", err)
	}
}

func postNewMessage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Failed to read body: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()

	//Specify http request method with url
	mux.HandleFunc("GET /messages/{id}", getMessage)
	mux.HandleFunc("POST /messages", postNewMessage)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()
}
