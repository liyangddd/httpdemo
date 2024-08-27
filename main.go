package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func timer_log(interval int) {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		log.Printf("Just tell you I am still alive")
	}
}

func get_randam(max int) int {
	return rand.Intn(max)
}

func randam(w http.ResponseWriter, req *http.Request) {
	r := get_randam(100)
	log.Printf("%s access /random, return %d", req.RemoteAddr, r)
	fmt.Fprintf(w, "%d\n", r)
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s access /hello", req.RemoteAddr)
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s access /headers", req.RemoteAddr)
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	f, err := os.OpenFile("http.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalln("open log file fail")
	}
	defer func() {
		f.Close()
	}()
	log.SetOutput(f)
	go timer_log(10)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/random", randam)

	http.ListenAndServe(":8090", nil)
}
