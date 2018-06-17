package sorts

import "github.com/Go-Practice/CLRS/trees"

func HeapSort(input *[]int) {
	trees.InitHeap(input)
	tmp := *input
	for i:= len(tmp) - 1; i >= 1; i-- {
		tmp[0], tmp[i] = tmp[i], tmp[0]
		println(tmp[i])
		tmp = tmp[:i]
		trees.MaxHeapify(&tmp, 0)
	}
}