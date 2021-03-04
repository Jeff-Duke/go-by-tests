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

// SumAllTails sums the tails of slices
func SumAllTails(numsToSum ...[]int) (sums []int) {
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
