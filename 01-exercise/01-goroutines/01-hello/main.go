package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("go-routines-1")

	// goroutine with anonymous function
	go func() {
		fun("go-routines-2")
	}()

	// goroutine with function value call
	call := fun
	go call("go-routines-3")

	// wait for goroutines to end
	fmt.Println("Waiting for go-routines..")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done..")
}
