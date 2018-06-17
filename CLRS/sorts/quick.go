package sorts

func partion(input *[]int, l, r int) int {
	tmp := *input

	x, i := tmp[r], l - 1

	for j := l; j < r ; j++ {
		if tmp[j] <= x {
			i++
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}
	}

	tmp[i + 1], tmp[r] = tmp[r], tmp[i + 1]
	return i + 1
}

func Quicksort(input *[]int, l, r int) {
	if l < r {
		q := partion(input, l, r)
		Quicksort(input, l, q - 1)
		Quicksort(input, q + 1, r)
	}
}