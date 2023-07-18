package main

import "fmt"

/**
 * @Time    : 2023/7/17 14:14
 * @File    : interface1.go
 * @Project : chapter3
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Golang 中的接口 interface
 */

// Action 接口
type Action interface {
	run()
	eat()
}

// NullInterface 空接口
type NullInterface interface {
}

// Human 结构体
type Human struct {
	weight float32
	name   string
}

func (h Human) run() {
	fmt.Println("run run run")
}

func (h Human) eat() {
	fmt.Println("eat eat eat")
}

func main() {
	// 空接口可以被赋予任何类型
	var n NullInterface = 1
	fmt.Println("null interface", n)

	var act Action = Human{
		weight: 120,
		name:   "Bob",
	}

	act.run()
	act.eat()
}
