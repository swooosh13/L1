package main

import "fmt"

type Human struct {
	Name string
	Age int
}

func (h *Human) Speak() {
	fmt.Println("Human speak")
}

func (h *Human) Eat() {
	fmt.Println("Human eat")
}

type Action struct {
	Human
}

// Пример переопределения методов
func (a *Action) Eat() {
	fmt.Println("Action eat")
}

func (a *Action) Sit() {
	fmt.Println("Action sit")
}

func main() {
	h := Human{}
	h.Speak()		// Human speak
	h.Eat()			// Human eat

	a := &Action{}
	a.Speak()		// Human speak
	a.Eat()			// Action eat
	a.Sit()			// Action sit
}

