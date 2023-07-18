package main

import "fmt"

/**
 * @Time    : 2023/7/17 14:54
 * @File    : method1.go
 * @Project : chapter3
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Golang 方法
 */

type Person struct {
	name string
}

func main() {
	p := Person{name: "Alan"}
	p.SayMyName()
}

func (p Person) SayMyName() {
	fmt.Printf("My name is %s\n", p.name)
}
