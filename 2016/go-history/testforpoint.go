// +build OMIT
package main

import "fmt"

// START OMIT
type T int

func (T) M() {
	fmt.Println("Call M")
}

var i *T
var x **T

func main() {
  i = new(T)
  x = &i

  i.M()
  x.M()
}
// STOP OMIT
