package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/positive", Positive)
	http.HandleFunc("/negative", Negative)
	http.HandleFunc("/updatehashtags", Updated)

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
