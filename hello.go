package main

import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "os"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
    instanceId := rand.Intn(1000000)

    message := os.Getenv("MESSAGE")
    if message == "" {
      message = "Hello world"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "'%s' from %d.", message, instanceId)
    })

    port := os.Getenv("PORT")
    if port == "" {
      port = "8080"
    }
    listenAddr := fmt.Sprintf(":%s", port)
    log.Printf("Binding to %s...\n", listenAddr)
    if err := http.ListenAndServe(listenAddr, logRequest(http.DefaultServeMux)); err != nil && err != http.ErrServerClosed {
        log.Fatal("Could not listen on %s: %v", listenAddr, err)
    }

    log.Printf("Server stopped.\n")
}
