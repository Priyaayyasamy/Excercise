package main

import "fmt"

func add(x float64, y float64) float64 {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
