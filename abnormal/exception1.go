package main

import "fmt"

/**
 * @Time    : 2023/7/25 15:09
 * @File    : exception1.go
 * @Project : chapter5
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Golang-异常处理
 */

func main() {
	m := make(map[int]int)
	m[0] = 1
	m[1] = 2
	m[2] = 4
	for k, v := range m {
		res := divide(k, v)
		fmt.Printf("%v/%v = %v\n", v, k, res)
	}
}

func divide(k int, v int) (ret int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("除数不能为0，返回-1以作表示")
			ret = -1
		}
	}()
	return v / k
}
