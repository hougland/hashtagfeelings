package main

import (
	"fmt"
	"html"
	"net/http"
)

func Positive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Positive, %q", html.EscapeString(r.URL.Path))
}

func Negative(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Negative, %q", html.EscapeString(r.URL.Path))
}
