package sorts


func BucketSort(input []int) []int { //TODO only works if all elems are present
	b := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		b[input[i]] = input[i]
	}
	for i := 0; i < len(input) - 1; i++ {
		InsertionSort(&b)
	}
	return b
}