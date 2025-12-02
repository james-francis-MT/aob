package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/James-Francis-MT/aob/internal/advent"
	"github.com/James-Francis-MT/aob/internal/server"
)

func main() {
	// Create the advent calendar for current year
	currentYear := time.Now().Year()
	calendar := advent.NewCalendar(currentYear, "content")

	// Create the server
	srv, err := server.New(calendar, "templates", "static")
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Start the HTTP server
	addr := ":8080"
	fmt.Printf("Starting Advent of Beckie server on %s\n", addr)
	fmt.Printf("Visit http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, srv); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
