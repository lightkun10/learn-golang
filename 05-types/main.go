package main

import (
	"fmt"
)

var y = 43

var z string = "Shaken, not stirred"
var a string = `James said, "Shaken, not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
