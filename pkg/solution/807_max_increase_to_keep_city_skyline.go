package solution

import (
	"fmt"
	"math/rand"
)

// Input: grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
// Output: 35
// Explanation: The building heights are shown in the center of the above image.
// The skylines when viewed from each cardinal direction are drawn in red.
// The grid after increasing the height of buildings without affecting skylines is:
// gridNew = [ [8, 4, 8, 7],
//
//	[7, 4, 7, 7],
//	[9, 4, 8, 7],
//	[3, 3, 3, 3] ]
func maxIncreaseKeepingSkyline(grid [][]int) int {
	//find max value of row and column
	l := len(grid)
	maxRow := make([]int, l)
	maxCol := make([]int, l)

	for i, row := range grid {
		for j, value := range row {
			maxRow[i] = Max(maxRow[i], value)
			maxCol[j] = Max(maxCol[j], value)
		}
	}

	fmt.Printf("max col: %v \n", maxCol)
	fmt.Printf("max row: %v \n", maxRow)

	var result int
	for i, row := range grid {
		for j, value := range row {
			result += Min(maxRow[i], maxCol[j]) - value
		}
	}
	return result
}

func findMaxColumnValue(grid [][]int) []int {
	var result []int
	var max = 0
	var index = 0
	var length = len(grid)

	for length > index {
		for _, row := range grid {
			if row[index] > max {
				max = row[index]
			}
		}
		index++
		result = append(result, max)
		max = 0
	}

	return result
}

func findMaxRowValue(grid [][]int) []int {
	var result []int
	for _, row := range grid {
		result = append(result, MaxInArray(row))
	}
	return result
}

func Run807() {
	fmt.Println("Run 807 solution begin:")

	//Generate a random array of length n
	var m [][]int
	fmt.Println("[")
	for i := 0; i < 4; i++ {
		fmt.Print("[")
		var m1 []int
		for j := 0; j < 4; j++ {
			m1 = append(m1, rand.Intn(20))
			fmt.Printf("%v ", m1[j])
		}
		m = append(m, m1)
		fmt.Println("]")
	}
	fmt.Println("]")

	fmt.Printf("Max increase is %v \n", maxIncreaseKeepingSkyline(m))
}
