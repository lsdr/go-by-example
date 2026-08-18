// https://gobyexample.com/arrays

package main

import "fmt"

func main() {
	var a [7]int
	fmt.Println("emp:", a)

	a[3] = 100
	a[6] = 50
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// cool - changing array size to another number (like 9)
	// will make typecheck complain about size below
	b := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("dcl:", b)

	b = [...]int{100, 4: 400, 500, 600}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
