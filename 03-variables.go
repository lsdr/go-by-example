// https://gobyexample.com/variables
package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// declared vars w/o initialization are zero-valued
	var e int
	fmt.Println(e)

	// zero-value for a bool is False
	//var d bool = true
	var d bool
	fmt.Println(d)

	// short for var f string = "short"
	f := "short"
	fmt.Println(f)
}