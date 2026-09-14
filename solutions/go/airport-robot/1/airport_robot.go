package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello (name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}