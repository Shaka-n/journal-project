package main

import (
	"io/ioutil"
)

func (p *Entry) saveEntry() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadEntry(title string) (*Entry, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Entry{Title: title, Body: body}, nil
}
