package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type FileDatabase struct {
}

var database *sql.DB

func initializeDatabase() {
	database, _ = sql.Open("sqlite3", "./local-dev.db")

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS journal_entry (id UUID PRIMARY KEY, author TEXT, title TEXT, entry TEXT)")
	something, err := statement.Exec()
	fmt.Println("Created the schema for the table")
	fmt.Println(err)
	fmt.Println(something)
	/**
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
	**/
}

// TODO - add a concept of users to the application
func (fileDatabase FileDatabase) saveJournalEntry(entry *Entry) error {
	primaryKey := (uuid.New()).String()
	_, err := database.Exec("INSERT INTO journal_entry VALUES($1, $2, $3, $4)", primaryKey, "root", entry.Title, entry.Body)
	if err != nil {
		log.Fatal("Error writing to the database")
		log.Fatal(err)
	}
	log.Println("Successfully stored entry")
	return nil
}

func (fileDatabase FileDatabase) loadJournalEntriesForAuthor(author string) (*[]Entry, error) {
	//TODO understand slices
	var results []Entry
	var title string
	var body string
	rows, err := database.Query("select title, entry FROM journal_entry WHERE author = \"root\"", 1)
	if err != nil {
		log.Fatal("Failed reading from the database")
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&title, &body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Entry found:")
		log.Println(title, body)
		results = append(results, Entry{Title: title, Body: body})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &results, nil
}
