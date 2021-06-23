package main

import (
	"log"
	"net/http"

	"gopkg.in/olahol/melody.v1"
)

func main() {
	m := melody.New()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		m.HandleRequest(rw, r)
	})

	m.HandleMessage(func(s *melody.Session, b []byte) {
		m.Broadcast(b)
	})

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
