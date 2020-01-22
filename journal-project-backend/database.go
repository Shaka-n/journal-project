package main

import (
	"io/ioutil"
)

type FileDatabase struct {
}

func (fileDatabase FileDatabase) saveJournalEntry(p *Entry) error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func (fileDatabase FileDatabase) loadJournalEntry(title string) (*Entry, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	var stringBody = string(body)
	if err != nil {
		return nil, err
	}
	return &Entry{Title: title, Body: stringBody}, nil
}
