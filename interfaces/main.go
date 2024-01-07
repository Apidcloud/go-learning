package main

import "fmt"

type test interface {
	getGreeting() string
}

type bot interface {
	test
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	e, s := englishBot{}, spanishBot{}
	printGreeting(e)
	printGreeting(s)
}

func (englishBot) getGreeting() string {
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola que tal"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
