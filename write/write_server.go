// Test with...
// curl -X POST http://localhost:8080/api/v1/write -H "Content-Type: application/json" -d '{"content":"Save to DB"}'

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"snippetstore/storage"
)

const DOMAIN string = "http://barteksnippet/"

type WritePayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type WriteResponse struct {
	URL string `json:"url"`
}

func handleWrite(w *Writer, rw http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ReturnError(rw, "Error reading body", err)
		return
	}

	var payload WritePayload

	// Parse the JSON data into the struct
	err = json.Unmarshal(body, &payload)
	if err != nil {
		ReturnError(rw, "Error parsing JSON", err)
		return
	}

	// For now, Title is ignored
	if payload.Content == "" {
		ReturnError(rw, "Missing content", nil)
		return
	}

	ID, err := w.Write(payload.Content)
	if err != nil {
		ReturnError(rw, "Error storing content", err)
		return
	}

	response := WriteResponse{
		URL: DOMAIN + ID,
	}
	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(response)
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
	s := storage.NewDummy()
	w := New(s)
	address := ":8080"

	handler := func(rw http.ResponseWriter, r *http.Request) {
		handleWrite(w, rw, r)
	}

	http.HandleFunc("/api/v1/write", handler)
	fmt.Println("Write Server Listening on " + address)
	log.Fatal(http.ListenAndServe(address, nil))
}
