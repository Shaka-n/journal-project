package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Entry struct {
	Title string
	Body  []byte
}

func main() {
	http.HandleFunc("/entries/view/", viewEntry)
	http.HandleFunc("/entries/edit/", editEntry)

	p1 := &Entry{Title: "Hello", Body: []byte("This is a sample entry in a journal")}
	p1.saveEntry()
	p2, _ := loadEntry("Hello")
	fmt.Println(string(p2.Body))

	log.Fatal(http.ListenAndServe(":8080", nil))
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
