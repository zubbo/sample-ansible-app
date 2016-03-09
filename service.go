package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<body> %s</body>", "This is a test")
}

func main() {
	log.Println("Starting service on port 80")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
