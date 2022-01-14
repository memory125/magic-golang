package function

import (
	"fmt"
	"reflect"
	"testing"
	splitstr "wing.com/magic-golang/unittest/splitstring"
)

// 将测试用力封装到map中
/*
	=== RUN   TestSplitStrByStructMapGroup
	=======TestSplitStrByStructMapGroup======
	=== RUN   TestSplitStrByStructMapGroup/case_1
	=== RUN   TestSplitStrByStructMapGroup/case_2
	=== RUN   TestSplitStrByStructMapGroup/case_3
	=== RUN   TestSplitStrByStructMapGroup/case_4
	=== RUN   TestSplitStrByStructMapGroup/case_5
	--- PASS: TestSplitStrByStructMapGroup (0.00s)
		--- PASS: TestSplitStrByStructMapGroup/case_1 (0.00s)
		--- PASS: TestSplitStrByStructMapGroup/case_2 (0.00s)
		--- PASS: TestSplitStrByStructMapGroup/case_3 (0.00s)
		--- PASS: TestSplitStrByStructMapGroup/case_4 (0.00s)
		--- PASS: TestSplitStrByStructMapGroup/case_5 (0.00s)
	PASS
*/
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
			got := splitstr.SplitStr(smg.input, smg.sep)
			if !reflect.DeepEqual(got, smg.expected) {
				t.Errorf("Expected value: #%v, but got the value %#v.\n", smg.expected, got)
			}
		})
	}
}
