package greet

import (
	"fmt"
	"io"
	"net/http"
)

// Greet gives a hearty greeting
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler is used to greet a user on an http connection
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "pickles")
}

func main() {
	http.ListenAndServe(
		":5000", http.HandlerFunc(MyGreeterHandler))
}
