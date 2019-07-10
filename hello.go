package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(Hello())

	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
} // .main

func Hello() string {
	return "Hello, world."
} // .Hello

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(Hello()))
} // .home
