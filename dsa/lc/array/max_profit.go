func maxProfit(prices []int) int {
    //you can try new combos and keep track of max profit
    min, profit, max_profit := prices[0], 0, 0
    for i:=1; i<len(prices);i++ {
        if prices[i] < min {
            min = prices[i]
        }
        profit = prices[i] - min
        if profit > max_profit {
            max_profit = profit
        }
    }
    return max_profit
}