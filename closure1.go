package main

import "fmt"

/**
 * @Time    : 2023/7/17 09:41
 * @File    : closure1.go
 * @Project : chapter3
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    :
 */

func main() {
	f1 := func1()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
}

func func1() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
