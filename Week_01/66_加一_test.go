package Week_01

//
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i]%10 != 0 {
			return digits
		}
	}
	res := make([]int, len(digits)+1)
	res[0] = 1
	return res
}
