package question

func largestAltitude(gain []int) int {
	var sum, max int
	for _, v := range gain {
		sum += v
		if sum > max {
			max = sum
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
