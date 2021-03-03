package main

import "fmt"

const es = "Spanish"
const fr = "French"

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "
const frenchGreeting = "Bonjour, "

// Hello prints a hello world message
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == es {
		return spanishGreeting + name
	}

	if language == fr {
		return frenchGreeting + name
	}

	return englishGreeting + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
