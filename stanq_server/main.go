// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DQFile - Our struct for all articles
type DQFile struct {
	Name     string `json:"name"`
	CheckSum string `json:"chksum"`
}

var Files []DQFile

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllFIles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllFIles")
	json.NewEncoder(w).Encode(Files)
}

func returnSingleFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Files {
		if article.Name == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createFileList(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new DQFile struct
	// append this to our Files array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article DQFile
	json.Unmarshal(reqBody, &article)
	// update our global Files array to include
	// our new DQFile
	Files = append(Files, article)

	json.NewEncoder(w).Encode(article)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/files", returnAllFIles)
	myRouter.HandleFunc("/files", createFileList).Methods("POST")
	myRouter.HandleFunc("/files/{name}", returnSingleFile)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Files = []DQFile{
		{Name: "agent", CheckSum: "21f951ed89a4a54064c12ee781df4bb167db578951e940790748cfa8da558021"},
		{Name: "pid_checker", CheckSum: "445523ede4db9704353fa95a39328abda530f0e9d7c809049ea46564e069524b"},
	}
	handleRequests()
}
