package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cinfinit/gowatcher"
)

func main() {
	// Start watching for file changes (only active with -tags dev)
	gowatcher.Watch(".")

	// Simple example app
	port := 8080
	if len(os.Args) > 1 {
		if p, err := strconv.Atoi(os.Args[1]); err == nil {
			port = p
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from GoWatch! Time: %s\n", time.Now().Format(time.RFC3339))
	})

	log.Printf("Server starting on :%d", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Fatal(err)
	}
}
