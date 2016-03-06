package main

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func Positive(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	// db := OpenDBConnection()
	// hashtag := SelectRandomHashtag(db, "positive")
	//
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(hashtag); err != nil {
	// 	panic(err)
	// }
}

func Negative(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Ricky", []string{"cats", "cats"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	// db := OpenDBConnection()
	// hashtag := SelectRandomHashtag(db, "negative")
	//
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(hashtag); err != nil {
	// 	panic(err)
	// }
}
