package unitTest

import "strings"

/**
 * @Time    : 2023/7/26 14:12
 * @File    : split.go
 * @Project : chapter5
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Golang-单元测试-函数
 */

func split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
