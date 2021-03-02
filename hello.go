package main

import "fmt"

// Hello prints a hello world message
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
