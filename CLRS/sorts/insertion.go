package sorts

func InsertionSort(input *[]int) {
	tmp := *input

	for i := 1; i < len(tmp); i++ {
		key := tmp[i]

		j := i - 1
		for ; j >= 0 && tmp[j] > key; j-- {
			tmp[j + 1] = tmp[j]
		}

		tmp[j + 1] = key
	}
}
