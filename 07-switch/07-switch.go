// https://gobyexample.com/switch

package main

import (
	"fmt"
	"time"
)

func main() {
	weekday := time.Now().Weekday()

	switch weekday {
	case time.Monday:
		fmt.Println("Monday :(")
	case time.Thursday:
		fmt.Println("It's Friday Eve!!")
	case time.Friday:
		fmt.Println("Friday, baby!")
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!!")
	default:
		fmt.Println("It's a weekday, carry on...")
	}
}
