package palindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	originX := x

	var reverted int

	for x > 0 {
		reverted = reverted*10 + x%10
		x /= 10
	}

	return originX == reverted
}
