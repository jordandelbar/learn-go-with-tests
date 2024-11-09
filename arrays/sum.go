package arrays

func Sum(array [5]int) int {
	var sum int

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return sum
}
