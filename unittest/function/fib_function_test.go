package function

import (
	"fmt"
	"reflect"
	"testing"
	"wing.com/magic-golang/unittest/fib"
)

/*
	=== RUN   TestFib
	=======TestFib======
	=== RUN   TestFib/5!
	=== RUN   TestFib/6!
	=== RUN   TestFib/7!
	=== RUN   TestFib/1!
	=== RUN   TestFib/2!
	=== RUN   TestFib/3!
	=== RUN   TestFib/4!
	--- PASS: TestFib (0.00s)
		--- PASS: TestFib/5! (0.00s)
		--- PASS: TestFib/6! (0.00s)
		--- PASS: TestFib/7! (0.00s)
		--- PASS: TestFib/1! (0.00s)
		--- PASS: TestFib/2! (0.00s)
		--- PASS: TestFib/3! (0.00s)
		--- PASS: TestFib/4! (0.00s)
	PASS
*/
func TestFib(t *testing.T) {
	fmt.Println("=======TestFib======")
	// 将要测试的所有情况抽象为结构体
	type FibTestUnit struct {
		input    uint64
		expected uint64
	}

	// 使用map统一管理测试用例
	structMapGroup := map[string]FibTestUnit{
		"1!": {1, 1},
		"2!": {2, 2},
		"3!": {3, 6},
		"4!": {4, 24},
		"5!": {5, 120},
		"6!": {6, 720},
		"7!": {7, 5040},
	}

	for name, smg := range structMapGroup {
		// 调用testing.T中的Run方法测试
		t.Run(name, func(t *testing.T) {
			got := fib.Fib(smg.input)
			if !reflect.DeepEqual(got, smg.expected) {
				t.Errorf("Expected value: #%v, but got the value %#v.\n", smg.expected, got)
			}
		})
	}

}
