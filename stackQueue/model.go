package stackQueue

//Item placeholder for type
type Item interface{}

//Node represent node item
type Node struct {
	item Item
	next *Node
}

//Iterator provide mechanism to iterate item
type Iterator interface {
	hasNext() bool
	next() (Item, error)
	remove() (bool, error)
}
