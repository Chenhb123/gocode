package main

import "fmt"

/*
给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，
问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。算法的函数签名如下：

// coins 中是可选硬币面值，amount 是目标金额
int coinChange(int[] coins, int amount);

比如说 k = 3，面值分别为 1，2，5，总金额 amount = 11。那么最少需要 3 枚硬币凑出，即 11 = 5 + 5 + 1。
*/

type dp struct {
	Amount int   // 需要硬币数
	Coin   []int // 具体硬币数值
}

func main() {
	//count := coinChange([]int{1, 2, 5}, 11)
	//println(count)
	dp := coinChange2([]int{1, 2, 5}, 11)
	fmt.Println(dp)
}

/*


状态转移方程: dp(n) = { 0, n=0; -1, n<0; min(1+dp(n-coin)),coin∈coins }

*/
func coinChange(coins []int, amount int) int {
	// 定义数组
	dptable := make([]int, amount+1)
	for i := 1; i < len(dptable); i++ {
		dptable[i] = amount + 1
	}
	for i := 0; i < len(dptable); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dptable[i-coin]+1 < dptable[i] {
				dptable[i] = dptable[i-coin] + 1
			}
		}
	}
	if dptable[amount] == amount+1 {
		return -1
	}
	return dptable[amount]
}

func coinChange2(coins []int, amount int) dp {
	// 定义数组
	dps := make([]dp, amount+1)
	for i := 1; i < len(dps); i++ {
		dps[i].Amount = amount + 1
		dps[i].Coin = []int{}
	}
	for i := 1; i < len(dps); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dps[i-coin].Amount+1 < dps[i].Amount {
				dps[i].Amount = dps[i-coin].Amount + 1
				dps[i].Coin = append(dps[i-coin].Coin, coin)
			}
		}
	}
	if dps[amount].Amount == amount+1 {
		dps[amount].Amount = -1
		dps[amount].Coin = []int{}
	}
	return dps[amount]
}
