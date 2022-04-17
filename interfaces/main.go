package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type hindiBot struct {}

func (eb englishBot) getGreeting() string{
	return "Hi There!"
}

func (hb hindiBot) getGreeting() string{
	return "नमस्ते!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	hb := hindiBot{}

	printGreeting(eb)
	printGreeting(hb)
}