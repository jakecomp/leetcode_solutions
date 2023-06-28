package main

func removeElement(nums []int, val int) int {

	var index int = 0
	for _, num := range nums {
		if num != val {
			nums[index] = num
			index++
		}
	}
	return index
}
