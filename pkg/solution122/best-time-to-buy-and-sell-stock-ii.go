package solution122

import "math"

// DP
func maxProfit(prices []int) int {
	// 定义sell代表不持有股票时的最大收益，初始化为0
	// 定义hold代表持有股票时的最大收益，初始化为一个非常小的数，代表一开始没有股票
	sell, hold := 0, math.MinInt32

	// 遍历所有的股票价格
	for _, price := range prices {
		// 保存上一次的不持有股票时的最大收益，用于后面计算
		prevSell := sell
		// 当前不持有股票时的最大收益，是上一次不持有股票的最大收益和（上一次持有股票的最大收益+当前股票价格）之间的较大值
		sell = max(sell, hold+price)
		// 当前持有股票时的最大收益，是上一次持有股票的最大收益和（上一次不持有股票的最大收益-当前股票价格）之间的较大值
		hold = max(hold, prevSell-price)
	}

	// 返回最后不持有股票时的最大收益，即为能买卖股票的最大收益
	return sell
}
