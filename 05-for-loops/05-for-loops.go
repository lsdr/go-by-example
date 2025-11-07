// https://gobyexample.com/for

package main

import "fmt"

func main() {
	fmt.Println("loops are vibes!")

	fmt.Println("1. old school C-like loop:")
	i := 1
	for i <= 3 {
		fmt.Println("[old school]:", i)
		i = i + 1
	}

	fmt.Println("2. for loop:")
	for j := 0; j < 3; j++ {
		fmt.Println("[for loop]:", j)
	}

	fmt.Println("3. using range in loop:")
	for i := range 3 {
		fmt.Println("[range]:", i)
	}

	fmt.Println("4. FizzBuzz time:")
	for n := range 20 {
		if n == 20 {
			break
		}

		if n%3 == 0 {
			fmt.Print("fizz")
		}

		if n%5 == 0 {
			fmt.Print("buzz")
		}

		continue
	}
}
