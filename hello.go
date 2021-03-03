package main

import "fmt"

const englishGreeting = "Hello, "

// Hello prints a hello world message
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	return englishGreeting + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
