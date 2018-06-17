package sorts


func CountingSort(input []int) []int { //TODO fix
	b := make([]int, len(input))
	c := make([]int, len(input))

	for i := 0; i < len(input); i++{
		c[input[i]] = c[input[i]] + 1
	}
	for i := 1; i < len(input); i++{
		c[i] = c[i] + c[i - 1]
	}
	for i := len(input) - 1; i > 0 ; i-- {
		b[c[input[i]]] = input[i]
		c[input[i]] = c[input[i]] - 1
	}

	return b
}