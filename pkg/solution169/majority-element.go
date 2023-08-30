package solution169

// "Boyer–Moore Majority Vote"算法利用了“大多数元素”的定义。
// 在一个数组中，如果一个元素出现的次数超过了该数组长度的一半，那么它就是大多数元素，也就是说这样的元素在数组中的出现次数比其他所有元素的出现次数总和还要多。
// 我们先设定第一个元素作为主要元素，计数器设置为1，遍历数组时：
// 如果遇到的元素和主要元素相同，计数器增加1；
// 如果遇到的元素和主要元素不同，计数器减少1；
// 当计数器为0时，更换主要元素为当前元素，并且设定计数器为1。
// 想象一下，如果存在大多数元素，即它的数量超过数组长度的一半，那么即使在遍历过程中遇到其他元素时计数器不断减少，因为大多数元素的数量众多，到最后计数器仍然会被主要元素"撑起"。
// 因此，便可以采用这种方式在遍历一次数组后找到大多数元素。
func majorityElement(nums []int) int {
	majority, counter := 0, 0

	for _, num := range nums {
		if counter == 0 {
			majority, counter = num, 1
		} else if num == majority {
			counter++
		} else {
			counter--
		}
	}
	return majority
}
