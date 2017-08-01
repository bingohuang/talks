// +build OMIT

package main

import "fmt"

// START OMIT
func main() {
	in := make(chan int)
	go Generate(in)
	for i := 0; i < 10; i++ {
		prime := <-in
		fmt.Println(prime)
		out := make(chan int)
		go Filter(in, out, prime)
		in = out
	}
}
func Generate(in chan<- int) {
	for i := 2; ; i++ {
		in <- i
	}
}
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
// STOP OMIT
