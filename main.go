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
	http.HandleFunc("/", returnStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnStatus(w http.ResponseWriter, r *http.Request) {
	status := Current{
		Status{CommitID: getCommitID()},
	}
	json.NewEncoder(w).Encode(status)
}

func getCommitID() string {

	//store log in file
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "prefix", log.LstdFlags)
	dir := "/tmp/foo"
	err = os.RemoveAll(dir)
	logger.Println(err)

	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      "https://github.com/pratikjagrut/status-rest-api.git",
		Progress: os.Stdout,
	})
	logger.Println(err)

	head, err := repo.Head()
	logger.Println(err)

	arr := head.Strings()
	return arr[1]
}

func main() {
	handleRequests()
}
