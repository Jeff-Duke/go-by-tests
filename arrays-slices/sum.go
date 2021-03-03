package arrsum

// Sum takes a list of numbers and adds them together, returning the sum
func Sum(nums [5]int) int {
	var sum int
	for i := 0; i < 5; i++ {
		sum += nums[i]
	}
	return sum
}
