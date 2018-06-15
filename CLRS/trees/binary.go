package trees

type Node struct {
	Val int
	Left, Right *Node
}

func BuildTree(input []int) *Node {
	if len(input) == 0 {
		return nil
	}

	root := buildTreeLevel(nil, 0, input)

	return root
}

func buildTreeLevel(root *Node, level int, input []int) *Node {
	if level < len(input) {
		tmp := &Node{Val: input[level]}
		root = tmp

		root.Left = buildTreeLevel(root.Left, level * 2 + 1, input )
		root.Right = buildTreeLevel(root.Right, level * 2 + 2, input)
		return root
	}
	return nil
}