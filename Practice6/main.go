package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

type Response struct {
    Message  string `json:"message"`
    Hostname string `json:"hostname"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    resp := Response{
        Message:  "Hello from Go!",
        Hostname: hostname,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", handler)
    fmt.Printf("Listening on :%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}