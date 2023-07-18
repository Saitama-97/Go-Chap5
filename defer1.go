package main

import "fmt"

/**
 * @Time    : 2023/7/17 09:28
 * @File    : defer1.go
 * @Project : chapter3
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    :
 */

func main() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Println(i) }()
	}
}
