package splitstring

import (
	"reflect"
	"testing"
)

func TestSplitStr(t *testing.T) {
	ret := SplitStr("babcbef", "b")
	expect := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(ret, expect) {
		t.Errorf("Expected value: %v, but got the value %v.\n", expect, ret)
	}
}
