package main

import (
	"fmt"
	"time"
)

//START OMIT
func printer(print chan string) {
	for {
		s := <- print
		fmt.Println("I print " + s)
	}
}

func main() {
	print := make(chan string)
	go printer(print)
	fmt.Println("from main")
	print <- "something"
	print <- "something else"
	time.Sleep(1 * time.Second)
	fmt.Println("from main")
}
//END OMIT
