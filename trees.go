package trees

type Tree interface {
	Parent() *Tree
	Children() []*Tree
}

func Traversal[T Tree](tree T, activation func(T, int)) {
	var path []int = []int{0}
	var direction int = 0
	for len(path) != 0 {
		activation(tree, direction)
		if direction < len(tree.Children()) {
			path = append(path, direction)
			tree = tree.Children()[direction]
			direction = 0
		}
	}
}
