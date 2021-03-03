package main

import "fmt"

const englishGreeting = "Hello, "

// Hello prints a hello world message
func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishGreeting + name
}

func main() {
	fmt.Println(Hello("world"))
}
