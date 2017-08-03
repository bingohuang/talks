// +build OMIT
package main

import "fmt"

// START OMIT
func generate(ch chan<- int) {
    for i := 2; ; i++ {
        ch <- i // Send 'i' to channel 'ch'.
    }
}
func filter(src <-chan int, dst chan<- int, prime int) {
    for i := range src { // Loop over values received from 'src'.
        if i%prime != 0 {
            dst <- i // Send 'i' to channel 'dst'.
        }
    }
}
func main() {
	ch := make(chan int)
	go generate(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		out := make(chan int)
		go filter(ch, out, prime)
		ch = out
	}
}
// STOP OMIT
