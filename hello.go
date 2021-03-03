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

	return greeting(language) + name
}

func greeting(language string) (greeting string) {
	switch language {
	case fr:
		return frenchGreeting
	case es:
		return spanishGreeting
	default:
		return englishGreeting
	}
}

func main() {
	fmt.Println(Hello("world", ""))
}
