// +build OMIT
package main

import "fmt"

type Point struct {
	x float64
	y float64
}

func main() {
	// START OMIT
	var x1 = map[Point]int{
		Point{1.0, 2.1}: 1,
	}
	var x2 = map[Point]int{
		{2.1, 1.0}: 2,
	}
	// STOP OMIT
	fmt.Println("x1", x1)
	fmt.Println("x2", x2)
}
