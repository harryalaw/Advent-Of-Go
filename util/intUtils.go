package util

func IntMax(nums ...int) int {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func IntMin(nums ...int) int {
	min := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func IntSum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func IntAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
