package main

import (
	"fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Go Modules with Gorilla Mux!")
    })

    http.Handle("/", router)

    fmt.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil)
}
