package main

import (
	"fmt"
	"time"
)

func whatAmI(i any) {
	switch i.(type) {
	case bool:
		fmt.Println("Boolean")
	case int:
		fmt.Println("Integer")
	default:
		fmt.Printf("Not sure: %T\n", i)
	}
}

func main() {
	whatAmI(1)
	whatAmI(true)
	whatAmI("What?")
	whatAmI(time.Sunday)
}
