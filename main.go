package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func get_randam(max int) int {
	return rand.Intn(max)
}

func randam(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%d\n", get_randam(100))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/random", randam)

	http.ListenAndServe(":8090", nil)
}
