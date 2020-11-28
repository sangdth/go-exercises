package main

import (
	"fmt"

	"github.com/sangdth/go-excercises/pointer"
)

func main() {
	hello := "Hello World"
	hello = "something" // valid
	// hello := "something" // invalid as := only assign at initialization.
	fmt.Println(hello)

	card := newCard()
	fmt.Println(card)
	pointer.Pointer()
}

func newCard() string {
	return "this is new card"
}
