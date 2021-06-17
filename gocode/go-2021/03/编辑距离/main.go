package main

import "fmt"

/*
给定两个字符串s1和s2，计算将s1转换成s2所使用的最少操作数。
你可以对一个字符串进行如下三种操作：
1.插入一个字符
2.删除一个字符
3.替换一个字符

示例1：
输入： s1 = "horse", s2 = "ros"
输出：3
解释：
horse -> rorse (将'h'替换为 'r')
rorse -> rose （删除 'r'）
rose -> ros （删除 'e'）

示例2：
输入：s1 = "intention", s2 = "execution"
输出： 5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
*/

/*
解决两个字符串的动态规划问题，一般都是用两个指针 i,j 分别指向两个字符串的最后，然后一步步往前走，缩小问题的规模。
*/

type dp struct {
	val    int
	choice []int // 0: 啥也不做，1：代表插入，2：代表删除，3：代表替换
}

func main() {
	//min := minDistance("intention", "execution")
	//println(min)
	dp := minDistance2("intention", "execution")
	fmt.Println(dp)
}

func minDistance(s1, s2 string) int {
	rune1, rune2 := []rune(s1), []rune(s2)
	l1, l2 := len(rune1), len(rune2)
	var dp [][]int
	for i := 0; i < l1+1; i++ {
		dp = append(dp, make([]int, l2+1))
	}
	for i := 1; i <= l1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= l2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if rune1[i-1] == rune2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 对s1进行操作
				// 假设先进行插入操作
				dp[i][j] = dp[i][j-1] + 1
				// 与删除操作做对比，取最小值
				if dp[i][j] > dp[i-1][j]+1 {
					dp[i][j] = dp[i-1][j] + 1
				}
				// 与替换操作做对比，取最小值
				if dp[i][j] > dp[i-1][j-1]+1 {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
		}
	}
	return dp[l1][l2]
}

func minDistance2(s1, s2 string) dp {
	r1, r2 := []rune(s1), []rune(s2)
	l1, l2 := len(r1), len(r2)
	var dps [][]dp
	for i := 0; i < l1+1; i++ {
		dps = append(dps, make([]dp, l2+1))
	}
	for i := 1; i <= l1; i++ {
		var d dp
		d.val = i
		for c := 0; c < i; c++ {
			// s2为空字符串，s2长度即为执行删除操作次数
			d.choice = append(d.choice, 2)
		}
		dps[i][0] = d
	}
	for j := 1; j <= l2; j++ {
		var d dp
		d.val = j
		for c := 0; c < j; c++ {
			// s1为空字符串，s1长度即为执行插入操作次数
			d.choice = append(d.choice, 1)
		}
		dps[0][j] = d
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if r1[i-1] == r2[j-1] {
				dps[i][j] = dps[i-1][j-1]
			} else {
				// 对s1进行操作
				// 假设先进行插入操作
				dps[i][j].val = dps[i][j-1].val + 1
				dps[i][j].choice = append(dps[i][j-1].choice, 1)
				if dps[i][j].val > dps[i-1][j].val+1 {
					// 与删除操作做对比，取最小值
					dps[i][j].val = dps[i-1][j].val + 1
					dps[i][j].choice = append(dps[i-1][j].choice, 2)
				}
				if dps[i][j].val > dps[i-1][j-1].val+1 {
					// 与替换操作做对比，取最小值
					dps[i][j].val = dps[i-1][j-1].val + 1
					dps[i][j].choice = append(dps[i-1][j-1].choice, 3)
				}
			}
		}
	}
	return dps[l1][l2]
}
