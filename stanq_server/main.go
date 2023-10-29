// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	_, err := fmt.Fprintf(w, "Welcome to the HomePage!")
	if err != nil {
		return
	}
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
	reqBody, _ := io.ReadAll(r.Body)
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
		{Name: "file_checker", CheckSum: "f62f0949ffa5a041e31b2fc706d93708ce93b62b7099f516380f8d39d15380c6"},
		{Name: "pid_checker", CheckSum: "539181b57717e6f3430d599ca3b11fe18c9d405c8cfe3ff163037694e4864b1a"},
		{Name: "agent", CheckSum: "b97ab6fabafba27199d50a190a2ad6513ccf8ee722558e86d2a45fd2ac535c67"},
	}
	handleRequests()
}
