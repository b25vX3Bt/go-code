package main

import (
    "fmt"
    "net/http"
    "log"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {

    log.Printf("response start\n")
    w.Header().Set("X-Content-Type","text/plain")
    w.WriteHeader(http.StatusOK)

    time.Sleep(5000 * time.Millisecond)

    w.Write([]byte("Hello, world"))

    //fmt.Fprintf(w, "Hello, World")
    log.Printf("response end\n")
}

func main() {

    //// no call
    //http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //    fmt.Fprintf(w, "Hello, World")
    //})

    http.HandleFunc("/", handler)

    fmt.Printf("listen start :8080\n")
    http.ListenAndServe(":8080", nil)
}

