package solution_lcr_170

func reversePairs(record []int) int {
	return mergeSort(record, 0, len(record)-1)
}

func mergeSort(record []int, start int, end int) int {
	if start >= end {
		return 0
	}
	// 计算中间值
	mid := start + (end-start)/2
	// 递归排序左右两边的数组
	count := mergeSort(record, start, mid) + mergeSort(record, mid+1, end)
	temp := make([]int, end-start+1)
	i, j, k := start, mid+1, 0
	// 计算逆序对
	for i <= mid && j <= end {
		// 如果 record[i] <= record[j]，说明 record[i] 之后的元素都大于 record[j]
		if record[i] <= record[j] {
			temp[k] = record[i]
			i++
			count += j - (mid + 1)
		} else {
			temp[k] = record[j]
			j++
		}
		k++
	}
	// 处理剩余的元素
	for ; i <= mid; i++ {
		temp[k] = record[i]
		k++
		count += j - (mid + 1)
	}
	for ; j <= end; j++ {
		temp[k] = record[j]
		k++
	}
	// 将排序后的数组放回原数组
	// 这里不能直接 record = temp，因为 record 是传入的参数，这样只是修改了 temp 的指向
	for i := 0; i < len(temp); i++ {
		record[start+i] = temp[i]
	}
	return count
}
