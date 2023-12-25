package main

import (
	"credit_card_validator/internal/http-server/handlers/card"
	"log"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/check", card.New())
	log.Fatal(s.ListenAndServe())
}
