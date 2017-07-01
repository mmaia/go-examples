package main

import (
	"github.com/mmaia/go-playground/mathapp"
	"fmt"
)

func main() {
	fmt.Println("sqrt of 3 = ", mathapp.Sqrt(3))
	fmt.Printf("sum of 3 + 11 = %v, while 255 - 173,5 = %v", 3 + 11, 255 - 173.5)
}