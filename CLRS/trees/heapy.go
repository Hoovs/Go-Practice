package trees

func maxHeapify(input *[]int, index int) {
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
		maxHeapify(input, largest)
	}
}

func InitHeap(input *[]int) {
	for i := len(*input) / 2 - 1; i >= 0; i-- {
		maxHeapify(input, i)
	}

}

func HeapSort(input *[]int) {
	InitHeap(input)
	tmp := *input
	for i:= len(tmp) - 1; i >= 1; i-- {
		tmp[0], tmp[i] = tmp[i], tmp[0]
		println(tmp[i])
		tmp = tmp[:i]
		maxHeapify(&tmp, 0)
	}
}

