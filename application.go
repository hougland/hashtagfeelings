package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/positive", Positive)
	http.HandleFunc("/negative", Negative)

	http.ListenAndServe(":5000", nil)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
