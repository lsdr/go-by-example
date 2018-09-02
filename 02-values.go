// https://gobyexample.com/values
package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0 / 3.0 =", 7.0/3.0)

	// bool ops
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// coersion
	// both output "5" as integer
	fmt.Println("10 / 2.0 =", 10/2.0)
	fmt.Println("10 / 2 =", 10.0/2)

	// the following line raises "invalid operation"
	// * GoLand complains and prevent run/build
	//fmt.Println("go" + 3)
}