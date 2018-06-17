package trees

func MaxHeapify(input *[]int, index int) {
	tmp := *input
	left := 2 * index + 1
	largest := index
	if left <= len(tmp) - 1 && tmp[left] > tmp[index] {
		largest = left
	}

	right := 2 * index + 2
	if right <= len(tmp) - 1 && tmp[right] > tmp[largest] {
		largest = right
	}

	if largest != index {
		tmp[index], tmp[largest] = tmp[largest], tmp[index]
		MaxHeapify(input, largest)
	}
}

func InitHeap(input *[]int) {
	for i := len(*input) / 2 - 1; i >= 0; i-- {
		MaxHeapify(input, i)
	}

}

