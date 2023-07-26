package unitTest

import (
	"reflect"
	"testing"
)

/**
 * @Time    : 2023/7/26 14:16
 * @File    : split_test.go
 * @Project : chapter5
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Golang-单元测试-测试函数
 */

func TestSplit(t *testing.T) {
	got := split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v,got:%v\n", want, got)
	}
}
