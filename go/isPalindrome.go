package main

func isPalindrome(x int) bool {

	// Negative numbers are not palindrome numbers
	if x < 0 {
		return false
	}
	// Reverse the integer
	new_num := x
	var rev int = 0
	for new_num > 0 {

		last_digit := new_num % 10
		rev = (rev * 10) + last_digit
		new_num = new_num / 10
	}

	return rev == x
}
