package internal

type QueueMethods interface {
	Check(str string)
	Remove() *Node
	Add(node Node)
	Display()
}

func NewCache() *Cache {
	return &Cache{}
}

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node
