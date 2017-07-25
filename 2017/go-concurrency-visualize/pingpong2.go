// +build OMIT

package main

import "time"

// START OMIT
func main() {
	var Ball int
	table := make(chan int)

	// 2个玩家
	go player(table)
	go player(table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int) {
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
// STOP OMIT
