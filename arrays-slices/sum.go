package arrsum

// Sum takes a list of numbers and adds them together, returning the sum
func Sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

// SumAll takes any number of collections and returns a slice of their sums
func SumAll(numsToSum ...[]int) (sums []int) {
	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}
