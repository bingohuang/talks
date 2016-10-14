// +build OMIT
package main

import "fmt"

func main() {
	// START OMIT
	var a [5]int = [5]int{2, 3, 4, 5, 6}
	for k, v := range a {
		fmt.Println(k, v)
	}
	for k := range a {
		fmt.Println(k)
	}
	for _ = range a {
		fmt.Println("print without care about the key and value")
	}
	for range a {
		fmt.Println("new syntax – print without care about the key and value")
	}
	// STOP OMIT
}
