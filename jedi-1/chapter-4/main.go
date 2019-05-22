package main

import (
	"fmt"
)

type clutch int

var x clutch

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
