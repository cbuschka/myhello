package main

import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
)

func main() {
    instanceId := rand.Intn(1000000)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world from %d.", instanceId)
    })

    listenAddr := ":8080"
    log.Printf("Binding to %s...\n", listenAddr)
    if err := http.ListenAndServe(listenAddr, nil); err != nil && err != http.ErrServerClosed {
        log.Fatal("Could not listen on %s: %v", listenAddr, err)
    }

    log.Printf("Server stopped.\n")
}
