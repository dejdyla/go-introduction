package main

import (
	"fmt"
	"time"
)

//START OMIT
func printer(print chan string, done chan bool) {
	for {
		select {
		case s := <- print:
			fmt.Println("I print " + s)
		case <- done:
			fmt.Println("I am done!")
			return
		}
	}
}

func main() {
	print := make(chan string)
	done := make(chan bool)
	go printer(print, done)
	fmt.Println("from main")
	print <- "something"
	done <- true
	time.Sleep(1 * time.Second)
	fmt.Println("from main")
}
//END OMIT