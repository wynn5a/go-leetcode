package solution825

//在社交媒体网站上有 n 个用户。给你一个整数数组 ages ，其中 ages[i] 是第 i 个用户的年龄。
//如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：
//age[y] <= 0.5 * age[x] + 7
//age[y] > age[x]
//age[y] > 100 && age[x] < 100
//否则，x 将会向 y 发送一条好友请求。
//注意，如果 x 向 y 发送一条好友请求，y 不必也向 x 发送一条好友请求。另外，用户不会向自己发送好友请求。
//
//返回在该社交媒体网站上产生的好友请求总数。
func numFriendRequests(ages []int) int {
	//满足下列2个条件，就可以发送
	//1. ages[y] <= ages[x] < (ages[y] - 7) * 2 && ages[y] <= 100
	//2. 0.5 * ages[x] + 7 < ages[y] <= ages[x] && ages[x] >= 100
	//使用前缀和
	ans := 0
	cnts, presum := make([]int, 120), make([]int, 121)
	for _, age := range ages {
		cnts[age-1]++
	}
	for i := 1; i < 121; i++ {
		presum[i] = presum[i-1] + cnts[i-1]
	}
	for _, age := range ages {
		if v := presum[age] - presum[age/2+7] - 1; v > 0 {
			ans += v
		}
	}
	return ans
}

//暴力循环不可取
//func numFriendRequests(ages []int) int {
//	res := 0
//
//	for x := 0; x < len(ages); x++ {
//		ax := ages[x]
//		for y := 0; y < len(ages); y++ {
//			if x == y {
//				continue
//			}
//			ay := ages[y]
//			if ay > 100 && ax < 100 {
//				continue
//			}
//			if ay > ax {
//				continue
//			}
//
//			if ay <= int(float32(ax)*0.5)+7 {
//				continue
//			}
//
//			res ++
//		}
//	}
//
//	return res
//}
