package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Greet writes a greeting to any io.Writer using dependency injection
// Supports both testing (bytes.Buffer) and real output (os.Stdout)
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// http://localhost:5001/
func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
	// Greet(os.Stdout, "Elodie")
}
