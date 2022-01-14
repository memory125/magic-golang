package benchmark

import (
	"testing"
	splitstr "wing.com/magic-golang/unittest/splitstring"
)

// benchmark 基准测试
/*
测试函数：
    1. 必须以Test开头，参数也必须是 *testing.B类型
      a. benchmark基准测试
      b. memory基准测试，如：申请了多少次内存
      c. 性能比较
	  d. 并行测试，示例如下：
			b.RunParallel(func(pb *testing.PB) {
				// 测试代码
			})
    2. 测试文件必须以*_test.go格式。
*/
/*
	goos: windows
	goarch: amd64
	pkg: wing.com/magic-golang/unittest/benchmark
	cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
	BenchmarkSplitStr
	BenchmarkSplitStr-8   	 4836988	       250.1 ns/op
	PASS
*/
func BenchmarkSplitStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		splitstr.SplitStr("桃花庵里桃花仙", "桃")
	}
}
