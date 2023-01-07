package main

type englishBot struct {
}

type spanishBot struct {
}

type bot interface {
	getGreeting() string
}
