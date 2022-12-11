package util

// provided by GPT3
func Lcm(a, b int) int {
	// Calculate the absolute value of a and b.
	absA := a
	if absA < 0 {
		absA = -absA
	}
	absB := b
	if absB < 0 {
		absB = -absB
	}

	// Calculate the greatest common divisor of a and b.
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// Use the formula lcm(a,b) = (a * b) / gcd(a, b) to calculate the
	// lowest common multiple.
	return (absA * absB) / gcd(absA, absB)
}

// Returns the lowest common multiple of a list of integers.
func LcmOfList(numbers []int) int {
	// Initialize the lowest common multiple to be the first number in the list.
	lcm := numbers[0]

	// Calculate the lcm of the remaining numbers in the list.
	for _, number := range numbers[1:] {
		lcm = Lcm(lcm, number)
	}

	return lcm
}
