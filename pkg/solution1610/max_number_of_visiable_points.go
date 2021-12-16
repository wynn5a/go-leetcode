package solution1610

import (
	"math"
	"sort"
)

//给你一个点数组 points 和一个表示角度的整数 angle ，你的位置是 location ，其中 location = [posx, posy] 且 points[i] = [xi, yi] 都表示 X-Y 平面上的整数坐标。
//你 不能 进行移动改变位置，但可以通过 自转 调整观测角度
//对于每个点，如果由该点、你的位置以及从你的位置直接向东的方向形成的角度 位于你的视野中 ，那么你就可以看到它。
//
//同一个坐标上可以有多个点。你所在的位置也可能存在一些点，但不管你的怎么旋转，总是可以看到这些点。同时，点不会阻碍你看到其他点。
//
//返回你能看到的点的最大数目
//
//用到了 极角 的知识
func visiblePoints(points [][]int, angle int, location []int) int {
	sameCnt := 0
	var polarDegrees []float64
	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			sameCnt++
		} else {
			polarDegrees = append(polarDegrees, math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0])))
		}
	}
	sort.Float64s(polarDegrees)

	n := len(polarDegrees)
	for _, deg := range polarDegrees {
		polarDegrees = append(polarDegrees, deg+2*math.Pi)
	}

	maxCnt := 0
	degree := float64(angle) * math.Pi / 180
	for i, deg := range polarDegrees[:n] {
		j := sort.Search(n*2, func(j int) bool { return polarDegrees[j] > deg+degree })
		if j-i > maxCnt {
			maxCnt = j - i
		}
	}
	return sameCnt + maxCnt
}
