package main

import "fmt"

/*
你面前有一栋从 1 到 N 共 N 层的楼，然后给你 K 个鸡蛋（K 至少为 1）。
现在确定这栋楼存在楼层 0 <= F <= N，在这层楼将鸡蛋扔下去，鸡蛋恰好没摔碎（高于 F 的楼层都会碎，低于
F 的楼层都不会碎）。现在问你，最坏情况下，你至少要扔几次鸡蛋，才能确定这个楼层 F 呢？
*/

/*
状态转移方程： dp(K, N) = min{0<=i<=N,max(dp(K-1, i-1), dp(K, N-i))+1}
*/

func main() {
	count := superEggDrop2(100, 1000)
	println(count)
}

/*
转换为已知鸡蛋数和最大扔鸡蛋次数，求最多可以测试多高楼层的问题
*/
func superEggDrop2(k, n int) int {
	var dp [][]int
	for i := 0; i <= k; i++ {
		dp = append(dp, make([]int, n+1))
	}
	var m int
	for dp[k][m] < n {
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = dp[i-1][m-1] + dp[i][m-1] + 1
		}
	}
	return m
}

func superEggDrop(k, n int) int {
	dict := make(map[string]int)
	return dp2(k, n, dict)
}

func dp(k, n int, dict map[string]int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	if val, ok := dict[fmt.Sprintf("%d-%d", k, n)]; ok {
		return val
	}
	min := n
	for i := 1; i <= n; i++ {
		var res int
		broke := dp(k-1, i-1, dict)
		unbroke := dp(k, n-i, dict)
		if broke > unbroke {
			res = broke + 1
		} else {
			res = unbroke + 1
		}
		if res < min {
			min = res
		}
	}
	dict[fmt.Sprintf("%d-%d", k, n)] = min
	return min
}

func dp2(k, n int, dict map[string]int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	str := fmt.Sprintf("%d-%d", k, n)
	if val, ok := dict[str]; ok {
		return val
	}
	res := n
	lo, hi := 1, n
	for lo <= hi {
		mid := (lo + hi) / 2
		broke := dp2(k-1, mid-1, dict)
		unbroke := dp2(k, n-mid, dict)
		if broke > unbroke {
			if broke+1 < res {
				res = broke + 1
			}
			hi = mid - 1
		} else {
			if unbroke+1 < res {
				res = unbroke + 1
			}
			lo = mid + 1
		}
	}
	dict[str] = res
	return res
}
