package searches

func BinarySearch(input []int, val int) int {
	l , r := 0, len(input)

	for  l < r {
		mid := l +  (r - l)  / 2 // or int(uint(l+r)>>1)
		if val == input[mid] {
			return mid
		} else if val >= input[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
