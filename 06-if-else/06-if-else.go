// https://gobyexample.com/if-else

package main

import "fmt"

func main() {
	fmt.Println("FizzBuzz w/if-else:")

	for n := range 31 {
		if n == 0 {
			continue
		}

		if n%3 == 0 && n%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if n%3 == 0 {
			fmt.Println("fizz")
		} else if n%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(n)
		}
	}
}
