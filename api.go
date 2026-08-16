package main

import (
    "fmt"
    "net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "API is working successfully!")
}

func main() {
    http.HandleFunc("/api", apiHandler)
    http.ListenAndServe(":8080", nil)
}
