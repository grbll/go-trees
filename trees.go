package trees

type ArbitraryTree interface {
	Parent() *ArbitraryTree
	Children() []*ArbitraryTree
}
