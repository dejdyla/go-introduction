package main

import (
	"fmt"
	"time"
)

//START OMIT
func printer(s string) {
	for {
		fmt.Println("I print " + s)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printer("something")
	fmt.Println("from main")
	time.Sleep(1 * time.Second)
	fmt.Println("from main")
}
//END OMIT
