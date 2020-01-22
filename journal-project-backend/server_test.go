package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPost(t *testing.T) {
	resp := httptest.NewRecorder()

	path := "/entries/"

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	http.DefaultServeMux.ServeHTTP(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println(p)
		t.Fail()
	}
}
