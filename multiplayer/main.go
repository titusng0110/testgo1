package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s is visiting at %s", r.Header.Get("User-Agent"), time.Now().UTC().Format(time.RFC3339))
    fmt.Println(r.Header.Get("User-Agent"), "is visiting at", time.Now().UTC().Format(time.RFC3339))
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
