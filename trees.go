package trees

type Tree[T any] interface {
	Parent() Tree[T]
	Children() []Tree[T]
}

func Traversal[B any, T Tree[B]](tree T, activation func(T, int)) {
	var currentNode Tree[B] = tree
	if currentNode != nil {
		var direction int = 0
		var path []int = []int{0}
		for len(path) > 0 {
			activation(tree, direction)
			if direction < len(currentNode.Children()) {
				path = append(path, direction)
				direction = 0
			} else {
				direction = path[len(path)-1]
				path = path[:len(path)-1]
			}
		}
	}
}
