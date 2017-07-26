// +build OMIT

package main

import "fmt"

// START OMIT
func main() {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		out := make(chan int)
		go Filter(ch, out, prime)
		ch = out
	}
}
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
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
