package main

import (
    "fmt"
    "log"
    "net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s is visiting %s at %s", r.Header.Get("User-Agent"), r.URL.Path, time.Now().UTC().Format(time.RFC3339))
	fmt.Println(r.Header.Get("User-Agent"), "is visiting", r.URL.Path, "at", time.Now().UTC().Format(time.RFC3339))
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}