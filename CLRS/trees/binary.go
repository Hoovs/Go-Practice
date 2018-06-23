package trees


type BinarySearchTree struct {
}


func (t BinarySearchTree) Add(val int){
	println("BST")
}

func (t BinarySearchTree) Remove(val int) int {
	return 0
}

func (t BinarySearchTree) Search(val int) bool {
	return false
}

func NewBST() *BinarySearchTree {
	return &BinarySearchTree{}
}


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