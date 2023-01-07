package main

import "fmt"

func (englishBot) getGreeting() string {
	// very custom logic
	return "hi there!"
}

// delete variable because we don't use it
func (spanishBot) getGreeting() string {
	// very custom logic
	return "hola!"
}
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
