package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

// create a hash of the input string
func hashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// create a handler function that returns a sha256 hash of the input string at endpoint /hash/
func hashHandler(w http.ResponseWriter, r *http.Request) {
	// get the input string
	s := r.URL.Path[len("/hash/"):]
	// create a hash of the input string
	h := hashString(s)
	// return the hash in the browser window as plain text
	fmt.Fprintf(w, h)
}

// create a handler function that returns the hash of the input string
func main() {
	http.HandleFunc("/hash/", hashHandler)
	http.ListenAndServe(":8080", nil) // start the server on port 8080
}
