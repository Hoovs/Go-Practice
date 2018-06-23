package trees


type BinaryTree interface {
	Add(val int)
	Remove(val int) int
	Search(val int) bool
}