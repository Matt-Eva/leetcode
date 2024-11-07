package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

func main(){


}

func maxProfit(prices []int) int {
    
	profit := 0
	min := prices[0]
	
	for i := 0; i < len(prices); i++{
		if prices[i] < min {
			min = prices[i]
			continue
		} 

		profit = max(profit, prices[i] - min)
	}

	return profit
}