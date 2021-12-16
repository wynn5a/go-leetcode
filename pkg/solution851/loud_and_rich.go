package solution851

//有一组 n 个人作为实验对象，从 0 到 n - 1 编号，其中每个人都有不同数目的钱，以及不同程度的安静值（quietness）。为了方便起见，我们将编号为x的人简称为 "personx"。
//
//给你一个数组 richer ，其中 richer[i] = [ai, bi] 表示 personai比 personbi更有钱。另给你一个整数数组 quiet ，其中quiet[i] 是 person i 的安静值。
//richer 中所给出的数据 逻辑自恰（也就是说，在 person x 比 person y 更有钱的同时，不会出现 person y 比 person x 更有钱的情况 ）。
//
//现在，返回一个整数数组 answer 作为答案，其中answer[x] = y的前提是，在所有拥有的钱肯定不少于person x的人中，person y是最安静的人（也就是安静值quiet[y]最小的人）。
func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)

	//a simple graph
	g := make([][]int, n)
	for _, r := range richer {
		g[r[1]] = append(g[r[1]], r[0])
	}

	//init answer with -1
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	//depth first search
	for i := 0; i < n; i++ {
		dfs(i, ans, quiet, g)
	}
	return ans
}

func dfs(x int, ans []int, quiet []int, graph [][]int) {
	if ans[x] != -1 {
		return
	}
	ans[x] = x
	for _, y := range graph[x] {
		dfs(y, ans, quiet, graph)
		if quiet[ans[y]] < quiet[ans[x]] {
			ans[x] = ans[y]
		}
	}
}
