func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}
	buy := prices[len(prices)-1]
	var sell int
	var max_profit int = 0

	for i := (len(prices) - 2); i > -1; i-- {

		sell = prices[i]
		if (buy - sell) <= 0 {
			buy = sell
			continue
		}

		if (buy - sell) > max_profit {
			max_profit = (buy - sell)
		}
	}

	return max_profit
}