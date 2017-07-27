// +build OMIT

// START OMIT
package main

func main() {
  // 创建一个 int 类型的 channel
  ch := make(chan int)

  // 启动一个新的匿名 goroutine
  go func() {
      // 发送 42 给 channel
      ch <- 42
  }()
  // 从 channel 读取
  <-ch
}
// STOP OMIT
