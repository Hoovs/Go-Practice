package main

import "github.com/Go-Practice/CLRS/searches"

func main() {
	//toSort := []int{1,3,5,7,9,2,4,6,8,0}
	//sorts.Quicksort(&toSort, 0, len(toSort) - 1)
	//output := sorts.BucketSort(toSort)

	println(searches.BinarySearch([]int{1,3,5,7,9,11,12,13,14,15,16,20,30,40,50,60,70,80,90,100}, 3))

}