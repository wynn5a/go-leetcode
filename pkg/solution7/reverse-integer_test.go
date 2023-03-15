package solution7

import "testing"

func TestReverse(t *testing.T) {
	// 测试用例1：正常情况，输入整数为正数
	input1 := 123
	expectedOutput1 := 321
	actualOutput1 := reverse(input1)
	if actualOutput1 != expectedOutput1 {
		t.Errorf("TestReverse(%d) = %d; want %d", input1, actualOutput1, expectedOutput1)
	}

	// 测试用例2：正常情况，输入整数为负数
	input2 := -123
	expectedOutput2 := -321
	actualOutput2 := reverse(input2)
	if actualOutput2 != expectedOutput2 {
		t.Errorf("TestReverse(%d) = %d; want %d", input2, actualOutput2, expectedOutput2)
	}
}
