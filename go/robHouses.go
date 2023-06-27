package main

func max_num(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	var option_1 int = nums[0]
	var option_2 int = max_num(nums[0], nums[1])
	var previous_max int

	for i := 2; i < len(nums); i++ {
		previous_max = option_2
		option_2 = max_num(option_1+nums[i], option_2)
		option_1 = previous_max
	}
	return option_2
}
