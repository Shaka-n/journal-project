package main

// TODO
// Connect to actual database
// Write basic unit tests
// Create CI/CD process with Github actions and terraform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Title string
	Body  string
}

type JournalEntry interface {
	saveJournalEntry() error
	loadJournalEntry(string) *Entry
}

var fileDatabase FileDatabase

// TODO secure the cookie for our application

func main() {
	initializeDatabase()
	// TODO - Also need TLS, via lets encrypt
	http.HandleFunc("/entries/", handleEntriesRequest)
	// http.HandleFunc("/entries/view/", viewEntry)
	// http.HandleFunc("/entries/edit/", editEntry)
	http.Handle("/", http.FileServer(http.Dir("./build")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleEntriesRequest(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		createEntry(response, request)
	case "GET":
		getEntry(response, request)
	}
}

func createEntry(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	log.Println("Body received" + string(body))
	var entry Entry
	err = json.Unmarshal(body, &entry)
	if err != nil {
		panic(err)
	}
	log.Println("Entry received" + entry.Title)
	fileDatabase.saveJournalEntry(getUserIDFromCookie(request), &entry)
	response.WriteHeader(http.StatusCreated)
	// TODO return the one journal entry created from the database
}

func getEntry(response http.ResponseWriter, request *http.Request) {

	title := request.URL.Path[len("/entries/"):]
	log.Println("Title: " + title)
	journalEntries, err := fileDatabase.loadJournalEntriesForUser(getUserIDFromCookie(request))
	response.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "{error: \"Error getting entries\"}")
	} else {
		json.NewEncoder(response).Encode(journalEntries)
	}
}

// TODO add error handling
func getUserIDFromCookie(request *http.Request) uuid.UUID {
	userIDCookie, _ := request.Cookie("UserID")
	userID, _ := uuid.Parse(userIDCookie.Value)
	return userID
}
