package main

import (
   "fmt"
)

//START OMIT
type Animal interface {
	Speak() string
}

type Dog struct {}
type Cat struct {}
type Bernardine struct {
	Dog
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	animals := []Animal{ Dog{}, Cat{}, Bernardine{} }
	for _, a := range animals {
		fmt.Println(a.Speak())
	}
}
//END OMIT
