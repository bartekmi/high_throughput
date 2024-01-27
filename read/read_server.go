// Test with...
// curl -X POST http://localhost:8080/api/v1/write -H "Content-Type: application/json" -d '{"content":"Save to DB"}'

package main

import (
	"fmt"
	"log"
	"net/http"
	"snippetstore/storage"
)

func handleWrite(rdr *Reader, rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Path[1:]
	content, ok, err := rdr.Read(ID)
	if err != nil {
		ReturnError(rw, "Error reading content", err)
		return
	}

	if !ok {
		ReturnError(rw, fmt.Sprintf("ID '%s' does not exist", ID), err)
		return
	}

	rw.Header().Set("Content-Type", "text/plain")
	bytes := []byte(content.Content)
	_, err = rw.Write(bytes)
	if err != nil {
		ReturnError(rw, "Error writing response", err)
		return
	}
}

func ReturnError(w http.ResponseWriter, message string, err error) {
	log.Printf("%s: %v", message, err)
	http.Error(w, message, http.StatusBadRequest)
}

func main() {
	// Used for testing
	// s := storage.NewDummy()
	// s.Write(storage.KVPair{ID: "key1", Content: "Content 1"})
	// s.Write(storage.KVPair{ID: "key2", Content: "Content 2", Title: "Title 2"})

	s := storage.NewDynamoDB(storage.DYNAMODB_TABLE_PROD)

	w := New(s)
	address := ":8081"

	handler := func(rw http.ResponseWriter, r *http.Request) {
		handleWrite(w, rw, r)
	}

	http.HandleFunc("/", handler)
	fmt.Println("Read Server Listening on " + address)
	log.Fatal(http.ListenAndServe(address, nil))
}
