package main

import "fmt"

func main() {
	fmt.Println("Hello...")

	foo()

	fmt.Println("Something more")

	for i := 0; i < 40; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo!")
}

func bar() {
	fmt.Println("exit...")
}
