package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", Positive)
	http.HandleFunc("/n", Negative)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
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
}
