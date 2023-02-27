package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/speakeasy-api/speakeasy-auth-test-service/internal/auth"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth", auth.HandleAuth).Methods(http.MethodPost)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
