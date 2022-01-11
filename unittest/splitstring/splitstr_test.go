package splitstring

import (
	"reflect"
	"testing"
)

// 单元测试
/*
  测试函数：必须以Test开头，参数也必须是 *testing.T类型
*/
func TestSplitStr(t *testing.T) {
	ret := SplitStr("babcbef", "b")
	expect := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(ret, expect) {
		t.Errorf("Expected value: %v, but got the value %v.\n", expect, ret)
	}
}
