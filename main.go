package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


func main() {
	// Read in the .env file using godotenv
	err := godotenv.Load()

	if err != nil {
		// If there was an error, we're running without an env file
	}

	helloHandler := func() string {
		message := os.Getenv("HELLO_MESSAGE")

		if message == "" {
			message = "Hello, World!"
		}

		// Append a newline to the message
		message += "\n"

		return message
	}

	port := "80"

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		message := helloHandler()
		w.Write([]byte(message))
	})

	// Start the server
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
