package main

import (
	"GoLearning/chapter8"
	"testing"
)

// 测试文件的名字是北侧是的文件的名字_test
// 测试方法以Test开头
// go test -coverprofile=c.out将测试报告输出到c.out文件中
// go tool cover -html c.out以HTML查看
func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		// 測試用例
		{1, 2, 3},
		{2, 3, 5},
		{0, 0, 1},
	}

	for _, v := range tests {

		if v.a+v.b != v.c {
			// t有報錯方法
			t.Errorf("Error")
		}

	}
}

//cmd: go benchmark
func BenchmarkAdd(bc *testing.B) {
	a, b, c := 1, 2, 3
	for i := 0; i < bc.N; i++ {
		r := chapter8.Add(a, b)
		if r != c {
			bc.Errorf("No")
		}
	}
}
