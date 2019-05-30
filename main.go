package main

import (
	"encoding/json"
	git "gopkg.in/src-d/go-git.v4"
	"log"
	"net/http"
	"os"
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
		Status{CommitID: getCommitID()},
	}
	json.NewEncoder(w).Encode(status)
}

func getCommitID() string {
	_ = os.RemoveAll("/tmp/foo")
	repo, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/pratikjagrut/status-rest-api.git",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}
	head, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}
	arr := head.Strings()
	return arr[1]
}

func main() {
	handleRequests()
}
