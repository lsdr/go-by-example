// https://gobyexample.com/constants
package main

import "fmt"

const s string = "constant"

func main() {
	fmt.Println(s)

	const s = "wrong?"
	fmt.Println(s)

	// breaks compilation
	//const s = 2
	//fmt.Println(s)

	const n = 60
}