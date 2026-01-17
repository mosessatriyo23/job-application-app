package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Job Application API is running")
    })

    fmt.Println("Server running at :8080")
    http.ListenAndServe(":8080", nil)
}
