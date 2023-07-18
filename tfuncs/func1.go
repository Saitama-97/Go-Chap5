package main

import "fmt"

/**
 * @Time    : 2023/7/17 14:45
 * @File    : func1.go
 * @Project : chapter3
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Golang 函数
 */

func main() {
	sum := add(3, 4)
	fmt.Println(sum)
}

// add 函数，计算 i 和 i2 之和
func add(i int, i2 int) int {
	return i + i2
}
