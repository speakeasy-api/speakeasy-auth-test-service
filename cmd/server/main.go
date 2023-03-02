package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/speakeasy-api/speakeasy-auth-test-service/internal/auth"
	"github.com/speakeasy-api/speakeasy-auth-test-service/internal/requestbody"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	}).Methods(http.MethodGet)
	r.HandleFunc("/auth", auth.HandleAuth).Methods(http.MethodPost)
	r.HandleFunc("/requestbody", requestbody.HandleRequestBodyArrOfObj).Methods(http.MethodPost)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
