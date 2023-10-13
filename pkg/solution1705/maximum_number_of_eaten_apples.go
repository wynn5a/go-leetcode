package solution1705

import (
	"container/heap"
)

// 有一棵特殊的苹果树，一连 n 天，每天都可以长出若干个苹果。在第 i 天，树上会长出 apples[i] 个苹果，
// 这些苹果将会在 days[i] 天后（也就是说，第 i + days[i] 天时）腐烂，变得无法食用。
// 也可能有那么几天，树上不会长出新的苹果，此时用 apples[i] == 0 且 days[i] == 0 表示。
//
// 你打算每天 最多 吃一个苹果来保证营养均衡。注意，你可以在这 n 天之后继续吃苹果。
//
// 给你两个长度为 n 的整数数组 days 和 apples ，返回你可以吃掉的苹果的最大数目。
// --》先吃快过期的，优先队列
func eatenApples(apples, days []int) (ans int) {
	n := len(apples)
	h := hp{}
	for i := 0; i < n || len(h) > 0; i++ {
		if i < n && apples[i] > 0 {
			heap.Push(&h, pair{i + days[i], apples[i]})
		}
		for len(h) > 0 && (h[0].end <= i || h[0].left == 0) {
			heap.Pop(&h)
		}
		if len(h) > 0 {
			ans++
			h[0].left--
		}
	}
	return
}

type pair struct{ end, left int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) {
	a := *h
	v = a[len(a)-1]
	*h = a[:len(a)-1]
	return
}
