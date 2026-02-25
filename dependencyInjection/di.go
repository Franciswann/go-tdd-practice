package main

import (
	"fmt"
	"io"
	"os"
)

// Greet writes a greeting to any io.Writer using dependency injection
// Supports both testing (bytes.Buffer) and real output (os.Stdout)
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
