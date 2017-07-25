// +build OMIT

// START OMIT
package main

func main() {
    ch := make(chan int)

    go func() {
        ch <- 42
    }()

    <-ch
}
// STOP OMIT
