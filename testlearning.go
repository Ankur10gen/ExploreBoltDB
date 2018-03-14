package main

// Learning how &structName differs from structName while creating a variable
// of that type

import (
	"fmt"
)

type hello struct {
	message string
}

func (h *hello) say() {
	fmt.Println("HELLO", h.message)
}

func (h *hello) update() {
	h.message = "Gophers"
}

func main() {
	h1 := &hello{message: "World"}
	h1.update()
	h1.say()
}
