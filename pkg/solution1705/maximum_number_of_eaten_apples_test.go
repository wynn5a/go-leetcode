package solution1705

import "testing"

//解释：你可以吃掉 7 个苹果：
//- 第一天，你吃掉第一天长出来的苹果。
//- 第二天，你吃掉一个第二天长出来的苹果。
//- 第三天，你吃掉一个第二天长出来的苹果。过了这一天，第三天长出来的苹果就已经腐烂了。
//- 第四天到第七天，你吃的都是第四天长出来的苹果。
func Test_case_1(t *testing.T) {
	var (
		apples   = []int{1, 2, 3, 5, 2}
		days     = []int{3, 2, 1, 4, 2}
		expected = 7
	)
	actual := eatenApples(apples, days)
	if actual != expected {
		t.Errorf("eatenApples(%v, %v) = %d; expected %d", apples, days, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		apples   = []int{3, 0, 0, 0, 0, 2}
		days     = []int{3, 0, 0, 0, 0, 2}
		expected = 5
	)
	actual := eatenApples(apples, days)
	if actual != expected {
		t.Errorf("eatenApples(%v, %v) = %d; expected %d", apples, days, actual, expected)
	}
}
func Test_case_3(t *testing.T) {
	var (
		apples   = []int{2, 1, 10}
		days     = []int{2, 10, 1}
		expected = 4
	)
	actual := eatenApples(apples, days)
	if actual != expected {
		t.Errorf("eatenApples(%v, %v) = %d; expected %d", apples, days, actual, expected)
	}
}
