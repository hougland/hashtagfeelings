package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func Positive(w http.ResponseWriter, r *http.Request) {
	db := OpenDBConnection()
	hashtag := SelectRandomHashtag(db)
	fmt.Printf("hashtag: %v", hashtag)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(hashtag); err != nil {
		panic(err)
	}

}

func Negative(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Negative, %q", html.EscapeString(r.URL.Path))
}
