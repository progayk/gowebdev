package main

import "fmt"

// person defines a human being
type person struct {
	First string
	Last string
	Age int
}

// speak makes a person to speak
func (p *person) speak() {
	fmt.Printf("%s is speaking!", p.First)
}

// human implements the speak method
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		First: "Mayk",
		Last: "Jony",
		Age: 28,
	}

	// This won't work because person does not implement human
	// speak method has pointer reciever
	//saySomething(p1)

	saySomething(&p1)


}
