package main

// TODO
// Connect to actual database
// Write basic unit tests
// Create CI/CD process with Github actions and terraform

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Entry struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/entries/", handleEntriesRequest)
	http.HandleFunc("/entries/view/", viewEntry)
	http.HandleFunc("/entries/edit/", editEntry)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleEntriesRequest(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		createEntry(response, request)
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
	entry.saveEntry()
}

func viewEntry(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/entries/"):]
	journalEntry, err := loadEntry(title)
	if err != nil {
		fmt.Fprintf(w, "<h1>Error :(</h1><div>Looks like we couldn't find your page.  Sorry about that!</div>")
	} else {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", journalEntry.Title, journalEntry.Body)
	}
}

func editEntry(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/entries/edit/"):]
	entry, err := loadEntry(title)
	if err != nil {
		entry = &Entry{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, entry)
}
