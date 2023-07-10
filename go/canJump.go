func canJump(nums []int) bool {

	var can_jump bool = true
	var nums_length int = len(nums)
	var length_required int = 1
	for i := (nums_length - 2); i > -1; i-- {

		num := nums[i]
		if num < length_required {
			can_jump = false
			length_required++
		} else {
			can_jump = true
			length_required = 1
		}
	}
	return can_jump
}