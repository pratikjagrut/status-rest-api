package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//Status ...
type Status struct {
	CommitID string `json:"CommitID"`
}

//Current ...
type Current []Status

func handleRequests() {
	http.HandleFunc("/api/status", returnStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnStatus(w http.ResponseWriter, r *http.Request) {
	status := Current{
		Status{CommitID: "687c95e649f985c78a203e601bfb9a2822e0787c"},
	}
	json.NewEncoder(w).Encode(status)
}

func main() {
	handleRequests()
}
