package solution18

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums) // 首先将数组排序
	var res [][]int

	/* 双指针法，循环1 */

	for i := 0; i < len(nums)-3; i++ { // 从数组第一个数开始循环，到倒数第四个数结束
		if i > 0 && nums[i] == nums[i-1] { // 如果当前数和上一个数相同，则跳过
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target { // 若当前最小的和已经大于目标值，则跳出循环
			break
		}
		if nums[i]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target { // 同上，当前最大的和已经小于目标值
			continue
		}

		/* 双指针法，循环2 */

		for j := i + 1; j < len(nums)-2; j++ { // 从第二个数开始循环，到倒数第三个数结束
			if j > i+1 && nums[j] == nums[j-1] { // 如果当前数和上一个数相同，则跳过
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target { // 若当前最小的和已经大于目标值，则跳出循环
				break
			}
			if nums[i]+nums[j]+nums[len(nums)-1]+nums[len(nums)-2] < target { // 同上，当前最大的和已经小于目标值
				continue
			}

			/* 双指针法，循环3 */

			left, right := j+1, len(nums)-1 // 定义左右指针
			for left < right {              // 当左右指针不重合时
				sum := nums[i] + nums[j] + nums[left] + nums[right] // 计算4个数的和
				if sum == target {                                  // 如果和等于目标值
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]}) // 将结果加入到res中
					left++                                                              // 左指针右移
					right--                                                             // 右指针左移
					for left < right && nums[left] == nums[left-1] {                    // 跳过当前左指针指向的重复数字
						left++
					}
					for left < right && nums[right] == nums[right+1] { // 跳过当前右指针指向的重复数字
						right--
					}
				} else if sum < target { // 如果和比目标值小，则左指针右移
					left++
				} else { // 和比目标值大，则右指针左移
					right--
				}
			}
		}
	}

	return res // 返回结果
}
