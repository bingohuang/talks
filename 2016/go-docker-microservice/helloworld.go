// +build OMIT

// START OMIT
package main

import "fmt"

// 反转字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // 先计算出所有右值，然后再依次完成赋值操作
	}
	return string(r)
}

func main() {
	fmt.Println(Reverse("!界世 ,olleH"))
}
// STOP OMIT
