package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// blocking call
	f("direct")

	// this is a goroutine which will executed concurrently
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going") // anonymous goroutine function call

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
