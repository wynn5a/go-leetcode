package solution134

func canCompleteCircuit(gas []int, cost []int) int {
	total, subTotal, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		subTotal += gas[i] - cost[i]
		if subTotal < 0 {
			start = i + 1
			subTotal = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
