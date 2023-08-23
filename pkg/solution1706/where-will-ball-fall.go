package solution1706

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, n)
	// 对网格中的每一列进行运算
	for i := 0; i < n; i++ {
		// 初始位置设为网格的顶部，即第0行，第i列
		x, y := 0, i
		// 直到球移动到底部，或者被卡住为止
		for x < m {
			// 如果当前单元格是向右导向的挡板，且球在最右边或者右边的单元格是向左导向
			if grid[x][y] == 1 && (y == n-1 || grid[x][y+1] == -1) {
				// 球会被卡住，标记为-1，并跳出循环
				y = -1
				break
			}
			// 如果当前单元格是向左导向的挡板，且球在最左边或者左边的单元格是向右导向
			if grid[x][y] == -1 && (y == 0 || grid[x][y-1] == 1) {
				// 球会被卡住，标记为-1，并跳出循环
				y = -1
				break
			}
			// 否则，根据当前单元格的挡板，更新球的位置
			y += grid[x][y]
			// 然后球移动到下一行
			x++
		}
		ans[i] = y
	}
	return ans
}
