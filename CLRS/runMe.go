package main

import (
	"github.com/Go-Practice/CLRS/trees"
)

func main() {
	//toSort := []int{1,3,5,7,9,2,4,6,8,0}
	//sorts.Quicksort(&toSort, 0, len(toSort) - 1)
	//output := sorts.BucketSort(toSort)

	//println(searches.BinarySearch([]int{1,3,5,7,9,11,12,13,14,15,16,20,30,40,50,60,70,80,90,100}, 3))

	trie := trees.InitTrie()


	trie.Insert("matt")
	trie.Insert("matthew")
	trie.Insert("matthewhoovs")

	println(trie.FindPrefix("matt"))
	println(trie.FindPrefix("matthew"))
	println(trie.FindPrefix("john"))

	println(trie.Contains("matt"))
	println(trie.Contains("matth"))
	println(trie.Contains("matthew"))
	println(trie.Contains("john"))
}