package main

import (
	"fmt"
	"html"
	"net/http"
)

func Positive(w http.ResponseWriter, r *http.Request) {
	db := OpenDBConnection()
	hashtag := SelectRandomHashtag(db)
	fmt.Printf("hashtag: %v", hashtag)
}

func Negative(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Negative, %q", html.EscapeString(r.URL.Path))
}
