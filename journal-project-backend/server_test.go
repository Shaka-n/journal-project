package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO - how do we mock the database layer?
// I made a interface for the database, but have to go farther
func TestHappyPathPost(t *testing.T) {

	bodyMap := map[string]string{"Title": "Testing", "Body": "This is a body"}
	bodyJson, _ := json.Marshal(bodyMap)
	request, err := http.NewRequest("POST", "/entries/", bytes.NewBuffer(bodyJson))

	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(createEntry)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(responseRecorder, request)

	// Check the status code is what we expect.
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("createEntry returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// TODO once we have a response code
	// Check the response body is what we expect.
	// expected := `{"alive": true}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}

func TestPostNoBody(t *testing.T) {

}
