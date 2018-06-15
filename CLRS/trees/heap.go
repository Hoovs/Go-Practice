package trees


func Init(input *[]int) {
	length := len(*input)
	for i := length / 2 - 1; i >= 0; i-- {
		down(input, i)
	}
}

func down(in *[]int, index int) bool {
	input := *in
	left := 2 * index + 1
	if left > len(input) || left < 0 {
		return true
	}

	if right := left + 1; right < len(input) && right < left {
		left = right
	}
	if left > index {
		return false
	}

	input[index], input[left] = input[left], input[index]
	down(in, left)
	return true
}