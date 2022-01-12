package splitstring

import (
	"fmt"
	"reflect"
	"testing"
)

// 单元测试
/*
  测试函数：
	  1. 必须以Test开头，参数也必须是 *testing.T类型
		a. 接口测试
        b. 覆盖率等
	  2. 测试文件必须以*_test.go格式。
*/
func TestSplitStrBySingleInput(t *testing.T) {
	fmt.Println("=======TestSplitStrBySingleInput======")
	ret := SplitStr("babcbef", "b")
	expect := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(ret, expect) {
		t.Errorf("Expected value: %v, but got the value %v.\n", expect, ret)
	}
}

// 将测试用力封装到slice结构体中
func TestSplitStrByStructSliceGroup(t *testing.T) {
	fmt.Println("=======TestSplitStrByStructSliceGroup======")
	// 将要测试的所有情况抽象为结构体
	type TestUnit struct {
		input    string
		sep      string
		expected []string
	}

	// 使用slice统一管理测试用例
	structSliceGroup := []TestUnit{
		{"babcbef", "b", []string{"", "a", "c", "ef"}},
		{"abcdef", "b", []string{"a", "cdef"}},
		{"a:b:c:d", ":", []string{"a", "b", "c", "d"}},
		{"123456789", "456", []string{"123", "789"}},
		{"桃花坞里桃花庵", "桃", []string{"", "花坞里", "花庵"}},
	}

	for _, scg := range structSliceGroup {
		got := SplitStr(scg.input, scg.sep)
		if !reflect.DeepEqual(got, scg.expected) {
			t.Errorf("Expected value: #%v, but got the value %#v.\n", scg.expected, got)
		}
	}
}

// 将测试用力封装到map中
func TestSplitStrByStructMapGroup(t *testing.T) {
	fmt.Println("=======TestSplitStrByStructMapGroup======")
	// 将要测试的所有情况抽象为结构体
	type TestUnit struct {
		input    string
		sep      string
		expected []string
	}

	// 使用map统一管理测试用例
	structMapGroup := map[string]TestUnit{
		"case 1": {"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case 2": {"abcdef", "b", []string{"a", "cdef"}},
		"case 3": {"a:b:c:d", ":", []string{"a", "b", "c", "d"}},
		"case 4": {"123456789", "456", []string{"123", "789"}},
		"case 5": {"桃花坞里桃花庵", "桃", []string{"", "花坞里", "花庵"}},
	}

	for name, smg := range structMapGroup {
		// 调用testing.T中的Run方法测试
		t.Run(name, func(t *testing.T) {
			got := SplitStr(smg.input, smg.sep)
			if !reflect.DeepEqual(got, smg.expected) {
				t.Errorf("Expected value: #%v, but got the value %#v.\n", smg.expected, got)
			}
		})
	}
}

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
	pkg: wing.com/magic-golang/unittest/splitstring
	cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
	BenchmarkSplitStr
	BenchmarkSplitStr-8   	 4348641	       302.6 ns/op
	PASS
*/
func BenchmarkSplitStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitStr("桃花庵里桃花仙", "桃")
	}
}
