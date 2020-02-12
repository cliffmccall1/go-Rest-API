package main

import (
	"log"
	"net/http"
)

func apiResonse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"GET method requested"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Can't find method requested"}`))
	}
}

func main() {
	http.HandleFunc("/", apiResonse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
